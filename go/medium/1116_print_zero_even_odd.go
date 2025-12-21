type ZeroEvenOdd struct {
	n        int
  goZero, goOdd, goEven  chan int
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		goZero: make(chan int),
		goOdd: make(chan int),
		goEven: make(chan int),
        n:        n,
	}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
    for i := 1; i <= z.n; i++ {
        printNumber(0)
        if i % 2 == 0 { // even
            z.goEven <- i
            <- z.goZero
        } else {
            z.goOdd <- i
            <- z.goZero
        }
    }
    close(z.goEven)
    close(z.goOdd)
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
    for v := range z.goEven {
        printNumber(v)
        z.goZero <- 0
    }
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
    for v := range z.goOdd {
        printNumber(v)
        z.goZero <- 0
    }
}
