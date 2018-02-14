package docker

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// DockerImage represents the base namespace for the Docker image
const DockerImage = "mdelapenya/liferay-portal-nightlies"

// dockerContainerName represents the name of the container to be run
const dockerContainerName = "liferay-portal-nightly"

// dockerBinary represents the name of the binary to execute Docker commands
const dockerBinary = "docker"

// checkDockerContainerExists checks if the container is running
func check(cmdArgs []string) bool {
	cmd := exec.Command(dockerBinary, cmdArgs...)

	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}

// checkDockerContainerExists checks if the container is running
func checkDockerContainerExists() bool {
	cmdArgs := []string{
		"container",
		"inspect",
		dockerContainerName,
	}

	return check(cmdArgs)
}

// checkDockerImageExists checks if the image is already present
func checkDockerImageExists(dockerImage string) bool {
	cmdArgs := []string{
		"image",
		"inspect",
		dockerImage,
	}

	return check(cmdArgs)
}

// DownloadDockerImage downloads the image
func downloadDockerImage(dockerImage string) {
	if checkDockerImageExists(dockerImage) {
		log.Println("Skipping pulling [" + dockerImage + "] as it's already present locally.")
		return
	}

	cmdArgs := []string{
		"pull",
		dockerImage,
	}

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command(dockerBinary, cmdArgs...)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}

	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("%s", errStr)
	fmt.Printf("%s", outStr)
}

// removeDockerContainer removes the running container
func removeDockerContainer() {
	cmdArgs := []string{
		"rm",
		"-fv",
		dockerContainerName,
	}

	cmd := exec.Command(dockerBinary, cmdArgs...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}

// RunDockerImage runs the image
func RunDockerImage(dockerImage string) {
	downloadDockerImage(dockerImage)

	if checkDockerContainerExists() {
		removeDockerContainer()
	}

	cmdArgs := []string{
		"run",
		"-d",
		"-p", "8080:8080",
		"-p", "11311:11311",
		"--name", dockerContainerName,
		dockerImage,
	}

	cmd := exec.Command(dockerBinary, cmdArgs...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", stdoutStderr)
}
