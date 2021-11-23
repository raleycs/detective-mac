package main

import (
    "fmt"
    "os"
    "runtime"
    "strings"
    "github.com/raleycs/attackers-toolbox/internal"
)

// finds all .DS_Store files for a specific
// user and extracts information out of those files
func AnalyzeDsStore() {
    var username string // holds target's username

    // get user input
    fmt.Print("Enter a username: ")
    fmt.Scanln(&username)
    fmt.Println()

    // confirm that the user exists on the current system
    path := "/Users/" + username
    fmt.Println("Scanning DS_Stores for " + username)
    if fsService.FileExists(path) == false {
        fmt.Println("[*] User does not exist...")
        fmt.Println("[*] Exiting program")
        os.Exit(0)
    }

    // retrieve all .DS_Store files for the given user
    files := fsService.RetrieveFiles(".DS_Store", path)
    fmt.Println("[*] Retrived the following .DS_Store files:")
    fmt.Println(files)

    // re-direct user back to main menu
    Menu("darwin")
}

// menu displays all available tools for use
// it will prompt for user selection and return
// the numerical option that they have chosen
func Menu(OperatingSystem string) {
    var response string // holds user response

    fmt.Println("[*] Enter \"q\" or \"quit\" to exit")

    // mac-specific tools
    if OperatingSystem == "darwin" {

        // print tool options for user
        fmt.Println("1) DS_Store Explorer")
        fmt.Println()
        fmt.Print("Please select a tool: ")

        // retrieve user response
        fmt.Scanln(&response)
        fmt.Println()
        response = strings.ToLower(response)

        // execute appropriate tools based on 
        // user response
        if response == "1" {
            AnalyzeDsStore()
        }
    }

    // exit program gracefully
    if response == "q" || response == "quit" {
        os.Exit(0)
    } else {
        Menu(OperatingSystem)
    }
}

func main() {

    fmt.Println("*-------------------------------------*")
    fmt.Println("|                                     |")
    fmt.Println("|                                     |")
    fmt.Println("|  Welcome to the Attacker's Toolbox  |")
    fmt.Println("|                                     |")
    fmt.Println("|                                     |")
    fmt.Println("*-------------------------------------*")
    fmt.Print("\n\n")

    fmt.Println("[*] Retrieving system information")
    OperatingSystem := runtime.GOOS
    fmt.Println("[*] Current OS: " + OperatingSystem)
    fmt.Println()

    // retrieve user input via standard input
    Menu(OperatingSystem)
}
