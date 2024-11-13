package main

import (
	"fmt"
	"os"
)

func main() {

	/*
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("Error while creating file: ",err)
			return
		}
		defer file.Close()

		content := "Dwaipayan Biswas"
		byte, errors := io.WriteString(file,content+"\n")
		fmt.Println("Byte return is: ",byte)
		if errors != nil {
			fmt.Println("Error while writing file: ",errors)
			return
		}

		fmt.Println("Successfully created file")
	*/

	//Buffer Reader
	// file,err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while opening file: ",err)
	// 	return
	// }
	// defer file.Close()

	// //creating buffer
	// buffer := make([]byte, 1024)

	// for{
	// 	n,err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}

	// 	if err != nil {
	// 		fmt.Println("Error while reading file: ",err)
	// 		return
	// 	}

	// 	fmt.Println(string(buffer[:n]))
	// }

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	fmt.Println(string(content))

}
