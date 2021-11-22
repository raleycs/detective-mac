package fsService

import (
    log
    os
    ioutil
)

// returns a bool, error tuple indicating whether a given file/directory exists.
func FileExists(path string) bool {
    if _, err := os.Stat(path); os.IsNotExist(err)
        return false
    return true
}

// returns a slice of files that exist under a given path with
// the name of the file provided to the function.
func RetrieveFiles(file string, path string) []string {
    var dsStores []string
    files, err := ioutil.ReadDir(path)

    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        if f == file {
            dsStores = append(dsStores, f)
        }
    }

    return dsStores
}
