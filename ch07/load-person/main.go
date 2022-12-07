package main
import (
	"fmt"
	"io"
	"strings"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}
// Accept interface return struct
func loadPerson(r io.Reader) (Person, error) {
	var p Person
	// store value decoded in p
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func main() {
	s := `{"Name":"Joe", "Age":18}`
	p, err := loadPerson(strings.NewReader(s))
	fmt.Println(p)
	fmt.Println(err)
}



