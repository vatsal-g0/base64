package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func ToBase64(s string) (string, error) {
	bytes := []byte(s)

	var binstr strings.Builder

	for _, b := range bytes {
		fmt.Fprintf(&binstr, "%08b", b)
	}

	binary := binstr.String()

	for len(binary)%6 != 0 {
		binary += "0"
	}

	var result strings.Builder

	for i := 0; i < len(binary); i += 6 {
		chunk := binary[i : i+6]

		n, err := strconv.ParseInt(chunk, 2, 8)
		if err != nil {
			return "", err
		}

		result.WriteByte(base64Table[n])
	}

	padding := (3 - len(bytes)%3) % 3

	for range padding {
		result.WriteByte('=')
	}

	return result.String(), nil
}

func FromBase64(s string) (string, error) {
	s = strings.TrimRight(s, "=")

	var binary strings.Builder

	for _, ch := range s {
		index := strings.IndexRune(base64Table, ch)

		if index == -1 {
			return "", errors.New("invalid Base64 character")
		}

		bits := fmt.Sprintf("%06b", index)
		binary.WriteString(bits)
	}

	var result strings.Builder

	binStr := binary.String()

	for i := 0; i+8 <= len(binStr); i += 8 {
		byteStr := binStr[i : i+8]

		var value byte

		for _, bit := range byteStr {
			value <<= 1

			if bit == '1' {
				value |= 1
			}
		}

		result.WriteByte(value)
	}

	return result.String(), nil
}

func PrintUsage() {
	fmt.Printf("usage: %s <e|d> <string>\n", os.Args[0])
}

func main() {
	if len(os.Args) != 3 {
		PrintUsage()
		os.Exit(1)
	}

	option := os.Args[1]
	text := os.Args[2]

	switch strings.ToLower(option) {
	case "e":
		encoded, err := ToBase64(text)
		if err != nil {
			panic(err)
		}
		
		fmt.Println(encoded)
	
	case "d":
		decoded, err := FromBase64(text)
		if err != nil {
			panic(err)
		}

		fmt.Println(decoded)
	
	default:
		PrintUsage()
		os.Exit(1)
	}
}
