package main

import (
    "fmt"
    "os"
    "github.com/mohakkataria/gojek_takehome/lib"
)

func main() {
    argsWithoutProg := os.Args[1:]

    if len(argsWithoutProg) > 0 {
        // take commands from file
        fileName := argsWithoutProg[0]
        fmt.Print(fileName)
        lib.ReadAndProcessFromFile(fileName)
    } else {
        //We need to make it interactive session now
        lib.ReadAndProcessStdIn()
    }

}
