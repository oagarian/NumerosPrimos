package main

import (
	"os"
	"time"
	"log"
	"bufio"
	"fmt"
	"strconv"
)

func create() {
	file_time := time.Now().Format("02-01-06-15-04")
	file_name := fmt.Sprintf("%s.txt", file_time)
	file, err := os.Create(file_name) 

	if err != nil {
		log.Fatal(err)
	}
	
	
	writer := bufio.NewWriter(file)

	for num := 0; num > -1; num++ {
		resultado := 0
		for i := 2; i <= num / 2; i ++ {
			if num % i == 0 {
				resultado++
				break
			}
		}

		if resultado == 0 {
			number := strconv.Itoa(num)
			number_text := fmt.Sprintf("\n %s ", number)
			real_text := []byte(number_text)
			fmt.Fprintf(writer, string(real_text))
			
		} 
	}
	

	
}

func main() {
	create()
}