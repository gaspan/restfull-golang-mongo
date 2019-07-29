package routes

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"goMongodbMaster/models"
)

func Route() {
	e := echo.New()
	e.Use(middleware.Logger())
	api := e.Group("/api")
	api.POST("/student", models.CreateStudent)
	api.GET("/student", models.GetStudents)
	api.GET("/student/:id", models.GetStudent)
	api.POST("/student/:id/update", models.UpdateStudent)
	api.DELETE("/student/:id", models.DeleteStudent)
	s := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e.Logger.Fatal(e.StartServer(s))
}
