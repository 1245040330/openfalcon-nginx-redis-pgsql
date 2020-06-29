package pgsqlInfo

import (
	"awesomeProject/config"
	"awesomeProject/util"
	"database/sql"
	"encoding/json"
	"time"
)

//pgsql数据采集
func SqlInfo() string{
	db, err := sql.Open("postgres", "port="+config.Get().Pgsql.Port+" user=postgres password="+config.Get().Pgsql.Pass+" dbname="+config.Get().Pgsql.Dbname+" sslmode=disable")
	if err!=nil{
		config.Error.Println(err)
	}
	config.Info.Println("pgsqlOpen")
	count:=sqlSelect(db)
	config.Info.Println("*sqlSelect")
	db.Close()
	config.Info.Println("*sqlClose")

	res:=pgsqlPostData(count)
	return res
}

//查询连接数
func sqlSelect(db *sql.DB) int64{
	rows, err := db.Query("select count('count') from pg_stat_activity")
	if err!=nil{
		config.Error.Println(err)
	}
	var count int64
	for rows.Next() {
		err = rows.Scan(&count)
	}
	return count
}

func pgsqlPostData(count int64) string {
	data := []config.Data{
		config.Data{
			Metric:    "pgsql.maxconnection",
			Endpoint:  config.Get().Ip,
			Tags:      "name=maxconnection",
			Value:     int(count),
			Timestamp: int(time.Now().Unix()),
			Step:      60,
		},
	}
	jsonStr, err := json.MarshalIndent(data, "", " ")
	if err!=nil{
		config.Error.Println(err)
	}
	config.Info.Println("formated: ", string(jsonStr))
	res := util.Post(config.Get().Api, jsonStr, "application/json")
	return res
}
