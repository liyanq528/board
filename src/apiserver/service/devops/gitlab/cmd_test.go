package gitlab_test

import (
	"git/inspursoft/board/src/apiserver/service/devops/gitlab"
	"git/inspursoft/board/src/common/model"
	"git/inspursoft/board/src/common/utils"
	"os"
	"testing"

	"github.com/astaxie/beego/logs"
	"github.com/stretchr/testify/assert"
)

var user = model.User{
	Username: "testuser18",
	Email:    "testuser18@inspur.com",
	Password: "123456a?",
}
var project = model.Project{
	Name: "myrepo",
}
var createdUser gitlab.UserCreation
var token gitlab.ImpersonationToken
var addSSHKeyResponse gitlab.AddSSHKeyResponse
var createdProject gitlab.ProjectCreation

var adminAccessToken = "si1Z1eUZUVui7XFarUyW"
var sshPubKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDaa046MKqllR1bE0pfPYwcVYHmBx291OzeWj5VHS6FCsVeLnky99pJigp3uwDz68uTDOx1I+zUU3XE39o4591isCCbM9ba5l2hKvGnHUoTRdG6Pkc9gy+OdKIJMGFca58Bt1hhPCa5FT8cQadsSnr7rGmg1O5tfG6a9mjzKFjn3nNNlYi5U6BsJxD3ReV5mVkFea5wH2yMzrHCSxTQiyLM8owB9Dem7Mrqz799sfB9MjC6ryVGwJd8oZOxGCB7hNz/Eenb+EUjdevxLFAVZgakTk4vDm/ubVfQjrdxGg4MaAbD4+kYNezEfh9c5W2uC0QlZHQhItEoMqytmWmjeZF7 root@10.110.25.227"

func TestMain(m *testing.M) {
	utils.InitializeDefaultConfig()
	os.Exit(m.Run())
}

func TestUserCreation(t *testing.T) {
	var err error
	createdUser, err = gitlab.NewGitlabHandler(adminAccessToken).CreateUser(user)
	assert.New(t).Nilf(err, "Error occurred while creating user via Gitlab API: %+v", err)
}

func TestImpersonateToken(t *testing.T) {
	var err error
	token, err = gitlab.NewGitlabHandler(adminAccessToken).ImpersonationToken(createdUser)
	logs.Debug("Impersonated token: %+v", token)
	assert.New(t).Nilf(err, "Error occurred while impersonating token via Gitlab API: %+v", err)
}

func TestAddSSHKey(t *testing.T) {
	var err error
	addSSHKeyResponse, err = gitlab.NewGitlabHandler(token.Token).AddSSHKey("user-ssh-key", sshPubKey)
	assert := assert.New(t)
	assert.Nilf(err, "Error occurred while adding SSH key via gitlab API: %+v", err)
	assert.NotNilf(addSSHKeyResponse, "Failed to get response after adding SSH key.", nil)

}

func TestCreateRepo(t *testing.T) {
	var err error
	createdProject, err = gitlab.NewGitlabHandler(token.Token).CreateRepo(user, project)
	assert.New(t).Nilf(err, "Error occurred while creating repo via Gitlab API: %+v", err)
}

func TestDeleteRepo(t *testing.T) {
	err := gitlab.NewGitlabHandler(token.Token).DeleteProject(createdProject.ID)
	assert.New(t).Nilf(err, "Error occurred while deleting project via Gitlab API: %+v", err)
}

func TestUserDeletion(t *testing.T) {
	err := gitlab.NewGitlabHandler(adminAccessToken).DeleteUser(createdUser.ID)
	assert.New(t).Nilf(err, "Error occurred while deleting user via Gitlab API: %+v", err)
}
