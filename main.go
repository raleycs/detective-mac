package main

import (
    "fmt"
    "os/exec"
)

// menu displays all available tools for use
// it will prompt for user selection and return
// the numerical option that they have chosen
func menu() string {
    var response string

    fmt.Println("Please select a tool:")
    fmt.Println("1) Timestamp Changer")
    fmt.Println("2) Portscanner")
    fmt.Scanln(&response)
    return response
}

func main() {

    fmt.Println("Welcome to the Attacker's Toolbox!")

    response := menu()

    fmt.Println("You have selected " + response)

    cmd := "echo"

    cmd_arg0 := "Hello world!"

    stdout, err := exec.Command(cmd, cmd_arg0).Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(stdout))
}
