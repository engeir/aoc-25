package main

import (
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/engeir/aoc-25/utils"
)

type MachineConfig struct {
	Lights  []int
	Buttons [][]int
	Joltage []int
}

const (
	lights  = `\[.*\]`
	button  = `\(.*?\)`
	joltage = `\{.*\}`
)

var (
	reLights  = regexp.MustCompile(lights)
	reButtons = regexp.MustCompile(button)
	reJoltage = regexp.MustCompile(joltage)
)

func parseConfig(line string) MachineConfig {
	// Extract lights
	lightsIdx := reLights.FindAllStringIndex(line, -1)
	lght := line[lightsIdx[0][0]+1 : lightsIdx[0][1]-1]
	var lights []int
	for _, l := range lght {
		switch l {
		case '.':
			lights = append(lights, 0)
		case '#':
			lights = append(lights, 1)
		}
	}
	// Extract buttons
	buttonsIdx := reButtons.FindAllString(line, -1)
	var buttons [][]int
	for _, btnIdx := range buttonsIdx {
		elements := strings.Split(btnIdx[1:len(btnIdx)-1], ",")
		var numbers []int
		for _, e := range elements {
			i, err := strconv.Atoi(e)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, i)
		}
		buttons = append(buttons, numbers)
	}
	// Extract joltage
	joltageIdx := reJoltage.FindAllStringIndex(line, -1)
	joltageStr := line[joltageIdx[0][0]+1 : joltageIdx[0][1]-1]
	jolt := strings.Split(joltageStr, ",")
	var joltages []int
	for _, e := range jolt {
		i, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err)
		}
		joltages = append(joltages, i)
	}
	return MachineConfig{Lights: lights, Buttons: buttons, Joltage: joltages}
}

func getConfig(lines []string) []MachineConfig {
	var machineConfigs []MachineConfig
	for _, line := range lines {
		machineConfigs = append(machineConfigs, parseConfig(line))
	}
	return machineConfigs
}

func toggleLightInt(light []int, idx int) []int {
	light[idx] = (light[idx] + 1) % 2
	return light
}

func toggleLightStr(light string, idx int) string {
	b := []byte(light)
	switch b[idx] {
	case '.':
		b[idx] = '#'
	case '#':
		b[idx] = '.'
	}
	return string(b)
}

// generateCombinations generates all slices of length n with exactly k ones
func generateCombinations(n, k int) [][]int {
	var result [][]int
	current := make([]int, n)
	var backtrack func(pos, count int)

	backtrack = func(pos, count int) {
		// If we've placed k ones, add the current combination
		if count == k {
			combo := make([]int, n)
			copy(combo, current)
			result = append(result, combo)
			return
		}
		// If we've reached the end or can't place more ones, return
		if pos == n || count > k {
			return
		}
		// Try placing a 1 at the current position
		current[pos] = 1
		backtrack(pos+1, count+1)
		// Backtrack: remove the 1 and try with 0
		current[pos] = 0
		backtrack(pos+1, count)
	}

	backtrack(0, 0)
	return result
}

func pressButtons(mc MachineConfig, pressButton []int) []int {
	lights := make([]int, len(mc.Lights))
	for i, press := range pressButton {
		if press == 1 {
			for _, push := range mc.Buttons[i] {
				lights = toggleLightInt(lights, push)
			}
		}
	}
	return lights
}

func solvePart1(lines []string) int {
	machineConfigs := getConfig(lines)
	count := 0
	for _, mc := range machineConfigs {
		presses := 1
		var newLights []int
		for {
			combinations := generateCombinations(len(mc.Buttons), presses)
			for _, comb := range combinations {
				newLights = pressButtons(mc, comb)
				if slices.Compare(newLights, mc.Lights) == 0 {
					break
				}
			}
			if slices.Compare(newLights, mc.Lights) == 0 {
				break
			}
			presses += 1
		}
		count += presses
	}
	return count
}

func solvePart2(lines []string) int {
	return 0
}

func main() {
	lines, err := utils.ReadLines("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Day 10 ===")
	fmt.Println("Part 1: ", solvePart1(lines))
	fmt.Println("Part 2: ", solvePart2(lines))
}
