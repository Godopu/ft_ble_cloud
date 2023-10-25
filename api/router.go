package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewWebServer(addr string) *http.Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: createRouter(),
	}

	return srv
}

func createRouter() *gin.Engine {
	apiEngine := gin.New()
	apiGroup := apiEngine.Group("/api")
	{
		apiGroup.POST("/emg", func(c *gin.Context) {
			body := map[string]interface{}{}
			dec := json.NewDecoder(c.Request.Body)
			err := dec.Decode(&body)
			if err != nil {
				panic(err)
			}

			log.Println("storing and expecting EMG")

			c.JSON(http.StatusOK, map[string]interface{}{"computing_delay": 200})
		})

		apiGroup.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello world")
		})
	}

	r := gin.Default()
	// gin.
	// r can accept all messages from apiEngine and staticEngine
	r.Any("/*any", func(c *gin.Context) {
		defer handleError(c)
		w := c.Writer
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		path := c.Param("any")

		if strings.HasPrefix(path, "/api") {
			apiEngine.ServeHTTP(c.Writer, c.Request)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	return r
}
