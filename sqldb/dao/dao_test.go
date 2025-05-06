package dao

import (
	"root/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * hostIP:54320 --[portfowarding]--> 172.17.0.3:5432(docker postgres container)
 */
func TestPostgresInsertUser(t *testing.T) {
	dao := NewPostgres("postgres://postgres:1234@192.168.246.130:54320/postgres")
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	err := dao.InsertUser(model.User{UserName: "cjutesttest", Age: "111"})
	assert.NoError(t, err)
}

func TestPostgresSelectUser(t *testing.T) {
	dao := NewPostgres("postgres://postgres:1234@192.168.246.130:54320/postgres")
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	users, err := dao.SelectUser()
	assert.NoError(t, err)
	t.Log(users)
}

/**
 * hostIP:33061 --[portfowarding]--> 172.17.0.3:3306(docker mariadb container)
 */
func TestMariaInsertUser(t *testing.T) {
	dao := NewMariaDB("root:1234@192.168.246.130:33061?test")
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	err := dao.InsertUser(model.User{UserName: "cjutesttest", Age: "111"})
	assert.NoError(t, err)
}

func TestMariaSelectUser(t *testing.T) {
	dao := NewMariaDB("root:1234@192.168.246.130:33061?test")
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	users, err := dao.SelectUser()
	assert.NoError(t, err)
	t.Log(users)
}

/**
 * hostIP:33060 --[portfowarding]--> 172.17.0.2:3306(docker mysql container)
 */
func TestMySqlInsertUser(t *testing.T) {
	dao := NewMysqlDB("root:1234@192.168.246.130:33060?test")
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	err := dao.InsertUser(model.User{UserName: "cjutesttest", Age: "111"})
	assert.NoError(t, err)
}

func TestMySqlSelectUser(t *testing.T) {
	dao := NewMysqlDB("root:1234@192.168.246.130:33060?test")
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	users, err := dao.SelectUser()
	assert.NoError(t, err)
	t.Log(users)
}

/**
 * oralce 11g
 * requiremenets(build-time): C compiler(gcc) with CGO_ENABLED=1
 * requiremenets(run-time): oracle client instant with path( window )/LD_LIBRARY_PATH and PATH( linux )
 */
func TestOracleSelectUser(t *testing.T) {
	dao := NewOracleDB(`user="juchoi" password="1234" connectString="192.168.246.131:1521/xe"`)
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	users, err := dao.SelectUser()
	assert.NoError(t, err)
	t.Log(users)
}

func TestOracleInsertUser(t *testing.T) {
	dao := NewOracleDB(`user="juchoi" password="1234" connectString="THEROADORA11JE"`)
	assert.NotNil(t, dao)
	if dao == nil {
		return
	}
	defer dao.Close()
	user := model.User{
		UserName: "go-juchoi",
		Age:      "15211",
	}
	err := dao.InsertUser(user)
	assert.NoError(t, err)
}

/**
 * tibero7
 * requirement: tibero client instant and ODBC Manager configuration ( window )
 *              tibero client instant and unixODBC installation with *.ini file configuration ( linux/unix )
 */
func TestPingTibero(t *testing.T) {
	dao, err := NewTibero("tbodbc")
	assert.Nil(t, err)
	if err != nil {
		return
	}
	defer dao.Close()
	assert.Nil(t, dao.Ping())
}
