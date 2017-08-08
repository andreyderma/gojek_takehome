package lib

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "github.com/mohakkataria/gojek_takehome/parkingLot"
    "github.com/mohakkataria/gojek_takehome/car"
    "strconv"
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

var parkingLotInstance *parkingLot.ParkingLot

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

    if numberOfArguments, exists := allowedCommands[command]; exists {

        if len(arguments) != numberOfArguments {
            fmt.Println("Unsupported Command Arguments")
        }
        // after validation perform the necessary command
        switch command {
        case "create_parking_lot":
            if numberOfSlots, err := strconv.Atoi(arguments[0]); err != nil {
                fmt.Println("Error : " + err.Error())
            } else {
                parkingLotInstance.Initialize(numberOfSlots)
            }

        case "park":
            regNo := arguments[0]
            color := arguments[1]
            car := car.Create(regNo, color)
            parkingLotInstance.Park(car)

        case "leave":
            if slot, err := strconv.Atoi(arguments[0]); err != nil {
                fmt.Println("Error : " + err.Error())
            } else {
                parkingLotInstance.Leave(slot)
            }

        case "status":
            parkingLotInstance.Status()

        case "registration_numbers_for_cars_with_colour":
            color := arguments[0]
            parkingLotInstance.GetRegNosForCarsWithColor(color, true)

        case "slot_numbers_for_cars_with_colour":
            color := arguments[0]
            parkingLotInstance.GetSlotNosForCarsWithColor(color)

        case "slot_number_for_registration_number":
            regNo := arguments[0]
            parkingLotInstance.GetSlotNoForRegNo(regNo, true)
        }
    } else {
        fmt.Println("Unsupported Command")
    }
}

func init() {
    parkingLotInstance = parkingLot.GetInstance()
}
