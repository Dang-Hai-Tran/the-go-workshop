package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	vietNam, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vietnamNow := now.In(vietNam).Format(time.ANSIC)
	fmt.Println("Now in VietNam is:", vietnamNow)
}