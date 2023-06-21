package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Tank struct {
	ID           string
	volume       float64
	winesStocked []string
	winesUsed    []float64
	transferred  bool
	capacity     float64
}

type RemainingVolume struct {
	wine   string
	volume float64
	tankID int
}

func GenerateMemoKey(wines []string, index int) string {
	return strings.Join(wines, "_") + strconv.Itoa(index)
}

func FindTankCombination(tanks []Tank, remainingVolume float64, bestCombination []int, bestDifference *float64, currentCombination *[]int, index int, memo map[string]struct {
	difference  float64
	combination []int
}) {
	if remainingVolume < 0.1 {
		difference := math.Abs(remainingVolume)
		if difference < *bestDifference {
			*bestDifference = difference
			bestCombination = make([]int, len(*currentCombination))
			copy(bestCombination, *currentCombination)
		}
		return
	}

	memoKey := GenerateMemoKey([]string{fmt.Sprintf("%f", remainingVolume)}, index)
	if memoEntry, exists := memo[memoKey]; exists {
		if memoEntry.difference < *bestDifference {
			*bestDifference = memoEntry.difference
			bestCombination = make([]int, len(memoEntry.combination))
			copy(bestCombination, memoEntry.combination)
		}
		return
	}

	for i := index; i < len(tanks); i++ {
		if tanks[i].volume >= 0.1 && fmt.Sprintf("%v", tanks[i].winesUsed[0]) == "/" && !contains(convertToStringSlice(bestCombination), strconv.Itoa(i)) {
			*currentCombination = append(*currentCombination, i)

			FindTankCombination(tanks, remainingVolume-tanks[i].volume, bestCombination, bestDifference, currentCombination, i+1, memo)

			*currentCombination = (*currentCombination)[:len(*currentCombination)-1]
		}
	}
}

func convertToStringSlice(intSlice []int) []string {
	strSlice := make([]string, len(intSlice))
	for i, v := range intSlice {
		strSlice[i] = strconv.Itoa(v)
	}
	return strSlice
}

func FindBestTanksCombination(tanks []Tank, remainingVolume float64, bestCombination *[]int, bestDifference *float64) bool {
	currentCombination := make([]int, 0)
	memo := make(map[string]struct {
		difference  float64
		combination []int
	})
	FindTankCombination(tanks, remainingVolume, *bestCombination, bestDifference, &currentCombination, 0, memo)

	return len(*bestCombination) > 0
}

// Different data structures
func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

