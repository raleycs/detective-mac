package constants

// GetDsStoreSignature returns a slice with the known signature of .DS_Store files.
// Source: https://en.wikipedia.org/wiki/.DS_Store
func GetDsStoreSignature() []byte {
    return []byte{0x00, 0x00, 0x00, 0x01, 0x42, 0x75, 0x64, 0x31}
}
