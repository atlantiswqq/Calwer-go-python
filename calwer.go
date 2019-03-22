package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const URL = `https://www.1396j.com/xyft/?utp=topbar`

type Data struct {
	KaijiangDate string
	nums         []int
	caption      []string
}

func ParseURL() {
	client := &http.Client{}
	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("init request error.")
	}
	request.Header.Add("Cookie", "UM_distinctid=1695ddfc7b61a4-0ba07d0291e26b-133a6850-1aeaa0-1695ddfc7b785a; _ga=GA1.2.489354662.1552058143; Hm_lvt_dad24abebba647625189f407f7103e48=1552058141,1552059645; countdown_sound=0; ccsalt=e242c8ad8b1f0717dd46f88bb68cea50; CNZZDATA5418000=cnzz_eid%3D1631767431-1552053896-%26ntime%3D1553265064; Hm_lpvt_dad24abebba647625189f407f7103e48=1553267310; _gid=GA1.2.508773704.1553267310; _gat_gtag_UA_108446227_1=1")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("request error.")
	}
	if resp.StatusCode != 200 {
		fmt.Println("statuscode fail,", resp.StatusCode)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("parse html error,", err)
	}
	doc.Find("tr[class]").Each(func(i int, selection *goquery.Selection) {
		mydata := new(Data)
		caption := make([]string, 0)
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			ntext := strings.Replace(selection.Text(), " ", "", -1)
			ntext = strings.Replace(ntext, "\n", "", -1)
			caption = append(caption, ntext)
		})
		mydata.KaijiangDate = caption[0]
		mydata.nums = StringToIntarr(caption[1])
		mydata.caption = caption[2:]
		fmt.Printf("%+v\n", *mydata)
	})
}

func StringToIntarr(strnum string) []int {
	result := make([]int, 0)
	strarr := strings.Split(strnum, "")
	for _, item := range strarr {
		intval, err := strconv.Atoi(item)
		if err == nil {
			result = append(result, intval)
		}
	}
	return result
}

func main() {
	ParseURL()
}
