package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const endPoint = "https://www.city.aizuwakamatsu.fukushima.jp/index_php/gomical/index_i.php?typ=p"

type Schedule struct {
	Place      string
	PlaceID    string
	Categories []*Category
}

type Category struct {
	Genre string
	Date  time.Time
}

func main() {
	ch := make(chan *Schedule)
	ids := GetIDList()

	for _, id := range ids {
		go GetInfo(ch, id)
	}

	count := len(ids)

	for count != 0 {
		she, ok := <-ch
		if ok {
			printSchedule(she)
		} else {
			fmt.Printf("err\n")
		}
		count--
	}
}

func GetIDList() []string {
	doc, err := goquery.NewDocument(endPoint)

	if err != nil {
		fmt.Errorf("%s", err.Error())
		return []string{}
	}

	ids := doc.Find("form ul li.tri1 select").First().Children().Map(func(i int, s *goquery.Selection) string {
		return s.AttrOr("value", "000000")
	})

	return ids
}

func GetInfo(ch chan *Schedule, id string) {
	values := &url.Values{}
	values.Add("m", id)
	values.Add("d", "1")

	req, err := http.NewRequest("POST", endPoint, strings.NewReader(values.Encode()))
	if err != nil {
		close(ch)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		close(ch)
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		close(ch)
		return
	}

	schedule := &Schedule{}
	place := doc.Find("h2.title3").Last().Text()
	list := doc.Find("h2.title3 + ul")
	result := make(map[string]string)
	list.Find("li.tri1").Each(func(i int, sel *goquery.Selection) {
		date := sel.Children().Text()
		info := sel.Contents().Last().Text()

		result[date] = info
	})

	schedule.PlaceID = id
	schedule.Place = place
	for key, val := range result {
		cate := &Category{Genre: val, Date: parser(key)}
		schedule.Categories = append(schedule.Categories, cate)
	}

	ch <- schedule
}

func parser(date string) time.Time {
	now := time.Now()
	a := strings.Split(date, "/")
	b := strings.Split(a[1], "(")

	month, _ := strconv.Atoi(a[0])
	day, _ := strconv.Atoi(b[0])

	result := time.Date(now.Year(), time.Month(month), day, 0, 0, 0, 0, now.Location())

	return result
}

func printSchedule(schedule *Schedule) {
	fmt.Printf("%s %s\n", schedule.PlaceID, schedule.Place)
	for _, c := range schedule.Categories {
		fmt.Printf("%s %s\n", c.Date.String(), c.Genre)
	}
}
