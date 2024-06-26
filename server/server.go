package server

import (
	"app.myriadflow.com/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	Routes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:808
}

func Routes(r *gin.Engine) {
	// User routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetUser)
	r.GET("/users/all", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Brand routes
	r.POST("/brands", controllers.CreateBrand)
	r.GET("/brands/:id", controllers.GetBrand)
	r.GET("/brands/all", controllers.GetAllBrands)
	r.PUT("/brands/:id", controllers.UpdateBrand)
	r.DELETE("/brands/:id", controllers.DeleteBrand)

	// Collection routes
	r.POST("/collections", controllers.CreateCollection)
	r.GET("/collections/:id", controllers.GetCollection)
	r.GET("/collections/all", controllers.GetAllCollections)
	r.PUT("/collections/:id", controllers.UpdateCollection)
	r.DELETE("/collections/:id", controllers.DeleteCollection)

	// Phygital routes
	r.POST("/phygitals", controllers.CreatePhygital)
	r.GET("/phygitals/:id", controllers.GetPhygital)
	r.GET("/phygitals/all", controllers.GetAllPhygital)
	r.PUT("/phygitals/:id", controllers.UpdatePhygital)
	r.DELETE("/phygitals/:id", controllers.DeletePhygital)

	// WebXR routes
	r.POST("/webxr", controllers.CreateWebXR)
	r.GET("/webxr/:id", controllers.GetWebXR)
	r.GET("/webxr/all", controllers.GetAllWebXR)
	r.PUT("/webxr/:id", controllers.UpdateWebXR)
	r.DELETE("/webxr/:id", controllers.DeleteWebXR)

	// Avatar routes
	r.POST("/avatars", controllers.CreateAvatar)
	r.GET("/avatars/:id", controllers.GetAvatar)
	r.GET("/avatars/all", controllers.GetAllAvatars)
	r.PUT("/avatars/:id", controllers.UpdateAvatar)
	r.DELETE("/avatars/:id", controllers.DeleteAvatar)

	// Variant routes
	r.POST("/variants", controllers.CreateVariant)
	r.GET("/variants/:id", controllers.GetVariant)
	r.GET("/variants/all", controllers.GetAllVariant)
	r.PUT("/variants/:id", controllers.UpdateVariant)
	r.DELETE("/variants/:id", controllers.DeleteVariant)
}
