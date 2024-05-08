package main

import "fmt"

func main(){
	runApplication(10)
	runApp(false)
	runApps(false)
}

//defer
func logging(){
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("Run application ")
	result := 10 / value
	fmt.Println("Result :", result)
}
//

//panic
func endApp(){
	fmt.Println("Aplikasi selesai!!")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("Aplikasi ERROR!!")
	}
	fmt.Println("Aplikasi berjalan")
}
//

//recover

func endApps(){
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message :", message)
	}

	fmt.Println("Aplikasi selesai!!")
}

func runApps(error bool){
	defer endApps()
	if error {
		panic("Erorr!!")
	}
	fmt.Println("Aplikasi berhasil berjalan!")
}

 

