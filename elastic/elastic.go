package main

import (
	"net/http"
	"bytes"
	"github.com/nu7hatch/gouuid"
	"log"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//defer un(trace("Request"))

	u, err1 := uuid.NewV4()

	if err1 != nil {
		log.Fatal(err1)
	}

	url := "https://search-testgo-lxx2ybjkefof4w57ou5wg62csa.us-east-1.es.amazonaws.com/hello/foo/"+u.String()

	var jsonStr = []byte(`{"type" : "log", "post_date" : "`+time.Now().Format("2006-01-02")+`",  "message": "trying out Elasticsearch"}`)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func main() {
	http.HandleFunc("/foo", handler)
	http.ListenAndServe(":8081", nil)
}

