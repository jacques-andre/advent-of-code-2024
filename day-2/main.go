package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var FILE string = "./data/input.txt"

func main() {
	part1()
}

func part1() {
	levels, err := getLevels()
	if err != nil {
		log.Fatalf("error getting lists %s", err)
	}
	count := 0
	for _, level := range levels {
		safe := isLevelSafe(level)
		if safe {
			count++
		}
		log.Printf("safe: %+v, level: %+v", safe, level)
	}
	log.Printf("count: %d", count)
}

func isLevelSafe(level []int) bool {
	levelCpy := make([]int, len(level))
	copy(levelCpy, level)
	isInc := false
	isDec := false
	for i := 0; i < len(levelCpy)-1; i++ {
		currentFloor := levelCpy[i]
		floorNext := levelCpy[i+1]
		diff := math.Abs(float64(currentFloor - floorNext))

		if floorNext-currentFloor > 0 {
			isInc = true
		} else {
			isDec = true
		}

		// if we are doing action, check it is correct
		if isInc && currentFloor > floorNext {
			return false
		}
		if isDec && currentFloor < floorNext {
			return false
		}
		if diff == 0 || diff > 3 {
			return false
		}
	}
	return true
}

func getLevels() ([][]int, error) {
	levels := make([][]int, 0)
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatalf("unable to open file %s", err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		level := make([]int, 0)
		for i := range fields {
			value, err := strconv.Atoi(fields[i])
			if err != nil {
				log.Fatalf("unable to convert value to int %s", err)
				return nil, err
			}
			level = append(level, value)
		}
		levels = append(levels, level)
	}
	return levels, nil
}
