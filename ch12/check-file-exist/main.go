package main
import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Stat("test.txt")
	// os.Stat may return multiples errors
	if err != nil {
		fmt.Println(err)
		// Check error if file isn't exist
		if os.IsNotExist(err) {
			fmt.Println("test.txt file doesn't exist.")
			fmt.Println(f)
		}
	}
	// if file is exist print file infos
	fileName := f.Name()
	fileModifTime := f.ModTime()
	fileSize := f.Size()
	fmt.Printf("filename : %v, modtime : %v, filesize(bytes) : %v\n", fileName, fileModifTime, fileSize)
}