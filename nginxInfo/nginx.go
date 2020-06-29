package nginxInfo

import (
	"awesomeProject/config"
	"awesomeProject/util"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

func NginxInfo() string{
	var res = util.Get(config.Get().NginxStatus, "")
	resNginxArr := strings.Split(res, "\n")
	//config.Info.Println(resNginxArr[2])
	count := strings.Split(resNginxArr[2], " ")
	successCount, err := strconv.Atoi(count[1])
	if err!=nil{
		config.Error.Println(err)
	}
	handsCount, err := strconv.Atoi(count[2])
	if err!=nil{
		config.Error.Println(err)
	}
	failCount := handsCount - successCount

	data := []config.Data{
		config.Data{
			Metric:    "nginx.query_count",
			Endpoint:  config.Get().Ip,
			Tags:      "name=query_count",
			Value:     successCount,
			Timestamp: int(time.Now().Unix()),
			Step:      60,
		},
		config.Data{
			Metric:    "nginx.err_count",
			Endpoint:  config.Get().Ip,
			Tags:      "name=err_count",
			Value:     failCount,
			Timestamp: int(time.Now().Unix()),
			Step:      60,
		},
	}
	jsonStr, err := json.MarshalIndent(data, "", " ")
	if err!=nil{
		config.Error.Println(err)
	}
	//config.Info.Println("formated: ", string(jsonStr))
	resData := util.Post(config.Get().Api, jsonStr, "application/json")
	return resData
}

