type FooBar struct {
	c chan struct{}
    n int
}

func NewFooBar(n int) *FooBar {
    return &FooBar{
        c: make(chan struct{}),
        n: n,
    }
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
        // printFoo() outputs "foo". Do not change or remove this line.
        printFoo()
        <- fb.c
        fb.c <- struct{}{}
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
        fb.c <- struct{}{}
		// printBar() outputs "bar". Do not change or remove this line.
        printBar()
        <- fb.c
	}
}
