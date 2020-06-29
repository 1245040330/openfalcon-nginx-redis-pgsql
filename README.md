# openfalcon-nginx-redis-pgsql
openfalcon 插件  集成nginx redis pgsql

##文件拷贝到 home文件夹下
```cassandraql
执行 cd /home/didi ./didi 检查文件运行是否报错 
```

## 添加到自动任务中 
```cassandraql
* * * * * root cd /home/didi/ && ./didi 
```


##执行二进制文件如果没有执行权限，先赋给文件权限

```cassandraql
chmod 777 didi 
```

## 执行前修改 etc/*.yml文件

```cassandraql
服务器redis 账号密码 
pgsql 账号密码 dbname 
当前服务器ip 
监控台数据接收接口 api   
```

## nginxStatus nginx信息接口 

```cassandraql
配置方法 参照 
location /monitor/nginx_status 
{ 
stub_status on;
 access_log off ;
 allow 127.0.0.1;
 deny all; 
} 
```


## pgsql汇报字段

--------------------------------
| key |  tag | type | note |
|-----|------|------|------|
|pgsql.maxconnection|port|GAUGE|已连接客户端的数量|

## redis汇报字段

--------------------------------
| key |  tag | type | note |
|-----|------|------|------|
|redis.connected_clients|port|GAUGE|已连接客户端的数量|
|redis.blocked_clients|port|GAUGE|正在等待阻塞命令（BLPOP、BRPOP、BRPOPLPUSH）的客户端的数量|
|redis.used_memory|port|GAUGE|由 Redis 分配器分配的内存总量，以字节（byte）为单位|
|redis.used_memory_rss|port|GAUGE| 从操作系统的角度，返回 Redis 已分配的内存总量（俗称常驻集大小）|
|redis.used_memory_peak|port|GAUGE|redis内存使用最大值|
|redis.instantaneous_ops_per_sec|port|GAUGE|每秒处理命令条数|
|redis.rejected_connections|port|COUNTER|采集周期内拒绝连接总数|
|redis.db0|port|GAUGE|redis现存key数量|


## nginx汇报字段

--------------------------------
| key |  tag | type | note |
|-----|------|------|------|
|nginx.query_count|port|GAUGE|接口调用次数|
|nginx.err_count|port|GAUGE|接口调用失败次数|