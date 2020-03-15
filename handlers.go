package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func pageConv(c *gin.Context) {
	var Curencys []string= []string{
		"USD",
		"EUR",
		"GBP",
		"AUD",
		"AZN",
		"AMD",
		"BGN",
		"HUF",
		"RUR",
	}
	// Call the HTML method of the Context to render a template
	c.HTML(
		http.StatusOK,
		"convertor.html",
		gin.H{
			"title": "Convertor",
			"Curencys": Curencys,
		},
	)
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

func getRBK(c *gin.Context) {
	s := struct {
		Amount        string `json:"amount"`
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
func mainPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"mainPage.html",
		gin.H{},
	)
}
