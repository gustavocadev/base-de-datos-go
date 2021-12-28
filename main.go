package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/fatih/color"
)

type User struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func ReadJSON(name string) ([]*User, error) {
	var users []*User

	dataBytes, err := ioutil.ReadFile(name)

	if err != nil {
		return nil, errors.New("hay un error, a lo mejor el archivo no existe :(")
	}

	json.Unmarshal(dataBytes, &users)
	return users, nil
}

func WriteJSON(name string, mySlice []*User) error {

	dataBytes, _ := json.MarshalIndent(mySlice, "", "\t")

	err := ioutil.WriteFile(name, dataBytes, 0744)

	if err != nil {
		return errors.New("hay un error, a lo mejor el archivo no existe :(")
	}
	return nil
}

func Menu() uint8 {
	var option uint8
	fmt.Println("Ingrese una opcion:")
	fmt.Println("(1) Añadir datos")
	fmt.Println("(2) Eliminar datos")
	fmt.Println("(3) Mostrar dato de un usuario ")
	fmt.Println("(4) Listar todos los datos")
	fmt.Println("(5) Actualizar")
	fmt.Println("(6) Terminar")
	fmt.Print("Elige -> ")
	fmt.Scan(&option)

	return option
}

func Delete(option *uint8) {
	var users []*User
	users, _ = ReadJSON("data.json")

	if *option == 1 {
		var name string
		fmt.Print("Ingresa el nombre del usuario: ")
		fmt.Scan(&name)

		for idx, user := range users {
			if user.Name == name {
				users = append(users[:idx], users[idx+1:]...)
			}
		}
	} else if *option == 2 {
		users = []*User{}
	}

	WriteJSON("data.json", users)
}

func ShowInfoUser(users []*User, name string) {
	for _, user := range users {
		if user.Name == name {
			color.Yellow("Mi nombre es: %s", user.Name)
			color.Green("Mi edad es: %d", user.Age)
		}
	}
}

func ShowAllUsers(users []*User) {
	for _, user := range users {
		color.Cyan("Mi nombre es: %s", user.Name)
		color.Green("Mi edad es: %d", user.Age)
	}
}

func main() {
	color.Yellow("Hola, Bienvenido Bienvenida a nuestra base 😊 (Escrita en Go)")
	for {
		option := Menu()
		if option == 1 {
			fmt.Println("Vamos a añadir un dato 😃")
			var name string
			fmt.Print("Ingrese su Nombre: ")
			fmt.Scan(&name)

			var users []*User
			users, _ = ReadJSON("data.json")

			var age uint8
			fmt.Print("ingrese su edad: ")
			fmt.Scan(&age)

			users = append(users, &User{name, age})

			WriteJSON("data.json", users)

		} else if option == 2 {
			var option uint8
			fmt.Println("Qué quieres eliminar?")
			fmt.Println("(1) ¿Quieres eliminar un usuario? ")
			fmt.Println("(2) ¿Quieres eliminar toda los datos?")
			fmt.Print("Elige => ")
			fmt.Scan(&option)

			Delete(&option)
		} else if option == 3 {
			var name string
			fmt.Print("De qué usuario desea mostrar sus datos?: ")
			fmt.Scan(&name)

			users, _ := ReadJSON("data.json")

			ShowInfoUser(users, name)

		} else if option == 4 {

			users, _ := ReadJSON("data.json")
			ShowAllUsers(users)

		} else if option == 5 {
			var name string
			fmt.Print("De qué usuario desea mostrar sus datos?: ")
			fmt.Scan(&name)

			var key string
			fmt.Print("Qué dato desea actualizar?: ")
			fmt.Scan(&key)

			var newValue string
			fmt.Print("Ingrese el nuevo valor: ")
			fmt.Scan(&newValue)

			var users []*User
			users, _ = ReadJSON("data.json")

			for _, user := range users {
				if user.Name == name {
					if key == "name" {

						user.Name = newValue
					}
					if key == "age" {
						fmt.Println("Hola")
						myAge, _ := strconv.Atoi(newValue)
						user.Age = uint8(myAge)
					}
				}
			}
			fmt.Println(users)
			WriteJSON("data.json", users)

		} else if option == 6 {
			break
		}
	}
}
