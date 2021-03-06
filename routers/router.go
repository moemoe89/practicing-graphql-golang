//
//  Practicing GraphQL
//
//  Copyright © 2020. All rights reserved.
//

package routers

import (
	ap "github.com/moemoe89/go-graphql-gendhis/api"
	mw "github.com/moemoe89/go-graphql-gendhis/api/middleware"
	"github.com/moemoe89/go-graphql-gendhis/api/v1/user"
	usrGraphQL "github.com/moemoe89/go-graphql-gendhis/api/v1/user/delivery/graphql"
	usrHttp "github.com/moemoe89/go-graphql-gendhis/api/v1/user/delivery/http"

	"github.com/gin-gonic/gin"
	"github.com/moemoe89/go-localization"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// GetRouter will create a variable that represent the gin.Engine
func GetRouter(lang *language.Config, log *logrus.Entry, userSvc user.Service) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(mw.CORS)
	r.GET("/", ap.Ping)
	r.GET("/ping", ap.Ping)

	api := r.Group("/api")
	apiV1 := api.Group("/v1")

	usr := usrHttp.NewUserCtrl(lang, log, userSvc)

	apiV1.POST("/user", usr.Create)
	apiV1.GET("/user", usr.List)
	apiV1.GET("/user/:id", usr.Detail)
	apiV1.PUT("/user/:id", usr.Update)
	apiV1.DELETE("/user/:id", usr.Delete)

	apiV1.GET("/graphql/user", usrGraphQL.Handler(userSvc))
	apiV1.POST("/graphql/user", usrGraphQL.Handler(userSvc))

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
