package serv

const (
	env_config = `# etcd
ETCD_ENDPOINTS = 127.0.0.1:2379
ETCD_TIMEOUT   = 30
ETCD_USERNAME  = 
ETCD_PASSWORD  =

# config dir
CONFIG_NAMESPACE = default1,default2
CONFIG_PATH      = /path/to/config1,/path/to/config2

# cache
CACHE_OPEN         = On
CACHE_PATH         = /path/to/cache
USE_CACHE_WHEN_ERR = Off`
)
