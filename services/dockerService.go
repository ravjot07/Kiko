package services

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCodeInDocker(code string) (string, error) {
	// Save the code to a file
	userCodePath := "/tmp/user_code.mjs"
	err := os.WriteFile(userCodePath, []byte(code), 0644)
	if err != nil {
		return "", err
	}

	// Define Docker command
	cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/usr/src/app/user_code.mjs", userCodePath), "node:14")

	// Run the Docker container with the user's code
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}
