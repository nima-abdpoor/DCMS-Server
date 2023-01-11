package influx

import (
	"context"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

var i float32 = 1.0

func StartInfluxDB() {
	url := "http://localhost:8086"
	token := "JfQB7rL2FJ71o6KFdECe9SScc4IXuGk448K1tQAB58iejZpa-NRoqszIFL-_CGi3oVLv1Cs8ilCWg7Lu3UScbw=="
	org := "DCMS"
	bucket := "Bucket"
	client := influxdb2.NewClient(url, token)
	writeAPI := client.WriteAPIBlocking(org, bucket)
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"salam": i},
		time.Now())
	i++
	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("query parsing error: %s\n", err)
	}
	client.Close()
}
