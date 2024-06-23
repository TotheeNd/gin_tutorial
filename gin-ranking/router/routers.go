package router

import (
	// "gin-ranking/config"
	"gin-ranking/controllers"
	"gin-ranking/pkg/logger"

	// "github.com/gin-contrib/sessions"
	// sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(gin.LoggerWithConfig(logger.LoggerToFile())) //加载日志中间件
	r.Use(logger.Recover)                              //加载日志中间件相关
	// store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	// r.Use(sessions.Sessions("mysession", store))

	// user := r.Group("/user")
	// {
	// 	user.POST("/register", controllers.UserController{}.Register)
	// 	user.POST("/login", controllers.UserController{}.Login)
	// }
	player := r.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers) //查询信息的路由入口：http://XXX:8080/player/list
	}
	// vote := r.Group("/vote")
	// {
	// 	vote.POST("/add", controllers.VoteController{}.AddVote)
	// }
	// r.POST("/ranking", controllers.PlayerController{}.GetRanking)

	return r
}
