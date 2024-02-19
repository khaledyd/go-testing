package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type
	w.Header().Set("Content-Type", "text/html")
	// Write the HTML response with CSS
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Welcome to the Go Testing App</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					margin: 0;
					padding: 0;
					background-color: #f2f2f2;
				}
				.container {
					text-align: center;
					margin-top: 100px;
				}
				h1 {
					color: #333;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Welcome to the Go Testing App</h1>
			</div>
		</body>
		</html>
	`)
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server listening on port 9000")
	http.ListenAndServe(":9000", nil)
}
