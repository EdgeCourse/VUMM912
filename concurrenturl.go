/*
fetches issues from the xkcd comics website and downloads each URL to build an offline JSON index. At the time of writing, there are over 2500 comics (URLs) to download.

To do this sequentially (that is, one at a time), it would take a long time (probably hours), or the operation might fail

we will implement a Worker pool (to be explained later) to handle multiple HTTP requests at a time, keeping the connection alive and getting multiple results in a very short time.

A goroutine can be compared to a lightweight thread (although it’s not a thread, as many goroutines can work on a single thread) which makes it lighter, faster and reliable

When two or more goroutines are running, they need a way to communicate with each other: channels

The xkcd website has a JSON interface to allow external services use their API. Download the data from this interface to build our offline index. https://xkcd.com/info.0.json



*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

/*
Based on the JSON interface, design the struct to be used as a model for what data we want to extract for JSON handling:
*/

type Result struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

const Url = "https://xkcd.com"

/*
create a function that serves the core purpose of the application — fetching the comic.

create a custom HTTP client and set timeout to 5 seconds. Join the strings using the strings package, create a new request and send it using the previously created client. If the request is successful, we decode the data from JSON into our local struct. Then we close the response body and return a pointer to the struct.


*/

func fetch(n int) (*Result, error) {

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	// concatenate strings to get url; ex: https://xkcd.com/571/info.0.json

	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}

	var data Result

	// error from web service, empty struct to avoid disruption of process
	if resp.StatusCode != http.StatusOK {
		data = Result{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}

	resp.Body.Close()

	return &data, nil
}

type Job struct {
	number int
}

/*
Create a Worker pool to concurrently handle the operations by setting up buffered channels. A buffered channel is a channel with a specified capacity. With a buffered channel, send operations are blocked when the buffer is full and receive operations are blocked when the buffer is empty. We need this feature because in a Worker Pool, we assign multiple jobs to a number of workers and we want to ensure they are handled in an organized way.

If there are 6 workers in the worker pool, the buffered channel will ensure at every point in time, at most 6 jobs are given to the 6 workers.

*/

var jobs = make(chan Job, 100)
var results = make(chan Result, 100)
var resultCollection []Result

func allocateJobs(noOfJobs int) {
	for i := 0; i <= noOfJobs; i++ {
		jobs <- Job{i + 1}
	}
	close(jobs)
}

/*
After creating the buffered channels and setting up the final results variable, create a function to allocate jobs to the jobs channel. As expected, this function will block when i = 100, which means no new job will be added until a job has been received by the worker. After all available jobs have been allocated, the jobs channel will be closed to avoid further writes.

A worker pool maintains multiple threads (or in our case, goroutines) and waits for tasks (jobs) to be assigned to them. For example, let’s say we have 1000 jobs. We create a worker pool which spawns 100 workers. If the jobs channel is buffered at 100-capacity, the workers takes in the 100 jobs, and as some jobs are done processing, new jobs are being allocated, which goes to the workers, and so on.

A worker pool uses Go’s WaitGroup, a synchronization primitive (type) that tells the main goroutine to wait for a collection of goroutines to finish.

First define a worker function. The worker gets a job from the allocated jobs channel, processes the result, and passes the value to the results channel. In the createWorkerPool function, we use the WaitGroup primitive to set up a Worker pool. The wg.Add(1) call increments the WaitGroup counter. The counter must be zero if the program is to stop running (which is why we have the wg.Wait() call). The wg.Done() call in the worker function decrements the counter and if all is done, the control is returned to the main goroutine and the results channel is closed to prevent further writes.



*/

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		result, err := fetch(job.number)
		if err != nil {
			log.Printf("error in fetching: %v\n", err)
		}
		results <- *result
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i <= noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

/*
results are added to the results channel we created. But it is buffered and can only accept 100 at a time. We need a seperate goroutine to retrieve the results and give room for other results.

If the result from the results channel is valid, append it to the results collection. We have a boolean channel named “done”; use it to check if all the results have been collated.


*/

func getResults(done chan bool) {
	for result := range results {
		if result.Num != 0 {
			fmt.Printf("Retrieving issue #%d\n", result.Num)
			resultCollection = append(resultCollection, result)
		}
	}
	done <- true
}

/*
First, allocate jobs. Use 3000 because at the time of writing, xkcd has over 2500 comic issues, and we want to make sure we get all of them.


*/

func main() {
	// allocate jobs
	noOfJobs := 3000
	go allocateJobs(noOfJobs)

	// get results
	done := make(chan bool)
	go getResults(done)

	// create worker pool
	noOfWorkers := 100
	createWorkerPool(noOfWorkers)

	// wait for all results to be collected
	<-done

	// convert result collection to JSON
	data, err := json.MarshalIndent(resultCollection, "", "    ")
	if err != nil {
		log.Fatal("json err: ", err)
	}

	// write json data to file
	err = writeToFile(data)
	if err != nil {
		log.Fatal(err)
	}
}

func writeToFile(data []byte) error {
	f, err := os.Create("xkcd.json")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
