package main

import (
	"fmt"
	"test/dog"
	"test/printer"
	"test/user"
)

var (
	urepo = &user.UserRepository{}
	drepo = &dog.DogRepository{}
)

func main() {
	// create users
	u := user.New("luan")
	urepo.CreateUser(u)

	u2 := user.New("joao")
	urepo.CreateUser(u2)

	// create dogs
	d := dog.New(u.Id, "pudim")
	drepo.CreateDog(d)

	d2 := dog.New(u.Id, "nala")
	drepo.CreateDog(d2)

	d3 := dog.New(u2.Id, "rabito")
	drepo.CreateDog(d3)

	dsvc := dog.NewService(drepo)
	u.FeedMyDogs(dsvc)

	fmt.Println("--- USER DOGS ---")
	for _, each := range u.GetUserDogsNames(dsvc) {
		fmt.Println(each)
	}

	fmt.Println("--- ALL DOGS ---")
	for _, each := range drepo.ListDogs() {
		msg := "Fed"
		if each.IsHungry() {
			msg = "Hungry"
		}

		printer.Print(each, msg)
	}
}
