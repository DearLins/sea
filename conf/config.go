package conf

import (
	"flag"
	"io/ioutil"
	"github.com/go-yaml/yaml"
)

type Configuration struct {
	AppHost            string   `yaml:"app_host"`    // app_host
	AppPort            string   `yaml:"app_port"`   // app_port
	PageSize           int      `yaml:"page_size"`   // page_size

	MysqlHost          string   `yaml:"mysql_host"`
	MysqlPort 		   string  	`yaml:"mysql_port"`
	MysqlDatabase	   string   `yaml:"mysql_database"`
	MysqlUsername      string   `yaml:"mysql_username"`
	MysqlPassword      string   `yaml:"mysql_password"`
	MysqlCharset       string   `yaml:"mysql_charset"`
	MysqlCollation     string   `yaml:"mysql_collation"`
	MysqlPrefix        string   `yaml:"mysql_prefix"`

	RedisHost   	   string   `yaml:"redis_host"`
	RedisPwd   		   string   `yaml:"redis_pwd"`
	RedisPort 		   string   `yaml:"redis_port"`
	RedisDatabase 	   int  	`yaml:"redis_database"`

}

const (
	DEFAULT_PAGESIZE = 10
)

var configuration *Configuration

func LoadConfiguration(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	if config.PageSize <= 0 {
		config.PageSize = DEFAULT_PAGESIZE
	}
	configuration = &config
	return err
}

//获取配置
func GetConfiguration() *Configuration {
	return configuration
}

//读取配置
func init()  {
	configFilePath := flag.String("C", "conf/config.yaml", "config file path")
	flag.Parse()

	if err := LoadConfiguration(*configFilePath); err != nil {
		panic(err)
		return
	}
}

//conf.GetConfiguration().PageSize

