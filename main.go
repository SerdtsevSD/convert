package main

import (
	ctx"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

var router *gin.Engine
var conn *pgx.Conn

func main() {
	conn = connect()
	defer conn.Close(ctx.Background())
	router = gin.Default()
	router.Use(LiberalCORS)
	router.Delims("&{", "}&")
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/static", http.Dir("./static"))
	initializeRoutes()
	router.Run(":9090")
}

func connect() *pgx.Conn {
	conn, err := pgx.Connect(ctx.Background(), "user=postgres password=volter1973 host=localhost port=5432 dbname=convertor sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}

func LiberalCORS(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	if c.Request.Method == "OPTIONS" {
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			c.Header("Access-Control-Allow-Headers", c.Request.Header["Access-Control-Request-Headers"][0])
		}
		c.AbortWithStatus(http.StatusOK)
	}
}
