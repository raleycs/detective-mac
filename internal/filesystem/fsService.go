package fsService

import (
    "path/filepath"
    "fmt"
    "io/fs"
    "log"
    "os"
    "github.com/raleycs/detective-mac/internal/constants"
)

// FileExists returns a bool determining if a file exists.
// If the file exists then true is returned, otherwise false.
func FileExists(path string) bool {
    if _, err := os.Stat(path); os.IsNotExist(err) { return false }
    return true
}

// RetrieveFiles returns a slice of strings that exist under the string path.
// Each element in the slice of strings carries the full path with the file.
func RetrieveFiles(file string, path string) []string {
    var dsStores []string

    // function called for every .DS_Store file found
    var verifyFile = func(filePath string, dir fs.DirEntry, err error) error {

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

            // read file into memory
            var signature [8]byte // array of size 8
            buffer := make([]byte, 8) // read first 8 bytes of the file into a temporary buffer slice
            _, err = file.Read(buffer)
            if err != nil {
                return err
            }
            copy(signature[:], buffer) // copy contents of buffer slice into array

            // compare file signature
            if signature != constants.GetDsStoreSignature() {
                fmt.Printf("[*] %s does not match signature!\n", filePath)
            } else {
                dsStores = append(dsStores, filePath) // add file to confirmed .DS_Store slice
            }
        }
        return nil
    }

    // retrieve all .DS_Store files under given the given user's home directory
    err := filepath.WalkDir(path, verifyFile)

    // log errors from filepath
    if err != nil {
        log.Fatal(err)
    }

    return dsStores
}
