package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   int
	Name string
}

var items []Item
var idCounter int

func main() {
	r := gin.Default()

	for {
		showMenu()
		option := readOption()

		switch option {
		case 1:
			createItem()
		case 2:
			readItems()
		case 3:
			updateItem()
		case 4:
			deleteItem()
		case 5:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}

	r.Run(":8080")
}

func createItem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nombre: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	idCounter++
	newItem := Item{ID: idCounter, Name: name}
	items = append(items, newItem)
	fmt.Println("Ítem creado con ID:", idCounter)
}

func readItems() {
	if len(items) == 0 {
		fmt.Println("No hay ítems registrados.")
		return
	}

	fmt.Println("--- Ítems ---")
	for _, item := range items {
		fmt.Printf("ID: %d, Nombre: %s\n", item.ID, item.Name)
	}
}

func updateItem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID del ítem a actualizar: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	id, _ := strconv.Atoi(input)

	for i, item := range items {
		if item.ID == id {
			fmt.Print("Nuevo nombre: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			items[i].Name = name
			fmt.Println("Ítem actualizado.")
			return
		}
	}

	fmt.Println("Ítem no encontrado.")
}

func deleteItem() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("ID del ítem a eliminar: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	id, _ := strconv.Atoi(input)

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			fmt.Println("Ítem eliminado.")
			return
		}
	}

	fmt.Println("Ítem no encontrado.")
}

func showMenu() {
	fmt.Println("\n--- CRUD en Consola con Gin ---")
	fmt.Println("1. Crear")
	fmt.Println("2. Leer")
	fmt.Println("3. Actualizar")
	fmt.Println("4. Eliminar")
	fmt.Println("5. Salir")
	fmt.Print("Opción: ")
}

func readOption() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	option, _ := strconv.Atoi(input)
	return option
}