func main() {

	var configFilePath string

	fmt.Println("Enter the path of the config file: ")
	fmt.Scanln(&configFilePath)

	configFilePath = strings.ReplaceAll(configFilePath, "'", "")

	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalf("Impossible to open the config file selected or you have forgot to drag and drop it: %s", configFilePath)
	}

	defer configFile.Close()

	fmt.Printf("Reading config file: %s\n", configFilePath)

	var tanks []Tank
	var wines []string
	var percentages []float64
	var quantities []float64

	// tankIds := make(map[string]struct{})
	tanksLineCount := 0

	filePath := "./output.txt"
	file, err := os.Create(filePath)
	scanner := bufio.NewScanner(configFile)
	tankIds := make(map[string]struct{})

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '!' || line[0] == '\r' {
			continue // ignore empty lines and comments beginning with '!'
		} else if line[0] == '#' {
			arrSpl := strings.Split(line, ";")
			if len(arrSpl) != 3 {
				log.Fatalf("duh error")
				return
			}

			tankID := arrSpl[0][1:]

			if _, exists := tankIds[tankID]; exists {
				log.Fatalf("Error: Duplicate tank ID found: %s", tankID)
				return
			}

			v, _ := strconv.ParseFloat(arrSpl[1], 64)

			tank := Tank{
				ID:       tankID,
				volume:   v,
				capacity: v,
			}

			if tank.volume <= 0 || tank.capacity > 1000 {
				log.Fatalf("Error: Invalid tank volume for tank %s. Volume should be between 0 and 1000.", tank.ID)
				return
			}

			wineTank := strings.TrimSpace(strings.ToLower(arrSpl[2]))
			var wine string
			found := false
			if strings.Contains(wineTank, "champagne") || strings.Contains(wineTank, "/") {
				wine = wineTank
				found = true
			} else {
				found = false
				for _, existingWine := range wines {
					if existingWine == wineTank {
						println(existingWine) //! Remove when fixed
						found = true
						wine = existingWine
						break
					}
				}
			}

			if !found {
				log.Fatal("Error: Unknown wine type for the tank", wineTank)
				return
			}

			tank.winesUsed = append(tank.winesUsed, v)
			tank.winesStocked = append(tank.winesStocked, wine)
			tanks = append(tanks, tank)

			tankIds[tankID] = struct{}{}

			tanksLineCount++
		} else if unicode.IsDigit(rune(line[0])) {
			quantity, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
				return
			}
			quantities = append(quantities, float64(quantity))
			if len(quantities) >= 2 {
				log.Println("Warning: Only the first quantity value will be considered. Ignoring additional values.")
				return
			}
		} else {
			pos := strings.Index(line, ";")
			if pos != -1 {
				var wine string = strings.TrimSpace(strings.ToLower(line[:pos]))

				var wineExist bool = false
				for _, existingWine := range wines {
					if existingWine == wine {
						wineExist = true
						break
					}
				}
				if !wineExist {
					wines = append(wines, wine)
					percentage, _ := strconv.ParseFloat(line[pos+1:], 64)
					percentages = append(percentages, percentage)
				} else {
					log.Fatalf("Error: Duplicate wine name found: %s", wine)
					return
				}

			}
		}
	}
	if (len(tanks) == 0) || (len(wines) == 0) || (len(quantities) == 0) {
		log.Fatal("Error: Missing information in the config file.")
		return
	}

	var totalPercentage float64 = 0.0
	for _, percentage := range percentages {
		if percentage <= 0 {
			log.Fatal("Error: Invalid percentage value.")
			return
		}
		totalPercentage += percentage
	}
	if totalPercentage != 100 {
		log.Fatal("Error: The sum of the percentages must be equal to 100.")
		return
	}

	for _, tank := range tanks {
		if tank.volume > 1000 {
			log.Fatal("Error: Invalid tank volume.")
			return
		}
		for i := 0; i < len(wines); i++ {
			if contains(tank.winesStocked, wines[i]) {
				continue
			}
			if percentages[i] <= float64(tank.capacity)/100.0 {
				tank.winesStocked = append(tank.winesStocked, wines[i])
				tank.capacity -= percentages[i] * 100.0
			}
		}
	}
	println("END PARSING FINAL 2") //! Remove when fixed

	fmt.Println("Wine:")
	for _, wine := range wines {
		fmt.Println(wine)
	}

	fmt.Println("Percentage:")
	for _, percentages := range percentages {
		fmt.Println(percentages)
	}

	fmt.Println("Quantity needed:")
	for _, quantity := range quantities {
		fmt.Println(quantity, "hl")
	}

	var totalVolume float64 = 0.0
	for _, quantity := range quantities {
		totalVolume += quantity
	}

	//? calculating percentages per wine
	requiredVolume := make([]float64, len(wines))
	for i := 0; i < len(wines); i++ {
		requiredVolume[i] = totalVolume * percentages[i] / 100
	}

	tanksVolumes := make([][]float64, len(tanks))
	for i := 0; i < len(tanks); i++ {
		tanksVolumes[i] = make([]float64, len(wines))
		for j := 0; j < len(wines); j++ {
			tanksVolumes[i][j] = tanks[i].volume * percentages[j] / 100
		}
	}

	tanksToUse := make([][]int, len(wines))
	totalVolumes := make(map[string]float64)
	remainingVolumes := make([]RemainingVolume, len(wines))

	for i := 0; i < len(wines); i++ {
		remainingVolume := requiredVolume[i]
		totalAvailableVolume := 0.0

		for j := 0; j < len(tanks) && remainingVolume > 0; j++ {
			// println(j, i)
			//fmt.Println(tanksVolumes[j])
			// fmt.Println(tanks[j].winesStocked)
			if tanksVolumes[j][i] > 0 && tanks[j].winesStocked != nil && tanks[j].winesStocked[0] == "/" {
				totalAvailableVolume += tanksVolumes[j][i]
				// println(tanks[j].ID)
			}
		}

		if remainingVolume > totalAvailableVolume {

			fmt.Println("The volume needed (" + fmt.Sprintf("%f", remainingVolume) + ") is more than the volume you have inside your empty tanks" + fmt.Sprintf("%f", totalAvailableVolume))
			return
		}

		for j := 0; j < len(tanks) && remainingVolume > 0; j++ {
			if tanksVolumes[j][i] > 0 && tanks[j].winesStocked[0] == "/" {
				volumeToUse := math.Min(remainingVolume, tanksVolumes[j][i])
				tanksToUse[i] = append(tanksToUse[i], j)
				remainingVolume -= volumeToUse
				totalVolumes[wines[i]] += volumeToUse
			}
		}
	}

	totalUsedVolume := 0.0
	for i := 0; i < len(wines); i++ {
		for j := 0; j < len(tanksToUse[i]); j++ {
			var tankIndex int = tanksToUse[i][j]
			totalUsedVolume += tanksVolumes[tankIndex][i]
		}
	}

	if totalUsedVolume < totalVolume {
		for i := 0; i < len(wines); i++ {
			for (totalUsedVolume > totalVolume) && (len(tanksToUse[i]) > 0) {
				var lastTankIndex = tanksToUse[i][len(tanksToUse[i])-1]
				var lastTankVolume float64 = tanksVolumes[lastTankIndex][i]
				tanksToUse[i] = tanksToUse[i][:len(tanksToUse[i])-1]
				totalUsedVolume -= lastTankVolume
			}
		}
	}
	for i := 0; i < len(wines); i++ {
		var totalTanksWithWineVol float64 = 0.0
		var reqVol float64 = requiredVolume[i]

		// Check if enough of said wine exist (buffer)
		for j := 0; j < len(tanks); j++ {
			if contains(tanks[j].winesStocked, wines[i]) {
				totalTanksWithWineVol += tanks[j].volume
			}
		}

		fmt.Println("required volume for " + wines[i] + " is " + fmt.Sprintf("%f", reqVol) + "hl")

		// Checker of if enough of said wine exist
		if totalTanksWithWineVol < reqVol {
			log.Fatal("Error, the volume of " + wines[i] + " is not enough to fill a tank")
			return
		}

		var remainingVolume float64 = totalTanksWithWineVol - reqVol

		if remainingVolume < 0 {
			fmt.Println("Impossible to have " + wines[i] + " you need to empty " + fmt.Sprintf("%f", (remainingVolume)) + "hl more")
		} else {
			fmt.Println("Remaining volume for " + wines[i] + " is " + fmt.Sprintf("%f", remainingVolume))
		}

		file.Write([]byte("FR\n"))
		file.Write([]byte("Ce fichier est le résultat de la formule, il contiens toutes les étapes de la solution\n"))
		file.Write([]byte("\n"))
		file.Write([]byte(("EN\n")))
		file.Write([]byte([]byte("This file is the result of the formula, it contains all the steps of the solut)ion\n")))
		file.Write([]byte("\n"))

		// var transfertsChampagne map[string]float64
		// var transferedVolumes map[string]map[int]float64

		for i := 0; i < len(wines); i++ {
			file.Write([]byte("\n"))
			file.Write([]byte("Wine: " + wines[i] + ":\n"))

			var tanksAlreadyWritten map[int]bool = map[int]bool{}

			for j := 0; j < len(tanksToUse[i]); j++ {
				var tankIndex int = tanksToUse[i][j]
				tanks[tankIndex].winesStocked[0] = "Champagne"

				if !tanksAlreadyWritten[tankIndex] {
					file.Write([]byte("Tank: " + tanks[tankIndex].ID + "\n"))
					tanksAlreadyWritten[tankIndex] = true

					for k := 0; k < len(wines); k++ {
						var percentage float64 = tanksVolumes[tankIndex][k] / tanks[tankIndex].volume * 100
						var equivalentVolume float64 = tanksVolumes[tankIndex][k]
						file.Write([]byte("	" + wines[k] + ": " + fmt.Sprintf("%f", percentage) + "% (" + fmt.Sprintf("%f", equivalentVolume) + "hl)\n"))
					}

					file.Write([]byte("\n"))
					var volumeNeeded float64 = tanksVolumes[tankIndex][i]
					file.Write([]byte("Filling " + tanks[tankIndex].ID + " (" + wines[i] + ")\n"))

					var remainingVolume float64 = volumeNeeded
					var transferDone bool = false
					var k int = 0

					for (remainingVolume > 0) && (k < len(tanksToUse[i])) {
						if tanks[k].winesStocked[0] == wines[i] && tanks[k].volume > 0 && tanks[k].volume > 0.001 {
							var availableVolume float64 = tanks[k].volume
							var transferVolume float64 = math.Min(remainingVolume, availableVolume)
							file.Write([]byte(string("Transfering " + fmt.Sprintf("%f", transferVolume) + "hl from " + tanks[k].ID + "\n")))
							remainingVolume -= transferVolume
							tanks[k].capacity -= transferVolume
							transferDone = true

							if tanksVolumes[k][i] < 0.001 {
								tanksVolumes[k][i] = 0.0
							}
						}
						k++
					}

					if remainingVolume > 0 && !transferDone {
						file.Write([]byte("Insufficient available volume in other cuves of the same wine to fill \n"))
					} else if remainingVolume > 0 {
						file.Write([]byte("Insufficient available volume in other cuves of the same wine to fill \n"))
					} else {
						file.Write([]byte("\n"))
					}
				}
			}
		}

		var totalVolumeFromAllWines float64 = 0.0
		for i := 0; i < len(wines); i++ {
			for j := 0; j < len(tanksToUse[i]); j++ {
				var tanksIndex = tanksToUse[i][j]
				totalVolumeFromAllWines += tanksVolumes[tanksIndex][i]
			}
		}
		fmt.Println("Total volume from all wines: " + fmt.Sprintf("%f", totalVolumeFromAllWines) + "hl")

		volumesUsed := make([]float64, len(wines))
		for i := 0; i < len(tanks); i++ {
			tank := &tanks[i]
			for j := 0; j < len(wines); j++ {
				if contains(tank.winesStocked, wines[j]) {
					volumeUsed := math.Min(volumesUsed[j], float64(tank.volume))
					VolumeRemaining := math.Max(float64(tank.volume)-volumeUsed, 0.0)
					if VolumeRemaining == 0 {
						tank.volume = 0
						tank.winesStocked[0] = "/"
					} else {
						tank.volume = VolumeRemaining
					}
					volumesUsed[j] -= volumeUsed

					// Storage of the remaining volume for each wine and each tank
					volume := RemainingVolume{
						tankID: i,
						wine:   wines[j],
						volume: VolumeRemaining,
					}
					remainingVolumes = append(remainingVolumes, volume)
				}
			}
		}

		file.Write([]byte("\n"))
		file.Write([]byte("EN\n"))
		file.Write([]byte("Remaining volumes:\n"))
		file.Write([]byte("Steps are in the order of the wines\n"))
		file.Write([]byte("\n"))
		file.Write([]byte("FR\n"))
		file.Write([]byte("Volumes restants:\n"))
		file.Write([]byte("Les étapes sont dans l'ordre des vins\n"))

		foundCombination := false
		var bestCombination []int

		for i := 0; i < len(tanks); i++ {
			tank := &tanks[i]
			for _, wStckd := range tank.winesStocked {
				//! ENTIEREMENT À REFAIRE, JE PEUX AVOIR PLUSIEURS VINS
				if wStckd != "/" && tank.volume > 0 && wStckd != "Champagne" && wStckd != "champagne" && wStckd != "CHAMPAGNE" && !tank.transferred {
					fmt.Println("Tank " + tank.ID + " still has " + fmt.Sprintf("%f", tank.volume) + "hl of " + wStckd)

					var remainingVolume float64 = 0.0
					//! A REFAIRE, si j'ai deux vins il se passe quoi ???????????????
					for j := 0; j < len(wines); j++ {
						if remainingVolumes[j].wine == wStckd && remainingVolumes[j].tankID == i {
							remainingVolume = remainingVolumes[j].volume
							break
						}
					}
					fmt.Println("Remaining volume of this wine in this tank: " + fmt.Sprintf("%f", remainingVolume) + "hl")
					fmt.Println("Searching for a tank with the same wine to transfer the remaining volume")

					bestTankIndex := -1
					// bestTankScore := math.MaxFloat64
					// differenceMin := math.MaxFloat64
					epsilon := 0.0001

					for j := 0; j < len(tanks); j++ {
						if j != i && tanks[j].volume <= remainingVolume && tanks[j].winesStocked[0] == "/" && tanks[j].volume < tanks[j].volume {
							var score float64 = math.Round(tanks[j].volume*10) / 10.0
							if math.Abs(remainingVolume-score) < epsilon {
								bestTankIndex = j
								// bestTankScore = score
								// differenceMin = math.Abs(score - remainingVolume)
							}
						}
					}

					if bestTankIndex != -1 {
						tanks[bestTankIndex].winesStocked[0] = wStckd
						file.Write([]byte("Transfering " + fmt.Sprintf("%F", tanks[bestTankIndex].volume) + "hl from " + wStckd + " to " + tanks[bestTankIndex].ID + "\n"))
						file.Write([]byte("\n"))
						wStckd = "/"
						tank.transferred = true
						tanks[bestTankIndex].transferred = true
						tank.volume -= tanks[bestTankIndex].volume
					} else if tank.volume != tank.capacity {
						fmt.Println("No tank with the same wine found.")

						var bestDifference = math.MaxFloat64
						foundCombination = FindBestTanksCombination(tanks, remainingVolume, &bestCombination, &bestDifference)
					}
				}
			}
		}
		var tank Tank
		if foundCombination {
			fmt.Println("Combination of tank found")

			var tanksAlreadyTransferred map[int]struct{}
			bestCombinationLen := len(bestCombination)
			tanksAlreadyTransferred = make(map[int]struct{}, bestCombinationLen)

			for index := 0; index < bestCombinationLen; index++ {
				_, ok := tanksAlreadyTransferred[index]
				if tanks[index].capacity >= 0.1 && !ok {
					tank = tanks[index]
					volumeOfTransfert := math.Min(tank.capacity, remainingVolume)
					if volumeOfTransfert >= 0.1 && tank.volume >= 0.1 {
						tank.volume -= volumeOfTransfert
						if tank.volume < 0 {
							tank.volume += volumeOfTransfert
							break
						}
						tanks[index].winesStocked[0] = tank.winesStocked[0]
						tanks[index].transferred = true
						file.Write([]byte("Transfer of " + fmt.Sprintf("%f", volumeOfTransfert) + "hl of " + tanks[index].winesStocked[0] + " from tank " + tank.ID + " to tank " + tanks[index].ID + "\n"))
						tanksAlreadyTransferred[index] = struct{}{}
					}
				}
			}

			file.WriteString("\n")
		}

		if tank.volume < 0.001 {
			if tank.winesStocked == nil {
				tank.winesStocked = append(tank.winesStocked, "/")
			} else {
				tank.winesStocked[0] = "/"
			}
		}

		// for i := 0; i < len(tanks); i++ {
		// tank := &tanks[i]
		// if tank.winesStocked[0] == "/" {
		// var volume float64 = tank.volume - tank.volume
		// fmt.Println("Tank " + tank.ID + ": " + fmt.Sprintf("%f", volume) + "hl of " + tank.winesStocked[0])
		// } else {
		// fmt.Println("Tank " + tank.ID + ": " + fmt.Sprintf("%f", tank.volume) + "hl of " + tank.winesStocked[0])
		// }
		// }
		file.Write([]byte("\n"))
		file.Write([]byte("EN\n"))
		file.Write([]byte("Remaining volumes:\n"))
		file.Write([]byte("Leftover volumes are in the order of the wines\n"))
		file.Write([]byte("\n"))
		file.Write([]byte("FR\n"))
		file.Write([]byte("Volumes restants:\n"))
		file.Write([]byte("Les restes sont dans l'ordre des vins\n"))
		file.Write([]byte("\n"))

		for i := 0; i < len(tanks); i++ {
			tank := &tanks[i]
			if tank.volume > 0 && tank.volume < tank.volume {
				file.Write([]byte("Tank " + tank.ID + ": " + fmt.Sprintf("%f", tank.volume) + "hl of " + tank.winesStocked[0] + "\n"))
			}
		}

		file.Write([]byte("\n"))
		file.Write([]byte("Tank size containing champagne\n"))
		for i := 0; i < len(tanks); i++ {

			tank := &tanks[i]
			if tank.winesStocked[0] == "Champagne" {
				file.Write([]byte("Tank " + tank.ID + ": " + fmt.Sprintf("%f", tank.volume) + "hl of " + tank.winesStocked[0] + "\n"))
			}
		}
		file.Write([]byte("\n"))
		file.Write([]byte("Total champagne volume is " + fmt.Sprintf("%f", totalVolumeFromAllWines) + "hl\n"))

		fmt.Println("The file has been written successfully\n")
		file.Close()
	}
}
