package main

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

func extractNumbers(pos string) ([]int, error) {
	posIdx := strings.Split(pos, ",")
	posX, err := strconv.Atoi(posIdx[0])
	if err != nil {
		return []int{}, err
	}
	posY, err := strconv.Atoi(posIdx[1])
	if err != nil {
		return []int{}, err
	}
	posZ, err := strconv.Atoi(posIdx[2])
	if err != nil {
		return []int{}, err
	}
	return []int{posX, posY, posZ}, nil
}

func calcDistance(pos1, pos2 string) float64 {
	pos1Idx, err := extractNumbers(pos1)
	if err != nil {
		log.Fatal(err)
	}
	pos2Idx, err := extractNumbers(pos2)
	if err != nil {
		log.Fatal(err)
	}
	posX := pos2Idx[0] - pos1Idx[0]
	posY := pos2Idx[1] - pos1Idx[1]
	posZ := pos2Idx[2] - pos1Idx[2]
	return math.Sqrt(math.Pow(float64(posX), 2) + math.Pow(float64(posY), 2) + math.Pow(float64(posZ), 2))
}

func findClosestPair(lines []string) (float64, [2]string) {
	shortestDistance := math.MaxFloat64
	var shortestDistancePair [2]string
	for idx, current := range lines {
		for _, line := range lines[idx+1:] {
			dist := calcDistance(current, line)
			if dist < shortestDistance {
				shortestDistance = dist
				shortestDistancePair = [2]string{current, line}
			}
		}
	}
	return shortestDistance, shortestDistancePair
}

func RemoveIndex(s [][]string, index int) [][]string {
	return append(s[:index], s[index+1:]...)
}

type BoolFloatTuple struct {
	LocalOther bool
	Distance   float64
}

func uniqueSlice(slice []BoolFloatTuple) []BoolFloatTuple {
	uniqueMap := make(map[BoolFloatTuple]bool)
	result := []BoolFloatTuple{}

	for _, v := range slice {
		if !uniqueMap[v] {
			uniqueMap[v] = true
			result = append(result, v)
		}
	}

	return result
}

func connectClosestCircuits(allCircs [][]string, nth int) [][]string {
	shortestDistance := math.MaxFloat64
	localDistance := math.MaxFloat64
	var (
		distanceList        []BoolFloatTuple
		shortestDistanceIdx [2]int
	)
	for focusedIdx, focusedCircuit := range allCircs {
		for _, focusedPos := range focusedCircuit {
			for _, searchPos := range focusedCircuit {
				dist := calcDistance(focusedPos, searchPos)
				if dist != 0 {
					distanceList = append(distanceList, BoolFloatTuple{true, dist})
				}
				if dist < localDistance && dist != 0 {
					localDistance = dist
				}
			}
			for searchIdx, searchCircuit := range allCircs[focusedIdx+1:] {
				for _, searchPos := range searchCircuit {
					dist := calcDistance(focusedPos, searchPos)
					distanceList = append(distanceList, BoolFloatTuple{false, dist})
					if dist < shortestDistance {
						shortestDistance = dist
						shortestDistanceIdx = [2]int{focusedIdx, searchIdx + focusedIdx + 1}
					}
				}
			}
		}
	}
	slices.SortStableFunc(distanceList, func(a, b BoolFloatTuple) int { return cmp.Compare(a.Distance, b.Distance) })
	distanceList = uniqueSlice(distanceList)
	if !distanceList[nth].LocalOther {
		newCirc := append(allCircs[shortestDistanceIdx[0]], allCircs[shortestDistanceIdx[1]]...)
		allCircs[shortestDistanceIdx[0]] = newCirc
		allCircs = RemoveIndex(allCircs, shortestDistanceIdx[1])
	}
	return allCircs
}

func solvePart1(lines []string) int {
	var circuits [][]string
	for _, line := range lines {
		circuits = append(circuits, []string{line})
	}
	for nth := 1; nth <= 1000; nth++ {
		circuits = connectClosestCircuits(circuits, nth)
	}
	slices.SortStableFunc(circuits, func(a, b []string) int { return cmp.Compare(len(b), len(a)) })
	result := 1
	lastLenght := 0
	i := 0
	for _, circuit := range circuits {
		if i > 2 {
			break
		}
		if lastLenght != len(circuit) {
			i++
			lastLenght = len(circuit)
			result *= len(circuit)
		}
	}
	return result
}

func solvePart2(lines []string) int {
	return 0
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Day 8 ===")
	fmt.Println("Part 1:", solvePart1(lines))
	fmt.Println("Part 2:", solvePart2(lines))
}
