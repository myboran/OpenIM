package main

import (
	"flag"
	"fmt"
	"strconv"

	"open-im/internal/api/friend"
	"open-im/internal/api/user"
	"open-im/pkg/common/log"
	"open-im/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(utils.CorsHandler())
	userRouterGroup := r.Group("/user")
	{
		userRouterGroup.POST("/get_user_info", user.GetUsersInfo)
		userRouterGroup.POST("/update_user_info", user.UpdateUserInfo)
		userRouterGroup.POST("/get_self_user_info", user.GetSelfUserInfo)
	}

	friendRouterGroup := r.Group("/friend")
	{
		friendRouterGroup.POST("/add_friend", friend.AddFriend)
		friendRouterGroup.POST("/delete_friend", friend.DeleteFriend)
	}
	// TODO
	//apiThird.MinioInit()
	log.NewPrivateLog("api")
	ginPort := flag.Int("port", 10000, "get ginServerPort from cmd,default 10000 as port")
	flag.Parse()
	fmt.Println("go go go")
	r.Run(":" + strconv.Itoa(*ginPort))
}
