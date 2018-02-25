package main

import "encoding/json"
import "fmt"


type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	fmt.Println("\n\nEncoding Example1");
	fmt.Println("Encoding or Marshaling means here -> Converting struct to Json")
	// Example1: Encoding or Marshaling struct data to json
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	fmt.Println("\n\nEncoding Example2");


	// Example2: Encoding or Marshaling struct data to json with different json tag name defined in structure
	// Encoding or Marshaling means here -> Converting struct to Json
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))



	// Decoding example: From Json to Structure
	// Decoding or Unmarshaling -> Converting json to struct
	fmt.Println("\n\nDecoding Example");
	fmt.Println("Decoding or Unmarshaling -> Converting json to struct")
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])


}

/*
Encoding Example1
Encoding or Marshaling means here -> Converting struct to Json
{"Page":1,"Fruits":["apple","peach","pear"]}


Encoding Example2
{"page":1,"fruits":["apple","peach","pear"]}


Decoding Example
Decoding or Unmarshaling -> Converting json to struct
{1 [apple peach]}
apple
*/
