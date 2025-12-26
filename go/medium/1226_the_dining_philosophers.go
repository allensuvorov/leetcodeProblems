type DiningPhilosophers struct {
	mu sync.Mutex
}

func (this *DiningPhilosophers) wantsToEat(
	philosopher int,
	pickLeftFork func(),
	pickRightFork func(),
	eat func(),
	putLeftFork func(),
	putRightFork func(),
) {
	this.mu.Lock()

	pickLeftFork()
	pickRightFork()
	eat()
	putLeftFork()
	putRightFork()

	this.mu.Unlock()
}
