## web api of kman
#### Description
###### web api service
###### Usage
    go install github.com/kovey/kman/web
###### Configs
    create .env file
``` env
# kow config
APP_NAME       = shop
APP_TIME_ZONE  = Asia/Shanghai
APP_PPROF_OPEN = true
APP_ETCD_OPEN  = false
APP_NODE_ID    = 1001

# debug
DEBUG_LEVEL     = info
DEBUG_SHOW_FILE = 1

# listen
SERV_HOST = 0.0.0.0
SERV_PORT = 8081

# etcd
ETCD_ENDPOINTS = 127.0.0.1:2379
ETCD_TIMEOUT   = 30
ETCD_USERNAME  = 
ETCD_PASSWORD  =
ETCD_NAMESPACE = default

# service
SERVICE_NAME  = kman-service
SERVICE_GROUP = default

# database config
DB_DRIVER   = mysql
DB_HOST     = 127.0.0.1
DB_PORT     = 3306
DB_USER     = root
DB_PASSWORD = 123456
DB_NAME     = kman
DB_CHARSET  = utf8mb4
```
