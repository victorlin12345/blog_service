package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victorlin12345/blog_service/routers"
)

func main(){
	r := gin.Default()
	routers.InitRoute(r)
	r.Run(":3000")
}

