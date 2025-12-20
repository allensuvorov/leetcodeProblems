// channel signaling
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

// mutex ordering
type Foo struct {
    mu2, mu3 sync.Mutex
}

func NewFoo() *Foo {
    foo := &Foo{}
    foo.mu2.Lock()
    foo.mu3.Lock()
    return foo
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
    f.mu2.Unlock()
}

func (f *Foo) Second(printSecond func()) {
    f.mu2.Lock()
    /// Do not change this line
	printSecond()
    f.mu3.Unlock()
}

func (f *Foo) Third(printThird func()) {
    f.mu3.Lock()
    // Do not change this line
	printThird()
}
