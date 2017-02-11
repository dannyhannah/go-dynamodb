package main

import (
	"fmt"
	"net/http"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/nu7hatch/gouuid"
	"log"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "1", r.URL.Path[1:])
	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-1")})

	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	u, err := uuid.NewV4()
	if err != nil {
		log.Fatal(err)
	}

	svc := dynamodb.New(sess)
	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"ImpressionId": {
				S: aws.String(u.String()),
			},
			"Request": {
				S: aws.String("foo"),
			},
		},
		TableName: aws.String("impressions"),
	}

	result, error := svc.PutItem(params)
	fmt.Println(result)
	if error != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
