package days

import (
	"io"
	"math/big"
)

type Day interface {
	Input(input io.Reader)
	Run1() int
	Run2() int
	RunStr1() string
	RunStr2() string
	Run1BigInt() big.Int
	Run2BigInt() big.Int
}
