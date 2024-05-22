package main

import "log"

func main() {
	var stars1 []string
	stars2 := make([]string, 0, 10)
	for i := 0; i < 10; i++ {
		stars1 = append(stars1, "1")
		stars2 = append(stars2, "1")
		log.Printf("stars1 cap:%v len:%v", cap(stars1), len(stars1))
		log.Printf("stars2 cap:%v len:%v", cap(stars2), len(stars2))
	}
}
