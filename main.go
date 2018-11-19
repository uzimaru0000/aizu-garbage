package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const endPoint = "https://www.city.aizuwakamatsu.fukushima.jp/index_php/gomical/index_i.php?typ=p"

func main() {
	place, info := GetInfo("000800")

	fmt.Printf("%s\n", place)
	for k, v := range info {
		fmt.Printf("%s -> %s\n", k, v)
	}
}

func GetInfo(id string) (string, map[string]string) {
	values := &url.Values{}
	values.Add("m", id)
	values.Add("d", "1")

	req, err := http.NewRequest("POST", endPoint, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Printf("%v\n", err)
		return "", nil
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)

	doc, err := goquery.NewDocumentFromResponse(resp)

	if err != nil {
		fmt.Printf("%v\n", err)
		return "", nil
	}

	place := doc.Find("h2.title3").Last().Text()

	list := doc.Find("h2.title3 + ul")

	result := make(map[string]string)

	list.Find("li.tri1").Each(func(i int, sel *goquery.Selection) {
		date := sel.Children().Text()
		info := sel.Contents().Last().Text()

		result[date] = info
	})

	return place, result

}
