package main

import (
    "fmt"
    "os"
    "runtime"
    "strconv"
    "strings"
    "github.com/raleycs/detective-mac/internal/filesystem"
    "github.com/raleycs/detective-mac/internal/analysis"
)

// dsStore finds all .DS_Store files for a specific
// user and analyzes their contents. This function returns nothing.
func dsStore() {
    var username string // holds target's username

    // get user input
    fmt.Print("Enter a username: ")
    fmt.Scanln(&username)
    fmt.Println()

    // confirm that the user exists on the current system
    path := "/Users/" + username
    fmt.Println("[*] Scanning DS_Stores for " + username)
    if fsService.FileExists(path) == false {
        fmt.Println("[!] User does not exist...")
        Menu() // redirect user back to menu
    }

    // retrieve all .DS_Store files for the given user
    files := fsService.RetrieveFiles(".DS_Store", path)
    fmt.Println("[*] Found " + strconv.Itoa(len(files)) + " .DS_Store files")

    // analyze each .DS_Store file and output reports to reports directory
    fmt.Println("[*] Status: Running analysis on .DS_Store files")
    analysisService.AnalyzeDsStore(files)
    fmt.Println("[*] Status: Analysis Finished")

    // re-direct user back to main menu
    Menu()
}

// Menu displays all available tools for use it will prompt for user selection.
// The appropriate action will be taken based on the user's numerical response.
// This function returns nothing.
func Menu() {
    var response string // holds user response

    fmt.Println("\n\n[*] Enter \"q\" or \"quit\" to exit")

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
        dsStore()
    }

    // exit program gracefully
    if response == "q" || response == "quit" {
        os.Exit(0)
    } else {
        Menu()
    }
}

func main() {

    fmt.Println("*-------------------------------------*")
    fmt.Println("|                                     |")
    fmt.Println("|                                     |")
    fmt.Println("|    Welcome to the Detective Mac     |")
    fmt.Println("|                                     |")
    fmt.Println("|                                     |")
    fmt.Println("*-------------------------------------*")
    fmt.Print("\n\n")

    if runtime.GOOS != "darwin" {
        fmt.Println("This can only be run for Mac Machines!")
        os.Exit(0)
    }

    // retrieve user input via standard input
    Menu()
}
