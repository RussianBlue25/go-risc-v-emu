package bin

import (
    "os"
    "fmt"
    "io/ioutil"
)

func BinLoad(Memory [65536]uint8) [65536]uint8 {
    file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("can't open file")
		panic(err)
    }

    binary, errb := ioutil.ReadAll(file)
	if errb != nil {
		fmt.Println("can't read binary")
		panic(errb)
    }

    copy(Memory[:], []uint8(binary)[:])

    return Memory
}