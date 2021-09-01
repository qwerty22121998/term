package term

import "fmt"

func Move(x, y int) {
	fmt.Printf("\033[%v;%vH", x, y)
}

func Up(n int) {
	fmt.Printf("\033[%vA", n)
}

func Down(n int) {
	fmt.Printf("\033[%vB", n)
}

func Right(n int) {
	fmt.Printf("\033[%vC", n)
}

func Left(n int) {
	fmt.Printf("\033[%vD", n)
}

func UpBegin(n int) {
	fmt.Printf("\033[%vE", n)
}

func DownBegin(n int) {
	fmt.Printf("\033[%vF", n)
}

func Cls() {
	fmt.Print("\033[2J")
}

func ClsLine() {
	fmt.Print("\033[2K")
}
