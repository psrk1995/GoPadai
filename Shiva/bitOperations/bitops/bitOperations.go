package bitops

import (
	"fmt"
	"math/bits"
)

/*
CheckLSB recieves a number as int and
checks Least Significant Bit (LSB) of a number using bitwise operator.

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
		fmt.Printf("\nLSB of %d is set (1)", num)
	} else {
		fmt.Printf("\nLSB of %d is unset (0)", num)
	}
}

/*
CheckMSB recieves a number as int and
checks Most Significant Bit (MSB) of a number using bitwise operator.

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
	bitLength := bits.Len(num)
	var msb uint

	if bitLength > 0 {
		msb = 1 << (bits.Len(num) - 1)
	} else {
		msb = 0
	}

	if num&msb == 0 {
		fmt.Printf("\nMSB of %d is unset (0)", msb)
	} else {
		fmt.Printf("\nMSB of %d is set (1)", msb)
	}
}

/*
CheckBitStatus recieves a number as int and
check status of bit at the given position of number using bitwise operator.

1. Input a number from user. Store it in some variable say num and
 position of the bit which status has to be found as bitPosition

  => Let the number be 5 (101) and position as 1 i.e. of 0

2 Find number of bits used to represent the given number in memory.
	Use bits.Len() to find the bit length.

3. To get the status of bit, move first bit of 1 to highest order.
	 Left shift 1 bitPosition-1 times and store result in some variable
	 say msb = 1 << (bitPosition - 1)

	 msb = 1 << 1 results in msb = 2 (10)

4. If bitwise AND operation num & msb evaluate not equal to 0 then MSB of num is set otherwise not.
			101
		& 010
		------
		  000
		------
*/
func CheckBitStatus(num uint, bitPosition int) {
	bitStatus := num >> bitPosition
	if bitStatus&1 == 0 {
		fmt.Printf("\nStatus of bit at index %d of %d is unset (0)", bitPosition, num)
	} else {
		fmt.Printf("\nStatus of bit at index %d of %d is set (1)", bitPosition, num)
	}
}

/*
SetBitStatus recieves a number as int and
set status of bit at the given position of number to 1 (set) using bitwise operator.

1. Input a number from user. Store it in some variable say num and
 position of the bit which status has to be found as bitPosition

  => Let the number be 5 (101) and position as 1 i.e. of 0

2. Move first bit of 1 to highest order.
	 Left shift 1 bitPosition times
	 1 << bitPosition => 1 << 1  = 10

4. Perform bitwise OR on given number and value obtained after shifting.
			101 - 5
		| 010 - 2
		------
		  111 - 7
		------
*/
func SetBitStatus(num uint, bitPosition int) {
	newNum := 1<<bitPosition | num
	fmt.Printf("\nNumber before setting %d bit: %d (in decimal)", bitPosition, num)
	fmt.Printf("\nNumber after setting %d bit: %d (in decimal)", bitPosition, newNum)
}

/*
UnsetBitStatus recieves a number as int and
unset status of bit at the given position of number to 1 (set) using bitwise operator.

1. Input number and nth bit position to clear from user. Store it in some variable say num and bitPosition.
2. Left shift 1, n times i.e. 1 << n.
3. Perform bitwise complement with the above result. So that the nth bit becomes unset and rest of bit becomes set i.e. ~ (1 << n).
4. Finally, perform bitwise AND operation with the above result and num. The above three steps together can be written as num & (~ (1 << n));
*/
func UnsetBitStatus(num uint, bitPosition int) {
	newNum := num & (^(1 << bitPosition))
	fmt.Printf("\nNumber before unsetting (clear) %d bit: %d (in decimal)", bitPosition, num)
	fmt.Printf("\nNumber after unsetting (clear) %d bit: %d (in decimal)", bitPosition, newNum)
}
