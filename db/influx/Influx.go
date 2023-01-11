package influx

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var i float32 = 1.0

func StartInfluxDB() {
	makeAPostRequest()
}

func makeAPostRequest() {
	unixTime := time.Date(2023, 01, 01, 14, 28, 00, 1111111, time.Local).Unix()
	fmt.Println(unixTime)
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf("home,room=Living temp=21.1,hum=35.9,co=0i %d\n"+
		"home,room=Kitchen temp=14.1,hum=36.6,co=22i 1641063600\n"+
		"home,room=Living temp=18.2,hum=36.4,co=17i 1641067200\n"+
		"home,room=Kitchen temp=220.7,hum=36.5,co=26i 1641067200", unixTime))
	req, err := http.NewRequest("POST", "http://localhost:8086/api/v2/write?org=DCMS&bucket=Bucket&precision=s", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Token 0F5vE9GsVRDfRgP7fKr1V9XaHaR6mKJ15uAqUOsGv6htEc7ShJlQSp2XE0mXFMEP7sdAteMImSLWzF_53ky_aQ==")
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}
