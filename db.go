package main

import (
	"hash/crc32"
	"os"
)

type Header struct {
	crc     crc32.Table
	tstamp  uint32
	ksz     uint32
	valueSz uint32
}

type Record struct {
	Header Header
	Key    string
	Value  []byte
}

type Opts struct {
	Perm os.FileMode
}

type BitCaskHandle struct {
	FileName chan *os.File
}

func Open(DirectoryName string, opts ...Opts) (BitCaskHandle, error) {

	if opts == nil {
		opts[0].Perm = os.ModePerm
	}

	filePermission := opts[0].Perm

	err := Create(DirectoryName, filePermission)
	if err != nil {
		return BitCaskHandle{}, err
	}

	// var file chan *os.File
	file := make(chan *os.File)
	// defer file.Close()

	go CreateMultipleFiles(file, DirectoryName)

	// var newFile *os.File
	//newFile := <-file
	// stat, err := file.Stat()
	// if err != nil {
	// 	return BitCaskHandle{}, err
	// }

	// fmt.Printf("stat.Size(): %v\n", stat.Size())

	return BitCaskHandle{
		FileName: file,
	}, err

}
