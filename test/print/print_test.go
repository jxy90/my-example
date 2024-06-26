package main

import (
	"log"
	"testing"
)

type C struct {
	num int
}

func TestPrint(t *testing.T) {

	var (
		string1 = "123"
		int1    = 123
		float1  = 123.1
		bool1   = true
		array1  = [2]int{1, 2}
		map1    = map[string]string{"a": "1"}
		slice1  = []string{"1", "2"}
		struct1 = struct{ a int }{a: 1}
		struct2 = struct{ c C }{c: C{num: 1}}
	)
	output := []any{string1, int1, float1, bool1, array1, map1, slice1, struct1, struct2}
	log.Printf("%s, %d, %f, %t, %v, %v, %v, %v, %v", string1, int1, float1, bool1, array1, map1, slice1, struct1, struct2)
	log.Printf("%v, %v,%v, %v, %v, %v, %v, %v, %v", output...)
	log.Printf("%+v, %+v,%+v, %+v, %+v, %+v, %+v, %+v, %+v", output...)
	log.Printf("%#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v, %#v", output...)
}
