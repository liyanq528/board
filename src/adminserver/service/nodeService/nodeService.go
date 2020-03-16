package nodeService

import (
	"bufio"
	"bytes"
	"fmt"
	"git/inspursoft/board/src/adminserver/dao"
	"git/inspursoft/board/src/adminserver/dao/nodeDao"
	"git/inspursoft/board/src/adminserver/models/nodeModel"
	"git/inspursoft/board/src/adminserver/service"
	"git/inspursoft/board/src/adminserver/utils"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

func AddRemoveNodeByContainer(nodePostData *nodeModel.AddNodePostData,
	actionType nodeModel.ActionType, yamlFile string) (*nodeModel.NodeLog, error) {
	configuration, statusMessage := service.GetAllCfg("")
	if statusMessage == "BadRequest" {
		return nil, fmt.Errorf("failed to get the configuration")
	}
	hostName := configuration.Apiserver.Hostname
	masterIp := configuration.Apiserver.KubeMasterIP
	registryIp := configuration.Apiserver.RegistryIP

	hostFilePath := path.Join(nodeModel.BasePath, nodeModel.NodeHostsFile)
	hostFileName := fmt.Sprintf("%s@%s", hostFilePath, nodePostData.NodeIp)
	if err := GenerateHostFile(masterIp, nodePostData.NodeIp, registryIp, hostFileName); err != nil {
		return nil, err
	}

	var newLogId int64
	nodeLog := nodeModel.NodeLog{
		Ip: nodePostData.NodeIp, Completed: false, Success: false, CreationTime: time.Now().Unix(), LogType: actionType}
	if id, err := InsertLog(&nodeLog); err != nil {
		return nil, err
	} else {
		newLogId = id
	}

	var containerEnv = nodeModel.ContainerEnv{NodeIp: nodePostData.NodeIp,
		NodePassword:   nodePostData.NodePassword,
		HostIp:         hostName,
		HostPassword:   nodePostData.HostPassword,
		HostUserName:   nodePostData.HostUsername,
		HostFile:       hostFileName,
		MasterIp:       masterIp,
		MasterPassword: nodePostData.MasterPassword,
		InstallFile:    yamlFile,
		LogId:          newLogId}

	if err := LaunchAnsibleContainer(&containerEnv); err != nil {
		return nil, err
	}
	return &nodeLog, nil
}

func LaunchAnsibleContainer(env *nodeModel.ContainerEnv) error {
	logCache := dao.GlobalCache.Get(env.NodeIp).(*nodeModel.NodeLogCache)
	var secureShell *utils.SecureShell
	var err error
	secureShell, err = utils.NewSecureShell(&logCache.DetailBuffer, env.HostIp, env.HostUserName, env.HostPassword)
	if err != nil {
		return err
	}

	envStr := fmt.Sprintf(""+
		"--env MASTER_PASS=%s --env MASTER_IP=%s --env NODE_IP=%s --env NODE_PASS=%s --env LOG_ID=%d "+
		"--env ADMIN_SERVER_IP=%s --env ADMIN_SERVER_PORT=%d --env INSTALL_FILE=%s --env HOSTS_FILE=%s ",
		env.MasterPassword, env.MasterIp, env.NodeIp, env.NodePassword, env.LogId,
		env.HostIp, 8081, env.InstallFile, env.HostFile)

	LogFilePath := path.Join(nodeModel.BasePath, nodeModel.LogFileDir)
	PreEnvPath := path.Join(nodeModel.BasePath, nodeModel.PreEnvDir)

	cmdStr := fmt.Sprintf(`docker run -td -v %s:/tmp/log -v %s:/ansible_k8s/pre-env %s an:1`,
		LogFilePath, PreEnvPath, envStr)
	logs.Info(cmdStr)

	err = secureShell.ExecuteCommand(cmdStr)
	if err != nil {
		return err
	}
	return nil
}

func UpdateLog(putLogData *nodeModel.UpdateNodeLog) error {
	var logData *nodeModel.NodeLog
	var err error
	logData, err = nodeDao.GetNodeLog(putLogData.LogId);
	if err != nil {
		return err
	}

	logData.Success = putLogData.Success
	logData.Completed = true
	if errUpdate := nodeDao.UpdateNodeLog(logData); errUpdate != nil {
		return errUpdate
	}

	logFileName := path.Join(nodeModel.BasePath, putLogData.LogFile)
	if errInsert := InsertLogDetail(logData.Ip, logFileName, logData.CreationTime); errInsert != nil {
		return errInsert
	}

	if putLogData.Success {
		if putLogData.InstallFile == nodeModel.AddNodeYamlFile {
			if err := nodeDao.InsertNodeStatus(&nodeModel.NodeStatus{
				Ip:           putLogData.Ip,
				CreationTime: logData.CreationTime}); err != nil {
				logs.Info(err.Error())
			}
		}
		if putLogData.InstallFile == nodeModel.RemoveNodeYamlFile {
			if err := nodeDao.DeleteNodeStatus(&nodeModel.NodeStatus{Ip: putLogData.Ip}); err != nil {
				logs.Info(err.Error())
			}
		}
	}
	dao.GlobalCache.Delete(putLogData.Ip)
	return nil
}

func InsertLog(nodeLog *nodeModel.NodeLog) (int64, error) {
	var nodeLogCache = nodeModel.NodeLogCache{}
	nodeLogCache.NodeLogPtr = nodeLog
	if err := dao.GlobalCache.Put(nodeLog.Ip, &nodeLogCache, 3600*time.Second); err != nil {
		return 0, err
	}

	if nodeLog.LogType == nodeModel.ActionTypeAddNode {
		nodeLogCache.DetailBuffer.WriteString(fmt.Sprintf("---Begin add node:%s----", nodeLog.Ip))
	} else {
		nodeLogCache.DetailBuffer.WriteString(fmt.Sprintf("---Begin remove node:%s----", nodeLog.Ip))
	}

	if newId, err := nodeDao.InsertNodeLog(nodeLog); err != nil {
		return 0, err
	} else {
		return newId, nil
	}
}

func InsertLogDetail(ip, logFileName string, creationTime int64) error {
	if dao.GlobalCache.IsExist(ip) {
		logs.Info(fmt.Sprintf("No cache data for node:%s", ip))
		return nil
	}
	logCache := dao.GlobalCache.Get(ip).(*nodeModel.NodeLogCache)

	filePtr, _ := os.Open(logFileName)
	defer filePtr.Close()

	if fileContent, err := ioutil.ReadFile(logFileName); err != nil {
		return err
	} else {
		if _, writeErr := logCache.DetailBuffer.Write(fileContent); writeErr != nil {
			return nil
		}
		logCache.DetailBuffer.WriteString("---End---")
	}

	detail := nodeModel.NodeLogDetailInfo{
		CreationTime: creationTime,
		Detail:       logCache.DetailBuffer.String()}
	if _, err := nodeDao.InsertNodeLogDetail(&detail); err != nil {
		logs.Info(err.Error())
	}
	return nil
}

func GetNodeStatusList(nodeStatusList *[]nodeModel.NodeStatus) error {
	return nodeDao.GetNodeStatusList(nodeStatusList)
}

func GetPaginatedNodeLogList(v *nodeModel.PaginatedNodeLogList) error {
	offset := (v.Pagination.PageIndex - 1) * v.Pagination.PageSize
	if err := nodeDao.GetNodeLogList(v.LogList, v.Pagination.PageSize, offset); err != nil {
		return err
	}

	totalCount, err := nodeDao.GetLogTotalRecordCount()
	if err != nil {
		return err
	}
	v.Pagination.TotalCount = totalCount
	v.Pagination.PageCount = int(v.Pagination.TotalCount)/v.Pagination.PageSize + 1
	return nil
}

func GetNodeLogDetail(logTimestamp int64, nodeIp string, nodeLogDetail *[]nodeModel.NodeLogDetail) error {
	var reader *bufio.Reader
	if CheckExistsInCache(nodeIp) {
		logCache := dao.GlobalCache.Get(nodeIp).(*nodeModel.NodeLogCache)
		reader = bufio.NewReader(strings.NewReader(logCache.DetailBuffer.String()))
	} else {
		detailInfo, err := nodeDao.GetNodeLogDetail(logTimestamp)
		if err != nil {
			return err
		}
		reader = bufio.NewReader(bytes.NewBufferString(detailInfo.Detail))
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			errorMsg := fmt.Sprintf("Unexpected error occurred.%s", err.Error())
			return fmt.Errorf(errorMsg)
		}
		line = strings.TrimSpace(line)
		detail := setLogStatus(line)
		*nodeLogDetail = append(*nodeLogDetail, detail)
	}

	return nil
}

