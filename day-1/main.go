package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var FILE = "./data/input.txt"

func main() {
	// part1()
	part2()
}

func part2() {
	lValues, rValues, err := getLists()
	if err != nil {
		log.Fatalf("unable to get lists %s", err)
		return
	}
	simScore := 0
	lookupValues := make(map[int]int, 0)
	for _, l := range lValues {
		if lookupValues[l] != 0 {
			simScore += l * lookupValues[l]
			continue
		}
		occurances := countOccurances(rValues, l)
		if occurances > 0 {
			lookupValues[l] = occurances
			simScore += l * occurances
		}
	}
	log.Printf("got sim score of %d", simScore)
}

func countOccurances(values []int, value int) int {
	count := 0
	for _, i := range values {
		if i == value {
			count++
		}
	}
	return count
}

func part1() {
	lValues, rValues, err := getLists()
	if err != nil {
		log.Fatalf("unable to get lists %s", err)
		return
	}
	slices.Sort(lValues)
	slices.Sort(rValues)

	var sum int
	for i := 0; i < len(lValues); i++ {
		l := lValues[i]
		r := rValues[i]
		diff := 0
		if l > r {
			diff = l - r
		} else {
			diff = r - l
		}
		sum += diff
		fmt.Printf("checking l: %d, r: %d, diff:%d \n", l, r, diff)
	}
	fmt.Printf("got sum of %d \n", sum)
}

func getLists() ([]int, []int, error) {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatalf("unable to open file %s", err)
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lValues := make([]int, 0)
	rValues := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		l, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Fatalf("unable to convert left value to int")
			return nil, nil, err
		}
		r, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Fatalf("unable to convert right value to int")
			return nil, nil, err
		}
		lValues = append(lValues, l)
		rValues = append(rValues, r)
	}
	return lValues, rValues, nil
}
