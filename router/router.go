package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	baseGroup := r.Group("/v1")
	initActRouter(baseGroup)
	initAuthorRouter(baseGroup)
	initCommentRouter(baseGroup)
	initHotListRouter(baseGroup)
	initItemRouter(baseGroup)
	initRelationRouter(baseGroup)
	initSearchRouter(baseGroup)
	initUserRouter(baseGroup)
}

func initHotListRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/hot_list")
}

func initRelationRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/relation")
}

func initItemRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/item")
}

func initSearchRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/search")
}

func initUserRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/user")
}

func initCommentRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/comment")
}

func initAuthorRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/author")
}

func initActRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("/act")
}


