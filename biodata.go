package main

import (
	"fmt"
	"os"
	"strings"
)

type Biodata interface {
	Biodata(input string, memberList []string) string
}

type Person struct {
	ID        int
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (p Person) Biodata(input string, memberList []string) Person {

	switch strings.ToLower(input) {
	case "budi":
		person := Person{
			ID:        1,
			Name:      "Budi",
			Alamat:    "Jakarta",
			Pekerjaan: "Developer",
			Alasan:    "Mencari pengalaman baru",
		}
		return person
	case "thomas":
		person := Person{
			ID:        2,
			Name:      "Thomas",
			Alamat:    "Jakarta",
			Pekerjaan: "Developer",
			Alasan:    "Mencari pengalaman belajar golang",
		}
		return person
	case "rahmat":
		person := Person{
			ID:        3,
			Name:      "Rahmat",
			Alamat:    "Bandung",
			Pekerjaan: "Mobile Developer",
			Alasan:    "Belajar bahasa pemrograman baru",
		}
		return person
	case "rizal":
		person := Person{
			ID:        4,
			Name:      "Rizal",
			Alamat:    "Surabaya",
			Pekerjaan: "Front Developer",
			Alasan:    "Belajat golang",
		}
		return person
	case "rudi":
		person := Person{
			ID:        5,
			Name:      "Rudi",
			Alamat:    "Jakarta",
			Pekerjaan: "Backend Developer",
			Alasan:    "Belajar bahasa backend yang lain",
		}
		return person
	default:
		return Person{}
	}
}

func main() {

	input := os.Args[1]

	memberList := []string{"Budi", "Thomas", "Rahmat", "Rizal", "Rudi"}

	person := Person{}
	process := person.Biodata(input, memberList)
	if len(process.Name) == 0 {
		fmt.Println("Data tidak ditemukan")
		os.Exit(1)
	}
	result := fmt.Sprintf("\n ID : %d \n nama : %s \n alamat : %s \n pekerjaan : %s \n alasan : %s", process.ID, process.Name, process.Alamat, process.Pekerjaan, process.Alasan)
	fmt.Println(result)
}
