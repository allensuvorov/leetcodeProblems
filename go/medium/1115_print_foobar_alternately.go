type FooBar struct {
	gf, gb chan struct{}
    n int
}

func NewFooBar(n int) *FooBar {
	gf := make(chan struct{}, 1)
	gb := make(chan struct{}, 1)
    gf <- struct{}{}
    return &FooBar{
        gf: gf,
        gb: gb,
        n: n,
    }
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
        <- fb.gf
        // printFoo() outputs "foo". Do not change or remove this line.
        printFoo()
        fb.gb <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
        <- fb.gb
		// printBar() outputs "bar". Do not change or remove this line.
        printBar()
        fb.gf <- struct{}{}
	}
}
