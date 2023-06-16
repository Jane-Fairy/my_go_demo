package router

import (
	"errors"
	"go.uber.org/zap"
	"my_go/configs"
	"my_go/pkg/file"
)

//资源组
type resource struct {
	//mux          core.Mux
	logger *zap.Logger
	//db           mysql.Repo
	//cache        redis.Repo
	//interceptors interceptor.Interceptor
	//cronServer   cron.Server
}

//服务组
type Server struct {
	//Mux        core.Mux
	//Db         mysql.Repo
	//Cache      redis.Repo
	//CronServer cron.Server
}

func NewHTTPServer(logger *zap.Logger, cronLogger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger

	firstStartBrowserUri := configs.ProjectDomain + configs.ProjectPort //第一次启动访问的url地址

	_, ok := file.IsExists(firstStartBrowserUri)

	if !ok {
		firstStartBrowserUri += "/install"
	} else {
		//初始化数据库

	}
	return nil, nil
}
