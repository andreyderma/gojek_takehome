package lib

import (
    "github.com/mohakkataria/gojek_takehome/car"
    "strconv"
    "fmt"
    "github.com/mohakkataria/gojek_takehome/parkingLot"
    "os"
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


// Process the command taken in from file/stdin
// Separate the command and arguments for command
// Validate the command and then do the necessary action
func processCommand(command string) bool {
    commandDelimited := strings.Split(command, " ")
    lengthOfCommand := len(commandDelimited)
    arguments := []string{}
    if lengthOfCommand < 1 {
        fmt.Fprintln(os.Stderr, "Unsupported Command")
        return false
        os.Exit(1)
    } else if lengthOfCommand == 1 {
        command = commandDelimited[0]
    } else {
        command = commandDelimited[0]
        arguments = commandDelimited[1:]
    }

    // check if command is one of the allowed commands
    if numberOfArguments, exists := allowedCommands[command]; exists {

        if len(arguments) != numberOfArguments {
            fmt.Println("Unsupported Command Arguments")
            return false
        }

        // after validation of number of arguments per command, perform the necessary command
        switch command {
        case "create_parking_lot":
            if numberOfSlots, err := strconv.Atoi(arguments[0]); err != nil {
                fmt.Println("Error : " + err.Error())
                return false
            } else {
                parkingLot.Initialize(numberOfSlots)
                return true
            }

        case "park":
            regNo := arguments[0]
            color := arguments[1]
            car := car.Create(regNo, color)
            err := parkingLot.Park(car)
            if err != nil {
                return false
            }
            return true

        case "leave":
            if slot, err := strconv.Atoi(arguments[0]); err != nil {
                fmt.Println("Error : " + err.Error())
                return false
            } else {
                err := parkingLot.Leave(slot)
                if err != nil {
                    return false
                }
                return true
            }

        case "status":
            parkingLot.Status()
            return true

        case "registration_numbers_for_cars_with_colour":
            color := arguments[0]
            _, err := parkingLot.GetRegNosForCarsWithColor(color, true)
            if err != nil {
                return false
            }
            return true

        case "slot_numbers_for_cars_with_colour":
            color := arguments[0]
            err := parkingLot.GetSlotNosForCarsWithColor(color)
            if err != nil {
                return false
            }
            return true


        case "slot_number_for_registration_number":
            regNo := arguments[0]
            _, err := parkingLot.GetSlotNoForRegNo(regNo, true)
            if err != nil {
                return false
            }
            return true

        default:
            return false
        }
    } else {
        fmt.Println("Unsupported Command")
        return false
    }
}


