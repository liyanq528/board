package service

import (
	"encoding/json"
	"errors"
	"git/inspursoft/board/src/common/model"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/logs"
	"github.com/ghodss/yaml"
)

var loadPath string

const (
	serviceAPIVersion    = "v1"
	serviceKind          = "Service"
	nodePort             = "NodePort"
	deploymentAPIVersion = "extensions/v1beta1"
	deploymentKind       = "Deployment"
	maxPort              = 32765
	minPort              = 30000
)

var (
	pathErr                        = errors.New("ERR_DEPLOYMENT_PATH_NOT_DIRECTORY")
	emptyServiceNameErr            = errors.New("ERR_NO_SERVICE_NAME")
	portMaxErr                     = errors.New("ERR_SERVICE_NODEPORT_EXCEED_MAX_LIMIT")
	portMinErr                     = errors.New("ERR_SERVICE_NODEPORT_EXCEED_MIN_LIMIT")
	emptyDeployErr                 = errors.New("ERR_NO_DEPLOYMENT_NAME")
	invalidErr                     = errors.New("ERR_DEPLOYMENT_REPLICAS_INVAILD")
	emptyContainerErr              = errors.New("ERR_NO_CONTAINER")
	NameInconsistentErr            = errors.New("ERR_SERVICE_NAME_AND_DEPLOYMENT_NAME_INCONSISTENT")
	ServiceNameInconsistentErr     = errors.New("ERR_SERVICE_NAME_INCONSISTENT_WITH_YAML_FILE")
	ProjectNameInconsistentErr     = errors.New("ERR_PROJECT_NAME_INCONSISTENT_WITH_YAML_FILE")
	DeploymentNotFoundErr          = errors.New("ERR_DEPLOYMENT_NOT_FOUND")
	ServiceNotFoundErr             = errors.New("ERR_SERVICE_NOT_FOUND")
	ServiceYamlFileUnmarshalErr    = errors.New("ERR_SERVICE_YAML_FILE_UNMARSHAL")
	DeploymentYamlFileUnmarshalErr = errors.New("ERR_DEPLOYMENT_YAML_FILE_UNMARSHAL")
	deploymentKindErr              = errors.New("ERR_DEPLOYMENT_YAML_KIND")
	serviceKindErr                 = errors.New("ERR_SERVICE_YAML_KIND")
	deploymentAPIVersionErr        = errors.New("ERR_DEPLOYMENT_YAML_API_VERSION")
	serviceAPIVersionErr           = errors.New("ERR_SERVICE_YAML_API_VERSION")
)

func CheckDeploymentPath(loadPath string) error {
	if fi, err := os.Stat(loadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(loadPath, 0755); err != nil {
			return err
		}
	} else if !fi.IsDir() {
		return pathErr
	}

	return nil
}

func ServiceExists(serviceName string, projectName string) (bool, error) {
	var servicequery model.ServiceStatus
	servicequery.Name = serviceName
	servicequery.ProjectName = projectName
	s, err := GetService(servicequery, "name", "project_name")

	return s != nil, err
}

func GenerateYamlFile(name string, structdata interface{}) error {
	info, err := json.Marshal(structdata)
	if err != nil {
		logs.Error("Marhal json failed, err:%+v\n", err)
		return err
	}

	context, err := yaml.JSONToYAML(info)
	if err != nil {
		logs.Error("Generate yaml data failed, err:%+v\n", err)
		return err
	}

	err = ioutil.WriteFile(name, context, 0644)
	if err != nil {
		logs.Error("Generate yaml file failed, err:%+v\n", err)
		return err
	}
	return nil
}

func GenerateDeploymentYamlFileFromK8S(deployConfigURL string, absFileName string) error {
	deployConfig, err := GetDeployConfig(deployConfigURL)
	if err != nil {
		return err
	}
	return GenerateYamlFile(absFileName, &deployConfig)
}

func GenerateServiceYamlFileFromK8S(serviceConfigURL string, absFileName string) error {
	serviceConfig, err := GetServiceStatus(serviceConfigURL)
	if err != nil {
		return err
	}
	return GenerateYamlFile(absFileName, &serviceConfig)
}

func DeleteServiceConfigYaml(serviceConfigPath string) error {
	err := os.RemoveAll(serviceConfigPath)
	if err != nil {
		logs.Error("Failed to delete deployment files: %+v\n", err)
		return err
	}

	return nil
}

func getYamlFileData(serviceConfig interface{}, serviceConfigPath string, fileName string) error {
	serviceFileName := filepath.Join(serviceConfigPath, fileName)
	yamlData, err := ioutil.ReadFile(serviceFileName)
	if err != nil {
		return err
	}

	jsonData, err := yaml.YAMLToJSON(yamlData)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, serviceConfig)
	if err != nil {
		return err
	}

	return nil
}
