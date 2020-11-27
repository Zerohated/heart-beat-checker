package main

import (
	"syscall"

	conf "github.com/Zerohated/heart-beat-checker/configs"
	"github.com/Zerohated/heart-beat-checker/internal/controller"
	"github.com/Zerohated/heart-beat-checker/internal/model"
	"github.com/Zerohated/tools/pkg/logger"

	"github.com/fvbock/endless"
	"github.com/robfig/cron/v3"

	"github.com/gin-gonic/gin"
)

var (
	schedule *cron.Cron
	log      = logger.Logger
	config   = conf.Config
)

func init() {
	// Connect DB
	dbConf := config.PostgresConf
	model.Init(dbConf)
}

func main() {
	if config.Stage == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.Default()

	c := controller.NewController()
	// // Echo received message
	// app.POST("/echo", c.EchoHandler)
	app.GET("/users", c.GetUserList)
	app.POST("/user", c.RegisterUser)

	// 定时任务
	if schedule != nil {
		schedule.Stop()
	}
	schedule = cron.New()
	// schedule.AddFunc("0 4 * * *", func() { c.DealTodos() })
	// schedule.AddFunc("0 8 * * *", func() { c.HandleDailyReport() })
	schedule.Start()

	server := endless.NewServer(config.Port, app)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	server.ListenAndServe()
}
