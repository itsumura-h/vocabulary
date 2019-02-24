package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// JSONを返すサンプル
func Sample1(c *gin.Context) {
	response := gin.H{
		"str_value": "aaa",
		"int_value": 10,
		"arr_value": []int64{1, 2, 3},
	}

	c.JSON(200, response)
}

// URLパラメータを取得してJSONで返すサンプル
func Sample2(c *gin.Context) {
	name := c.Param("name")
	response := gin.H{
		"name": name,
	}

	c.JSON(200, response)
}

func Sample3(c *gin.Context) {
	type Request struct {
		Words []string `json:"words"`
	}

	var request Request
	c.BindJSON(&request)

	//fmt.Println(request.Words)

	for _, word := range request.Words {
		fmt.Println(word)
	}

	c.JSON(200, request)
}
