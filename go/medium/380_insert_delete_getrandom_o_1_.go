type RandomizedSet struct {
    buckets [][]int
    divider int
    bucketMaxLen int
}

func (this *RandomizedSet) evacuate() {
    // grow case
    this.divider = this.growSet()
    newBuckets := make([][]int, this.divider)
    for _, bucket := range this.buckets{
        for _, num := range bucket {
            bucket := this.hashFunc(num)
            newBuckets[bucket] = append(newBuckets[bucket], num)
        }
    }
    this.buckets = newBuckets
}


func (this *RandomizedSet) hashFunc(val int) int {
    if val < 0 {
        val = -val
    }
    return val % this.divider
}


func (this *RandomizedSet) growSet() int {
    val := len(this.buckets)
    return val + val/2
}



func Constructor() RandomizedSet {
    return RandomizedSet{make([][]int,10), 10, 1000}
}


func (this *RandomizedSet) Insert(val int) bool {
    present := false
    bucket := this.hashFunc(val)
    for _, v := range this.buckets[bucket] {
        if v == val {
            present = true
        }
    }
    // insert
    if !present {
        this.buckets[bucket] = append(this.buckets[bucket], val)
        if len(this.buckets[bucket]) > this.bucketMaxLen {
            this.evacuate()
        }
    }
    return !present
}


func (this *RandomizedSet) Remove(val int) bool {
    present := false
    bucket := this.hashFunc(val)
    for i, v := range this.buckets[bucket] {
        if v == val {
            present = true
            oldLen := len(this.buckets[bucket])
            part1 := this.buckets[bucket][:i]
            part2 := this.buckets[bucket][i+1:]
            this.buckets[bucket] = append(part1, part2...)
            newLen := len(this.buckets[bucket])
            if newLen != oldLen - 1 {
                fmt.Printf("error removing value %v from backet", v)
            }
            break
        }
    }
    return present
}


func (this *RandomizedSet) GetRandom() int {
    nonEmptyBackets := []int{}
    setLen := 0
    for i, v := range this.buckets {
        if len(v) > 0 {
            setLen += len(v)
            nonEmptyBackets = append(nonEmptyBackets, i)
        }
    }

    randVal := rand.Intn(setLen) + 1
    for _, v := range nonEmptyBackets {
        bucketLen := len(this.buckets[v])
        if randVal - bucketLen > 0 { // 1 - 1 = 0
            randVal -= bucketLen
        } else {
            return this.buckets[v][randVal-1]
        }
    }
    return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
