type Router struct {
    cap int
    q []packet // master queue, FIFO (composite key)
    destToTimestamps map[int][]packet // map of keyed queues (key = destination)
    b map[packet]struct{} // buffer (composite key)
}

type packet struct {
    source, destination, timestamp int
}

func Constructor(memoryLimit int) Router {
    return Router{memoryLimit, []packet{}, map[int][]packet{}, map[packet]struct{}{}}
}


func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
    p := packet{source, destination, timestamp}

    // check duplicates by key with a map
    if _, ok := r.b[p]; ok {
        return false
    }
    // evictin in FIFO order
    if len(r.q) == r.cap {
        p := r.q[0]
        if len(r.destToTimestamps[p.destination]) > 0 {
            r.destToTimestamps[p.destination]=r.destToTimestamps[p.destination][1:] // pop from keyed queues
        }
        r.q = r.q[1:] // pop from master queue
        delete(r.b, p)
    }
    // insertion
    r.q = append(r.q, p) // master q
    if _, ok := r.destToTimestamps[destination]; !ok {
        r.destToTimestamps[destination] = []packet{}
    }
    r.destToTimestamps[destination] = append(r.destToTimestamps[destination], p)
    r.b[p] = struct{}{}
    return true
}


func (r *Router) ForwardPacket() []int {
    if len(r.b) == 0 {
        return []int{}
    }

    p := r.q[0]
    if len(r.destToTimestamps[p.destination]) > 0 {
        r.destToTimestamps[p.destination]=r.destToTimestamps[p.destination][1:] // pop from keyed queues
    }
    r.q = r.q[1:] // pop from master queue
    delete(r.b, p)

    return []int{p.source, p.destination, p.timestamp}
}


func (r *Router) GetCount(destination int, startTime int, endTime int) int {
    if _, ok := r.destToTimestamps[destination]; !ok {
        return 0
    }
    
    data := r.destToTimestamps[destination]

    left := sort.Search(len(data), func(i int) bool { return data[i].timestamp >= startTime })
    right := sort.Search(len(data), func(i int) bool { return data[i].timestamp > endTime })
    return right - left
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
