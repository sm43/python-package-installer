package installer

import (
	"fmt"
	"os/exec"
)

func (i *Installer) installPackage(pkg string) {
	venv := getVenvName()
	defer cleanup(venv)

	// create venv
	cmd, err := exec.Command("/bin/sh", "installer/scripts/create.sh", venv).Output()
	if err != nil {
		fmt.Println("failed to create venv: ", err)
		return
	}
	fmt.Println("create output: ", string(cmd))

	// install package
	targetDir := pkg
	cmd, err = exec.Command("/bin/sh", "installer/scripts/install.sh", venv, targetDir, pkg).Output()
	if err != nil {
		fmt.Println("failed to install package: ", err)
		return
	}
	fmt.Println("install output: ", string(cmd))

	// zip and copy package
	cmd, err = exec.Command("/bin/sh", "installer/scripts/copy.sh", venv, targetDir,
		i.diskTargetLocation).Output()
	if err != nil {
		fmt.Println("failed to zip and copy package: ", err, string(cmd))
		return
	}
	fmt.Println("copy output: ", string(cmd))
}

func cleanup(venv string) {
	cmd, err := exec.Command("/bin/sh", "installer/scripts/cleanup.sh", venv).Output()
	if err == nil {
		fmt.Println(string(cmd))
	}
}
