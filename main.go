package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	//binary, lookErr := exec.LookPath("docker")
	//if lookErr != nil {
	//	fmt.Println("lookerr")
	//	panic(lookErr)
	//}

	args := []string{"build", "."}
	cmd := exec.Command("docker", args...)

	var out bytes.Buffer
	cmd.Stdout = &out

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Encountered an error while running command")
		log.Fatal(err)
	}

	fmt.Println("\nStdout:")
	stdOut := strings.Split(out.String(), "\n")
	for i, line := range stdOut {
		fmt.Printf("%d - %s\n", i, line)
	}

	fmt.Println("\nStdErr:")
	stdErr := strings.Split(errOut.String(), "\n")
	for i, line := range stdErr {
		fmt.Printf("%d - %s\n", i, line)
	}

}
