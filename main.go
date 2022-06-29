package main

import (
	"encoding/json"
	"fmt"
)

func main(){

	reviews := RefreshReviews()

	json, err := json.Marshal(reviews)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}
