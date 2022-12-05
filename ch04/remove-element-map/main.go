package main

import "fmt"

func main()  {
	user := map[string]string{
		"305" : "Sue",
		"204" : "Bob",
		"631" : "Jake",
		"234" : "Tracy",
	}
	delete(user, "631")
	fmt.Println(user)
}