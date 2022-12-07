package main

import "fmt"

type Record struct {
	key       string
	valueType string
	data      T
}

type Person struct {
	lastName  string
	age       int
	isMarried bool
}

type Animal struct {
	name     string
	caterogy string
}
type T interface{}

func newRecord(key string, i T) Record {
	r := Record{}
	r.key = key
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case Person:
		r.valueType = "Person"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}

func main() {
	m := make(map[string]T)
	a := Animal{name: "oreo", caterogy: "cat"}
	p := Person{lastName: "Anabell", isMarried: false, age: 19}
	m["person"] = p
	m["animal"] = a
	m["age"] = 34
	m["isMarried"] = true
	m["lastName"] = "Smith"
	listRecord := []Record{}
	for k, v := range m {
		r := newRecord(k, v)
		listRecord = append(listRecord, r)
	}
	for _, v := range listRecord {
		fmt.Printf("Key: %v, Type: %v, Data: %v\n", v.key, v.valueType, v.data)
	}
}
