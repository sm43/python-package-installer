package workerqueue

import (
	"fmt"

	"github.com/sm43/python-package-installer/installer"
)

var WorkerQueue chan chan string

func StartDispatcher(install *installer.Installer, nworkers int) {
	WorkerQueue = make(chan chan string, nworkers)

	for i := 0; i < nworkers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := NewWorker(i+1, WorkerQueue, install)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-installer.WorkQueue:
				fmt.Println("Received work request for package: ", work)
				go func() {
					worker := <-WorkerQueue

					fmt.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
