package logic

import (
	"math"
)

type IFibonacci interface {
	IsFi(num int) bool
	PreviousFi(num int) int
	ValuesAroundFi(num int) (int, int)
}

type FiService struct {
}

func (fs *FiService) IsFi(num int) bool {
	return fs.isPerfectSquare(5*num*num+4) || fs.isPerfectSquare(5*num*num-4)
}

func (fs *FiService) PreviousFi(num int) int {
	for i := 1; ; i++ {
		if fs.IsFi(num + i) {
			return num + i
		}
		if fs.IsFi(num - i) {
			return num - i
		}
	}
}

func (fs *FiService) ValuesAroundFi(num int) (int, int) {
	if num == 0 {
		return 0, 1
	}
	a, b := 0, 1
	for b < num {
		a, b = b, a+b
	}
	return a, b + a
}

func (fs *FiService) isPerfectSquare(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}
