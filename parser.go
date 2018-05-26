package main

import (
	"fmt"
	"log"
	"os"
	//"unsafe"
	//"bytes"
	"encoding/binary"
)

type Header struct {
	Text [28]byte
	FirstDataOffset uint32
	SizeOfCompressedFile uint32
	HeaderVersion uint32
	SizeOfDecompressedData uint32
	NumberOfCompressedDataBlocks uint32
}

type SubHeaderV0 struct {
	_ [2]byte
	Version [2]byte
	BuildNumber [2]byte
	Flags [2]byte
	LengthInMillisec [4]byte
	CRC32 [4]byte
}

type SubHeaderV1 struct {
	VersionId [4]byte
	VersionNumber [4]byte
	BuildNumber [2]byte
	Flags [2]byte
	LengthInMillisec [4]byte
	CRC32 [4]byte
}

func main() {
	path := "/tmp/w3replays/replay.php?m=DownloadReplay&rid=9"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	defer file.Close()

	header := Header{}
	//data := readBytes(file, int(unsafe.Sizeof(header)))
	
	// create an io.Reader object
	//buffer := bytes.NewBuffer(data)
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Header:\n%s\n", header.Text)

	version := header.HeaderVersion

    var subHeader SubHeaderV1

    // HOW?
	if int(version) == 0 {
		//subHeader = SubHeaderV0{}
	} else {
		subHeader = SubHeaderV1{}
	}

    err = binary.Read(file, binary.LittleEndian, &subHeader)
    if err != nil {
        log.Fatal(err)
    }

    // does not work:
    // fmt.Print(string(subHeader.VersionId)) // why?
    fmt.Print(string(subHeader.VersionId[:]))
}

func readBytes(file *os.File, n int) []byte {
	bytes := make([]byte, n)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
