func predictPartyVictory(senate string) string {
    size := len(senate)
    rq := NewQueue(size)
    dq := NewQueue(size)
    
    for i := range senate {
        if senate[i] == 'R' {
            rq.Enqueue(i)
        } else {
            dq.Enqueue(i)
        }
    }
    fmt.Println("1: ",rq, dq)
    for rq.Head != rq.Tail && dq.Head != dq.Tail {
        if rq.Data[rq.Head] < dq.Data[dq.Head] {
            Ban(&rq, &dq, size)
        } else {
            Ban(&dq, &rq, size)
        }

    }
    fmt.Println("2: ",rq, dq)

    if rq.Head == rq.Tail {
        return "Dire"
    } else {
        return "Radiant"
    }
}

func Ban (a, b *Queue, offset int) {
    a.Enqueue(a.Head + offset)
    a.Dequeue()
    b.Dequeue()
}

type Queue struct {
    Data []int
    Head int
    Tail int
}

func NewQueue (len int) Queue {
    data := make([]int, len)
    return Queue{data, 0, 0}
}


func (q *Queue) Enqueue(v int) {
    q.Data[q.Tail] = v
    if q.Tail == len(q.Data) {
        q.Tail = 0
    } else {
        q.Tail++
    }
}


func (q *Queue) Dequeue() int {
    res := q.Data[q.Head]
    if q.Head == len(q.Data) {
        q.Head = 0
    } else {
        q.Head++
    }
    return res
}
