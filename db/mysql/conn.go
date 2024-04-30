package mysql

import (
	"cloud_storage/configurator"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var dataBase *sql.DB

func MySqlConn(location string) *sql.DB {
	cmd := getConnCmd(location)
	dataBase, err := sql.Open("mysql", cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	dataBase.SetMaxOpenConns(1000)
	err = dataBase.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return dataBase
}

func getConnCmd(location string) string {
	mysqlConfig, err := configurator.NewMysqlConfig().Parse(location)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	user := mysqlConfig.GetAttr("User")
	password := mysqlConfig.GetAttr("Password")
	ip := mysqlConfig.GetAttr("IP")
	port := mysqlConfig.GetAttr("Port")
	databaseName := mysqlConfig.GetAttr("DatabaseName")
	cmd := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, ip, port, databaseName)
	return cmd
}
