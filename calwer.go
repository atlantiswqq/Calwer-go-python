package main

import (
<<<<<<< HEAD
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

const URL=`https://www.1396j.com/xyft/?utp=topbar`

type Data struct{
	KaijiangDate string
	nums []int
	caption []string
}


func ParseURL(){
	client:= &http.Client{}
	request,err:= http.NewRequest("GET",URL,nil)
	if err!=nil{
		fmt.Println("init request error.")
	}
	request.Header.Add("Cookie", "UM_distinctid=1695ddfc7b61a4-0ba07d0291e26b-133a6850-1aeaa0-1695ddfc7b785a; _ga=GA1.2.489354662.1552058143; Hm_lvt_dad24abebba647625189f407f7103e48=1552058141,1552059645; countdown_sound=0; ccsalt=e242c8ad8b1f0717dd46f88bb68cea50; CNZZDATA5418000=cnzz_eid%3D1631767431-1552053896-%26ntime%3D1553265064; Hm_lpvt_dad24abebba647625189f407f7103e48=1553267310; _gid=GA1.2.508773704.1553267310; _gat_gtag_UA_108446227_1=1")
	request.Header.Add("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	resp,err:= client.Do(request)
	if err !=nil{
		fmt.Println("request error.")
	}
	if resp.StatusCode!=200{
		fmt.Println("statuscode fail,",resp.StatusCode)
	}
	defer resp.Body.Close()

	doc,err:=goquery.NewDocumentFromReader(resp.Body)
	if err !=nil{
		fmt.Println("parse html error,",err)
	}
	doc.Find("tr[class]").Each(func(i int, selection *goquery.Selection) {
		mydata:=new(Data)
		caption:=make([]string,0)
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			ntext:= strings.Replace(selection.Text()," ","",-1)
			ntext = strings.Replace(ntext,"\n","",-1)
			caption = append(caption,ntext)
		})
		mydata.KaijiangDate= caption[0]
		mydata.nums = StringToIntarr(caption[1])
		mydata.caption = caption[2:]
		fmt.Printf("%+v\n",*mydata)
	})
}

func StringToIntarr(strnum string)[]int{
	result:= make([]int,0)
	strarr:=strings.Split(strnum,"")
	for _,item:=range strarr{
		intval,err:=strconv.Atoi(item)
		if err==nil {
			result = append(result, intval)
		}
	}
	return result
}

func main(){
	ParseURL()
=======
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url=`http://xx.xx.xx.xx:8001/audit/review`
type PrivateBody struct {
	Classifier  string   `json:"classifier"`
	Content     string   `json:"content"`
	CreatedAt   string   `json:"created_at"`
	ImgUrls     []string `json:"img_urls"`
	PatientName string   `json:"patient_name"`
	ProjectID   string   `json:"project_id"`
	Title       string   `json:"title"`
	UserID      string   `json:"user_id"`
}

type PrivateData struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ProjectID  string `json:"project_id"`
		UserID     string `json:"user_id"`
		ImgURL     string `json:"img_url"`
		Result     string `json:"result"`
		Reason     string `json:"reason"`
		TextResult string `json:"text_result"`
		TextReason string `json:"text_reason"`
		ImgResult  string `json:"img_result"`
		ImgReason  string `json:"img_reason"`
		TextData   struct {
			Per []string      `json:"per"`
			Loc []string      `json:"loc"`
			Org []string      `json:"org"`
			Med []interface{} `json:"med"`
			Occ []string      `json:"occ"`
			Dat []string      `json:"dat"`
			Fee []string      `json:"fee"`
		} `json:"text_data"`
		ImgData struct {
			Recognize struct {
				PrintForm    bool `json:"print_form"`
				HandWriting  bool `json:"hand_writing"`
				IDCard       bool `json:"id_card"`
				Face         bool `json:"face"`
				Title        bool `json:"title"`
				ColorfulSeal bool `json:"colorful_seal"`
				BlackSeal    bool `json:"black_seal"`
				WaterMark    bool `json:"water_mark"`
			} `json:"recognize"`
			Infor struct {
				Per []interface{} `json:"per"`
				Loc []interface{} `json:"loc"`
				Org []interface{} `json:"org"`
				Med []interface{} `json:"med"`
				Occ []interface{} `json:"occ"`
				Dat []interface{} `json:"dat"`
				Fee []interface{} `json:"fee"`
			} `json:"infor"`
		} `json:"img_data"`
	} `json:"data"`
}

func main(){
	var reqBody PrivateBody
	URLS :=[]string{"https://15527852422579018783g6dbfb3abff215bf96638d63da666dc93b.jpeg@!large.png"}
	reqBody.Classifier="total"
	reqBody.Content= "各位好心人，"
	reqBody.CreatedAt="1552785242"
	reqBody.PatientName="马超"
	reqBody.ProjectID= "3221xx806"
	reqBody.Title= "你要坚强，一切都会好起来的！"
	reqBody.UserID="111111"
	reqBody.ImgUrls=URLS
	client:= &http.Client{}
	body,err:= json.Marshal(reqBody)
	if err!=nil{
		fmt.Printf("%T",err)
	}
	reader:=bytes.NewReader(body)
	request,err:= http.NewRequest("POST",url,reader)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36")
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	if err!=nil{
		panic(err)
		os.Exit(0)
	}
	resp,err:= client.Do(request)
	if err!=nil{
		fmt.Println(err)
	}
	if resp.StatusCode!=200{
		fmt.Printf("statuscode:%d",resp.StatusCode)
	}
	defer resp.Body.Close()
	respBytes,err:=ioutil.ReadAll(resp.Body)
	mydata:=new(*PrivateData)
	if err:=json.Unmarshal(respBytes,mydata);err!=nil{
		fmt.Printf("parse error : %T",err)
	}
	fmt.Printf("%+v",*mydata)
>>>>>>> aa53638f59100e33b3f1c5b50b51551af331ba1c
}