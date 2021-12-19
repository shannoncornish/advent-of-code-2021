package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func floor(n int) int {
	return int(math.Floor(float64(n) / 2))
}

func ceil(n int) int {
	return int(math.Ceil(float64(n) / 2))
}

type Token struct {
	Kind   byte
	Number int
}

var (
	PairOpen      = byte('[')
	PairSeparator = byte(',')
	PairClose     = byte(']')
	Number        = byte('N')
)

type Tokens []Token

func (ts Tokens) string() string {
	var sb strings.Builder
	for _, t := range ts {
		switch t.Kind {
		case PairOpen, PairSeparator, PairClose:
			sb.WriteByte(t.Kind)
			break
		default:
			sb.WriteString(strconv.Itoa(t.Number))
			break
		}
	}

	return sb.String()
}

func add(left, right Tokens) Tokens {
	var ts Tokens
	ts = append(ts, Token{
		Kind: PairOpen,
	})
	ts = append(ts, left...)
	ts = append(ts, Token{
		Kind: PairSeparator,
	})
	ts = append(ts, right...)
	ts = append(ts, Token{
		Kind: PairClose,
	})

	return ts
}

type Action struct {
	Kind  byte
	Index int
}

var (
	Explode = byte('E')
	Split   = byte('S')
)

func reduce(tokens Tokens) Tokens {
	reduced := &tokens

	var action *Action
	for {
		reducing := *reduced

		action = nil

		var pairs int
		for i, t := range reducing {
			if t.Kind == PairOpen {
				pairs++
			}

			if t.Kind == PairClose {
				pairs--
			}

			if pairs > 4 {
				action = &Action{
					Kind:  Explode,
					Index: i,
				}
				break
			}
		}

		if action != nil {
			left := reducing[action.Index+1]
			for leftIndex := action.Index - 1; leftIndex >= 0; leftIndex-- {
				if reducing[leftIndex].Kind == Number {
					reducing[leftIndex].Number += left.Number
					break
				}
			}

			right := reducing[action.Index+3]
			for rightIndex := action.Index + 5; rightIndex < len(reducing); rightIndex++ {
				if reducing[rightIndex].Kind == Number {
					reducing[rightIndex].Number += right.Number
					break
				}
			}

			var modified Tokens
			modified = append(modified, reducing[0:action.Index]...)
			modified = append(modified, Token{
				Kind:   Number,
				Number: 0,
			})
			modified = append(modified, reducing[action.Index+5:]...)

			reduced = &modified
			continue
		}

		for i, t := range reducing {
			if t.Kind == Number && t.Number >= 10 {
				action = &Action{
					Kind:  Split,
					Index: i,
				}
				break
			}
		}

		if action != nil {
			t := reducing[action.Index]

			var modified Tokens
			modified = append(modified, reducing[0:action.Index]...)
			modified = append(modified, Token{
				Kind: '[',
			})
			modified = append(modified, Token{
				Kind:   Number,
				Number: floor(t.Number),
			})
			modified = append(modified, Token{
				Kind: ',',
			})
			modified = append(modified, Token{
				Kind:   Number,
				Number: ceil(t.Number),
			})
			modified = append(modified, Token{
				Kind: ']',
			})
			modified = append(modified, reducing[action.Index+1:]...)

			reduced = &modified
			continue
		}

		if action == nil {
			break
		}
	}

	return *reduced
}

func magnitude(tokens Tokens) int {
	reduced := &tokens

	for {
		reducing := *reduced

		var leafIndex int

		var pairs int
		for i, t := range reducing {
			switch t.Kind {
			case PairOpen:
				pairs++
				break
			case PairClose:
				pairs--
				break
			case PairSeparator:
				if pairs > 1 {
					if reducing[i-1].Kind == Number &&
						reducing[i+1].Kind == Number {
						leafIndex = i - 2
					}
				}
				break
			}

			if leafIndex > 0 {
				break
			}
		}

		if leafIndex > 0 {
			left := reducing[leafIndex+1]
			right := reducing[leafIndex+3]

			magnitude := 3*left.Number + 2*right.Number

			var modified Tokens
			modified = append(modified, reducing[0:leafIndex]...)
			modified = append(modified, Token{
				Kind:   Number,
				Number: magnitude,
			})
			modified = append(modified, reducing[leafIndex+5:]...)

			reduced = &modified
			continue
		} else {
			break
		}
	}

	left := (*reduced)[1]
	right := (*reduced)[3]

	return 3*left.Number + 2*right.Number
}

type Tuple struct {
	i, j int
}

func main() {
	tokens := parseInput()

	magnitudes := make(map[Tuple]int)

	for i := 0; i < len(tokens); i++ {
		for j := 0; j < len(tokens); j++ {
			if i == j {
				continue
			}

			magnitudes[Tuple{i, j}] = magnitude(reduce(add(tokens[i], tokens[j])))
		}
	}

	var largest int
	for _, magnitude := range magnitudes {
		if magnitude > largest {
			largest = magnitude
		}
	}

	fmt.Printf("The largest magnitude of any sum: %d\n", largest)

	// Answer: 4807
}

func parseNumber(s string) Tokens {
	var tokens Tokens
	for i := 0; i < len(s); i++ {
		r := s[i]
		switch r {
		case '[', ',', ']':
			tokens = append(tokens, Token{
				Kind: r,
			})
			break
		default:
			j := i + 1
			for ; j < len(s); j++ {
				var stop bool
				switch s[j] {
				case '[', ',', ']':
					stop = true
					break
				}

				if stop {
					break
				}
			}
			n, _ := strconv.Atoi(s[i:j])
			tokens = append(tokens, Token{
				Kind:   Number,
				Number: n,
			})
			i = j - 1
			break
		}
	}

	return tokens
}

func parseInput() []Tokens {
	name := "input.txt"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	log.Printf("Input file: %s\n", name)

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var tokens []Tokens
	for scanner.Scan() {
		tokens = append(tokens, parseNumber(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return tokens
}
