package services

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCodeInDocker(code string, tests string) (string, error) {
	// Save the user code to a file
	userCodePath := "/tmp/user_code.mjs"
	err := os.WriteFile(userCodePath, []byte(code), 0644)
	if err != nil {
		return "", err
	}

	// Save the test code to a file
	testCodePath := "/tmp/test_code.test.js"
	err = os.WriteFile(testCodePath, []byte(tests), 0644)
	if err != nil {
		return "", err
	}

	// Define Docker command
	cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/usr/src/app/user_code.mjs", userCodePath), "-v", fmt.Sprintf("%s:/usr/src/app/test_code.test.js", testCodePath), "react-practice-runner")

	// Run the Docker container with the user code and tests
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}
