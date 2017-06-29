package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"cloud.google.com/go/vision"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	cl "github.com/mpmlj/clarifai-client-go"
	"golang.org/x/net/context"
)

func rekognitionDetectLabels(url string) []*rekognition.Label {
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
		return nil
	}

	return resp.Labels
}

func visionDetectLabels(url string) []*vision.EntityAnnotation {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewClient(ctx)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	image := vision.NewImageFromURI(url)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	labels, err := client.DetectLabels(ctx, image, 100)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	return labels
}

func clarifaiDetectLabels(url string) []*cl.Output {
	sess, err := cl.Connect("uX94X_AnN7W68a9Ms7G55UoGE3KKH5073Kn2PSht", "e9NtNTLzoNBdXvQnWoPu32Xl-9yRhUgR7XU_6W9A")
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	data := cl.InitInputs()
	_ = data.AddInput(cl.NewImageFromURL(url), "")

	resp, err := sess.Predict(data).Do()
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return nil
	}

	return resp.Outputs

}

func DetectLabels(url string) {
	fmt.Println("Fetching labels for ", url)

	rekogLabels := rekognitionDetectLabels(url)
	fmt.Println("\r\n***************************************** Rekognition Labels :")
	for _, label := range rekogLabels {
		fmt.Printf("%s - %f\r\n", *label.Name, *label.Confidence)
	}

	visionLabels := visionDetectLabels(url)
	fmt.Println("\r\n***************************************** Vision Labels :")
	for _, label := range visionLabels {
		fmt.Printf("%s - %f\r\n", label.Description, label.Score)
	}

	clarifaiLabels := clarifaiDetectLabels(url)
	fmt.Println("\r\n***************************************** Clarifai Labels :")
	for _, concept := range clarifaiLabels[0].Data.Concepts {
		fmt.Printf("%s - %f\r\n", concept.Name, concept.Value)
	}
}

func main() {
	urls := []string {
		"http://cdn-thumbshot-ie.pearltrees.com/fa/a8/faa897cb9748e852545eae3b40d3c6d6-b52square.jpg",
		"https://static1.squarespace.com/static/574741382fe1310155f5d5b2/t/5767f8ac9c03e023b60ff3b1/1466433218590/image009.jpeg",
		"https://storage.googleapis.com/api-test-bucket-for-vision-api/WIN_20170119_12_22_15_Pro.jpg",
		"https://storage.googleapis.com/api-test-bucket-for-vision-api/IMG_3678%20(Edited).JPG"}

	for _, url := range urls {
		fmt.Println("\r\n\r\n")
		DetectLabels(url)
	}
}

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
