package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input")
	if err != nil {
		log.Fatalln("Cant read file")
	}
	defer f.Close()

	r := bufio.NewReader(f)
	line, err := r.ReadString(byte('\n'))
	if err != nil {
		log.Fatalln("Cant read line")
	}

	fmt.Printf("Before: %s\nAfter: %s\n", line, transferCode(line))
}

func transferCode(code string) string {
	var index uint
	codeArr := strArrToIntArr(strings.Split(code, ","))
	codeArr[1] = 12
	codeArr[2] = 2

	for {
		opCode := codeArr[index]
		val1Idx, val2Idx, outIdx := codeArr[index+1], codeArr[index+2], codeArr[index+3]
		valueBefore := codeArr[outIdx]

		switch opCode {
		case 1:
			codeArr[outIdx] = codeArr[val1Idx] + codeArr[val2Idx]
		case 2:
			codeArr[outIdx] = codeArr[val1Idx] * codeArr[val2Idx]
		case 99:
			fmt.Println("Finished")
			return strings.Join(intArrToStrArr(codeArr), ",")
		}

		fmt.Printf("Op: %d\t Val1: %d\t Val2: %d\t OutIdx: %d\t Before: %d\t After: %d\n", opCode, codeArr[val1Idx], codeArr[val2Idx], outIdx, valueBefore, codeArr[outIdx])

		index += 4
	}
}

func strArrToIntArr(s []string) []int {
	var err error
	temp := make([]int, len(s))

	for i, e := range s {
		temp[i], err = strconv.Atoi(strings.Trim(e, "\n"))
		if err != nil {
			log.Fatalf("Error converting %s to int\n", e)
		}
	}

	return temp
}

func intArrToStrArr(i []int) []string {
	temp := make([]string, len(i))

	for i, e := range i {
		temp[i] = strconv.Itoa(e)
	}

	return temp
}
