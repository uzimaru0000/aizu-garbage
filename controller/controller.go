package controller

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/uzimaru0000/aizu-garbage/model"
)

const endPoint = "https://www.city.aizuwakamatsu.fukushima.jp/index_php/gomical/index_i.php?typ=p"

func GetPlaceList() []*model.Place {
	places := []*model.Place{}
	// 	var places []*model.Placeこっちのが好き

	// このIDもここにべたがきより外部で定義しておきたい
	// 早期リターンはGood、あとはエラーとnilを返せるともっといいと思う
	req, err := createRequest("000000")
	if err != nil {
		log.Printf(err.Error())
		return places
	}

	// client作成も別でやりたい
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// このclient作成も外部でやって注入してあげれば、拡張性が上がる
	client := &http.Client{Transport: tr}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Printf(err.Error())
		return places
	}

	//NewDocumentFromResponseはdeprecated
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Printf(err.Error())
		return places
	}

	// ここでdocを直接触らない方がいい
	// GetPlaceListにはDBは関係ないし、データにアクセスできてしまうので、危険っていうのと、
	// ぱっと見なんの処理をしているのかがわからないので、分離して、適切な関数名をつけてほしい
	infos := doc.Find("form ul li.tri1 select").First().Children().Map(func(i int, s *goquery.Selection) string {
		// ここの処理も関数で分けたい
		return s.AttrOr("value", "000000") + " " + s.Text()
	})

	// ここも別関数
	for _, info := range infos {
		arr := strings.Split(info, " ")
		places = append(places, &model.Place{PlaceID: arr[0], Name: arr[1]})
	}

	return places
}

func GetInfo(place *model.Place) (*model.Schedule, error) {
	req, err := createRequest(place.PlaceID)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}

	//---------------------------------------------
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	// -----------------------------------------------共通化すべき

	// ここの処理も一つ一つ分離した方がいいのと、ここの処理はmodelにあるべき処理
	schedule := &model.Schedule{ID: place.PlaceID, Place: place}
	list := doc.Find("h2.title3 + ul") // セレクタもmodel側で定義してそれを呼び出して結果を返すメソッドを生やす
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

	return schedule, nil
}

func createRequest(id string) (*http.Request, error) {
	//ここのValuesの作成を別関数で切り出したい
	// このリクエストの作成に依存し切ってしまっているため
	values := &url.Values{}
	values.Add("m", id)
	values.Add("d", "1")

	req, err := http.NewRequest("POST", endPoint, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	// ここも外でやりたいお気持ちもある

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.122 Safari/537.36 Vivaldi/2.3.1440.61")

	return req, nil
}
