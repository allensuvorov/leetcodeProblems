type H2O struct {
    oxyChan chan int
    hydGo chan int
    hydReleased chan int
}

func NewH2O() *H2O {
	h := &H2O{
        oxyChan: make(chan int, 1),
        hydGo: make(chan int, 2),
        hydReleased: make(chan int, 2), 
    }
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
    h.hydGo <- 1 // 
    // releaseHydrogen() outputs "H". Do not change or remove this line.
	releaseHydrogen()
    h.hydReleased <- 1 // i released it. ok?
}

func (h *H2O) Oxygen(releaseOxygen func()) {
    h.oxyChan <- 1 // O
    // releaseOxygen() outputs "H". Do not change or remove this line.
	releaseOxygen()
    // by now we have the molecule - H2O

    for range 2 {
        <- h.hydGo // hydrogen release - allow
    }

    for range 2 { // wait for 2 hydrogen release acks
        <- h.hydReleased // i see you released it. ok acks
    }

    <- h.oxyChan
}

/*

output: HHO 
hydGo: HH
oxyChan: O

- H ->
- H -> 
- H ->

- O -> 
- O ->
- O ->
*/
