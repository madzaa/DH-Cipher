package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func readKeys() (int64, int64) {
	_, arr := findNumbers()
	g, _ := strconv.ParseInt(arr[0], 10, 32)
	p, _ := strconv.ParseInt(arr[1], 10, 32)
	return g, p
}

func findNumbers() (error, []string) {
	reader := bufio.NewReader(os.Stdin)
	title, _ := reader.ReadString('\n')
	re := regexp.MustCompile(`[^A-Za-z ]\d*`)
	arr := re.FindAllString(title, -1)
	return nil, arr
}

func readSharedSecret() int64 {
	_, arr := findNumbers()
	d, _ := strconv.ParseInt(arr[0], 10, 32)
	return d
}

func encrypt(input string, shift int32) string {
	output := strings.Builder{}
	for _, i := range input {
		if unicode.IsLetter(i) {
			if unicode.IsUpper(i) {
				output.WriteRune((i+shift-65)%26 + 65)
			} else {
				output.WriteRune((i+shift-97)%26 + 97)
			}
		} else {
			output.WriteRune(i)
		}
	}
	return output.String()
}

func decrypt(input string, shift int32) string {
	return encrypt(input, 26-shift)
}

func calculateSecrets(g, p, d int64, b int) (B, S int64) {
	B = 1
	S = 1
	for i := 0; i < b; i++ {
		B = (B * g) % p
	}
	for i := 0; i < b; i++ {
		S = (S * d) % p
	}
	return B, S
}

func promptProposal(S int64) {
	s := S % 26
	marriageProposal := encrypt("Will you marry me?", int32(s))
	fmt.Println(marriageProposal)
	input := getInput()
	if strings.Compare(decrypt(input, int32(s)), "Yeah, okay!") == 0 {
		fmt.Println(encrypt("Great!", int32(s)))
	}
	if strings.Compare(decrypt(input, int32(s)), "Let's be friends.") == 0 {
		fmt.Println(encrypt("What a pity!", int32(s)))
	}
}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}

func main() {
	b := 7
	g, p := readKeys()
	fmt.Println("OK")
	A := readSharedSecret()
	B, S := calculateSecrets(g, p, A, b)
	fmt.Printf("B is %d\n", B)

	promptProposal(S)
}
