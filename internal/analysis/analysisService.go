package analysisService

import(
    "fmt"
    "log"
    "os"
)

// AnalyzeDsStore takes in a slice of .DS_Store files with their full paths.
// Each file will be parsed individually. A singular report is created that
// will contain information about all of the .DS_Store files analyzed.
func AnalyzeDsStore(files []string) {

    // parse each .DS_Store individually
    for _, path := range(files) {

        // call anonymous function to close out files at the appropriate time
        func() {

            // open file
            file, err := os.Open(path)
            if err != nil {
                return
            }
            defer file.Close() // close file after completion of anonymous function

            // get file size
            fileInfo, err := file.Stat()
            if err != nil {
                log.Fatal(err)
            }

            // read file into a temporary buffer
            buffer := make([]byte, fileInfo.Size())
            _, err = file.Read(buffer)
            if err != nil {
                log.Fatal(err)
            }

            fmt.Printf("0x%X\n", buffer[0:26])

            return
        }()
    }
}
