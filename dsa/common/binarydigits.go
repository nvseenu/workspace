package common

import "fmt"

func BinaryDigits(n int) [][]byte {
	rs := make([]byte, n)
	bdList := &BinaryDigitList{elms: [][]byte{}}
	binaryDigits(0, n, rs, bdList)
	return bdList.elms
}

func binaryDigits(start int, n int, res []byte, bls *BinaryDigitList) {
	if start == n {
		fmt.Println("binaryDigits1 res:", res)
		rs := make([]byte, len(res))
		copy(rs, res)
		bls.add(rs)
		return
	}
	fmt.Println(">>> start: ", start, "path:", 0)
	res[start] = 0
	binaryDigits(start+1, n, res, bls)
	fmt.Println(">>> start: ", start, "path:", 1)
	res[start] = 1
	binaryDigits(start+1, n, res, bls)
}

type BinaryDigitList struct {
	elms [][]byte
}

func (b *BinaryDigitList) add(v []byte) {
	b.elms = append(b.elms, v)
}
