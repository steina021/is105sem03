package main

import (
	"fmt"

	"github.com/steina021/is105sem03/mycrypt"
)

func main() {

	kryptertRune := mycrypt.Krypter([]rune("Kjevik;SN39040;18.03.2022 01:50;6"), mycrypt.ALF_SEM03, 4)
	kryptertString := string(kryptertRune)
	fmt.Println(kryptertString)

	alfLength := len(mycrypt.ALF_SEM03)

	dekryptertRune := mycrypt.Krypter(kryptertRune, mycrypt.ALF_SEM03, alfLength-4)
	dekryptertString := string(dekryptertRune)
	fmt.Println(dekryptertString)

}
