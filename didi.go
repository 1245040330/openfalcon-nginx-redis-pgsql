package main

import (
	"awesomeProject/config"
	"awesomeProject/nginxInfo"
	"awesomeProject/pgsqlInfo"
	"awesomeProject/redisInfo"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/toolkits/pkg/file"
	"github.com/toolkits/pkg/runner"
	"os"
	"path"
)
var conf *string

// auto detect configuration file
func aconf() {
	conf = flag.String("f", "", "specify configuration file.")
	if *conf != "" && file.IsExist(*conf) {
		return
	}

	*conf = path.Join(runner.Cwd, "etc", "configuration.local.yml")
	if file.IsExist(*conf) {
		return
	}

	*conf = path.Join(runner.Cwd, "etc", "configuration.yml")
	if file.IsExist(*conf) {
		return
	}

	fmt.Println("no configuration file for sender")
	os.Exit(1)
}
func pconf() {
	if err := config.ParseConfig(*conf); err != nil {
		fmt.Println("cannot parse configuration file:", err)
		os.Exit(1)
	} else {
		fmt.Println("parse configuration file:", *conf)
	}
}

func main() {
	//读取静态资源
	aconf()
	pconf()
	//pg数据获取
	resPg:=pgsqlInfo.SqlInfo()
	config.Info.Println("pg数据提交返回\n"+resPg)
	//redis数据获取
	resRedis:=redisInfo.RedisInfo()
	config.Info.Println("redis数据提交返回\n"+resRedis)
	//nginx 获取数据
	resNginx:=nginxInfo.NginxInfo()
	config.Info.Println("nginx数据提交返回\n"+resNginx)
}
