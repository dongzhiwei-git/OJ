package conf

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Mysql NewMysql `yaml:"mysql" json:"mysql"`
}

type NewMysql struct {
	Master MysqlTest `yaml:"master"`
	Slave  MysqlTest `yaml:"slave"`
}

type MysqlTest struct {
	Dsn string `yaml:"dsn" json:"dsn"`
}

var conf Conf

func Get() Conf {
	return conf
}

func Init() error {
	config, err := ioutil.ReadFile("./conf/web.yaml")
	if err != nil {
		log.Println("yaml parse error", err)
	}
	yaml.Unmarshal(config, &conf)
	fmt.Print("qwe", conf.Mysql.Master.Dsn)
	return nil
}
