package installer

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (i *Installer) InstallPackage(pkg string) error {
	venv := getVenvName()
	defer cleanup(venv)

	// create venv
	cmd, err := exec.Command("/bin/sh", "installer/scripts/create.sh", venv).Output()
	if err != nil {
		fmt.Println("failed to create venv: ", venv, err)
		return err
	}
	fmt.Println("create output: ", string(cmd))

	// install package
	targetDir := getTargetDir(pkg)
	cmd, err = exec.Command("/bin/sh", "installer/scripts/install.sh", venv, targetDir, pkg).Output()
	if err != nil {
		fmt.Println("failed to install package: ", pkg, err)
		return err
	}
	fmt.Println("install output: ", string(cmd))

	// zip and copy package
	cmd, err = exec.Command("/bin/sh", "installer/scripts/copy.sh", venv, targetDir,
		i.diskTargetLocation).Output()
	if err != nil {
		fmt.Println("failed to zip and copy package: ", err, string(cmd))
		return err
	}
	fmt.Println("copy output: ", string(cmd))

	return nil
}

func cleanup(venv string) {
	cmd, err := exec.Command("/bin/sh", "installer/scripts/cleanup.sh", venv).Output()
	if err == nil {
		fmt.Println(string(cmd))
	}
}

func getTargetDir(pkg string) string {
	if !strings.Contains(pkg, "https://") {
		return pkg
	}
	targetDir := strings.Split(pkg, "https://")[1]
	return strings.ReplaceAll(targetDir, "/", "_")
}

func (i *Installer) isAlreadyInstalled(pkg string) bool {
	targetDir := getTargetDir(pkg)
	if _, err := os.Stat(filepath.Join(i.diskTargetLocation, targetDir+".zip")); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
