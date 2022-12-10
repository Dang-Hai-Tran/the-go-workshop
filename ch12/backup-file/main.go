package main
import (
	"fmt"
	"io"
	"errors"
	"os"
)

var (
	ErrWorkingFileNotFound = errors.New("the working file is not found")
)

func createBackup(working, backup string) error {
	// check the working file is exist before backing it up
	_, err := os.Stat(working)
	if err != nil {
		// return ErrWorkingFileNotFound error if file is not exist
		if os.IsNotExist(err) {
			return ErrWorkingFileNotFound
		}
		// Return other err in other case
		return err
	}
	workFile, err := os.Open(working)
	if err != nil {
		return err
	}
	defer workFile.Close()
	content, err := io.ReadAll(workFile)
	if err != nil {
		return err
	}
	err = os.WriteFile(backup, content, 0744)
	if err != nil {
		return err
	}
	return nil
}

func addNotes(workingFile, notes string) error {
	// This function will append notes to the workingFile
	f, err := os.OpenFile(workingFile, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0744)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err =f.WriteString(notes)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	backupFile := "backup.txt"
	workingFile := "working.txt"
	err := createBackup(workingFile, backupFile)
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 10; i++ {
		note := fmt.Sprintf("Note %v\n", i)
		err := addNotes(workingFile, note)
		if err != nil {
			panic(err)
		}
	}
}

