/*
* This file is to solve Problem Set 1
* CMPT 383
*
* @Yilun Qian
* ps.go
 */

package ps1

import (
	"io/ioutil"
	"fmt"
	"errors"
	"strconv"
	"strings"
)

// Q1
func countPrimes(n int) int {
	count := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++;
		}
	}
	return count
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if (n % i == 0) {
			return false
		}
	}
	return true
}

// Q2
func countStrings(filename string) map[string] int{

	stream, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("no word file to read!")
	}

	readString := string(stream)

	// split string into an array by multiple delimitters
	// reference: https://groups.google.com/forum/#!topic/golang-nuts/SCSDr8O20CQ
	words := strings.FieldsFunc(readString, func(r rune) bool{
		switch r {
		case '\n', ' ':
			return true
		}
		return false
	})

	stringMap := make(map[string] int)

	// initialization
	for _, s := range words {
		stringMap[s] = 1
	}

	count := 0
	for i, s := range words {
		for j := i + 1; j < len(words); j++ {
			if s == words[j] {
				count++
				stringMap[s] = count
			}
		}
	}
	return stringMap
}



//Q3
type Time24 struct {
	hour, minute, second uint8
}

func validTime24(t Time24) bool {
	if t.hour >= 0 && t.hour < 24 {
		if t.minute >= 0 && t.minute < 60 {
			if t.second >= 0 && t.second < 60 {
				return true
			}
		}
	}
	return false
}

func equalsTime24(t1 Time24, t2 Time24) bool {
	if (!validTime24(t1) || !validTime24(t2)) {
		fmt.Println("Check input! They were invalid!")
		return false
	}
	if t1.hour == t2.hour && t1.minute == t2.minute && t1.second == t2.second {
		return true
	}
	return false
}

func lessThanTime24(t1 Time24, t2 Time24) bool {
	if (!validTime24(t1) || !validTime24(t2)) {
		fmt.Println("Check input! They were invalid!")
		return false
	}

	if t1.hour > t2.hour {
		return false
	}
	if t1.hour < t2.hour {
		return true
	}

	if t1.minute > t2.minute {
		return false
	}
	if t1.minute < t2.minute {
		return true
	}

	if t1.second >= t2.second {
		return false
	}
	return true
}

func minTime24(times []Time24) (Time24, error)  {
	small := Time24{0, 0, 0}
	if (len(times) == 0) {
		return small, errors.New("The input slice is empty!")
	}
	small = times[0]
	for _, val := range times {
		if (!validTime24(val)) {
			return small, errors.New("One of element of input slice is invalid!")
		}
		if (lessThanTime24(val, small)) {
			small = val
		}
	}
	return small, nil
}

// Q4 method-1
func (t Time24) String() string  {
	hour := strconv.Itoa(int(t.hour))
	minute := strconv.Itoa(int(t.minute))
	second := strconv.Itoa(int(t.second))

	tToString := ""
	if (int(t.hour) < 10) {
		tToString = "0"
	}
	tToString += hour + ":"

	if (int(t.minute) < 10) {
		tToString += "0"
	}
	tToString += minute + ":"

	if (int(t.second) < 10) {
		tToString += "0"
	}
	tToString += second

	return tToString
}

// Q4 method-2
// This should be passing by reference
func (t *Time24) AddOneHour()  {
	if (t.hour == 23){
		t.hour = 0
	} else {
		t.hour++
	}
}

// Q5
func allBitSeqs(n int) [][]int {
	if n <= 0 {
		return make([][]int, 0)
	}

	// get the largest number that 3-digits can be presented on base 2
	N := 1
	for i := 0; i < n; i++ {
		N *= 2
	}
	seqs := [][]int{}
	largestInt := N - 1

	// adding the array that are all 1's
	row := sameBitLengthHelper(largestInt, n)
	seqs = append(seqs, row)

	tempNum := largestInt
	for i := 0; i < largestInt; i++ {
		tempNum --
		tempBinary := strconv.FormatInt(int64(tempNum), 2)
		tempLen := len(tempBinary)

		if tempLen == len(row) {
			seqs = append(seqs, sameBitLengthHelper(tempNum, n))
		} else {
			seqs = append(seqs, differBitLengthHelper(n, tempBinary, tempLen))
		}
	}
	return seqs
}

// The function is to convert number on base 2 to array if lengths are the same
func sameBitLengthHelper(input int, n int) []int {
	inputBinaryS := string(strconv.FormatInt(int64(input), 2))
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		numS, _ := strconv.Atoi(string(inputBinaryS[i]))
		arr[i] = numS
	}
	return arr
}

// This function is to convert number on base 2 to array if lengths are NOT the same
func differBitLengthHelper(n int, s string, length int) []int {
	tempArr := make([]int, length)
	for i := 0; i < length; i++ {
		numS, _ := strconv.Atoi(string(s[i]))
		tempArr[i] = numS
	}
	// assign tempArr to resultArr, meanwhile, adding 0 to fill the length
	result := make([]int, n)
	index := n-length
	for i := 0; i < length; i++ {
		result[index] = tempArr[i]
		index++
	}
	return result
}