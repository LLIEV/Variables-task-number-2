/* Golang BEGINNER.ТЕОРИЯ 
тема:переменные,задача № 2 
программa для перевода футов в метры 
(1 фут = 0.3048 метр).*/

package main

import "fmt"

func main() {
	fmt.Print("Enter length in feet: ")
	var foots float64
	fmt.Scanf("%f", &foots)
	meters := foots * 0.3048
	fmt.Println("Length in meters: ", meters)
}
