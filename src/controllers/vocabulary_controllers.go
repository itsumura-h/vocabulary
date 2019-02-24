package controllers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"

	"../service"
)

type WordRow struct {
	Key     int    `json:"key"`
	Word    string `json:"word"`
	Pron    string `json:"pron"`
	Meaning string `json:"meaning"`
}

func SearchByWords(c *gin.Context) {
	//リクエストで受け取る配列
	type Request struct {
		Words []string `json:"words"`
	}

	var request Request
	c.BindJSON(&request)

	//チャネルの定義
	finished := make(chan bool)
	results := []WordRow{}

	//配列の中身を並行処理
	for key, word := range request.Words {
		results = append(results, WordRow{key, word, "", ""})

		go func(key int, word string) {
			wordrow := service.Scraping(key, word)
			results[key] = wordrow
			finished <- true
		}(key, word)
	}

	//終わるまで待つ
	for i := 0; i < len(request.Words); i++ {
		<-finished
	}

	c.JSON(200, results)
}

//_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/_/

function downloadCSV(c *gin.Context) {
	
}