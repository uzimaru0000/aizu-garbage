package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/uzimaru0000/aizu-garbage/utils/model"
)

const endPoint = "https://www.city.aizuwakamatsu.fukushima.jp/index_php/gomical/index_i.php?typ=p"

func GetPlaceList() []*model.Place {
	places := []*model.Place{}
	doc, err := goquery.NewDocument(endPoint)

	if err != nil {
		return places
	}

	infos := doc.Find("form ul li.tri1 select").First().Children().Map(func(i int, s *goquery.Selection) string {
		return s.AttrOr("value", "000000") + " " + s.Text()
	})

	for _, info := range infos {
		arr := strings.Split(info, " ")
		places = append(places, &model.Place{PlaceID: arr[0], Name: arr[1]})
	}

	return places
}

func GetInfo(places []*model.Place) []*model.Schedule {
	result := make(chan []*model.Schedule)
	value := make(chan *model.Schedule)
	finish := make(chan bool)

	var wg sync.WaitGroup

	go func() {
		r := make([]*model.Schedule, 0)

		for {
			select {
			case v := <-value:
				if v != nil {
					r = append(r, v)
				}
			case <-finish:
				result <- r
				return
			}
		}
	}()

	for _, p := range places {
		wg.Add(1)
		go func(place *model.Place) {
			defer wg.Done()
			scrapeInfo(value, place)
		}(p)
	}
	wg.Wait()
	finish <- true

	r := <-result
	return r
}

func scrapeInfo(ch chan *model.Schedule, place *model.Place) {
	req, err := createRequest(place.PlaceID)
	if err != nil {
		fmt.Printf("%v\n", err)
		ch <- nil
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v\n", err)
		ch <- nil
		return
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Printf("%v\n", err)
		ch <- nil
		return
	}

	schedule := &model.Schedule{ID: place.PlaceID, Place: place}
	list := doc.Find("h2.title3 + ul")
	result := make(map[string]string)
	list.Find("li.tri1").Each(func(i int, sel *goquery.Selection) {
		date := sel.Children().Text()
		info := sel.Contents().Last().Text()
		result[date] = info
	})

	for key, val := range result {
		cate := &model.Category{PlaceID: place.PlaceID, Info: val, Date: model.Parser(key)}
		schedule.Categories = append(schedule.Categories, cate)
	}

	ch <- schedule
}

func createRequest(id string) (*http.Request, error) {
	values := &url.Values{}
	values.Add("m", id)
	values.Add("d", "1")

	req, err := http.NewRequest("POST", endPoint, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}
