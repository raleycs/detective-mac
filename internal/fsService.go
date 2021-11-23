package fsService

import (
    "path/filepath"
    "fmt"
    "io/fs"
    "log"
    "os"
)

// returns a bool, error tuple indicating whether a given file/directory exists.
func FileExists(path string) bool {
    if _, err := os.Stat(path); os.IsNotExist(err) { return false }
    return true
}

// returns a slice of files that exist under a given path with
// the name of the file provided to the function.
func RetrieveFiles(file string, path string) []string {
    var dsStores []string

    // function called for every .DS_Store file found
    var AnalyzeFile = func(filePath string, dir fs.DirEntry, err error) error {

        // handle errors from original dirwalk
        if err != nil {
            log.Fatal(err)
        }

        // if the file is .DS_Store add to dsStores slice
        if dir.Name() == ".DS_Store"{
            fullPath := filePath + dir.Name()

            // print to standard output if the file is .DS_Store
            fmt.Println("[*] " +  fullPath)
            dsStores = append(dsStores, fullPath)
        }

        return nil
    }

    // retrieve all .DS_Store files under given the given user's home directory
    fmt.Println("[*] Retrieved the following .DS_Store files:")
    err := filepath.WalkDir(path, AnalyzeFile)

    // log errors from filepath
    if err != nil {
        log.Fatal(err)
    }

    return dsStores
}
