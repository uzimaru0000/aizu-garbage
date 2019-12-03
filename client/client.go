package client

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/uzimaru0000/aizu-garbage/model"
	"github.com/uzimaru0000/aizu-garbage/usecase/port"
)

const (
	separator = ","
)

type client struct {
	client *http.Client
}

func NewClient(httpClient *http.Client) port.ScheduleFetcher {
	return &client{client: httpClient}
}

func (s *client) FetchSchedules(placeID string) ([]*model.Schedule, error) {
	req, err := createPostRequest(placeID)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	schedules, err := getSchedules(doc)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func getSchedules(doc *goquery.Document) ([]*model.Schedule, error) {
	list := doc.Find("h2.title3 + ul")
	scheduleInfos := list.Find("li.tri1").Map(makeScheduleInfo)

	schedules := make([]*model.Schedule, len(scheduleInfos))
	for i, info := range scheduleInfos {
		schedule, err := makeSchedule(info)
		if err != nil {
			return nil, err
		}
		schedules[i] = schedule
	}

	return schedules, nil
}

func makeScheduleInfo(i int, s *goquery.Selection) string {
	date := s.Children().Text()
	info := s.Contents().Last().Text()

	return date + separator + info
}

func makeSchedule(info string) (*model.Schedule, error) {
	data := strings.Split(info, separator)
	if len(data) != 2 {
		return nil, errors.New("bad arg")
	}

	date, err := timeFromString(data[0])
	if err != nil {
		return nil, err
	}
	ty := model.ScheduleTypeFromString(data[1])
	return &model.Schedule{Type: ty, Date: *date}, nil
}

func timeFromString(str string) (*time.Time, error) {
	now := time.Now()
	a := strings.Split(str, "/")
	b := strings.Split(a[1], "(")

	month, err := strconv.Atoi(a[0])
	if err != nil {
		return nil, err
	}
	day, err := strconv.Atoi(b[0])
	if err != nil {
		return nil, err
	}

	result := time.Date(now.Year(), time.Month(month), day, 0, 0, 0, 0, now.Location())

	return &result, nil
}
