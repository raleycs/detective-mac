package analysisService

import(
    "encoding/hex"
    "fmt"
    "io"
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

            // read header into a temporary buffer
            buffer := make([]byte, 20)
            _, err = file.Read(buffer)
            if err != nil {
                log.Fatal(err)
            }

            // get root block boundaries
            startString, err := strconv.ParseInt(hex.EncodeToString(buffer[8:12]), 16, 64)
            if err != nil {
                log.Fatal(err)
            }
            rootSize, err := strconv.ParseInt(hex.EncodeToString(buffer[12:16]), 16, 64)
            if err != nil {
                log.Fatal(err)
            }

            fmt.Println("[*] Extracting root block")

            // extract root block
            _, err = file.Seek(int64(startString), io.SeekStart)
            if err != nil {
                log.Fatal(err)
            }
            root := make([]byte, rootSize)
            n, err := file.Read(root[:cap(root)])
            if err != nil && err != io.EOF {
                log.Fatal(err)
            }
            root = root[:n]

            fmt.Println("[*] Reading root block")

            // get number of offsets
            numOffsets, err := strconv.ParseInt(hex.EncodeToString(root[4:8]), 16, 64)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("[*] Offsets found: %d\n", numOffsets)

            // extract offsets
            offsets := []int64{}
            for i := 0; i < int(numOffsets); i++ {
                address, err := strconv.ParseInt(hex.EncodeToString(root[(i * 4) + 12:(i * 4) + 16]), 16, 64)
                if err != nil {
                    log.Fatal(err)
                }
                offsets = append(offsets, address)
                fmt.Printf("[*] Offset %d: 0x%x\n", i, offsets[len(offsets)-1])
            }

            return
        }()
    }
}
