package elf

import (
    "os"
    "encoding/binary"
    "fmt"
)

type elf32Header struct {
    E_ident     [16]byte
    E_type      uint16
    E_machine   uint16
    E_version   uint32
    E_entry     uint32
    E_phoff     uint32
    E_shoff     uint32
    E_flags     uint32
    E_ehsize    uint16
    E_phentsize uint16
    E_phnum     uint16
    E_shentsize uint16
    E_shnum     uint16
    E_shstrndx  uint16
}

func ElfLoad() {
    file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("can't open file")
		panic(err)
    }

    var elf32Header elf32Header

    errb := binary.Read(file, binary.LittleEndian, &elf32Header)
    if errb != nil {
        fmt.Println("can't read binary")
    }
    fmt.Printf("%x\n", elf32Header)

    if elf32Header.E_ident[0] == 0x7f && elf32Header.E_ident[1] == 0x45 && elf32Header.E_ident[2] == 0x4c && elf32Header.E_ident[3] == 0x46 {
        fmt.Println("this is an ELF file")
    } else {
        fmt.Println("this is not an ELF file!!")
        os.Exit(1)
    }

    if elf32Header.E_ident[4] == 0x01 {
        fmt.Println("32bit")
    } else {
        fmt.Println("not 32bit")
        os.Exit(1)
    }
}