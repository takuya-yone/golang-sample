package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	DueDate time.Time `json:"date"`
}

var tasks = []Task{{
	ID:      1,
	Title:   "A",
	Content: "A-task",
	DueDate: time.Now(),
}, {
	ID:      2,
	Title:   "B",
	Content: "B-task",
	DueDate: time.Now(),
}, {
	ID:      3,
	Title:   "C",
	Content: "C-task",
	DueDate: time.Now(),
}}

func main() {
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		if err := enc.Encode(&tasks); err != nil {
			log.Fatal(err)
		}
		fmt.Println(buf.String())

		_, err := fmt.Fprint(w, buf.String())
		if err != nil {
			return
		}
	}

	// GET /tasks
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
