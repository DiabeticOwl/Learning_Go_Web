package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fullName := fmt.Sprintf(os.Args[1] + " " + os.Args[2])

	htmlString := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>Hello World!</title>
		</head>
		<body>
		<h1>` + fullName +
		`</h1>
		</body>
		</html>
	`)

	file, err := os.Create("index.html")
	if err != nil {
		fmt.Println("Error creating the file.", err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(htmlString))
}
