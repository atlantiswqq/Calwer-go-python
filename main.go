package main

import (
	"calwer/mydate"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

func main() {
	sdate, edate := "2019-01-01", time.Now().Format("2006-01-02")
	DateSlice := mydate.MakeDateSlice(sdate, edate)
	for _, SingleDate := range DateSlice {
		SearcheDate := "https://www.1396j.com/xyft/kaijiang?date=%c&_=%d"
		TimeStamp := time.Now().Unix()
		ComplexURL := fmt.Sprintf(SearcheDate, SingleDate, TimeStamp)
		err := ParseURL(ComplexURL)
		if err != nil {
			fmt.Println(err)
		}
	}
}

//ParseURL for parse url
func ParseURL(url string) error {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	request.Header.Add("Referer", "https://www.1396j.com/xyft/kaijiang")
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return errors.New("StatusError")
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return errors.New("HTMLParseError")
	}
	fmt.Println(doc)
	return nil
}

func redirect(OriginalURL, NewURL string) string {
	return ""
}
