package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// curl -X GET -H "Cache-Control: no-cache" "http://localhost:1323/start"
// curl -X POST -H "Accept: application/json" -H "Content-Type: application/json" -H "Cache-Control: no-cache" -d '{
// "data":[
// 	{"name":"1"},
// 	{"name":"2"},
// 	{"name":"3"},
// 	{"name":"4"},
// 	{"name":"5"}
// 	]
// }' "http://localhost:1323/"

var (
	// MaxWorker       = os.Getenv("MAX_WORKERS")
	// MaxQueue        = os.Getenv("MAX_QUEUE")

	MaxWorker = "10"
	MaxQueue  = "10"

	MaxLength int64 = 1024
)

// A buffered channel that we can send work requests on.
var JobQueue chan Job

func init() {
	maxQueue, _ := strconv.Atoi(MaxQueue)
	JobQueue = make(chan Job, maxQueue)
}

func main() {
	e := echo.New()
	e.Get("/start", func(c echo.Context) error {
		maxWorker, _ := strconv.Atoi(MaxWorker)
		dispatcher := NewDispatcher(maxWorker)
		dispatcher.Run()
		fmt.Println("start!")
		return c.String(http.StatusOK, "Start!!")
	})
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})
	e.POST("/", func(c echo.Context) error {
		content := &PayloadCollection{}
		if err := c.Bind(content); err != nil {
			return err
		}
		// Go through each payload and queue items individually to be posted to S3
		for _, payload := range content.Payloads {

			// let's create a job with the payload
			work := Job{Payload: payload}

			// Push the work onto the queue.
			JobQueue <- work
		}
		return c.String(http.StatusOK, "CARETE!")
	})
	e.Run(standard.New(":1323"))
}
