package utils

import (
	"math/rand"
	"time"
)

const hashtag = '#'
const questionmark = '?'

// Seed random. Setting seed to 0 will use time.Now().UnixNano()
func Seed(seed int64) {
	if seed == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
	} else {
		rand.Seed(seed)
	}
}

// ReplaceWithNumbers Replace # with numbers
func ReplaceWithNumbers(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == hashtag {
			bytestr[i] = byte(RandDigit())
		}
	}
	if bytestr[0] == '0' {
		bytestr[0] = byte(rand.Intn(8)+1) + '0'
	}

	return string(bytestr)
}

// ReplaceWithLetters Replace ? with ASCII lowercase letters
func ReplaceWithLetters(str string) string {
	if str == "" {
		return str
	}
	bytestr := []byte(str)
	for i := 0; i < len(bytestr); i++ {
		if bytestr[i] == questionmark {
			bytestr[i] = byte(RandLetter())
		}
	}

	return string(bytestr)
}

// Replace ? with ASCII lowercase letters between a and f
// func replaceWithHexLetters(str string) string {
// 	if str == "" {
// 		return str
// 	}
// 	bytestr := []byte(str)
// 	for i := 0; i < len(bytestr); i++ {
// 		if bytestr[i] == questionmark {
// 			bytestr[i] = byte(randHexLetter())
// 		}
// 	}
//
// 	return string(bytestr)
// }

// RandLetter Generate random lowercase ASCII letter
func RandLetter() rune {
	return rune(byte(rand.Intn(26)) + 'a')
}

// Generate random lowercase ASCII letter between a and f
// func randHexLetter() rune {
// 	return rune(byte(rand.Intn(6)) + 'a')
// }

// RandDigit Generate random ASCII digit
func RandDigit() rune {
	return rune(byte(rand.Intn(10)) + '0')
}

// RandIntRange Generate random integer between min and max
func RandIntRange(min, max int) int {
	if min == max {
		return min
	}
	return rand.Intn((max+1)-min) + min
}

// RandFloat32Range Generate random float32 between min and max
func RandFloat32Range(min, max float32) float32 {
	if min == max {
		return min
	}
	return rand.Float32()*(max-min) + min
}

// RandFloat64Range Generate random float64 between min and max
func RandFloat64Range(min, max float64) float64 {
	if min == max {
		return min
	}
	return rand.Float64()*(max-min) + min
}
