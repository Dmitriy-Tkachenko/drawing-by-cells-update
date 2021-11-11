package main

import (
	"net/http"
	"os"
)

func main()  {
	wh := NewWebhook()
	wh.OnEvent(processingPhrasesAndClicks)

	http.HandleFunc("/",wh.HandleFunc)
	http.ListenAndServe(":" + os.Getenv("PORT"),nil)
}
