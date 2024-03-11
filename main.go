package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	r := *gin.Defaul()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"Message": "Pang",
		})
	})
	r.run()

    http.HandleFunc("/", hello)
    http.ListenAndServe(":8080", nil)
	
}
