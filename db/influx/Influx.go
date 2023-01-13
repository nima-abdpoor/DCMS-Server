package influx

import (
	"DCMS/parser"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var i float32 = 1.0

func StartInfluxDB(log parser.ParsedLog, name string) {
	makeAPostRequest(createLogAsInfluxDBLineProtocolFormat(log, name))
}

func makeAPostRequest(influxPoints []string) {
	client := &http.Client{}
	for _, point := range influxPoints {
		var data = strings.NewReader(point)
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
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createLogAsInfluxDBLineProtocolFormat(log parser.ParsedLog, fileInfo string) (influxData []string) {
	logType, uniqueId := getUserInfoFromFileInfo(fileInfo)
	unixTime := time.Date(2023, 01, 10, 14, 28, 00, 1111111, time.Local).Unix()
	if log.HasRequest {
		influxData = append(influxData, uniqueId+",LogType="+logType+",httpLog=Request url=\""+log.Request.URL+"\",Body=\""+log.Request.Body+"\",Headers=\""+log.Request.Header+"\" "+strconv.FormatInt(unixTime, 10))
	}
	if log.HasResponse {
		influxData = append(influxData, uniqueId+",LogType="+logType+",httpLog=Response url=\""+log.Response.URL+"\",Body=\""+log.Response.Body+"\",Headers=\""+log.Response.Header+"\",Code="+log.Response.Code+" "+strconv.FormatInt(unixTime, 10))
	}
	return
}

func getUserInfoFromFileInfo(fileInfo string) (logType string, uniqueCode string) {
	parts := strings.Split(fileInfo, "/")
	uniqueCode = parts[len(parts)-2]
	logType = parts[len(parts)-3]
	return
}
