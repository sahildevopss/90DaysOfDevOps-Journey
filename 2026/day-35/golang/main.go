package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
<!DOCTYPE html>
<html>
<head>
	<title>DevOps Container App</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background: #f4f4f4;
			text-align: center;
			padding-top: 80px;
		}

		.container {
			background: white;
			width: 600px;
			margin: auto;
			padding: 40px;
			border-radius: 12px;
			box-shadow: 0 4px 12px rgba(0,0,0,0.1);
		}

		h1 {
			color: #222;
		}

		.status {
			color: green;
			font-weight: bold;
		}
	</style>
</head>

<body>
	<div class="container">
		<h1>DevOps Container App</h1>

		<p class="status">
			Application is running successfully!
		</p>

		<p>Built with Go</p>
		<p>Containerized with Docker</p>
		<p>Optimized using Multi-Stage Builds</p>
	</div>
</body>
</html>
`)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server failed:", err)
	}
}
