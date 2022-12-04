package main

import (
	"fmt"
)

var soglia int

func main() {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&soglia)

	if soglia <= 0 {
		fmt.Print("La soglia inserita non è positiva.")
	} else {
		fmt.Print("Numeri abbondanti: ")
		NumeriAbbondanti(soglia)
	}

}

func SommaDivisoriPropri(n int) int {
	var sum int
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum = sum + i
			//fmt.Println(i)    FUNZIONA
		} else if n%1 != 0 {
			return 0
		}
	}
	return sum
}
func ÈAbbondante(n int) bool {
	var status bool
	if SommaDivisoriPropri(n) > n {
		status = true
	} else if SommaDivisoriPropri(n) < n {
		status = false
	}
	return status
}

func NumeriAbbondanti(limite int) {
	for i := 0; i < limite; i++ {
		if ÈAbbondante(i) == true {
			fmt.Print(i, " ")
		}
	}
}
