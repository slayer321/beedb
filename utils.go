package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Create(folderName string, filePermission os.FileMode) error {

	if _, err := os.Stat(folderName); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(folderName, filePermission)
		if err != nil {
			log.Fatalf("Failed to create the directory %v\n", err)
			return err
		}
	}

	return nil

}

func createASingleFile(DirectoryName string) (*os.File, error) {
	t := time.Now()

	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fileName := fmt.Sprintf("%s-%s.txt", DirectoryName, formatted)

	file, err := os.Create(filepath.Join(DirectoryName, fileName))
	if err != nil {
		return nil, err
	}

	return file, nil
}

func CreateMultipleFiles(newFile chan *os.File, DirectoryName string) {
	var file *os.File
	time := time.NewTicker(time.Second)

	file, err := createASingleFile(DirectoryName)
	if err != nil {
		log.Fatalf("Failed to create single file for first instance %v\n", err)
	}
	fmt.Printf("newFile channel before the first send %v\n", newFile)
	newFile <- file

	for newFile != nil {
		select {
		case <-time.C:
			// fmt.Printf("file.Name(): CreateMultipleFiles %v\n", file.Name())
			stat, err := file.Stat()
			if err != nil {
				log.Fatal("stat failed")
			}
			// fmt.Printf("stat.Size(): %v\n", stat.Size())

			// if stat.Size() > 10000 {
			// 	fmt.Println("Is greater")
			// } else if stat.Size() < 10000 {
			// 	fmt.Println("Is less then")
			// }Size
			// stat, err := file.Stat(); err == nil &&
			if stat.Size() > 10000 {
				file.Close()
				// fmt.Println("Inside the size greater then")
				file, err := createASingleFile(DirectoryName)
				if err != nil {
					log.Fatalf("Failed to create another file once the 10KB size was execeed %v\n", err)
				}
				// fmt.Printf("New file created before sending it to channel %v\n", file.Name())
				// fmt.Printf("newFile channel after the new file creation %v\n", newFile)

				newFile <- file
				// fmt.Println("After sending to new channel")

			}

			// case <-time.C:
			// 	if newFile == nil {
			// 		file, err := createASingleFile(DirectoryName)
			// 		if err != nil {
			// 			log.Fatalf("Failed to create single file %v\n", err)
			// 		}
			// 		newFile <- file
			// 	} else if stat, err := file.Stat(); err == nil && stat.Size() > 10000 {
			// 		file, err := createASingleFile(DirectoryName)
			// 		if err != nil {
			// 			log.Fatalf("Failed to create single file %v\n", err)
			// 		}
			// 		newFile <- file

			// 	}

		}
	}

}
