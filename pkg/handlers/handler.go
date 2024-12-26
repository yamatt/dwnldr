package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/yamatt/dwnldr/pkg/config"
)

func Download(url string, outputDir string, config config.Config) error {
	// Find the matching handler for the URL
	handler := findHandler(url, config.Handlers)
	if handler == nil {
		return fmt.Errorf("no handler found for URL: %s", url)
	}

	// Print which handler is being used
	fmt.Printf("Using handler: %s for URL: %s\n", handler.Name, url)

	// Run the command associated with the handler
	err := runCommand(handler.Command, url, outputDir)
	if err != nil {
		return fmt.Errorf("error running command: %w", err)
	}

	return nil
}

// findHandler tries to find a handler that matches the URL
func findHandler(url string, handlers []config.Handler) *config.Handler {
	for _, handler := range handlers {
		if matchesPattern(url, handler.Pattern) {
			return &handler
		}
	}
	return nil
}

// matchesPattern matches the URL against the pattern (basic prefix match)
func matchesPattern(url, pattern string) bool {
	// Compile the regex pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		// If the regex pattern is invalid, print an error and return false
		fmt.Printf("Invalid regex pattern: %s\n", pattern)
		return false
	}

	// Match the URL against the compiled regex pattern
	return re.MatchString(url)
}

// runCommand runs the command specified by the handler
func runCommand(command []string, url string, outputDir string) error {
	// Append the URL to the command
	cmdArgs := append(command, url)

	commandWithArgs := replacePlaceholders(command, url, outputDir)

	// Create the command using exec.Command
	cmd := exec.Command(commandWithArgs[0], commandWithArgs[1:]...)

	// Print the command that will be run
	fmt.Printf("Running command: %s %s\n", cmdArgs[0], strings.Join(cmdArgs[1:], " "))

	// Set the command's output to the terminal (stdout and stderr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("command failed: %w", err)
	}

	return nil

}

func replacePlaceholders(command []string, url string, outputDir string) []string {
	// Create a new slice to hold the updated command
	var updatedCommand []string

	// Iterate through each part of the command
	for _, part := range command {
		// Replace {url} with the actual URL
		part = strings.ReplaceAll(part, "{url}", url)

		// Replace {outputDir} with the actual output directory
		part = strings.ReplaceAll(part, "{outputDir}", outputDir)

		// Append the updated part to the updated command
		updatedCommand = append(updatedCommand, part)
	}

	return updatedCommand
}
