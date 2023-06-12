package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Tank struct {
	ID   string
	wine string
}

func main() {
	fmt.Println("WARNING : IF A QUANTITY GIVES DECIMAL NUMBERS TO WINE QUANTITIES THEN\nTHE NUMBER WILL BE ROUNDED")
	var formula map[string]int = make(map[string]int)
	var usedTanks map[int][]Tank = make(map[int][]Tank)
	var emptyTanks map[int][]Tank = make(map[int][]Tank)

	if len(os.Args) < 2 {
		panic("please provide path to file")
	}
	path := os.Args[1]

	fmt.Println("\n---- HI, Please give the total Hl you want as a result ----")
	var tempTTL string
	// Taking input from user
	fmt.Scanln(&tempTTL)

	total, err := strconv.Atoi(tempTTL)
	if err != nil {
		println(err.Error())
	}

	fmt.Println("\n Please enter the wine and its quantity (%) separated by a space (ex : Chardonnay 80)")
	fmt.Println("If you write the same wine name twice, the first will be overriden")
	fmt.Println("if quantity is not equal to 100% after you are done, the program will exit")
	fmt.Println("WRITE the word 'done' when you are finished")
	println()

	buffQuantity := 0
	inputF := bufio.NewReader(os.Stdin)
	//will exit only on "done" written
	for true {
		inputF, err := inputF.ReadString('\n')
		if err != nil {
			panic("error reading terminal, verify your authorisations")
		}

		inputF = strings.TrimSpace(inputF)
		if inputF == "done" {
			break
		} else {
			fullInput := strings.Split(inputF, " ")
			if len(fullInput) < 2 {
				fmt.Println("there is not enough arguments, 2 are needed !")
			} else {
				convVal, err := strconv.Atoi(fullInput[1])
				if err != nil {
					fmt.Println("your second value is not a number, please try again.")
				} else {
					quantity := total * convVal / 100
					fmt.Println("Quantity: ", quantity, "/", total)

					formula[fullInput[0]] = quantity
					buffQuantity += convVal
				}
			}
		}
		fmt.Println("\n---- type your next value ----")
	}
	if buffQuantity != 100 {
		panic("quantity is not 100% in formula")
	}

	start := time.Now()

	file, err := os.Open(path)
	if err != nil {
		println(err.Error())
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		println(err.Error())
	}
	if len(records) < 1 {
		panic("there is no tank")
	}

	var detectQuantity map[string]int = formula

	for _, record := range records {
		rec := strings.Split(record[0], ";")

		cast, err := strconv.Atoi(strings.TrimSpace(rec[1]))
		if err != nil {
			panic("wrong format")
		}

		if strings.Contains(rec[2], "/") {
			emptyTanks[cast] = append(emptyTanks[cast], Tank{rec[0], rec[2]})
		} else {
			usedTanks[cast] = append(usedTanks[cast], Tank{rec[0], rec[2]})
			detectQuantity[rec[2]] -= cast
		}
	}

	for key, d := range detectQuantity {
		if d > 0 {
			panic("Not enough of wine " + key)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
	println()

	fmt.Println("total : ", total)
	println()
	fmt.Println("formula  : ", formula)
	println()
	fmt.Println("used tanks  :\n", usedTanks)
	println()
	fmt.Println("empty tanks  :\n ", emptyTanks)

}
