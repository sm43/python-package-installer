package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sm43/python-package-installer/installer"
	"github.com/sm43/python-package-installer/workerqueue"
)

func main() {
	NWorkers := flag.Int("number-of-workers", 2, "The number of workers to start")
	Port := flag.String("port", "8080", "The port on which the server would be running")
	flag.Parse()

	log.Println("Initializing Installer...")
	installer := installer.NewInstaller()

	log.Println("Starting Dispatcher...")
	workerqueue.StartDispatcher(installer, *NWorkers)

	// TODO: the API cam return an ID which user can use to check
	// the status of installation
	http.HandleFunc("/", installer.Handler())

	log.Printf("Starting Server on %v...", *Port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":"+*Port), nil))
}
