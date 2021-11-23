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
            // confirm file signature by reading first 6 bytes
            // source: https://wiki.mozilla.org/DS_Store_File_Format
            file, err := os.Open(filePath)
            if err != nil {
                return err
            }

            defer file.Close()

            // get file size
            // fileInfo, err := file.Stat()
            // if err != nil {
            //     return err
            // }
            // fileSize := fileInfo.Size()
            // buffer := make([]byte, fileSize)

            // read file into memory
            buffer := make([]byte, 10) // read first 10 bytes of the file
            _, err = file.Read(buffer)
            if err != nil {
                return err
            }

            fmt.Printf("File signature:\n0x")
            for _, b := range buffer {
                fmt.Printf("%X", b)
            }
            fmt.Println()

            // print to standard output if the file is .DS_Store
            fmt.Println("[*] " +  filePath)
            dsStores = append(dsStores, filePath)
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
