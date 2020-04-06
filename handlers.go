package main

import (
	ctx "context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func pageConv(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"convertor.html",
		gin.H{
			"title": "Convertor",
		},
	)
}

func Login(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login",
		},
	)
}

func getRBK(c *gin.Context) {
	s := struct {
		Amount       string `json:"amount"`
		СurrencyFrom string `json:"selected_user_curency"`
		СurrencyTo   string `json:"selected_out_curency"`
	}{}
	c.BindJSON(&s)
	URL := "https://cash.rbc.ru/cash/json/converter_currency_rate/?currency_from=" + s.СurrencyFrom + "&currency_to=" + s.СurrencyTo + "&source=cbrf&sum=" + s.Amount + "&date="
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	var result map[string]map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	data := result["data"]
	DataNew := fmt.Sprintf("%v", data["sum_result"])
	c.JSON(200, gin.H{
		"result": DataNew,
	})
}

func auth(c *gin.Context) {
	var user struct {
		Id       int    `json:"id"`
		Username string `json:"username",sql:"username"`
	}
	c.BindJSON(&user)
	err := conn.QueryRow(ctx.Background(), "select id from Users where username = $1", user.Username).Scan(&user.Id)
	if err != nil {
		fmt.Printf("%s", err)
	}
	if user.Id != 0 {
		c.JSON(200, gin.H{
			"user": user,
		})
	} else {
		_, err = conn.Exec(ctx.Background(), "insert into Users(username) VALUES($1)", user.Username)
		if err != nil {
			fmt.Printf("%s", err)
		}
		err = conn.QueryRow(ctx.Background(), "select id from Users where username = $1", user.Username).Scan(&user.Id)
		c.JSON(200, gin.H{
			"user": user,
		})
	}
	c.SetCookie("User", "user", 3600, "/", "localhost", http.SameSiteLaxMode, false, false)
	cookie, err := c.Cookie("User")

	if err != nil {
		cookie = "NotSet"
		c.SetCookie("User", "testhello", 3600, "/", "localhost", http.SameSiteLaxMode, false, false)
	}

	fmt.Printf("Cookie value: %s \n", cookie)

}
