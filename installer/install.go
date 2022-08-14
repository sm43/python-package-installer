package installer

import (
	"fmt"
	"net/http"
	"os/exec"
)

type Installer struct {
	diskTargetLocation string
}

func NewInstaller() *Installer {
	return &Installer{
		diskTargetLocation: "",
	}
}

func (i *Installer) Handler() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// TODO: validate input

		go i.installPackage("install")

		fmt.Fprint(response, "request received!")
	})
}

func (i *Installer) installPackage(pkg string) {
	venv := getVenvName()
	defer cleanup(venv)

	// create venv
	cmd, err := exec.Command("/bin/sh", "installer/create.sh", venv).Output()
	if err != nil {
		fmt.Println("failed to create venv: ", err)
		return
	}
	fmt.Println("create output: ", string(cmd))

	// install package
	targetDir := pkg
	cmd, err = exec.Command("/bin/sh", "installer/install.sh", venv, targetDir, pkg).Output()
	if err != nil {
		fmt.Println("failed to install package: ", err)
		return
	}
	fmt.Println("install output: ", string(cmd))

	// zip and copy package
	cmd, err = exec.Command("/bin/sh", "installer/copy.sh", venv, targetDir, fmt.Sprint(pkg+".zip")).Output()
	if err != nil {
		fmt.Println("failed to zip and copy package: ", err, string(cmd))
		return
	}
	fmt.Println("copy output: ", string(cmd))
}

func cleanup(venv string) {
	cmd, err := exec.Command("/bin/sh", "installer/cleanup.sh", venv).Output()
	if err == nil {
		fmt.Println(string(cmd))
	}
}
