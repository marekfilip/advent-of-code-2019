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

	transferCode(line)
}

func transferCode(code string) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			codeArr := strArrToIntArr(strings.Split(code, ","))
			codeArr[1] = i
			codeArr[2] = j
			index := 0

		Loop:
			for {
				opcode := codeArr[index]
				val1Idx, val2Idx, outIdx := codeArr[index+1], codeArr[index+2], codeArr[index+3]

				switch opcode {
				case 1:
					codeArr[outIdx] = codeArr[val1Idx] + codeArr[val2Idx]
				case 2:
					codeArr[outIdx] = codeArr[val1Idx] * codeArr[val2Idx]
				case 99:
					break Loop
				}

				index += 4
			}

			if codeArr[0] == 19690720 {
				fmt.Printf("Found %d %d\n", i, j)
			}
		}
	}
}

func strArrToIntArr(s []string) []int {
	var err error
	temp := make([]int, 10000000)

	for i, e := range s {
		temp[i], err = strconv.Atoi(strings.Trim(e, "\n"))
		if err != nil {
			log.Fatalf("Error converting %s to int\n", e)
		}
	}

	return temp
}
