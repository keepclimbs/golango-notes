/*
	1 调试gin首先配置 环境变量
		-  go env -w GO111MODULE=on
		-  go env -w GOPROXY=https://goproxy.cn,direct
	2 切换到当前目录下 配置 go mod
		- go mod init gindemo(当前目录名)
		- go get -u github.com/gin-gonic/gin
	3 执行获取包的操作
		- go get -u github.com/gin-gonic/gin
	4 启动main.go 浏览器访问 localhost:8080/ping
		- go run .\main.go

	可以跟着 https://gin-gonic.com/zh-cn/docs/examples/pure-json/ 示例跑一变
*/

package main

import (
	"fmt"
	"gin/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default 使用 Logger 和 Recovery 中间件
	router := gin.Default()

	routers.InitRouters(router)
	gin.ForceConsoleColor()

	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()

	router.GET("/ping", func(c *gin.Context) {
		log.Printf("test log")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/*
		POST /post?id=1234&page=1 HTTP/1.1
		Content-Type: application/x-www-form-urlencoded

		name=manu&message=this_is_great

		可以看image.png
	*/
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	/*
		使用 postman工具实现
		http://localhost:8080/login

		{
			"user":"songbiaowei",
			"password":"liuxinru"
		}
	*/
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var form LoginForm
		// 在这种情况下，将自动选择合适的绑定
		if c.ShouldBind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	/*
		单文件上传
		为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	*/
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dst := "./" + file.Filename
		// 上传文件至指定的完整文件路径
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	/*
		多文件上传
		为 multipart forms 设置较低的内存限制 (默认是 32 MiB)

		curl -X POST http://localhost:8080/upload \
			-F "files=@C:\Users\songbw\Desktop\logo.png" \
			-F "files=@C:\Users\songbw\Desktop\logo1.png" \
			-H "Content-Type: multipart/form-data"
	*/
	router.POST("/mulUpload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)
			dst := "./" + file.Filename
			// 上传文件至指定目录
			c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	/*
		从 reader 读取数据
	*/
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.Run(":8080")
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
