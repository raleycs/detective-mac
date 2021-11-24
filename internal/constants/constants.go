package constants

// GetDsStoreSignature returns a byte array with the known signature of .DS_Store files.
// Source: https://en.wikipedia.org/wiki/.DS_Store
func GetDsStoreSignature() [8]byte {
    return [8]byte{0x00, 0x00, 0x00, 0x01, 0x42, 0x75, 0x64, 0x31}
}
