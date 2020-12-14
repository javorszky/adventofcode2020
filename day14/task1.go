package day14

import "strconv"

func task1() {
	_ = getInputs()
}

func decimalToBinary(i int64) string {
	b := strconv.FormatInt(i, 2)

	return b[len(b)-36:]
}
