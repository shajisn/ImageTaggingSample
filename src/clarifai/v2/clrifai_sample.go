package main 

import (
	cl "github.com/mpmlj/clarifai-client-go"
)

func main() {

	sess, err := cl.Connect("uX94X_AnN7W68a9Ms7G55UoGE3KKH5073Kn2PSht", "e9NtNTLzoNBdXvQnWoPu32Xl-9yRhUgR7XU_6W9A")
	if err != nil {
		panic(err)
	}
	
	data := cl.InitInputs()
	
	urls := []string{"http://cdn-thumbshot-ie.pearltrees.com/fa/a8/faa897cb9748e852545eae3b40d3c6d6-b52square.jpg", 
		"https://static1.squarespace.com/static/574741382fe1310155f5d5b2/t/5767f8ac9c03e023b60ff3b1/1466433218590/image009.jpeg"}
	
	// Option A. Adding an image from URL.
	_ = data.AddInput(cl.NewImageFromURL(urls[0]), "")
	// 2nd file...
	//_ = data.AddInput(cl.NewImageFromURL(urls[1]), "")
	// Option B. Adding an image from a local file.
	// NOTE. Currently API does not accept a mix of URL and base64 - based images!
	//i, err := cl.NewImageFromFile("../Dave_Gahan_New_York_2015-10-22.jpg")
	//if err != nil {
	//	panic(err)
	//}
	//_ = data.AddInput(i, "")

	// As per https://developer-preview.clarifai.com/guide/predict#predictBy,
	// general model is used by default (ID "aaa03c23b3724a16a56b629203edc62c"),
	// but you can also set your own model using a SetModel() method on your input.
	// data.SetModel("music-model-id-1")
	//data.SetModel(cl.PublicModelTravel)

	resp, err := sess.Predict(data).Do()
	if err != nil {
		panic(err)
	}

	cl.PP(resp)
}