package dao

import (
	"root/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

/***
* hostIP:33060 --[포트포워딩]--> 172.17.0.2:3306(docker mysql container)
***/
func TestMySqlInsertUser(t *testing.T) {
	dao := NewMysqlDB("root:1234@192.168.246.130:33060?test")
	defer dao.Close()
	err := dao.InsertUser(model.User{UserName: "cjutesttest", Age: 111})
	assert.NoError(t, err)
}

func TestMySqlSelectUser(t *testing.T) {
	dao := NewMysqlDB("root:1234@192.168.246.130:33060?test")
	defer dao.Close()
	users, err := dao.SelectUser()
	assert.NoError(t, err)
	t.Log(users)
}

/***
* oracle 11g
* requiremenets(build-time): C compiler with CGO_ENABLED=1
* requiremenets(run-time): oracle client instant with path( window )/LD_LIBRARY_PATH and PATH( linux )
***/
func TestOracleSelectUser(t *testing.T) {
	dao := NewOracleDB(`user="juchoi" password="1234" connectString="192.168.246.131:1521/xe"`)
	assert.NotNil(t, dao)
	defer dao.Close()
	users, err := dao.SelectUser()
	assert.NoError(t, err)
	t.Log(users)
}

func TestOracleInsertUser(t *testing.T) {
	dao := NewOracleDB(`user="juchoi" password="1234" connectString="THEROADORA11JE"`)
	assert.NotNil(t, dao)
	defer dao.Close()
	user := model.User{
		UserName: "go-juchoi",
		Age:      15211,
	}
	err := dao.InsertUser(user)
	assert.NoError(t, err)
}
