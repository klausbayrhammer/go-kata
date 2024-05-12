package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// At this point you should be able to create a web server from scratch, or by using `task-4` as a template.
// An interesting use of goroutines could be to have a pool of worker that can run tasks received via a HTTP POST in parallel.
// The following code can get you started to test it out you can use: `curl -X POST -d '{"name":"hello","parallel":2}' 127.0.0.1:8081/task/`
// What about getting something more done in each task?

// Task represent a task to execute in parallel, if you did not try out JSON in
// the previous tasks check out the docs.
// The weird strings after the field declaration are struct tags, and tell the
// json package how to match the json fields with the struct fields.
type Task struct {
	Name     string `json:"name"`
	Parallel int    `json:"parallel"`
}

// Here is a handler that uses some goroutines
func HandleTask(w http.ResponseWriter, r *http.Request) {
	var (
		task   Task
		result string
		start  = time.Now()
	)
	// read all the body
	buff, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "ERROR reading body: %+v", err)
	}
	r.Body.Close()

	// ummarshal the body in a variable
	err = json.Unmarshal(buff, &task)
	if err != nil {
		fmt.Fprintf(w, "ERROR umarshaling JSON %+v", err)
	}

	done := make(chan bool)
	// based on how parallel we need to be we run a goroutine
	for i := range task.Parallel {
		go func() { // func literals are closures
			fmt.Printf("executing task %q goroutine: %d\n", task.Name, i)
			done <- true
		}()
	}
	// now wait for all the goroutines we started to finish
	for range task.Parallel {
		<-done
	}
	result = fmt.Sprintf("task %q execute in %d parallel goroutines\nelapsed time %v\n", task.Name, task.Parallel, time.Now().Sub(start))
	fmt.Fprintf(w, "result: %s", result)
}
