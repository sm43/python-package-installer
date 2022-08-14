package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	log.Println("Starting Server...")

	http.HandleFunc("/", installHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func installHandler() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// TODO: validate input

		installPackage("install")

		fmt.Fprint(response, "package installed!")
	})
}

func installPackage(pkg string) {
	venv := "venv"
	defer cleanup(venv)

	// create venv
	cmd, err := exec.Command("/bin/sh", "installer/create.sh", venv).Output()
	if err != nil {
		fmt.Println("failed to create venv: ", err)
	}
	fmt.Println("create output: ", string(cmd))

	// install package
	targetDir := pkg
	cmd, err = exec.Command("/bin/sh", "installer/install.sh", venv, targetDir, pkg).Output()
	if err != nil {
		fmt.Println("failed to install package: ", err)
	}
	fmt.Println("install output: ", string(cmd))

	// zip and copy package
	cmd, err = exec.Command("/bin/sh", "installer/copy.sh", venv, targetDir, fmt.Sprint(pkg+".zip")).Output()
	if err != nil {
		fmt.Println("failed to zip and copy package: ", err, string(cmd))
	}
	fmt.Println("copy output: ", string(cmd))
}

func cleanup(venv string) {
	cmd, err := exec.Command("/bin/sh", "installer/cleanup.sh", venv).Output()
	if err == nil {
		fmt.Println(string(cmd))
	}
}
