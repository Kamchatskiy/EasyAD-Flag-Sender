package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

var teamToken = os.Getenv("TEAM_TOKEN")
var judgeIp = os.Getenv("JUDGE_IP")
var judgePort = ":31337"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			flag := r.Form.Get("flag")
			sendTCPRequest(flag)
			fmt.Fprintf(w, "<html>"+
				"<head>"+
				"<style>"+
				"body {"+
				"font-family: Arial, sans-serif;"+
				"/* Consistent font */"+
				"background-color: #f4f4f9;"+
				"/* Matching background color */"+
				"display: flex;"+
				"justify-content: center;"+
				"align-items: center;"+
				"height: 100vh;"+
				"margin: 0;"+
				"text-align: center;"+
				"/* Center text */"+
				"}"+
				"div {"+
				"padding: 20px;"+
				"background-color: #ffffff;"+
				"box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);"+
				"/* Matching form shadow */"+
				"border-radius: 8px;"+
				"/* Matching form border-radius */"+
				"width: fit-content;"+
				"/* Auto-adjust width */"+
				"max-width: 80%;"+
				"/* Maximum width to prevent too wide boxes on large screens */"+
				"}"+
				"p {"+
				"color: #333333;"+
				"font-size: 16px;"+
				"/* Consistent font size */"+
				"margin: 10px 0;"+
				"/* Consistent spacing */"+
				"}"+
				"</style>"+
				"</head>"+
				"<body>"+
				"<div>"+
				"<p>Флаг %s отправлен</p>"+
				"</html>", flag)
			fmt.Fprintf(w, "<html>"+
				"<p>Токен команды %s</p>"+
				"</div>"+
				"</body>", teamToken)
		} else {
			fmt.Fprintf(w,
				"<html>"+
					"<head>"+
					"<style>"+
					"body {"+
					"font-family: Arial, sans-serif;"+
					"/* Set a more modern font */"+
					"background-color: #f4f4f9;"+
					"/* Soft background color */"+
					"display: flex;"+
					"justify-content: center;"+
					"align-items: center;"+
					"height: 100vh;"+
					"margin: 0;"+
					"}"+
					"form {"+
					"padding: 20px;"+
					"background-color: #ffffff;"+
					"box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);"+
					"/* Subtle shadow for the form */"+
					"border-radius: 8px;"+
					"/* Rounded corners */"+
					"}"+
					"label {"+
					"font-size: 16px;"+
					"/* Larger font size for readability */"+
					"color: #333333;"+
					"display: block;"+
					"/* Makes label take the full width */"+
					"margin-bottom: 10px;"+
					"/* Uniform margin for elements */"+
					"}"+
					"input[type='text'],"+
					"input[type='submit'] {"+
					"width: 100%;"+
					"/* Full width for text input and button */"+
					"padding: 8px;"+
					"margin-bottom: 10px;"+
					"/* Uniform margin for elements */"+
					"box-sizing: border-box;"+
					"/* Include padding and border in the element's width */"+
					"}"+
					"input[type='text'] {"+
					"border: 1px solid #cccccc;"+
					"border-radius: 4px;"+
					"}"+
					"input[type='submit'] {"+
					"background-color: #4CAF50;"+
					"/* Green background */"+
					"color: white;"+
					"border: none;"+
					"cursor: pointer;"+
					"border-radius: 4px;"+
					"transition: background-color 0.3s;"+
					"/* Smooth transition for hover effect */"+
					"}"+
					"input[type='submit']:hover {"+
					"background-color: #45a049;"+
					"/* Slightly darker green when hovered */"+
					"}"+
					"</style>"+
					"</head>"+
					"<body>"+
					"<form method='post'>"+
					"<label for='flag'>Введите флаг:</label>"+
					"<input type='text' id='flag' name='flag'><br>"+
					"<input type='submit' value='Отправить'>"+
					"</form>"+
					"</body>"+
					"</html>")
		}
	})

	http.ListenAndServe(":80", nil)
}

func sendTCPRequest(flag string) {
	conn, err := net.Dial("tcp", judgeIp+judgePort)
	if err != nil {
		fmt.Println("tcp connection error", err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte(teamToken + "\n" + flag))
	if err != nil {
		fmt.Println("error on writing data", err)
		return
	}
}
