package services

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCodeInDocker(code string) (string, error) {
	// Save the user-provided code to a file
	userCodePath := "/tmp/user_code.js"
	err := os.WriteFile(userCodePath, []byte(code), 0644)
	if err != nil {
		return "", err
	}

	// Save the corresponding test file
	testCode := `
    const React = require('react');
    const { render, fireEvent } = require('@testing-library/react');
    const Counter = require('./user_code').default;

    test('increments and decrements the counter', () => {
      const { getByText } = render(<Counter />);
      const incrementButton = getByText('Increment');
      const decrementButton = getByText('Decrement');
      const counter = getByText('0');

      fireEvent.click(incrementButton);
      expect(counter).toHaveTextContent('1');

      fireEvent.click(decrementButton);
      expect(counter).toHaveTextContent('0');
    });
    `
	testFilePath := "/tmp/user_code.test.js"
	err = os.WriteFile(testFilePath, []byte(testCode), 0644)
	if err != nil {
		return "", err
	}

	// Define Docker command
	cmd := exec.Command("docker", "run", "--rm", "-v", fmt.Sprintf("%s:/usr/src/app/user_code.js", userCodePath), "-v", fmt.Sprintf("%s:/usr/src/app/user_code.test.js", testFilePath), "node:14")

	// Run the Docker container with the user's code
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}
