package config

import (
	"fmt"

	"github.com/toolkits/pkg/file"
)

type Config struct {
	Redis    redisSection    `yaml:"redis"`
	Ip       string         `yaml:"ip"`
	Api       string         `yaml:"api"`
	NginxStatus       string         `yaml:"nginxStatus"`
	Pgsql    pgsqlSection  `yaml:"pgsql"`

}
type pgsqlSection struct {
	Port    string         `yaml:"port"`
	Pass    string         `yaml:"pass"`
	Dbname  string          `yaml:"dbname"`
}
type ipSection struct {
	thisIp       string `yaml:"thisIp"`
}
type loggerSection struct {
	Dir       string `yaml:"dir"`
	Level     string `yaml:"level"`
	KeepHours uint   `yaml:"keepHours"`
}

type redisSection struct {
	Addr    string         `yaml:"addr"`
	Pass    string         `yaml:"pass"`
}

type timeoutSection struct {
	Conn  int `yaml:"conn"`
	Read  int `yaml:"read"`
	Write int `yaml:"write"`
}

type consumerSection struct {
	Queue  string `yaml:"queue"`
	Worker int    `yaml:"worker"`
}

type smtpSection struct {
	User               string `yaml:"user"`
	Pass               string `yaml:"pass"`
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	InsecureSkipVerify bool   `yaml:"insecureSkipVerify"`
}

var yaml Config

func Get() Config {
	return yaml
}

func ParseConfig(yf string) error {
	err := file.ReadYaml(yf, &yaml)
	if err != nil {
		return fmt.Errorf("cannot read yml[%s]: %v", yf, err)
	}
	return nil
}
