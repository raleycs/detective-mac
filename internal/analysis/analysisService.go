package analysisService

import(
    "encoding/hex"
    "fmt"
    "log"
    "os"
    "strconv"
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

            // read entire file into a temporary buffer
            buffer := make([]byte, fileInfo.Size())
            _, err = file.Read(buffer)
            if err != nil {
                log.Fatal(err)
            }

            // get root block
            startString, err := strconv.ParseInt(hex.EncodeToString(buffer[8:12]), 16, 64)
            if err != nil {
                log.Fatal(err)
            }
            endString, err := strconv.ParseInt(hex.EncodeToString(buffer[16:20]), 16, 64)
            if err != nil {
                log.Fatal(err)
            }
            rootSize, err := strconv.ParseInt(hex.EncodeToString(buffer[12:16]), 16, 64)
            if err != nil {
                log.Fatal(err)
            }
            startRoot, err := hex.DecodeString(fmt.Sprintf("%x", startString + 0x04))
            if err != nil {
                log.Fatal(err)
            }
            endRoot, err := hex.DecodeString(fmt.Sprintf("%x", rootSize + endString + 0x04))
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("0x%x\n", startRoot)
            fmt.Printf("0x%x\n", endRoot)

            return
        }()
    }
}