func GenerateHostFile(masterIp, nodeIp, registryIp, nodePathFile string) error {
	addHosts, err := os.Create(nodePathFile)
	defer addHosts.Close()
	if err != nil {
		return err
	}
	addHosts.WriteString("[masters]\n")
	addHosts.WriteString(fmt.Sprintf("%s\n", masterIp))
	addHosts.WriteString("[etcd]\n")
	addHosts.WriteString(fmt.Sprintf("%s\n", masterIp))
	addHosts.WriteString("[nodes]\n")
	addHosts.WriteString(fmt.Sprintf("%s\n", nodeIp))
	addHosts.WriteString("[registry]\n")
	addHosts.WriteString(fmt.Sprintf("%s\n", registryIp))
	return nil
}

func CheckExistsInCache(nodeIp string) bool {
	return dao.GlobalCache.IsExist(nodeIp)
}

func GetLogInfoInCache(nodeIp string) *nodeModel.NodeLog {
	logCache := dao.GlobalCache.Get(nodeIp).(*nodeModel.NodeLogCache)
	return logCache.NodeLogPtr
}

func checkIsEndingLog(log string) bool {
	return len(log) > 0 &&
		strings.Contains(log, "ok") &&
		strings.Contains(log, "changed") &&
		strings.Contains(log, "unreachable") &&
		strings.Contains(log, "failed")
}

