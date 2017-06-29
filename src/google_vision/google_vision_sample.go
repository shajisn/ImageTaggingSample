package main

import (
	"fmt"
	"log"
	//"os"

	// Imports the Google Cloud Vision API client package.
	"cloud.google.com/go/vision"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := vision.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	
	urls := []string{"http://cdn-thumbshot-ie.pearltrees.com/fa/a8/faa897cb9748e852545eae3b40d3c6d6-b52square.jpg", 
		"https://static1.squarespace.com/static/574741382fe1310155f5d5b2/t/5767f8ac9c03e023b60ff3b1/1466433218590/image009.jpeg"}
	

	// Sets the name of the image file to annotate.
	//filename := "vision/testdata/cat.jpg"

//	file, err := os.Open(filename)
//	if err != nil {
//		log.Fatalf("Failed to read file: %v", err)
//	}
//	defer file.Close()
	image := vision.NewImageFromURI(urls[0])
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	fmt.Println("Labels:")
	for _, label := range labels {
		fmt.Printf("%s - %f\r\n", label.Description, label.Score)
	}
}
