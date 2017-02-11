package main

import (
	"net/http"
	"fmt"
	"bytes"
	"github.com/nu7hatch/gouuid"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	u, err1 := uuid.NewV4()

	if err1 != nil {
		log.Fatal(err1)
	}

	url := "https://search-testgo-lxx2ybjkefof4w57ou5wg62csa.us-east-1.es.amazonaws.com/hello/foo/"+u.String()
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"user" : "kimchy", "post_date" : "2017-02-11T14:12:12","message" : "trying out Elasticsearch"}`)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
}

func main() {
	http.HandleFunc("/foo", handler)
	http.ListenAndServe(":8081", nil)
}

