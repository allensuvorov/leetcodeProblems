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
    if f == this.minFreq && l.Len() == 0 {
        delete(this.freqToKeysList, f)
        this.minFreq++
    }
    this.keyToFreq[key]++
    f = this.keyToFreq[key]
    if _, ok := this.freqToKeysList[f]; !ok {
        this.freqToKeysList[f] = list.New()
    }
    this.freqToKeysList[f].PushBack(key)
}


func (this *LFUCache) Get(key int) int {
    if val, ok := this.keyToVal[key]; ok {
        this.incFreq(key)
        return val
    }
    return -1
}


func (this *LFUCache) Put(key int, value int)  {
    if _, ok := this.keyToElement[key]; !ok {
        if len(this.keyToElement) == this.cap { // eviction case
            l := this.freqToKeysList[this.minFreq]
            frontElement := l.Front()
            lfuKey := frontElement.Value.(int)
            delete(this.keyToElement, lfuKey)
            delete(this.keyToFreq, lfuKey)
            delete(this.keyToVal, lfuKey)
            l.Remove(frontElement)
        }
        
        if _, ok := this.freqToKeysList[1]; !ok {
            this.freqToKeysList[1] = list.New()
        }
        
        e := this.freqToKeysList[1].PushBack(key)
        
        this.keyToElement[key] = e
        this.keyToFreq[key] = 1
        this.minFreq = 1
    } else {
        this.incFreq(key)
    }
    this.keyToVal[key] = value
}



/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
