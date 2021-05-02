package main

import (
	"fmt"
	"log"
	"net/http"
)

var jsonString =`
{
	"response": {
		"header": {
			"status": "ok",
			"startIndex": 0,
			"pageSize": 10,
			"totalRecords": 20
		},
		"schedules": [{
			"id":1
			}]
	}
}
`
func main() {
	fmt.Println("Mock Server started")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        //w.Write([]byte(jsonString))
		 w.Write([]byte("<html>"))
    })
 
    log.Fatal(http.ListenAndServe(":1081", nil))
}