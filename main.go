package main

import (
	"fmt"
	"main.go/electronics"
)

func printCharacteristics(phone electronics.Phone) {
	fmt.Printf("Brand: %v\n", phone.Brand())
	fmt.Printf("Model: %v\n", phone.Model())
	fmt.Printf("Type: %v\n", phone.Type())
	if sp, ok := phone.(electronics.Smartphone); ok {
		fmt.Printf("Operating System: %v\n", sp.OS())
	}
	if sp, ok := phone.(electronics.StationPhone); ok {
		fmt.Printf("Buttons Count: %v\n", sp.ButtonsCount())
	}
}
func main() {
	phones := [3]electronics.Phone{}

	phones[0] = electronics.NewApplePhone("iPhone 13")
	phones[1] = electronics.NewAndroidPhone("Xiaomi", "Poco 5F Pro")
	phones[2] = electronics.NewRadioPhone("Panasonic", "XYZ-123", 20)

	for i := 0; i < len(phones); i++ {
		fmt.Println("\tPhone ID: ", i+1)
		printCharacteristics(phones[i])
	}
}
