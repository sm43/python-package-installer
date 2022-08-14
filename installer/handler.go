package installer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type Installer struct {
	diskTargetLocation string
}

type response struct {
	Message string `json:"message"`
}

func NewInstaller() *Installer {
	targetLocation, _ := os.UserHomeDir()
	if targetLocation != "" {
		targetLocation = filepath.Join(targetLocation, "python-packages")

		if err := os.MkdirAll(targetLocation, os.ModePerm); err != nil {
			fmt.Println("failed to create target directory: ", err)
		}
	}
	return &Installer{
		diskTargetLocation: targetLocation,
	}
}

func (i *Installer) Handler() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		pkgs, ok := request.URL.Query()["install"]
		if !ok {
			responseWriter(http.StatusBadRequest, "package name missing!", response)
			return
		}

		if i.isAlreadyInstalled(pkgs[0]) {
			responseWriter(http.StatusCreated, fmt.Sprintf("package %s already installed!", pkgs[0]), response)
			return
		}

		go i.installPackage(pkgs[0])

		responseWriter(http.StatusCreated, fmt.Sprintf("installing package: %s", pkgs[0]), response)
	})
}

func responseWriter(statusCode int, message string, resp http.ResponseWriter) {
	body := response{
		Message: message,
	}
	resp.WriteHeader(statusCode)
	if err := json.NewEncoder(resp).Encode(body); err != nil {
		fmt.Println("failed to encode response : ", err)
	}
}
