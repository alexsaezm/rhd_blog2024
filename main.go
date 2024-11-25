package main

import (
	"fmt"
	"os"
	"runtime"
)

func getRelease() error {
	content, err := os.ReadFile("/etc/redhat-release")
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not found")
		}
		return fmt.Errorf("failed to read file: %w", err)
	}
	fmt.Printf("OS: %s", string(content))
	return nil
}

func main() {
	getRelease()
	fmt.Printf("Arch: %s\n", runtime.GOARCH)
	fmt.Printf("Go version: %s\n", runtime.Version())
}
