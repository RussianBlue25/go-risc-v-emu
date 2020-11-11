package elf

import (
    "os"
    "encoding/binary"
    "fmt"
    "io/ioutil"
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

type elf32Pheader struct {
    P_type      uint32
    P_offset    uint32
    P_vaddr     uint32
    P_paddr     uint32
    P_filesz    uint32
    P_memsz     uint32
    P_flags     uint32
    P_align     uint32
}

func ElfLoad(Memory [65536]uint8) [65536]uint8 {

    file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("can't open file")
		panic(err)
    }

    file2, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("can't open file")
		panic(err)
    }

    binary2, errb := ioutil.ReadAll(file2)
	if errb != nil {
		fmt.Println("can't read binary")
		panic(errb)
	}

    var elf32Header elf32Header
    var elf32Pheaders [128]elf32Pheader //better way?

    for i := 0;; i++ {
        if i == 0 {
            errb := binary.Read(file, binary.LittleEndian, &elf32Header)
            if errb != nil {
                fmt.Println("can't read elf header")
            }
            fmt.Printf("%x\n", elf32Header)
        } else if i <= int(elf32Header.E_phnum) {
            errpb := binary.Read(file, binary.LittleEndian, &elf32Pheaders[i-1])
            if errpb != nil {
                fmt.Println("can't read program header")
            }
            fmt.Printf("%x\n", elf32Pheaders[i-1])
        } else {
            break
        }
    }

    checkElfFormat(elf32Header, elf32Pheaders)

    //write to memory
    for i := 0; i < int(elf32Header.E_phnum); i++ {
        start := elf32Pheaders[i].P_vaddr - elf32Header.E_entry
        end := start + elf32Pheaders[i].P_memsz
        copy(Memory[start:end], []uint8(binary2)[elf32Pheaders[i].P_offset:elf32Pheaders[i].P_offset+elf32Pheaders[i].P_filesz])
    }

    return Memory
}

func checkElfFormat(elf32Header elf32Header, elf32Pheaders [128]elf32Pheader) {
    //TODO: add error check
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