package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

func fetchByteImage(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}
	
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}
	
	return bytes
}

func Rekognition_DetectLabels(url string) {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-west-2")}))

	svc := rekognition.New(sess)

	params := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: fetchByteImage(url),
		},
		//MaxLabels:     aws.Int64(1),
		//MinConfidence: aws.Float64(1.0),
	}
	resp, err := svc.DetectLabels(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Rekognition_DetectModerationLabels() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.DetectModerationLabelsInput{
		Image: &rekognition.Image{ // Required
			Bytes: []byte("PAYLOAD"),
			S3Object: &rekognition.S3Object{
				Bucket:  aws.String("S3Bucket"),
				Name:    aws.String("S3ObjectName"),
				Version: aws.String("S3ObjectVersion"),
			},
		},
		MinConfidence: aws.Float64(1.0),
	}
	resp, err := svc.DetectModerationLabels(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func main() {
	urls := []string{"http://cdn-thumbshot-ie.pearltrees.com/fa/a8/faa897cb9748e852545eae3b40d3c6d6-b52square.jpg", 
		"https://static1.squarespace.com/static/574741382fe1310155f5d5b2/t/5767f8ac9c03e023b60ff3b1/1466433218590/image009.jpeg"}
	
	Rekognition_DetectLabels(urls[0])
}
