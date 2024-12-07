package days

import (
	"io"
)

type Day interface {
	Input(input io.Reader)
	Run1() int
	Run2() int
	RunStr1() string
	RunStr2() string
}
