package main

import (
	"GoPadai/Shiva/bitOperations/bitops"
)

func main() {

	//to check the lsb
	bitops.CheckLSB(18)

	//to check the msb of numbers 0 and greater
	bitops.CheckMSB(0)

	//to check the status of the bit at position (0-31) from right
	bitops.CheckBitStatus(18, 0)

	//to set the bit status at position of give number
	bitops.SetBitStatus(4, 0)

	//to unset(clear) the bit status at position of give number
	bitops.UnsetBitStatus(11, 1)
}
