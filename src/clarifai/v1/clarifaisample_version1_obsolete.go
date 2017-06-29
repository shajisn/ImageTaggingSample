package main

import (
	"fmt"
	"github.com/clarifai/clarifai-go"
)

func main() {
	client := clarifai.NewClient("uX94X_AnN7W68a9Ms7G55UoGE3KKH5073Kn2PSht", "e9NtNTLzoNBdXvQnWoPu32Xl-9yRhUgR7XU_6W9A")
	// Get the current status of things
	info, err := client.Info()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n%+v\n", info)
	}
	// Let's get some context about these images
	urls := []string{"http://cdn-thumbshot-ie.pearltrees.com/fa/a8/faa897cb9748e852545eae3b40d3c6d6-b52square.jpg", 
		"https://static1.squarespace.com/static/574741382fe1310155f5d5b2/t/5767f8ac9c03e023b60ff3b1/1466433218590/image009.jpeg"}
	// Give it to Clarifai to run their magic
	tag_data, err := client.Tag(clarifai.TagRequest{URLs: urls})

	if err != nil {
		fmt.Println("\nSome erorr..........")
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", tag_data) // See what we got!
	}

	feedback_data, err := client.Feedback(clarifai.FeedbackForm{
		URLs:    urls,
		AddTags: []string{"cat", "animal"},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", feedback_data)
	}
}
