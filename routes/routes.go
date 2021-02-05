package routes

import (
	"crud_gin/crudproject/controllers"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func SetupRoutes(db *mgo.Session) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.Index)
	r.POST("/insert", controllers.Insert)
	r.GET("/show", controllers.Show)
	r.GET("/new", controllers.New)
	r.GET("/edit", controllers.Edit)
	r.POST("/update", controllers.UpdateProfile)
	r.GET("/delete", controllers.DeleteProfile)

	return r
}
