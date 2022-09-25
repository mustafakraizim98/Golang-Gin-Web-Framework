# Golang-Gin-Web-Framework
Gin is a high-performance micro-framework that can be used to build web applications and microservices. It makes it simple to build a request handling pipeline from modular, reusable pieces. It does this by allowing you to write middleware that can be plugged into one or more request handlers or groups of request handlers.

## Download and install Gin:
> go get github.com/gin-gonic/gin

## Run the server:
> go run main.go

## To use POST request:
You just need to test the request through JSON object, as shown below:
```
{
    "sender": "Alice",
    "receiver": "Bob",
    "message": "Hey Bob, I'm sending a new message to you."
}
```

## Features of the project:
- Gin Framework >>> `#0969DA` Done
- Endpoint declaration >>> `#0969DA` Done
- Validating response for status of requests >>> `#0969DA` Done
- Passing parameters for GET request >>> `#0969DA` Done
- Returns an array of objects with sender, receiver and message content in GET request >>> `#0969DA` Done
- Chronological descending order >>> `#0969DA` Done

## The most advantage in the project:
Using chronological to sort the object in a descending order based on the time and date of messages, a slice of code shown here:
```
CreatedAt   time.Time `json:"created_at"`
```

```
CreatedAt:   time.Now().Add(24 * time.Hour),
```

```
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
```

```
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
```
