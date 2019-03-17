package Controllers

import (
	"github.com/gin-gonic/gin"

	"../service"
	"../type"
)

func SearchByWords(c *gin.Context) {
	//リクエストで受け取る配列
	var request Type.Request
	c.BindJSON(&request)

	//チャネルの定義
	finished := make(chan bool)
	results := []Type.WordRow{}

	//配列の中身を並行処理
	for key, word := range request.Words {
		results = append(results, Type.WordRow{key, word, "", ""})

		go func(key int, word string) {
			wordrow := Service.Scraping(key, word)
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
