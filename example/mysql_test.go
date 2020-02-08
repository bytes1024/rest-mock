package example

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

type MysqlConfig struct {
	UserName string
	Password string
	Host     string
	Port     uint
	DataBase string
	Charset  string

	MaxOpenConnection int
	MaxIdleConnection int

}

var DB *sql.DB
var MysqlDbErr error

//初始信息
func init() {

	mysqlConfig := loadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DataBase, mysqlConfig.Charset)

	// 打开连接失败
	log.Info("mysql dsn: ",dsn)
	DB, MysqlDbErr = gorm.Open("mysql", dsn)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dsn: " + dsn)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	DB.SetMaxOpenConns(mysqlConfig.MaxOpenConnection)
	// 闲置连接数
	DB.SetMaxIdleConns(mysqlConfig.MaxIdleConnection)
	// 最大连接周期
	DB.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = DB.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}

func loadConfig() MysqlConfig {
	return MysqlConfig{
		UserName:          "root",
		Password:          "123456",
		Host:              "127.0.0.1",
		Port:              3306,
		DataBase:          "test",
		Charset:           "UTF-8",
		MaxOpenConnection: 50,
		MaxIdleConnection: 20,
	}
}

func TestClientMysql(t *testing.T) {

}
