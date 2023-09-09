package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Cornflake %v is up and running", os.Getenv("INDEX"))
		if err != nil {
			return
		}
	})

	fmt.Println("Backend server listening on :" + os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}
