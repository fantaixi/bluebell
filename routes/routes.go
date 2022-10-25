package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"bluebell/settings"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Setup(mode string) *gin.Engine {
	// gin 设置成发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// middlewares.RateLimitMiddleware(2*time.Second,1)  每两秒钟放入一个令牌，限流
	r.Use(logger.GinLogger(), logger.GinRecovery(true),middlewares.RateLimitMiddleware(2*time.Second,1))


	v1 := r.Group("api/v1")
	//注册业务路由
	v1.POST("/signup",controller.SignUpHandler)

	//登录
	v1.POST("/login",controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware(),)  //应用jwt

	{
		v1.GET("/community",controller.CommunityHandler)
		v1.GET("/community/:id",controller.CommunityDetailHandler)

		v1.POST("/post",controller.CreatePostHandler)
		v1.GET("/post/:id",controller.PostDetaliHandler)
		v1.GET("/posts",controller.GetPostListHandler)
		//根据时间或分数获取帖子列表
		v1.POST("/post2",controller.GetPostListHandler2)

		//投票
		v1.POST("/vote",controller.PostVoteController)
	}

	//pprof相关
	pprof.Register(r)
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		//如果是登录的用户，判断请求头中是否有token
		c.String(http.StatusOK,"pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg": "404",
		})
	})
	return r
}

