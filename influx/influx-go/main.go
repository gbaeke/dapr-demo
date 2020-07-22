package main

import (
	"context"
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go"
)

func main() {
	client := influxdb2.NewClient("http://localhost:9999", "token")
	writeAPI := client.WriteAPIBlocking("Inity", "InityBucket")

	jsond := map[string]interface{}{
		"tags":   "room=dorm,building=buildingA",
		"values": "avg=10.0,max=12.0",
	}

	line := fmt.Sprintf("stat,%s %s", jsond["tags"], jsond["values"])

	//p := influxdb2.NewPoint("stat",
	//	map[string]string{"unit": "temperature"},
	//	map[string]interface{}{"avg": 100.5, "max": 45},
	//	time.Now())
	err := writeAPI.WriteRecord(context.Background(), line)
	if err != nil {
		panic(err)
	}

	// query
	queryAPI := client.QueryAPI("Inity")
	result, err := queryAPI.Query(context.Background(), `from(bucket:"InityBucket")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		// Iterate over query response
		for result.Next() {
			// Notice when group key has changed
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// Access data
			fmt.Printf("value: %v\n", result.Record().Value())
		}
		// check for an error
		if result.Err() != nil {
			fmt.Printf("query parsing error: %s\n", result.Err().Error())
		}
	} else {
		panic(err)
	}

	client.Close()
}
