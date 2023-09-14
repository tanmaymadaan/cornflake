package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	cornflakeIndex := os.Getenv("INDEX")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Cornflake %s is up and running", cornflakeIndex)
		if err != nil {
			return
		}
	})

	fmt.Printf("Cornflake %s is up and running on port: %s", cornflakeIndex, os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}
