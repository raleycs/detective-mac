package fsService

import (
    "path/filepath"
    "fmt"
    "io/fs"
    "log"
    "reflect"
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
// The files are checked with their own file signatures before being added to the slice.
func RetrieveFiles(file string, path string) []string {
    var verified []string

    // verifyFile returns an error if file verification did not succeed. For each file
    // found under the dirwalk, it is compared with a known file signature to verify
    // that the file is what it claims to be.
    var verifyFile = func(filePath string, dir fs.DirEntry, err error) error {

        // handle errors from original dirwalk
        if err != nil {
            log.Fatal(err)
        }

        // open file
        f, _ := os.Open(filePath)

        defer f.Close() // close file after completion of verifyFile

        // verify file signatures
        if dir.Name() == ".DS_Store" {

            // read first 24 bytes of the file into a temporary buffer
            buffer := make([]byte, 20)
            _, err = f.Read(buffer)
            if err != nil {
                return err
            }

            // validate file signature and ensure equality of root offsets
            if reflect.DeepEqual(buffer[0:8], constants.GetDsStoreSignature()) &&
            reflect.DeepEqual(buffer[8:12], buffer[16:20]) {
                verified = append(verified, filePath) // add file to confirmed .DS_Store slice
            } else {
                fmt.Printf("[*] %s does not match signature!\n", filePath)
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

    return verified
}
