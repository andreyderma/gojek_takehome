package lib

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

// Function to read from stdin and process the command with arguments
func ReadAndProcessStdIn() {
    fmt.Println("Input:")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        command := strings.ToLower(scanner.Text())
        processCommand(command)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}


// Function to read from file line by line and process the command with arguments
func ReadAndProcessFromFile(filePath string) {
    f, err := os.Open(filePath)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        processCommand(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

