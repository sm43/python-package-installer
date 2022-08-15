package workerqueue

import (
	"fmt"

	"github.com/sm43/python-package-installer/installer"
)

func NewWorker(id int, workerQueue chan chan string, installer *installer.Installer) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan string),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
		Installer:   installer,
	}

	return worker
}

type Worker struct {
	ID          int
	Work        chan string
	WorkerQueue chan chan string
	QuitChan    chan bool
	Installer   *installer.Installer
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				fmt.Printf("worker %d: processing request for: %s\n", w.ID, work)
				if err := w.Installer.InstallPackage(work); err != nil {
					// TODO: if installation failed due to non permanent error, then requeue
					fmt.Printf("worker %d: failed to install package %s", w.ID, work)
					continue
				}
				fmt.Printf("worker %d: processing completed for %s!\n", w.ID, work)

			case <-w.QuitChan:
				fmt.Printf("worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
