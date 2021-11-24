package constants

// return byte array with signature of .DS_Store files
func GetDsStoreSignature() [8]byte {
    return [8]byte{0x00, 0x00, 0x00, 0x01, 0x42, 0x75, 0x64, 0x31} // source: https://en.wikipedia.org/wiki/.DS_Store
}
