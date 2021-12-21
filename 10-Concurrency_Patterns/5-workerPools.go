// Worker pools demo

package main

import (
	"fmt"
	"time"
)

func main() {

	const numJobs = 100                           // We have this number of jobs to complete
	jobsChan := make(chan int, numJobs)          // creates a buffered channel large enough to hold all jobs
	completedJobsChan := make(chan int, numJobs) // creates a buffered channel large enough to hold all completed jobs

	for w := 1; w <= 3; w++ { // this is number of workers.  Each worker will be a goroutine.
		go worker(w, jobsChan, completedJobsChan) // create worker w - pass worker number and both channels
	}

	for j := 1; j <= numJobs; j++ {
		jobsChan <- j // This loads the jobsChan channel with job numbers
	}
	close(jobsChan) // Close jobsChan channel for input after all jobs have been loaded.  Channel must be closed in order to call "range" function.

	for a := 1; a <= numJobs; a++ {
		<-completedJobsChan // Reads the completedJobsChan channel and does nothing with the contents.  Point is to clear the channel and to delay termination of the program until all jobs are reported as finished.
	}
}
func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) { // this syntax restricts the direction of each channel.  For THIS specific function, we will only SEND to completedJobsChan and RECEIVE from jobsChan

	for j := range jobsChan { // iterates (and RECEIVES) each and all the jobs in the channel.  Interesting that range seems to have its own receiver.
		fmt.Println("worker", id, "started  job", j, "with", len(jobsChan), "jobs left to process")
		time.Sleep(time.Second * 2) // simulates "work" that takes sleep time to complete
		fmt.Println("worker", id, "             finished job", j)
		completedJobsChan <- j // Loads finished job numbers into the completedJobsChan channel.
	}
}
