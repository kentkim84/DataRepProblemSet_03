package main

import (
	//"bytes"
	"fmt"
	"bufio"
	"os"
	"math/rand"
	"time"
	"regexp"
)

// generate a random number
func genRandInt(minNum int, maxNum int) int {
    return rand.Intn(maxNum) + minNum
}

func elizaResponse(message string) (string) {
	var response string
	responseList := []string{
		"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
		"Why don’t you tell me more about your father?"}

	randNum := genRandInt(0, 2)
	
	r, _ := regexp.Compile("\b[Ff](ather)")
	if r.MatchString(message) {
			response = "Why don’t you tell me more about your father?"
	} else {
		r, _ = regexp.Compile("[Ii](m|[\t\n\f\r ][Aa][Mm]|[^\t\n\f\r ]m)")
		if r.MatchString(message){
			response = r.ReplaceAllString(message, "How do you know you are")
		} else {
			response = responseList[randNum]
		}
	}

	return response
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Say something: ")
	message, _ := reader.ReadString('\n')
	fmt.Printf("%s", elizaResponse(message))
}