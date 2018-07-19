package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// import "math/rand"

func generateRandomString(size int) string {

	rand.Seed(time.Now().UTC().UnixNano())

	availableLetter := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var response string
	var i int
	for i = 0; i < size; i++ {
		response = response + availableLetter[rand.Intn(len(availableLetter))]
	}

	//fmt.Println(response)

	return response
}

func main() {

	startTime := time.Now()
	// Block Header
	blockHeader := "A6gsk08HTRyqVKLPLcdrRA90QA"

	randomStringSize := 30

	Found := false

	for Found == false {

		nounce := generateRandomString(randomStringSize)
		challenge := blockHeader + nounce
		sha256Value := sha256.Sum256([]byte(challenge))
		sha256ValueString := fmt.Sprintf("%x", sha256Value)

		if strings.HasPrefix(sha256ValueString, "00000") {

			timeTook := time.Now().Sub(startTime)

			fmt.Println("256 Hash : ", sha256ValueString)
			fmt.Println("Nounce : ", nounce)
			fmt.Println("BlockHeader : ", blockHeader)
			fmt.Println("Time Took : ", timeTook)
			Found = true
		}
	}

}
