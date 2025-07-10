package main

import (
	"encoding/json"
	"fmt"
)

func json_() {

	user1 := User{Id: 1, Name: "John", Email: "john@example.com", Password: "123456", Created_at: "2021-01-01"}
	user2 := User{Id: 2, Name: "Harry", Email: "harry@example.com", Password: "123456", Created_at: "2021-01-01"}

	// marshal
	jsonData1, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("JSON Data:", string(jsonData1))
	}

	jsonData2, err := json.MarshalIndent(user2, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("JSON Data:", string(jsonData2))
	}

	// unmarshal
	jsonData := []byte(`{"id": 101, "name": "Laptop","price": 100.5, "category": "Electronics"}`)

	var product1 Product

	err2 := json.Unmarshal(jsonData, &product1)
	if err != nil {
		fmt.Println("Error:", err2)
	}
	fmt.Println("Product:", product1)

	// generic json
	var genericData map[string]interface{}

	err3 := json.Unmarshal(jsonData, &genericData)
	if err != nil {
		fmt.Println("Error:", err3)
	}
	fmt.Println("Generic Data:", genericData)

}

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"-"`
	Created_at string `json:"created_at"`
}

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}
