package main

import (
	"log"
	"net/http"

	"school1/config"
	"school1/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
        r := gin.Default()
	config.CreateDatabaseAndTable() //生成数据库
	config.InitDB()                 //连接数据库
	defer config.DB.Close()         //关闭数据库

	// 创建Gin引擎
	router := gin.New()

	// 自定义日志中间件，过滤掉OPTIONS请求
	router.Use(func(c *gin.Context) {
		if c.Request.Method != "OPTIONS" {
			gin.Logger()(c)
		}
		c.Next()
	})

	// 启用CORS中间件
       router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	// 设置路由
       api := r.Group("/api")
    {
        api.GET("/getCounts", controllers.NewUserController().GetCounts)
        api.POST("/register", controllers.NewUserController().Register)
        api.GET("/getStudents", controllers.NewUserController().GetStudents)
        api.GET("/getSingleUser", controllers.NewUserController().GetSingleUser)
        api.GET("/getTeachers", controllers.NewUserController().GetTeachers)
        api.POST("/sendMessage", controllers.NewUserController().SendMessage)
        api.GET("/getMessageByMessageId", controllers.NewUserController().GetMessageByMessageId)
        api.POST("/setMessageRead", controllers.NewUserController().SetMessageRead)
        api.POST("/registerVisitor", controllers.NewUserController().RegisterVisitor)
        api.POST("/loginVisitor", controllers.NewUserController().LoginVisitor)
        api.POST("/reserveVisitor", controllers.NewUserController().ReserveVisitor)
        api.POST("/aiQuestionAnswer", controllers.NewUserController().AIQuestionAnswer)
    }

       r.Run(":8080")

	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
