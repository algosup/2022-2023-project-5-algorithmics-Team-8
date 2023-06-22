package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Action struct {
	Name string
	from []Tank
	to   []Tank
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
var formula2 map[string]float64 = make(map[string]float64)
var total float64 = 0
var solverTanksQ map[string]bool = make(map[string]bool)
var solverTanksID []Tank

func main() {
	var instructions []string

	if len(os.Args) < 2 {
		println("please provide path to file")
		return
	}
	path := os.Args[1]

	var rootNode Node = parseconfig(path)
	var emptySmallest float64 = -1
	for _, v := range rootNode.State.emptyTanks {
		if v.Capacity < emptySmallest {
			emptySmallest = v.Capacity
		}
	}

	start := time.Now()

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
		for _, key := range rootNode.State.emptyTanks {
			txt := fmt.Sprintf("There is 1 empty tank %vhL for a formula total of %vhL, we can only make do with the existing one ", key.Capacity, total)
			instructions = append(instructions, txt)
		}
	}

	for i := 0; i < len(rootNode.State.usedTanks); i++ {
		exist := false
		for wineN := range formula {
			_, exist = rootNode.State.usedTanks[i].Wines[wineN]
			if exist {
				break
			}
		}
		if !exist {
			rootNode.State.usedTanks = append(rootNode.State.usedTanks[:i], rootNode.State.usedTanks[i+1:]...)
			i--
		}
	}

	sort.SliceStable(rootNode.State.emptyTanks, func(i, j int) bool {
		return rootNode.State.emptyTanks[i].Capacity < rootNode.State.emptyTanks[j].Capacity
	})

	sort.SliceStable(rootNode.State.usedTanks, func(i, j int) bool {
		return rootNode.State.usedTanks[i].Capacity < rootNode.State.usedTanks[j].Capacity
	})

	search, found := findTanks(rootNode.State.emptyTanks, total)
	if !found {
		println("no solution possible because of wrong add up")
	}
	for _, se := range search {
		solverTanksQ[rootNode.State.emptyTanks[se].ID] = true
		solverTanksID = append(solverTanksID, rootNode.State.emptyTanks[se])
	}
	//now the algorithm
	result, err := Solve(formula, total, rootNode, instructions)
	if err != nil {
		println(err.Error())
		return
	}

	elapsed := time.Since(start)
	os.Create("output")
	file, err := os.OpenFile("output", os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	timeTaken := fmt.Sprintf("sebastien took %s \n\n", elapsed)
	file.WriteString(timeTaken)
	println()
	for _, instruction := range result {
		fmt.Println(instruction)
		file.WriteString(instruction + "\n")
	}
	file.WriteString("\nFinal tanks are :\n")
	for _, info := range solverTanksID {
		file.WriteString(info.ID + "\n")
	}
}

func findTanks(tanks []Tank, quantity float64) ([]int, bool) {

	for i := len(tanks) - 1; i >= 0; i-- {
		if solverTanksQ[tanks[i].ID] || tanks[i].Capacity > quantity {
			continue
		} else if tanks[i].Capacity == quantity {
			return []int{i}, true
		}
	}
	for i := len(tanks) - 1; i >= 0; i-- {
		if solverTanksQ[tanks[i].ID] || tanks[i].Capacity > quantity {
			continue
		}
		tks, found := findTanksRec(tanks[:i], quantity, tanks[i].Capacity, []int{i})
		if found {
			tks = append(tks, i)
			return tks, true
		}
	}
	return nil, false
}

func findTanksRec(tanks []Tank, searchedQ float64, quantity float64, ids []int) ([]int, bool) {
	var indices []int
	var sum float64
	for j, tank := range tanks {
		if solverTanksQ[tanks[j].ID] {
			continue
		}
		// println(sum, searchedQ, searchedQ == sum)
		if j >= ids[0] {
			return nil, false
		}
		for id := range ids {
			if j == id {
				continue
			}
		}

		sum = quantity + tank.Capacity
		ids = append(ids, j)
		if searchedQ > sum {
			indices, found := findTanksRec(tanks[j:], searchedQ, sum, ids)
			if found {
				var realInd []int
				for r := range realInd {
					realInd[r] += j
				}

				return indices, true
			} else if searchedQ == sum {
				var realInd []int
				for r := range realInd {
					realInd[r] += j
				}
				return indices, true
			} else {
				if len(ids) > 2 {
					ids = ids[:len(ids)-2]
				} else {
					ids = []int{ids[0]}
				}
			}
		} else if searchedQ == sum {
			indices = append(indices, j)
			return indices, true
		} else {
			continue
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
		pass := false
		quant := 0.
		wineN := ""
		for key, q := range formula {
			for key2 := range currentNode.State.usedTanks[i].Wines {
				if key2 == key && formula[key2] > 0 || key2 == key && formula2[key2] > 0 {
					pass = true
					quant = q
					wineN = key2
				}
			}
		}
		if !pass {
			continue
		}

		actual := currentNode.State.usedTanks[i]

		// Check if actual's quantity is equal to the content of a group of tanks
		receiverIndices, found := findTanks(currentNode.State.emptyTanks, quant)
		restInTank := actual.Capacity - quant

		found2 := false
		var receiverIndices2 []int
		if restInTank > 0 {
			// Check if actual's quantity is equal to the content of a group of tanks
			receiverIndices2, found2 = findTanks(currentNode.State.emptyTanks, actual.Capacity-quant)
		} else {
			found = true
		}

		if found && found2 {
			var receivers []Tank
			for _, receiver := range receiverIndices {
				receivers = append(receivers, currentNode.State.emptyTanks[receiver])
			}

			for _, receiver := range receiverIndices2 {
				receivers = append(receivers, currentNode.State.emptyTanks[receiver])
			}

			newStateDivide := TanksState{
				usedTanks:  []Tank{},
				emptyTanks: []Tank{},
			}
			newStateDivide.usedTanks = append(currentNode.State.usedTanks[:i], currentNode.State.usedTanks[i+1:]...)

			allIndices := receiverIndices
			allIndices = append(allIndices, receiverIndices2...)
			newStateDivide.emptyTanks = removeTanksByIndices(currentNode.State.emptyTanks, allIndices)

			for _, r := range receivers {
				r.Wines[wineN] = r.Capacity
				newStateDivide.usedTanks = append(newStateDivide.usedTanks, r)
			}
			newStateDivide.emptyTanks = append(newStateDivide.emptyTanks, actual)

			childNodeDivide := &Node{
				State:          newStateDivide,
				Transformation: Action{Name: "Divide", from: []Tank{actual}, to: receivers},
				Parent:         &currentNode,
				Children:       []*Node{},
				Score:          0,
			}
			Rate(childNodeDivide.Transformation, &childNodeDivide.Score)
			formula2[wineN] = 0
			bestChildNode = childNodeDivide
		} else {
			//add to found ttl tanks
			// for tk := range cur
			for k, tk := range solverTanksID {
				countW := ""
				for wn := range currentNode.State.usedTanks[i].Wines {
					countW = wn
				}
				if currentNode.State.usedTanks[i].Wines[countW] < formula[countW] {
					continue
				}
				try := tk.Capacity - formula[countW]
				try2 := currentNode.State.usedTanks[i].Wines[countW] - formula[countW]
				if try >= 0 {
					getTk, success := findTanks(currentNode.State.emptyTanks, try2)
					if success {
						formula[wineN] -= try
						for kley, wnes := range currentNode.State.usedTanks[i].Wines {
							solverTanksID[k].Wines[kley] += wnes
						}
						solverTanksID[k].Capacity = try
						var receivers []Tank
						receivers = append(receivers, solverTanksID[k])
						for _, receiver := range getTk {
							receivers = append(receivers, currentNode.State.emptyTanks[receiver])
						}

						newStateDivide := TanksState{
							usedTanks:  []Tank{},
							emptyTanks: []Tank{},
						}
						newStateDivide.usedTanks = append(currentNode.State.usedTanks[:i], currentNode.State.usedTanks[i+1:]...)

						newStateDivide.emptyTanks = removeTanksByIndices(currentNode.State.emptyTanks, getTk)

						for _, r := range receivers {
							r.Wines[wineN] = r.Capacity
							newStateDivide.usedTanks = append(newStateDivide.usedTanks, r)
						}

						childNodeDivide := &Node{
							State:          newStateDivide,
							Transformation: Action{Name: "Divide To Final", from: []Tank{actual}, to: receivers},
							Parent:         &currentNode,
							Children:       []*Node{},
							Score:          0,
						}
						Rate(childNodeDivide.Transformation, &childNodeDivide.Score)
						formula2[wineN] = 0
						bestChildNode = childNodeDivide
						break
					} else {
						continue
					}
				}
				//  else if try == 0 {

				// 	println("EQUAL 0", wineN)
				// }
			}
		}
		if bestChildNode.Score == -1 {
			continue
		}
		from := ""
		receivers := ""
		for _, tank := range bestChildNode.Transformation.from {
			from += tank.ID + " "
		}
		for _, tank := range bestChildNode.Transformation.to {
			receivers += tank.ID + " "
		}
		instruction := fmt.Sprintf("%s: from: %s, to: %s", bestChildNode.Transformation.Name, from, receivers)
		instructions = append(instructions, instruction)

		sort.SliceStable(bestChildNode.State.emptyTanks, func(i, j int) bool {
			return bestChildNode.State.emptyTanks[i].Capacity < bestChildNode.State.emptyTanks[j].Capacity
		})

		sort.SliceStable(bestChildNode.State.usedTanks, func(i, j int) bool {
			return bestChildNode.State.usedTanks[i].Capacity < bestChildNode.State.usedTanks[j].Capacity
		})
		return Solve(formula, total, *bestChildNode, instructions)
	}
	return instructions, nil
}

func Rate(act Action, score *float64) {
	//mix of two wines is closer than random moving a wine
	//right quantity of a wine in tank is closer than random moving a wine
	for _, i := range act.to {
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

func parseconfig(filepath string) Node {
	var rootNode Node

	file, err := os.Open(filepath)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.TrimSpace(line) == "" || line[0] == '!' || line[0] == '\r' { //Line is a comment so ignore it while parsing
			continue
		}

		tokens := strings.Split(line, ";")

		switch len(tokens) {

		case 1: // Case Desired Quantity
			quantity, err := strconv.Atoi(tokens[0])
			if err != nil {
				log.Fatal("abc is to crash at quantity")
				os.Exit(1)
			}

			if total != 0 {
				log.Fatal("Error: Desired quantity was defined more than once")
				os.Exit(1)
			}
			total = float64(quantity)

		case 2: //Case Wine ; Percentage
			var wine string = strings.TrimSpace(tokens[0])

			if formula[wine] != 0 {
				log.Fatalf("Error: Duplicate wine name found: %s", wine)
				os.Exit(1)
			}

			percentage, err := strconv.ParseFloat(tokens[1], 64)
			if err != nil {
				log.Fatalf("Error: Wine in fomula lacks percentage")
				os.Exit(1)
			}

			formula[wine] = percentage

			formula2[wine] = percentage

		case 3: //Case TankID ; Capacity ; Content
			tankID := tokens[0][1:]

			capacity, _ := strconv.ParseFloat(tokens[1], 64)

			content := strings.TrimSpace(tokens[2])

			if strings.Contains(content, "/") {
				rootNode.State.emptyTanks = append(rootNode.State.emptyTanks, Tank{tankID, make(map[string]float64), capacity})
				continue
			}
			rootNode.State.usedTanks = append(rootNode.State.usedTanks, Tank{tankID, map[string]float64{content: capacity}, capacity})

		}
	}
	return rootNode
}
