package main

import (
	"fmt"
	"os"
)

type Person struct {
	ID        int
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func (p Person) Biodata() string {
	return fmt.Sprintf("\n ID : %d \n nama : %s \n alamat : %s \n pekerjaan : %s \n alasan : %s", p.ID, p.Name, p.Alamat, p.Pekerjaan, p.Alasan)
}

func main() {

	input := os.Args[1]

	memberList := []string{"Budi", "Thomas", "Rahmat", "Rizal", "Rudi"}

	for i, member := range memberList {
		if input == member {
			person := Person{
				ID:        i,
				Name:      input,
				Alamat:    "Jakarta",
				Pekerjaan: "Developer",
				Alasan:    "Mencari pengalaman baru",
			}
			fmt.Println(person.Biodata())
		}
	}

}
