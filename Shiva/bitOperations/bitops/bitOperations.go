package bitops

import (
	"fmt"
	"math/bits"
)

//CheckLSB recieves a number as int and
//checks Least Significant Bit (LSB) of a number using bitwise operator.
/*
To check LSB of a number we need to perform bitwise ANDing.
The bitwise AND operation number & 1 will evaluate to 1 if LSB of number is set i.e. 1
otherwise evaluates to 0.

Even odd using bitwise operator

            12 ------------ decimal ---------------- 15
		 0000 1100 ------------  binary ---------------- 0000 1111
		 0000 0001 ----------  binary of 1 ------------- 0000 0001
		 ---------                                       ---- ----
		 0000 0000                                       0000 0001

As you can see above 12 & 1 evaluate to 0. Since, LSB of 12 is 0.
Whereas, 15 & 1 evaluate to 1 since LSB of 15 is 1.
*/
func CheckLSB(num int) {
	if num&1 == 1 {
		fmt.Printf("LSB of %d is set (1)", num)
	} else {
		fmt.Printf("LSB of %d is unset (0)", num)
	}
}

//checkMSB recieves a number as int and
//checks Most Significant Bit (MSB) of a number using bitwise operator.
/*
Step by step descriptive logic to check MSB of a number.

1. Input a number from user. Store it in some variable say num.
  => Let the number be 5 (101)

2 Find number of bits used to represent the given number in memory.
	Use bits.Len() to find the bit length.

3. To get MSB of the number, move first bit of 1 to highest order.
	 Left shift 1 bits - 1 times and store result in some variable
	 say msb = 1 << (bits - 1) and bits = 3

	 msb = 1 << 2 results in msb = 4 (100)

4. If bitwise AND operation num & msb evaluate not equal to 0 then MSB of num is set otherwise not.
			101
		& 100
		------
		 1000
		------
*/
func CheckMSB(num uint) {
	var b uint = 1 << (bits.Len(num) - 1)

	if num&1 == 1 {
		fmt.Printf("LSB of %d is set (1)", b)
	} else {
		fmt.Printf("LSB of %d is unset (0)", b)
	}
}
