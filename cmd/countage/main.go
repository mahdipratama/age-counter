package main

import (
	"agecounter"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// take user input: birthdate
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter date (DD/MM/YYYY): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	cleanInput := strings.TrimSpace(input) // remove newline

	var day, month, year int

	n, err := fmt.Sscanf(cleanInput, "%d/%d/%d", &day, &month, &year)
	if err != nil || n != 3 {
		fmt.Println("Error: Invalid format. Please enter numbers as DD/MM/YYYY.")
	}

	dateInput := agecounter.Age{
		Tanggal: day,
		Bulan:   month,
		Tahun:   year,
	}

	now := time.Now()

	dateNow := agecounter.Age{
		Tanggal: now.Day(),
		Bulan:   int(now.Month()),
		Tahun:   now.Year(),
	}

	age, err := agecounter.AgeCounter(dateInput, dateNow)

	fmt.Printf("Your age: %d tahun - %d bulan - %d hari", age.Tahun, age.Bulan, age.Tanggal)

}
