package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Message struct {
	Sender      string    `json:"sender"`
	Receiver    string    `json:"receiver"`
	MessageBody string    `json:"message"`
	CreatedAt   time.Time `json:"created_at"`
}

var messages = []Message{
	{
		Sender:      "Bob",
		Receiver:    "Alice",
		MessageBody: "Goodbye Alice.",
		CreatedAt:   time.Now().Add(24 * time.Hour),
	},
	{
		Sender:      "Bob",
		Receiver:    "Alice",
		MessageBody: "Hey Alice, What a great profile you have.",
	},
	{
		Sender:      "Bob",
		Receiver:    "Alice",
		MessageBody: "Hey Alice, I'm glad to connect with you.",
		CreatedAt:   time.Now(),
	},
	{
		Sender:      "Bob",
		Receiver:    "Alice",
		MessageBody: "Hey Alice, That's a great thing to keep in touch with you.",
		CreatedAt:   time.Now().Add(12 * time.Hour),
	},
	{
		Sender:      "Alice",
		Receiver:    "Bob",
		MessageBody: "Hey Bob, I'm glad to connect with you.",
		CreatedAt:   time.Now(),
	},
}

type timeSlice []Message

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].CreatedAt.After(p[j].CreatedAt)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getMessageByParameters(sender, receiver string) (timeSlice, error) {
	counter := 0
	var founded []Message

	for _, obj := range messages {
		if strings.EqualFold(sender, obj.Sender) && strings.EqualFold(receiver, obj.Receiver) {
			founded = append(founded, obj)
			counter++
		}
	}

	if counter == 0 {
		return nil, errors.New("message not found")
	} else {
		dateSortedReviews := make(timeSlice, 0, len(founded))

		for _, d := range founded {
			dateSortedReviews = append(dateSortedReviews, d)
		}

		sort.Sort(dateSortedReviews)

		return dateSortedReviews, nil
	}
}

func getHttpRequestJson(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, messages)
}

func getHttpRequestHtml(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"content": "404",
	})
}

func getHttpRequestJsonByParameters(context *gin.Context) {
	sender := context.Param("sender")
	receiver := context.Param("receiver")
	message, err := getMessageByParameters(sender, receiver)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"response": "message not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, message)
}

func postHttpRequestJson(context *gin.Context) {
	var newMessage Message

	if err := context.BindJSON(&newMessage); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"response": "bad request"})
	} else {
		messages = append(messages, newMessage)
		context.IndentedJSON(http.StatusOK, newMessage)
	}
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")

	router.POST("/message", postHttpRequestJson)
	router.GET("/", getHttpRequestHtml)
	router.GET("/message/list", getHttpRequestJson)
	router.GET("/message/list/:sender/:receiver", getHttpRequestJsonByParameters)

	err := router.Run("localhost:8081")
	if err != nil {
		return
	}
}
