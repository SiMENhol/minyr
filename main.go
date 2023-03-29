package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input = scanner.Text()
		if input == "q" || input == "exit" {
			fmt.Println("exit")
			os.Exit(0)
		} else if input == "convert" {
			fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")
			// funksjon som åpner fil, leser linjer, gjør endringer og lagrer nye linjer i en ny fil

			// flere else-if setninger
		} else {
			fmt.Println("Venligst velg convert, average eller exit:")

		}

	}
	src, err := os.Open("table.csv")
	//src, err := os.Open("/home/janisg/minyr/kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
        	log.Fatal(err)
	}
	defer src.Close()
        log.Println(src)
        
	
	var buffer []byte
	var linebuf []byte // nil
	buffer = make([]byte, 1)
        bytesCount := 0
	for {
		_, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesCount++
		//log.Printf("%c ", buffer[:n])
		if buffer[0] == 0x0A {
	           log.Println(string(linebuf))
		   // Her
		   elementArray := strings.Split(string(linebuf), ";")
		   if len(elementArray) > 3 {
			 celsius := elementArray[3]
			 fahr := conv.CelsiusToFahrenheit(celsius)
		         log.Println(elementArray[3])
	   	   }
                   linebuf = nil		   
		} else {
                   linebuf = append(linebuf, buffer[0])
		}	
		//log.Println(string(linebuf))
		if err == io.EOF {
			break
		}
	}

}

	//if err != nil {
	//		log.Fatal(err)
	//	}
}
