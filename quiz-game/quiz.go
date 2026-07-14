package main

import (
	"fmt"
)

func main() {
	fmt.Println("__________________________")
	fmt.Println("VINCE Maths Quiz Game📚 ")
	fmt.Println("__________________________")

	var name string
	var lname string
	fmt.Print("EnterName: ")
	fmt.Scanln(&name, &lname)

	fmt.Printf("Hello %v %s welcome to vince maths quiz Game \n", name, lname)

	var age uint
	fmt.Print("Enter Age:")
	fmt.Scan(&age)

	if age < 20 && age > 3 {
		fmt.Printf("yay🤩, congratulation %v %v! \n", name, lname)
		fmt.Println()
		fmt.Println("Lets crack the brain🤔💭")
	} else {
		fmt.Println("oops, you can't play the game")
		return
	}
	fmt.Println("👨‍💻Questions:(choose the correct answer: Dont cheat👀)")
	fmt.Println()
	fmt.Print("1.what is the square root of 100?")
	fmt.Println("(a)1000", "(b)10", "(c)50")
	var answer string
	fmt.Print("Answer:")
	fmt.Scan(&answer)
	scores := 0
	numberq := 5

	if answer == "b" || answer == "B" {
		fmt.Println(answer, "✅")
		scores++
	} else {
		fmt.Println(answer, "❌")
	}
	fmt.Print("2.what is the square of 10?")
	fmt.Println("(a)1", "(b)1000", "(c)100")
	var answer2 string
	fmt.Print("Answer:")
	fmt.Scan(&answer2)
	if answer2 == "c" || answer2 == "C" {
		fmt.Println(answer, "✅")
		scores++
	} else {
		fmt.Println(answer, "❌")
	}
	fmt.Print("3.what is the factorial of 4?")
	fmt.Println("(a)24", "(b)12", "(c)16")
	var answer3 string
	fmt.Print("Answer:")
	fmt.Scan(&answer3)
	if answer3 == "c" || answer3 == "C" {
		fmt.Println(answer3, "✅")
		scores++
	} else {
		fmt.Println(answer3, "❌")
	}
	fmt.Print("4.Convert 1111 to base 10?")
	fmt.Println("(a)15", "(b)40", "(c)10")
	var answer4 string
	fmt.Print("Answer:")
	fmt.Scan(&answer4)
	if answer4 == "a" || answer4 == "A" {
		fmt.Println(answer4, "✅")
		scores++
	} else {
		fmt.Println(answer4, "❌")
	}
	fmt.Print("5.Conver 23 to base 2?")
	fmt.Println("(a)46", "(b)10111", "(c)1011")
	var answer5 string
	fmt.Print("Answer:")
	fmt.Scan(&answer2)
	if answer5 == "b" || answer5 == "B" {
		fmt.Println(answer5, "✅")
		scores++
	} else {
		fmt.Println(answer5, "❌")
	}
	fmt.Println("_______________________")
	fmt.Printf("scores:= %d/5\n", scores)
	fmt.Printf("percentage:= %2.f%% \n", (float64(scores)/float64(numberq))*100)

	if scores > 3 {
		fmt.Println("💥weldone", name, lname)
	} else if scores >= 1 && scores <= 3 {
		fmt.Println("😴work hard ", name, lname)
	} else {
		fmt.Println("😡you are playing too much", name, lname)
	}
	fmt.Println("_______________________")
}
