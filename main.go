package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func readfile() {
	// f, err := os.Open("./bitcask/bitcask-2025-01-28T16:26:43.txt")
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// }

	// stat, err := f.Stat()
	// if err != nil {
	// 	fmt.Printf("err: %vcreateASingleFile\n", err)
	// }

	// fmt.Printf("stat.Size(): %v\n", stat.Size())
}

func main() {

	opts := Opts{
		Perm: os.ModePerm,
	}

	bitcask, err := Open("bitcask", opts)
	if err != nil {
		log.Fatal("err", err)
	}

	// for i := 0; i < math.MaxInt16; i++ {
	// 	_, err := file.WriteString(fmt.Sprintf("this is the file line count %d\n", i))
	// 	if err != nil {
	// 		log.Fatalf("Not able to write to the file %v\n", err)
	// 	}
	// 	// _, err := bitcask.FileName.WriteString(fmt.Sprintf("this is the file line count %d\n", i))
	// 	// if err != nil {
	// 	// 	log.Fatalf("Not able to write to the file %v\n", err)
	// 	// }
	// }
	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:
			file := <-bitcask.FileName

			fmt.Printf("file.Name() main: %v\n", file.Name())

			for i := 0; i < math.MaxInt16; i++ {
				_, err := file.WriteString(fmt.Sprintf("this is the file line count %d\n", i))
				if err != nil {
					log.Fatalf("Not able to write to the file %v\n", err)
				}
			}
		}
	}

	time.Sleep(time.Minute * 5)
}
