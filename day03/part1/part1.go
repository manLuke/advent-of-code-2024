package part1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 3 - Part 1")
	result := getResultOfMultiplication(getMemory())
	fmt.Println("Result:", result)
}

func getResultOfMultiplication(src string) int {
	var validOrder = [6]TokenType{Mul, OpenParen, Int, Comma, Int, CloseParen}
	tokens := tokenize(src)
	i := 1
	sum := 0

	for i < len(tokens) {
		match := true
		for j, tokenType := range validOrder {
			if i+j >= len(tokens) || tokens[i+j].Type != tokenType {
				match = false
			}
		}
		if match {
			sum += tokens[i+2].Value * tokens[i+4].Value
			i += len(validOrder)
		} else {
			i++
		}
	}

	return sum
}

func tokenize(src string) []Token {
	i := 0
	var tokens []Token

	for i < len(src) {
		newToken := parseToken(src, &i)
		tokens = append(tokens, newToken)
		i++
	}

	return tokens
}

func parseToken(src string, i *int) Token {
	if token, status := tokenizeInt(src, i); status {
		return token
	} else if token, status := tokenizePattern(src, i, "mul", Mul); status {
		return token
	} else if token, status := tokenizePattern(src, i, "(", OpenParen); status {
		return token
	} else if token, status := tokenizePattern(src, i, ",", Comma); status {
		return token
	} else if token, status := tokenizePattern(src, i, ")", CloseParen); status {
		return token
	}
	return Token{Type: Invalid}
}

func tokenizeInt(src string, index *int) (Token, bool) {
	numString := ""

	for *index < len(src) && isNumber(src[*index]) {
		numString = numString + string(src[*index])
		*index++
	}

	if len(numString) > 0 {
		n, err := strconv.Atoi(numString)
		common.Check(err)
		*index--
		return Token{Type: Int, Value: n}, true
	}

	return Token{}, false
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func tokenizePattern(src string, index *int, pattern string, tokenType TokenType) (Token, bool) {
	if *index+len(pattern) > len(src) {
		return Token{}, false
	}

	for i := 0; i < len(pattern); i++ {
		if *index+i > len(src) || pattern[i] != src[*index+i] {
			return Token{}, false
		}
	}

	*index += len(pattern) - 1
	return Token{Type: tokenType}, true
}

type TokenType int

type Token struct {
	Type  TokenType
	Value int
}

func (t TokenType) String() string {
	switch t {
	case Invalid:
		return "Invalid"
	case Mul:
		return "Mul"
	case OpenParen:
		return "OpenParen"
	case Int:
		return "Int"
	case Comma:
		return "Comma"
	case CloseParen:
		return "CloseParen"
	default:
		return "Unknown"
	}
}

const (
	Invalid TokenType = iota
	Mul
	OpenParen
	Int
	Comma
	CloseParen
)

func getMemory() string {
	file, err := os.Open("./day03/input.txt")
	common.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var src string
	for scanner.Scan() {
		line := scanner.Text()
		src += line
	}

	return src
}
