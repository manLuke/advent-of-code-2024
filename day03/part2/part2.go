package part2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/manLuke/advent-of-code-2024/common"
)

func Solve() {
	fmt.Println("Advent of Code - Day 3 - Part 2")
	result := processCorruptedMemory(getMemory())
	fmt.Println("Result:", result)
}

func processCorruptedMemory(src string) int {
	tokens := tokenize(src)
	return calculateEnabledMultiplications(tokens)
}

func calculateEnabledMultiplications(tokens []Token) int {
	i := 1
	sum := 0
	enabled := true

	for i < len(tokens) {
		switch {
		case isTokenSequence(tokens, i, []TokenType{Do, OpenParen, CloseParen}):
			enabled = true
			i += 3
		case isTokenSequence(tokens, i, []TokenType{Do, Not, OpenParen, CloseParen}):
			enabled = false
			i += 4
		case !enabled:
			i++
		case isTokenSequence(tokens, i, []TokenType{Mul, OpenParen, Int, Comma, Int, CloseParen}):
			sum += tokens[i+2].Value * tokens[i+4].Value
			i += 6
		default:
			i++
		}
	}

	return sum
}

func isTokenSequence(tokens []Token, i int, sequence []TokenType) bool {
	for j, expectedType := range sequence {
		if i+j >= len(tokens) || tokens[i+j].Type != expectedType {
			return false
		}
	}
	return true
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
	tokenParsers := []func(string, *int) (Token, bool){
		tokenizeInt,
		func(s string, idx *int) (Token, bool) {
			return tokenizePattern(s, idx, "mul", Mul)
		},
		func(s string, idx *int) (Token, bool) {
			return tokenizePattern(s, idx, "(", OpenParen)
		},
		func(s string, idx *int) (Token, bool) {
			return tokenizePattern(s, idx, ",", Comma)
		},
		func(s string, idx *int) (Token, bool) {
			return tokenizePattern(s, idx, ")", CloseParen)
		},
		func(s string, idx *int) (Token, bool) {
			return tokenizePattern(s, idx, "do", Do)
		},
		func(s string, idx *int) (Token, bool) {
			return tokenizePattern(s, idx, "n't", Not)
		},
	}

	for _, parser := range tokenParsers {
		if token, ok := parser(src, i); ok {
			return token
		}
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
	case Do:
		return "Do"
	case Not:
		return "Not"
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
	Do
	Not
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
