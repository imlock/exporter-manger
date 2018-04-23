package config

import (
	"github.com/astaxie/beego"
	"time"
	"strings"
)

type GlobalConfig struct {
	NodeCfg Node
	NginxCfg Nginx
	PhpfpmCfg Phpfpm
	RedisCfg Redis
	MemcachedCfg Memcached
}

type BaseConfig struct {
	Enable bool
	UnixSockFile string
	MetricsPath string
	MetricsRouter string
	Timeout time.Duration
	Filters map[string]string
}

type Node struct {
	BaseCfg BaseConfig
}

type Nginx struct {
	BaseCfg BaseConfig
}

type Phpfpm struct {
	BaseCfg BaseConfig
}

type Redis struct {
	BaseCfg BaseConfig
}

type Memcached struct {
	BaseCfg BaseConfig
}

var globalCfg GlobalConfig

func Init() {
	// node exporter
	globalCfg.NodeCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("node_exporter", false)
	if NodeConfig().BaseCfg.Enable {
		globalCfg.NodeCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("node_exporter.unix_sock",
			"/dev/shm/node_exporter.sock")
		globalCfg.NodeCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("node_exporter.metrics_path",
			"/metrics")
		globalCfg.NodeCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("node_exporter.metrics_router",
			"/node")
		globalCfg.NodeCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("node_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("node_exporter.filters", ""),
			",")
		if len(filters) > 0 {
			var m map[string]string
			m = make(map[string]string)
			for _, s := range filters {
				if len(s) == 0 {
					continue
				}
				m[s] = s
			}
			globalCfg.NodeCfg.BaseCfg.Filters = m
		}
	}

	// nginx exporter
	globalCfg.NginxCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("nginx_exporter", false)
	if NginxConfig().BaseCfg.Enable {
		globalCfg.NginxCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("nginx_exporter.unix_sock",
			"/dev/shm/nginx_exporter.sock")
		globalCfg.NginxCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("nginx_exporter.metrics_path",
			"/metrics")
		globalCfg.NginxCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("nginx_exporter.metrics_router",
			"/nginx")
		globalCfg.NginxCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("nginx_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("nginx_exporter.filters", ""),
			",")
		if len(filters) > 0 {
			var m map[string]string
			m = make(map[string]string)
			for _, s := range filters {
				if len(s) == 0 {
					continue
				}
				m[s] = s
			}
			globalCfg.NginxCfg.BaseCfg.Filters = m
		}
	}

	// php-fpm exporter
	globalCfg.PhpfpmCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("phpfpm_exporter", false)
	if PhpfpmConfig().BaseCfg.Enable {
		globalCfg.PhpfpmCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("phpfpm_exporter.unix_sock",
			"/dev/shm/php-fpm_exporter.sock")
		globalCfg.PhpfpmCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("phpfpm_exporter.metrics_path",
			"/metrics")
		globalCfg.PhpfpmCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("phpfpm_exporter.metrics_router",
			"/phpfpm")
		globalCfg.PhpfpmCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("phpfpm_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("phpfpm_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.PhpfpmCfg.BaseCfg.Filters = m
	}

	// redis exporter
	globalCfg.RedisCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("redis_exporter", false)
	if RedisConfig().BaseCfg.Enable {
		globalCfg.RedisCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("redis_exporter.unix_sock",
			"/dev/shm/redis_exporter.sock")
		globalCfg.RedisCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("redis_exporter.metrics_path",
			"/metrics")
		globalCfg.RedisCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("redis_exporter.metrics_router",
			"/redis")
		globalCfg.RedisCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("redis_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("redis_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.RedisCfg.BaseCfg.Filters = m
	}

	// memcached exporter
	globalCfg.MemcachedCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("memcached_exporter", false)
	if MemcachedConfig().BaseCfg.Enable {
		globalCfg.MemcachedCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("memcached_exporter.unix_sock",
			"/dev/shm/memcached_exporter.sock")
		globalCfg.MemcachedCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("memcached_exporter.metrics_path",
			"/metrics")
		globalCfg.MemcachedCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("memcached_exporter.metrics_router",
			"/memcached")
		globalCfg.MemcachedCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("memcached_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("memcached_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.MemcachedCfg.BaseCfg.Filters = m
	}
}

func NodeConfig() Node {
	return globalCfg.NodeCfg
}

func NginxConfig() Nginx {
	return globalCfg.NginxCfg
}

func PhpfpmConfig() Phpfpm {
	return globalCfg.PhpfpmCfg
}

func RedisConfig() Redis {
	return globalCfg.RedisCfg
}

func MemcachedConfig() Memcached {
	return globalCfg.MemcachedCfg
}