func checkIsSuccessExecuted(log string) bool {
	return len(log) > 0 &&
		log[strings.Index(log, "unreachable")+12:strings.Index(log, "unreachable")+13] == "0" &&
		log[strings.Index(log, "failed")+7:strings.Index(log, "failed")+8] == "0"
}

func setLogStatus(log string) nodeModel.NodeLogDetail {
	if strings.Index(log, "fatal") == 0 || strings.Index(log, "failed") == 0 {
		return nodeModel.NodeLogDetail{Message: log, Status: nodeModel.NodeLogResponseError}
	} else if strings.Index(log, "TASK") == 0 {
		return nodeModel.NodeLogDetail{Message: log, Status: nodeModel.NodeLogResponseStart}
	} else if strings.Index(log, "ok") == 0 || strings.Index(log, "changed") == 0 {
		return nodeModel.NodeLogDetail{Message: log, Status: nodeModel.NodeLogResponseWarning}
	} else if strings.Index(log, "---") == 0 {
		return nodeModel.NodeLogDetail{Message: log, Status: nodeModel.NodeLogResponseStart}
	} else if checkIsEndingLog(log) {
		if checkIsSuccessExecuted(log) {
			return nodeModel.NodeLogDetail{Message: log, Status: nodeModel.NodeLogResponseSuccess}
		}
		return nodeModel.NodeLogDetail{Message: log, Status: nodeModel.NodeLogResponseFailed}
	} else {
		return nodeModel.NodeLogDetail{Message: log, Status: nodeModel.NodeLogResponseNormal}
	}
}
