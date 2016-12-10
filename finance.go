
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Currencylayer struct {
	Success   bool   `json:"success"`
	Terms     string `json:"terms"`
	Privacy   string `json:"privacy"`
	Timestamp int    `json:"timestamp"`
	Source    string `json:"source"`
	Quotes    struct {
		Usdeur float64 `json:"USDEUR"`
		Usdinr float64 `json:"USDINR"`
	} `json:"quotes"`
}

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s { // _ = key, v = values
		sum += v
	}
	c <- sum // send sum to c
}

func convert_currency(input int){
    //Taken from https://currencylayer.com/quickstart
	url := fmt.Sprintf("http://apilayer.net/api/live?access_key=091577bcb9cb07feefdc94a978d54231&currencies=EUR,INR&source=USD")

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// & refers to a pointer
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record Currencylayer

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	// We want to know what 100 Indian Rupee (INR) is worth in Euros (EUR).
	// The formula is USDEUR/USDINR*100
	fmt.Println(input,"rupiah is worth in Euros = ", (record.Quotes.Usdeur/record.Quotes.Usdinr*float64(input)))
}


func main() {
    //Start with an array of integers
    // figure out how to build an array of randomly generated integers
    s := []int{327, 112, 82, 239, 214, 100, 99, 2, 123, 432, 73, 49}

	c := make(chan int) //make channel
	go sum(s[:len(s)/2], c) //sum 1st half of array
	go sum(s[len(s)/2:], c) //sum 2nd half of array
	x, y := <-c, <-c // receive from c

    convert_currency(x)
    convert_currency(y)
    convert_currency(x+y)

}
