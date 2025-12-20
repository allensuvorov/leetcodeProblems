type Foo struct {
    c1, c2 chan int
}

func NewFoo() *Foo {
	return &Foo{
        c1: make(chan int),
        c2: make(chan int),
	}
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
    f.c1 <- 1
}

func (f *Foo) Second(printSecond func()) {
	<- f.c1
    /// Do not change this line
	printSecond()
    f.c2 <- 2
}

func (f *Foo) Third(printThird func()) {
	<- f.c2
    // Do not change this line
	printThird()
}
