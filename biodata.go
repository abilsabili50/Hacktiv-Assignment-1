package main

import (
	. "assignment-1/data"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Arguments can't be Empty")
		return
	}
	nameNoLowerCase := os.Args[1]
	nameLowerCase := strings.ToLower(nameNoLowerCase)
	// fmt.Println(index)

	find(Students, nameLowerCase)
}

func find(data []Student, name string) {
	for _, student := range data {
		if name == strings.ToLower(student.Nama) {
			fmt.Println("Nama :", student.Nama)
			fmt.Println("Alamat :", student.Alamat)
			fmt.Println("Pekerjaan :", student.Pekerjaan)
			fmt.Println("Alasan :", student.Alasan)
			return
		}
	}
}
