package event

const (
	env_config = `# kom config
APP_NAME       = kman
APP_TIME_ZONE  = Asia/Shanghai
APP_PPROF_OPEN = true
APP_ETCD_OPEN  = false
APP_TEST_OPEN  = false
APP_NODE_ID    = 1001

# debug
DEBUG_LEVEL     = info
DEBUG_SHOW_FILE = 1

# etcd
ETCD_ENDPOINTS = 172.16.26.155:2379
ETCD_TIMEOUT   = 30
ETCD_USERNAME  = 
ETCD_PASSWORD  =
ETCD_NAMESPACE = default

# database config
DB_DRIVER   = mysql
DB_HOST     = 127.0.0.1
DB_PORT     = 3306
DB_USER     = root
DB_PASSWORD = 123456
DB_NAME     = kman
DB_CHARSET  = utf8mb4

# models path
MODELS_PATH   = module/models

# listen
SERV_HOST    = 0.0.0.0
SERV_PORT    = 8081
SERV_TTL     = 10
SERV_NAME    = kom 
SERV_GROUP   = default 
SERV_WEIGHT  = 1 
SERV_VERSION = 1.0.0`
)
