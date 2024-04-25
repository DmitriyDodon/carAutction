package server

import (
	"app/server/controller"
	"github.com/swaggo/echo-swagger"

	_ "app/docs"
)

func (s *server) initRoutes(c *controller.Controller) {
	s.e.GET("/swagger*", echoSwagger.WrapHandler)
	s.e.POST("/car", c.CreateCar)
	s.e.PUT("/car/:carID", c.UpdateCar)
	s.e.DELETE("/car/:carID", c.DeleteCar)
	s.e.GET("/car/:carID", c.GetCar)
	s.e.GET("/car", c.ListCars)
}
