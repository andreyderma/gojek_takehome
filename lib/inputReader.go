package lib

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

//map of allowed commands along with the arguments to read
var allowedCommands = map[string]int{
    "create_parking_lot" : 1,
    "park" : 2,
    "leave" : 1,
    "status" : 0,
    "registration_numbers_for_cars_with_colour" : 1,
    "slot_numbers_for_cars_with_colour" : 1,
    "slot_number_for_registration_number" : 1,
}

func ReadAndProcessStdIn() {
    fmt.Println("Input:")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        command := strings.ToLower(scanner.Text())
        commandDelimited := strings.Split(command, " ")
        lengthOfCommand := len(commandDelimited)
        if lengthOfCommand < 1 {
            fmt.Fprintln(os.Stderr, "Unsupported Command")
            os.Exit(1)
        } else if lengthOfCommand == 1 {
            processCommand(commandDelimited[0], nil)
        } else {
            processCommand(commandDelimited[0], commandDelimited[1:])
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}

func ReadAndProcessFromFile(filePath string) {

}


func processCommand(command string, arguments []string) {
    fmt.Println(command, arguments)
    if numberOfArguments, exists := allowedCommands[command]; exists {
        fmt.Println(numberOfArguments)
        if len(arguments) != numberOfArguments {
            fmt.Println("Unsupported Command Arguments")
        }
    } else {
        fmt.Println("Unsupported Command")
    }
}
