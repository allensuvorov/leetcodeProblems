type H2O struct {
    mu sync.Mutex
    hydGo chan int
    hydReleased chan int
}

func NewH2O() *H2O {
	h := &H2O{
        hydGo: make(chan int, 2),
        hydReleased: make(chan int, 2), 
    }
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
    <- h.hydGo // 
    // releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
    h.hydReleased <- 1
}

func (h *H2O) Oxygen(releaseOxygen func()) {
    h.mu.Lock()
    for range 2 { h.hydGo <- 0}
    for range 2 { <- h.hydReleased }
    // releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
    h.mu.Unlock()
}
