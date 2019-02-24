package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"

	"../type"
)

func Scraping(key int, word string) type.WordRow {
	// バイト型を定義してから文字列を結合してurlを作る
	// Goでは文字列連結はコストの高い操作
	// https://qiita.com/ruiu/items/2bb83b29baeae2433a79
	b := make([]byte, 0, 10)
	b = append(b, "https://eow.alc.co.jp/search?q="...)
	b = append(b, word...)
	b = append(b, "&ref=sa"...)

	url := string(b)

	//アクセスしてhtml取得
	html, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("アクセスエラー")
	}

	//発音記号
	pron := html.Find("span.pron").Text()
	//発音記号から末尾「、」を削除
	rep := regexp.MustCompile(`、`)
	pron = rep.ReplaceAllString(pron, "")

	//意味
	meaningByte := make([]byte, 0)
	meaningCount := 0
	meaningDom := html.Find("div#resultsList > ul > li").First()
	meaningDom = meaningDom.Find("div > ol > li")

	meaningDom.Each(func(_ int, s *goquery.Selection) {
		meaningText := s.Text()

		if meaningCount < 4 {
			if !strings.Contains(meaningText, "俗") &&
				!strings.Contains(meaningText, "卑") &&
				!strings.Contains(meaningText, "蔑") &&
				!strings.Contains(meaningText, "＝") {
				meaningText = Replace(meaningText)
				meaningByte = append(meaningByte, meaningText...)
				meaningCount++
			}
		}
	})
	meaning := string(meaningByte)

	return WordRow{Key: key, Word: word, Pron: pron, Meaning: meaning}
}

//正規表現で置換
func Replace(meaningText string) string {
	//【語源】や【学名】削除
	rep := regexp.MustCompile(`(◆.+)+`)
	meaningText = rep.ReplaceAllString(meaningText, "")

	//漢字の読みがな削除
	rep = regexp.MustCompile(`｛(\p{Hiragana}|\s|（|）|／)+｝`)
	meaningText = rep.ReplaceAllString(meaningText, "")

	//例文削除
	rep = regexp.MustCompile(`・[a-zA-Z\s]+.+`)
	meaningText = rep.ReplaceAllString(meaningText, "")

	//《things》削除
	rep = regexp.MustCompile(`《.+》`)
	meaningText = rep.ReplaceAllString(meaningText, "")

	//〔〕内のアルファベット削除
	// rep = regexp.MustCompile(`[a-zA-Z]|'`)
	// meaningText = rep.ReplaceAllString(meaningText, "-")

	return meaningText
}