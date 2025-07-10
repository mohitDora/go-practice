package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func http_() {
	/*
	  - http is a package that provides HTTP client and server implementations.
	  type Handler interface{
	    ServeHTTP(ResponseWriter, *Request)
	  }

	  func(w http.ResponseWriter, r *http.Request)

	  http..ResponseWriter is an interface that has methods to write HTTP response
	  - Header() Header
	  - Write([]byte) (int, error)
	  - WriteHeader(statusCode int)

	  http.Request is a struct that represents an HTTP request received by a server or to be sent by a client.

	  http.ServeMux is basically a router that matches the URL path to a handler function and redirects the request to the handler function.
	  It is implicitly used by http.HandleFunc but we can make our own router using http.ServeMux.
	*/
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/item", ItemHandler)
	http.Handle("/secure", LoggingMiddleware(AuthMiddleware(http.HandlerFunc(SecureHandler))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!, %s, %v, %v", r.URL.Path, r.Method, r.Proto)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
}

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		itemId := r.URL.Query().Get("id")
		fmt.Fprintf(w, "Get Item %s", itemId)
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var item Item
		err = json.Unmarshal(body, &item)
		if err != nil {
			http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Create Item %v", item)

	default:
		fmt.Fprintf(w, "Method not allowed")
	}
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secure Page")
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Logging middleware")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Auth middleware")
		if r.Header.Get("Authorization") != "secret" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
