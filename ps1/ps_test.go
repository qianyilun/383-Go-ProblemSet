/*
* This file is to test Problem Set 1 code
* CMPT 383
*
* @Yilun Qian
* ps_test.go
 */

package ps1

import (
	"testing"
	"fmt"
	"os"
	"log"
)

// Q1
func Test_countPrimes(t *testing.T){
	testValue := map[int] int{-6: 0, 4: 2, 109: 29, 1000: 168, 10000: 1229}

	// The following value is NOT correct, the right answer should be 25
	testExValue := map[int] int{100: 24}

	for key := range testValue {
		value := countPrimes(key)
		if (testValue[key] == value){
			fmt.Println("Correct so far!")
		} else {
			t.Errorf("Right value should be:", testValue[key], "instead of:", value)
		}
	}

	fmt.Println("Testing wrong answer case...")
	if (testExValue[100] != countPrimes(100)){
		fmt.Println("Correct so far!")
	} else {
		t.Errorf("Right value should be:", 25, "instead of:", 24)
	}
}

// Q2
func Test_countStrings(t *testing.T){
	// Create test files
	s1 := "The big big dog\nate the big apple"
	s2 := "This is a test string\n" +
			"a test string with a lot of lines\n" +
			"a lot of lines\n" +
			"a lot of lines\n" +
			"a lot of lines\n" +
			"a lot of lines\n" +
			"a lot of lines\n" +
			"a lot of lines\n" +
			"a lot of lines\n" +
			"The end of string"
	s3 := "1"

	createFiles("Q2-1.txt", s1)
	createFiles("Q2-2.txt", s2)
	createFiles("Q2-3.txt", s3)

	map1 := countStrings("Q2-1.txt")
	map2 := countStrings("Q2-2.txt")
	map3 := countStrings("Q2-3.txt")

	if len(map1) != 6 {
		t.Errorf("%v words expected but the program counts %v words. ", 6, len(map1))
	}
	if len(map2) != 11 {
		t.Errorf("%v words expected but the program counts %v words. ", 11, len(map2))
	}
	if len(map3) != 1 {
		t.Errorf("%v words expected but the program counts %v words. ", 1, len(map3))
	}
}

// Q2 helper function for creating test text files
func createFiles(filename string, s string){
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(s)
	file.Close()
}

// Q3 + Q4
func Test_Time24_funcs(t *testing.T) {
	t1, t2, t3 := Time24{25, 30, 30}, Time24{1, 70, 30}, Time24{12, 60, 90}
	t4, t5, t6 := Time24{12, 30, 59}, Time24{12, 12, 12}, Time24{25, 60, 60}

	t7, t8 := Time24{0, 0, 0}, Time24{0, 0, 0}
	/*
	* Testing validTime24...
	* Output:
	** false false false
	** true true false
	*/
	fmt.Println("Testing validTime24...")
	fmt.Printf( "%v %v %v\n", validTime24(t1), validTime24(t2), validTime24(t3) )
	fmt.Printf( "%v %v %v\n", validTime24(t4), validTime24(t5), validTime24(t6) )

	/*
	* Testing equalsTime24...
	* Output:
	** false
	** true
	*/
	fmt.Println("Testing equalsTime24...")
	fmt.Println(equalsTime24(t4, t5))
	fmt.Println(equalsTime24(t7, t8))

	/*
	* Testing lessThanTime24...
	* Output:
	** false
	** true
	** false
	*/
	fmt.Println("Testing lessThanTime24...")
	fmt.Println(lessThanTime24(t4, t5))
	fmt.Println(lessThanTime24(t5, t4))
	fmt.Println(lessThanTime24(t7, t8))

	/*
	* Testing minTime24...
	* Output:
	** {00:00:01} <nil>
	** {12:39:40} <nil>
	** {00:00:00} The input slice is empty!
	*/
	tArr := []Time24{ {12,45,30},{23,19,8},{0,30,45},{0,0,1},{0,0,2} }
	tArr2 := []Time24{ {12,39,40} }
	fmt.Println("Testing minTime24...")
	fmt.Println(minTime24(tArr))
	fmt.Println(minTime24(tArr2))
	fmt.Println(minTime24(make([]Time24, 0)) )

	/*
	* Testing AddOneHour() method
	* Output
	** 21:15:00 13:30:59 13:12:12
	** 22:15:00 14:30:59 14:12:12
	** 23:15:00 15:30:59 15:12:12
	** 00:15:00 16:30:59 16:12:12
	** 01:15:00 17:30:59 17:12:12
	 */

	time := Time24{20, 15, 0}
	fmt.Println(time)
	for i := 0; i < 5; i++ {
		time.AddOneHour()
		t4.AddOneHour()
		t5.AddOneHour()
		fmt.Println(time, t4, t5)
	}
}

// Q5
/*
* Output
** []
** []
** [[1] [0]]
** [[1 1] [1 0] [0 1] [0 0]]
** [[1 1 1 1] [1 1 1 0] [1 1 0 1] [1 1 0 0] [1 0 1 1] [1 0 1 0] [1 0 0 1] [1 0 0 0] [0 1 1 1] [0 1 1 0] [0 1 0 1] [0 1 0 0] [0 0 1 1] [0 0 1 0] [0 0 0 1] [0 0 0 0]]
 */
func Test_tester(t *testing.T)  {
	t1 := allBitSeqs(-6)
	t2 := allBitSeqs(0)
	t3 := allBitSeqs(1)
	t4 := allBitSeqs(2)
	t5 := allBitSeqs(4)

	fmt.Println( t1 )
	fmt.Println( t2 )
	fmt.Println( t3 )
	fmt.Println( t4 )
	fmt.Println( t5 )
}


