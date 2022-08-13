package main

import (
	pk "014_access_modifiers/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	fmt.Println(myCar)

	pk.PrintMessage("Hello world!")

	// pk.printMessage("Hello world private!")
}
