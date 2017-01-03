// Package server defines all things related to requests and handling them
package server

import (
	"github.com/cloudcloud/shortoftime/config"
	"github.com/gin-gonic/gin"
)

// Serv is an interface that applies to a runnable Server
type Serv interface {
	Serve() error
	SetConfig(config.Conf) Serv
	SetListener(string) Serv
	SetMode(string) Serv
}

// Server defines an instance by which a Server is configured and runnable
type Server struct {
	Config   config.Conf `json:"config"`
	GinServ  *gin.Engine `json:"-"`
	Listener string      `json:"listener"`
	Mode     string      `json:"mode"`
}

// Init instantiates and sets up the server
func Init(c config.Conf, m string) Serv {
	var s Serv
	s = new(Server)

	s.SetConfig(c).SetMode(m).SetListener(c.GetListener())
	gin.SetMode(m)

	srv := s.(*Server)
	srv.GinServ = gin.New()
	srv.GinServ.Use(
		gin.Recovery(),
		PrettyLogger(gin.DefaultWriter),
		PushConfig(srv.Config),
		// ConnectDatabase(s.Config),
		// CaptureRequest(),
	)

	idx := srv.GinServ.Group("/")
	idx.StaticFile("/", srv.Config.GetRootDir()+"public/index.html")

	//idx.GET("/login", GetLogin())
	//idx.GET("/settings", GetSettings())

	//idx.POST("/login", PostLogin())
	//idx.POST("/logout", PostLogout())
	//idx.POST("/settings", PostSettings())

	res := srv.GinServ.Group("/resource")
	res.StaticFile("/app.js", srv.Config.GetRootDir()+"public/app.js")
	res.StaticFile("/app.css", srv.Config.GetRootDir()+"public/app.css")
	res.StaticFile("/core.css", srv.Config.GetRootDir()+"public/core.css")

	app := srv.GinServ.Group("/app")
	app.StaticFile("/", srv.Config.GetRootDir()+"public/app.html")

	return srv
}

// Serve will start the serving process for a Server
func (s *Server) Serve() error {
	return s.GinServ.Run(s.Listener)
}

// SetConfig will accept a global configuration object
func (s *Server) SetConfig(c config.Conf) Serv {
	s.Config = c

	return s
}

// SetListener will give the server the bind address for Gin
func (s *Server) SetListener(l string) Serv {
	s.Listener = l

	return s
}

// SetMode will accept a Mode type for Gin to run as
func (s *Server) SetMode(m string) Serv {
	s.Mode = m

	return s
}
