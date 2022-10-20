package datasource

import (
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
	"webProject/config"
)

func NewRedis() *redis.Database {
	var databases *redis.Database

	cmsConfig := config.InitAppConfig()
	if cmsConfig != nil {
		//获取redis配置
		rd := cmsConfig.Redis
		databases = redis.New(redis.Config{
			Network: rd.NetWork,
			//Addr默认是120.0.0.1:6379
			//Addr:      rd.Addr+":"+rd.Port,
			Password:  rd.Password,
			Database:  "",
			Prefix:    rd.Prefix,
			MaxActive: 10,
			Timeout:   redis.DefaultRedisTimeout,
		})
	}

	return databases
}
