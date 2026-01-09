type LFUCache struct {
    cap int
    freqToKeysList map[int]*list.List
    keyToElement map[int]*list.Element
    keyToFreq map[int]int
    keyToVal map[int]int
    minFreq int
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        cap: capacity,
        freqToKeysList: make(map[int]*list.List),
        keyToElement: make(map[int]*list.Element),
        keyToFreq: make(map[int]int),
        keyToVal: make(map[int]int),
        minFreq: 0,
    }
}


func (this *LFUCache) incFreq(key int) { 
    e := this.keyToElement[key]
    f := this.keyToFreq[key]
    l := this.freqToKeysList[f]
    l.Remove(e)
    
    // Clean up the empty list
    if l.Len() == 0 {
        delete(this.freqToKeysList, f)
        if f == this.minFreq {
            this.minFreq++
        }
    }
    // Increment frequency
    this.keyToFreq[key]++
    newFreq := this.keyToFreq[key]
    
    // Ini new list fot new freq, if needed
    if _, ok := this.freqToKeysList[newFreq]; !ok {
        this.freqToKeysList[newFreq] = list.New()
    }

    newElement := this.freqToKeysList[newFreq].PushBack(key)
    this.keyToElement[key] = newElement
}


func (this *LFUCache) Get(key int) int {
    if val, ok := this.keyToVal[key]; ok {
        this.incFreq(key)
        return val
    }
    return -1
}

func (this *LFUCache) evict() {
    l := this.freqToKeysList[this.minFreq]
    frontElement := l.Front()
    lfuKey := frontElement.Value.(int)
    
    // Remove from list
    l.Remove(frontElement)
    if l.Len() == 0 {
        delete(this.freqToKeysList, this.minFreq)
    }
    // Clean up all key maps
    delete(this.keyToElement, lfuKey)
    delete(this.keyToFreq, lfuKey)
    delete(this.keyToVal, lfuKey)
}

func (this *LFUCache) Put(key int, value int)  {
    if _, ok := this.keyToVal[key]; ok {
        this.keyToVal[key] = value
        this.incFreq(key)
        return
    }
    // add new element
    if len(this.keyToVal) == this.cap { // eviction case
        this.evict()
    }
    
    // Init a freq list for freq = 1, if needed
    if _, ok := this.freqToKeysList[1]; !ok { 
        this.freqToKeysList[1] = list.New()
    }
    
    this.keyToVal[key] = value
    e := this.freqToKeysList[1].PushBack(key)
    this.keyToElement[key] = e
    this.keyToFreq[key] = 1
    this.minFreq = 1
}



/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

 /*
 
 */
