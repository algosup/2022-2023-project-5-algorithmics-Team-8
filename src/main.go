package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Action struct {
	Name         string
	Transmitters []Tank
	Receivers    []Tank
}

type Tank struct {
	ID       string
	Wines    map[string]float64
	Capacity float64
}

type TanksState struct {
	usedTanks  []Tank
	emptyTanks []Tank
}

type Node struct {
	State          TanksState
	Transformation Action
	Score          float64
	Parent         *Node
	Children       []*Node
}

var formula map[string]float64 = make(map[string]float64)
var total float64

func main() {
	var instructions []string

	var rootNode Node
	var emptySmallest float64 = 0

	if len(os.Args) < 2 {
		println("please provide path to file")
		return
	}
	path := os.Args[1]

	fmt.Println("\n---- HI, Please give the total Hl you want as a result ----")
	var tempTTL string
	// Taking input from user
	fmt.Scanln(&tempTTL)

	totalt, err := strconv.ParseFloat(tempTTL, 64)
	if err != nil {
		println(err.Error())
		return
	}
	total = totalt

	fmt.Println("\n Please enter the wine and its quantity (%) separated by a space (ex : Chardonnay 80)")
	fmt.Println("If you write the same wine name twice (uppercase = lowercase), the first input will be overridden")
	fmt.Println("if quantity is not equal to 100% after you are done, the program will Exit")
	fmt.Println("WRITE the word 'done' when you are finished")
	println()

	buffQuantity := 0.0
	inputF := bufio.NewReader(os.Stdin)
	//will Exit only on "done" written
	for {
		inputF, err := inputF.ReadString('\n')
		if err != nil {
			println("error reading terminal, verify your authorizations")
			return
		}

		inputF = strings.TrimSpace(inputF)
		if inputF == "done" {
			break
		} else {
			fullInput := strings.Split(inputF, " ")
			if len(fullInput) < 2 {
				fmt.Println("there is not enough arguments, 2 are needed !")
			} else {
				convVal, err := strconv.ParseFloat(fullInput[1], 64)
				if err != nil {
					fmt.Println("your second value is not a number, please try again.")
				} else {
					quantity := float64(total) * convVal / 100
					fmt.Println("Quantity: ", quantity, "/", total)

					formula[strings.ToLower(fullInput[0])] = quantity
					buffQuantity += convVal
				}
			}
		}
		fmt.Println("\n---- type your next value ----")
	}
	if buffQuantity != 100 {
		println("quantity is not 100% in formula")
		return
	}

	start := time.Now()

	file, err := os.Open(path)
	if err != nil {
		println(err.Error())
		return
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		println(err.Error())
		return
	}

	if len(records) < 1 {
		println("there is no tank")
		return
	}

	var detectQuantity map[string]float64 = make(map[string]float64)
	for key, value := range formula {
		detectQuantity[key] = value
	}

	for _, record := range records {
		rec := strings.Split(record[0], ";")
		rec[2] = strings.TrimSpace(strings.ToLower(rec[2]))
		cast, err := strconv.ParseFloat(strings.TrimSpace(rec[1]), 64)
		if err != nil {
			println("wrong format")
			return
		}

		if strings.Contains(rec[2], "/") {
			if emptySmallest == 0 || emptySmallest > cast {
				emptySmallest = cast
			}
			rootNode.State.emptyTanks = append(rootNode.State.emptyTanks, Tank{rec[0], make(map[string]float64), cast})
		} else {
			rootNode.State.usedTanks = append(rootNode.State.usedTanks, Tank{rec[0], map[string]float64{rec[2]: cast}, cast})
			if detectQuantity[rec[2]] != 0 {
				detectQuantity[rec[2]] -= cast
			}
		}
	}

	//Obligatory error
	for key, d := range detectQuantity {
		if d > 0 {
			println("Not enough of wine " + key)
			return
		}
	}

	if len(rootNode.State.emptyTanks) == 0 {
		println("Not enough empty tanks to attempt finding a solution")
		return
	}

	if emptySmallest > float64(total) {
		println("the smallest empty tank is bigger than the total, the operation is impossible")
		return
	}

	//conditional possibility
	if len(rootNode.State.emptyTanks) == 1 {
		for key, _ := range rootNode.State.emptyTanks {
			txt := fmt.Sprintf("There is 1 empty tank %vhL for a formula total of %vhL, we can only make do with the existing one ", key, total)
			instructions = append(instructions, txt)
		}
	}

	//now the algorithm
	result, err := Solve(formula, total, rootNode, instructions)
	if err != nil {
		println(err.Error())
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("page took %s", elapsed)
	println()
	for _, instruction := range result {
		fmt.Println(instruction)
	}
}

func findTanks(tanks []Tank, quantity float64) ([]int, bool) {
	var indices []int
	var sum float64

	for i, tank := range tanks {
		if sum+tank.Capacity <= quantity {
			sum += tank.Capacity
			indices = append(indices, i)

			if sum == quantity {
				return indices, true
			}
		}
	}

	return nil, false
}

func removeTanksByIndices(tanks []Tank, indices []int) []Tank {
	sort.Ints(indices)
	offset := 0
	for _, index := range indices {
		adjustedIndex := index - offset
		if adjustedIndex >= 0 && adjustedIndex < len(tanks) {
			tanks = append(tanks[:adjustedIndex], tanks[adjustedIndex+1:]...)
			offset++
		}
	}
	return tanks
}

func Solve(formula map[string]float64, total float64, currentNode Node, instructions []string) ([]string, error) {
	if len(currentNode.State.emptyTanks) == 0 {
		return instructions, nil
	}

	bestChildNode := &Node{
		State:          TanksState{},
		Transformation: Action{},
		Score:          -1,
		Parent:         &currentNode,
		Children:       []*Node{},
	}

	for i := 0; i < len(currentNode.State.usedTanks); i++ {
		transmitter := currentNode.State.usedTanks[i]

		for j := 0; j < len(currentNode.State.emptyTanks); j++ {
			receiver := currentNode.State.emptyTanks[j]

			// Check if transmitter's quantity is equal to the content of a group of tanks
			receiverIndices, found := findTanks(currentNode.State.emptyTanks, transmitter.Capacity)
			if found {
				newStateDivide := TanksState{
					usedTanks:  make([]Tank, len(currentNode.State.usedTanks)),
					emptyTanks: make([]Tank, len(currentNode.State.emptyTanks)),
				}
				copy(newStateDivide.usedTanks, currentNode.State.usedTanks)
				copy(newStateDivide.emptyTanks, currentNode.State.emptyTanks)

				// Remove transmitter from used tanks and add it to empty tanks
				newStateDivide.usedTanks = append(newStateDivide.usedTanks[:i], newStateDivide.usedTanks[i+1:]...)

				// Remove receiver tanks from empty tanks and add them to used tanks
				for _, index := range receiverIndices {
					newStateDivide.usedTanks = append(newStateDivide.usedTanks, newStateDivide.emptyTanks[index])
				}
				newStateDivide.emptyTanks = removeTanksByIndices(newStateDivide.emptyTanks, receiverIndices)

				childNodeDivide := &Node{
					State:          newStateDivide,
					Transformation: Action{Name: "Divide", Transmitters: []Tank{transmitter}, Receivers: newStateDivide.usedTanks[i:]},
					Parent:         &currentNode,
					Children:       []*Node{},
					Score:          0,
				}

				Rate(childNodeDivide.Transformation, &childNodeDivide.Score)

				if childNodeDivide.Score > bestChildNode.Score {
					bestChildNode = childNodeDivide
				}

				if childNodeDivide.Score >= 10 {
					return instructions, nil
				}
			}

			// Combine action
			for k := 0; k < len(currentNode.State.emptyTanks); k++ {
				if k == j {
					continue
				}
				receiver2 := currentNode.State.emptyTanks[k]

				newStateCombine := TanksState{
					usedTanks:  make([]Tank, len(currentNode.State.usedTanks)),
					emptyTanks: make([]Tank, len(currentNode.State.emptyTanks)),
				}
				copy(newStateCombine.usedTanks, currentNode.State.usedTanks)
				copy(newStateCombine.emptyTanks, currentNode.State.emptyTanks)

				// Remove transmitter and receiver from used tanks and add them to empty tanks
				newStateCombine.usedTanks = append(newStateCombine.usedTanks[:i], newStateCombine.usedTanks[i+1:]...)

				// Find the right group for transmitter and receiver
				transmitterIndex, found := findTanks(newStateCombine.emptyTanks, transmitter.Capacity)
				if found {
					for t := range transmitterIndex {
						if t > k {
							newStateCombine.emptyTanks = append(newStateCombine.emptyTanks[:t], newStateCombine.emptyTanks[t+1:]...)
							newStateCombine.emptyTanks = append(newStateCombine.emptyTanks[:k], append([]Tank{transmitter, receiver}, newStateCombine.emptyTanks[k:]...)...)
						} else {
							newStateCombine.emptyTanks = append(newStateCombine.emptyTanks[:k], append([]Tank{transmitter, receiver}, newStateCombine.emptyTanks[k:]...)...)
							newStateCombine.emptyTanks = append(newStateCombine.emptyTanks[:t], newStateCombine.emptyTanks[t+1:]...)
						}
					}
				} else {
					newStateCombine.emptyTanks = append(newStateCombine.emptyTanks[:k], append([]Tank{receiver}, newStateCombine.emptyTanks[k:]...)...)
					newStateCombine.emptyTanks = append(newStateCombine.emptyTanks[:j], append([]Tank{transmitter}, newStateCombine.emptyTanks[j:]...)...)
				}

				childNodeCombine := &Node{
					State:          newStateCombine,
					Transformation: Action{Name: "Combine", Transmitters: []Tank{transmitter, receiver}, Receivers: []Tank{receiver2}},
					Parent:         &currentNode,
					Children:       []*Node{},
					Score:          0,
				}

				Rate(childNodeCombine.Transformation, &childNodeCombine.Score)

				if childNodeCombine.Score > bestChildNode.Score {
					bestChildNode = childNodeCombine
				}

				if childNodeCombine.Score >= 10 {
					return instructions, nil
				}
			}
		}
	}

	transmitters := ""
	receivers := ""
	for _, tank := range bestChildNode.Transformation.Transmitters {
		transmitters += tank.ID + " "
	}
	for _, tank := range bestChildNode.Transformation.Receivers {
		receivers += tank.ID + " "
	}
	instruction := fmt.Sprintf("%s: Transmitters: %s, Receivers: %s", bestChildNode.Transformation.Name, transmitters, receivers)
	instructions = append(instructions, instruction)

	return Solve(formula, total, *bestChildNode, instructions)
}

func Rate(act Action, score *float64) {
	//mix of two wines is closer than random moving a wine
	//right quantity of a wine in tank is closer than random moving a wine
	for _, i := range act.Receivers {
		var buff float64 = 0
		for j, k := range i.Wines {
			if formula[j] > 0 {
				buff += (k * i.Capacity / 100) / (formula[j] * total / 100)
			} else {
				*score = 0
				return
			}
		}
		*score += buff
	}
}

func proposition(prop string) {
	println("\n", prop, ", should we proceed with Y/N ?")
	var agreement string
	// Taking input from user
	fmt.Scanln(&agreement)
	if strings.Contains(strings.ToLower(agreement), "y") {
		println("ok proceeding")
	} else {
		println("Abort mission!")
		os.Exit(0)
	}
}
