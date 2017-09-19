package service

import (
	"bufio"
	"errors"
	"fmt"
	"git/inspursoft/board/src/common/model"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

const (
	dockerTemplatePath  = "templates"
	dockerfileName      = "Dockerfile"
	templateNameDefault = "dockerfile-template"
)

func str2execform(str string) string {
	sli := strings.Split(str, " ")
	for num, node := range sli {
		sli[num] = "\"" + node + "\""
	}
	return strings.Join(sli, ", ")
}

func exec2str(str string) string {
	line := strings.TrimSpace(str)
	line = strings.TrimLeft(strings.TrimRight(line, "]"), "[")
	split := strings.Split(line, ",")
	for num, node := range split {
		node = strings.TrimSpace(node)
		split[num] = strings.TrimLeft(strings.TrimRight(node, "\""), "\"")
	}
	return strings.Join(split, " ")
}

func checkStringHasUpper(str ...string) error {
	for _, node := range str {
		isMatch, err := regexp.MatchString("[A-Z]", node)
		if err != nil {
			return err
		}
		if isMatch {
			errString := fmt.Sprintf(`string "%s" has upper charactor`, node)
			return errors.New(errString)
		}
	}
	return nil
}

func checkStringHasEnter(str ...string) error {
	for _, node := range str {
		isMatch, err := regexp.MatchString(`^\s*\n?(?:.*[^\n])*\n?\s*$`, node)
		if err != nil {
			return err
		}
		if !isMatch {
			errString := fmt.Sprintf(`string "%s" has enter charactor`, node)
			return errors.New(errString)
		}
	}
	return nil
}

func fixStructEmptyIssue(obj interface{}) {
	if f, ok := obj.(*[]string); ok {
		if len(*f) == 1 && len((*f)[0]) == 0 {
			*f = nil
		}
		return
	}
	if f, ok := obj.(*[]model.CopyStruct); ok {
		if len(*f) == 1 && len((*f)[0].CopyFrom) == 0 && len((*f)[0].CopyTo) == 0 {
			*f = nil
		}
		return
	}
	if f, ok := obj.(*[]model.EnvStruct); ok {
		if len(*f) == 1 && len((*f)[0].EnvName) == 0 && len((*f)[0].EnvValue) == 0 {
			*f = nil
		}
		return
	}
	if f, ok := obj.(*[]int); ok {
		if len(*f) == 1 && (*f)[0] == 0 {
			*f = nil
		}
	}
	return
}

func changeDockerfileStructItem(dockerfile *model.Dockerfile) {
	dockerfile.Base = strings.TrimSpace(dockerfile.Base)
	dockerfile.Author = strings.TrimSpace(dockerfile.Author)
	dockerfile.EntryPoint = strings.TrimSpace(dockerfile.EntryPoint)
	dockerfile.Command = strings.TrimSpace(dockerfile.Command)

	for num, node := range dockerfile.Volume {
		dockerfile.Volume[num] = strings.TrimSpace(node)
	}
	fixStructEmptyIssue(&dockerfile.Volume)

	for num, node := range dockerfile.Copy {
		dockerfile.Copy[num].CopyFrom = strings.TrimSpace(node.CopyFrom)
		dockerfile.Copy[num].CopyTo = strings.TrimSpace(node.CopyTo)
	}
	fixStructEmptyIssue(&dockerfile.Copy)

	for num, node := range dockerfile.RUN {
		dockerfile.RUN[num] = strings.TrimSpace(node)
	}
	fixStructEmptyIssue(&dockerfile.RUN)

	for num, node := range dockerfile.EnvList {
		dockerfile.EnvList[num].EnvName = strings.TrimSpace(node.EnvName)
		dockerfile.EnvList[num].EnvValue = strings.TrimSpace(node.EnvValue)
	}
	fixStructEmptyIssue(&dockerfile.EnvList)

	for num, node := range dockerfile.ExposePort {
		dockerfile.ExposePort[num] = strings.TrimSpace(node)
	}
	fixStructEmptyIssue(&dockerfile.ExposePort)
}

func changeImageConfigStructItem(reqImageConfig *model.ImageConfig) {
	reqImageConfig.ImageName = strings.TrimSpace(reqImageConfig.ImageName)
	reqImageConfig.ImageTag = strings.TrimSpace(reqImageConfig.ImageTag)
	reqImageConfig.ProjectName = strings.TrimSpace(reqImageConfig.ProjectName)
	reqImageConfig.ImageTemplate = strings.TrimSpace(reqImageConfig.ImageTemplate)
	reqImageConfig.ImageDockerfilePath = strings.TrimSpace(reqImageConfig.ImageDockerfilePath)
}

func CheckDockerfileItem(dockerfile *model.Dockerfile) error {
	changeDockerfileStructItem(dockerfile)

	if len(dockerfile.Base) == 0 {
		return errors.New("Baseimage in dockerfile should not be empty")
	}

	if err := checkStringHasUpper(dockerfile.Base); err != nil {
		return err
	}

	if err := checkStringHasEnter(dockerfile.EntryPoint, dockerfile.Command); err != nil {
		return err
	}

	for _, node := range dockerfile.ExposePort {
		if _, err := strconv.Atoi(node); err != nil {
			return err
		}
	}

	return nil
}

func CheckDockerfileConfig(config *model.ImageConfig) error {
	changeImageConfigStructItem(config)

	if err := checkStringHasUpper(config.ImageName, config.ImageTag); err != nil {
		return err
	}

	return CheckDockerfileItem(&config.ImageDockerfile)
}

func BuildDockerfile(reqImageConfig model.ImageConfig, wr ...io.Writer) error {
	var templatename string

	if len(reqImageConfig.ImageTemplate) != 0 {
		templatename = reqImageConfig.ImageTemplate
	} else {
		templatename = templateNameDefault
	}

	tmpl, err := template.New(templatename).Funcs(template.FuncMap{"str2exec": str2execform}).ParseFiles(filepath.Join(dockerTemplatePath, templatename))
	if err != nil {
		return err
	}

	if len(wr) != 0 {
		if err = tmpl.Execute(wr[0], reqImageConfig.ImageDockerfile); err != nil {
			return err
		}
		return nil
	}

	if fi, err := os.Stat(reqImageConfig.ImageDockerfilePath); os.IsNotExist(err) {
		if err := os.MkdirAll(reqImageConfig.ImageDockerfilePath, 0755); err != nil {
			return err
		}
	} else if !fi.IsDir() {
		return errors.New("Dockerfile path is not dir")
	}

	dockerfile, err := os.OpenFile(filepath.Join(reqImageConfig.ImageDockerfilePath, dockerfileName), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer dockerfile.Close()

	err = tmpl.Execute(dockerfile, reqImageConfig.ImageDockerfile)
	if err != nil {
		return err
	}

	return nil
}

func ImageConfigClean(path string) error {
	//remove Tag dir
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	//remove Image dir
	if fi, err := os.Stat(filepath.Dir(path)); os.IsNotExist(err) {
		return nil
	} else if !fi.IsDir() {
		errMsg := fmt.Sprintf(`%s is not dir`, filepath.Dir(path))
		return errors.New(errMsg)
	}

	parent, err := os.Open(filepath.Dir(path))
	if err != nil {
		return err
	}
	defer parent.Close()

	_, err = parent.Readdirnames(1)
	if err == io.EOF {
		return os.RemoveAll(filepath.Dir(path))
	}

	return nil
}

func GetDockerfileInfo(path string) (*model.Dockerfile, error) {
	var Dockerfile model.Dockerfile
	var fulline string
	dockerfile, err := os.Open(filepath.Join(path, "Dockerfile"))
	if err != nil {
		return nil, err
	}
	defer dockerfile.Close()

	scanner := bufio.NewScanner(dockerfile)
	for scanner.Scan() {
		if strings.HasPrefix(strings.TrimSpace(scanner.Text()), "#") {
			continue
		}
		fulline += string(scanner.Text())
		if strings.HasSuffix(scanner.Text(), "\\") {
			fulline = fulline[:len(fulline)-1]
			continue
		}
		split := strings.SplitN(strings.TrimSpace(fulline), " ", 2)
		fulline = ""

		// ignore empty line and lines with only one field
		if len(split) < 2 {
			continue
		}
		split[1] = strings.TrimSpace(split[1])
		switch split[0] {
		case "FROM":
			Dockerfile.Base = split[1]
		case "MAINTAINER":
			Dockerfile.Author = split[1]
		case "VOLUME":
			Dockerfile.Volume = append(Dockerfile.Volume, exec2str(split[1]))
		case "COPY":
			{
				var node model.CopyStruct
				var copyfrom, copyto string
				copystring := exec2str(split[1])
				split_copy := strings.Split(strings.TrimSpace(copystring), " ")
				copyfrom = strings.Join(split_copy[:len(split_copy)-1], " ")
				copyto = split_copy[len(split_copy)-1]
				node.CopyFrom = copyfrom
				node.CopyTo = copyto
				Dockerfile.Copy = append(Dockerfile.Copy, node)
			}
		case "RUN":
			Dockerfile.RUN = append(Dockerfile.RUN, split[1])
		case "ENTRYPOINT":
			Dockerfile.EntryPoint = exec2str(split[1])
		case "CMD":
			Dockerfile.Command = exec2str(split[1])
		case "ENV":
			{
				var node model.EnvStruct
				envstring := split[1]
				split_env := strings.SplitN(envstring, " ", 2)
				node.EnvName = split_env[0]
				node.EnvValue = strings.TrimSpace(split_env[1])
				Dockerfile.EnvList = append(Dockerfile.EnvList, node)
			}
		case "EXPOSE":
			Dockerfile.ExposePort = append(Dockerfile.ExposePort, split[1])
		}
	}

	return &Dockerfile, nil
}