package mydate

import (
	"fmt"
	"time"
)

//MakeDateSlice used for make date string slice
func MakeDateSlice(StartDate, EndDate string) []string {
	d, _ := time.ParseDuration("24h")
	tStartDte := StringToDate(StartDate)
	tEndDate := StringToDate(EndDate)
	datediff := tEndDate.Sub(tStartDte).Hours() / 24
	result := []string{}
	fmt.Println(datediff)
	for datediff >= 0 {
		tStartDte = tStartDte.Add(d)
		datediff = tEndDate.Sub(tStartDte).Hours() / 24
		SingleDate := DateToString(tStartDte)
		result = append(result, SingleDate)
	}
	return result
}

//StringToDate for chanege string to date
func StringToDate(Sdate string) time.Time {
	layout := "2006-01-02"
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tdate, _ := time.ParseInLocation(layout, Sdate, loc)
	return tdate
}

//DateToString for change date to string
func DateToString(tdate time.Time) string {
	sdate := tdate.Format("2006-01-02")
	return sdate
}
