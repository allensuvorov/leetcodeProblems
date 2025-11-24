/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iter *Iterator
    temp int
    iterPtrAtNext bool
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{
        iter: iter,
        temp: 0,
        iterPtrAtNext: false,
    }
}

func (this *PeekingIterator) hasNext() bool {
    if !this.iterPtrAtNext { // if synced
        return this.iter.hasNext()
    }
    return true
}

func (this *PeekingIterator) next() int {
    if !this.iterPtrAtNext { // if synced
        next := this.iter.next()
        fmt.Println("next()-> ", next)
        return next
    }

    // else - sync pointers
    this.iterPtrAtNext = false
    return this.temp

}

func (this *PeekingIterator) peek() int {
    if !this.iterPtrAtNext { // if synced
        this.temp = this.iter.next()
        this.iterPtrAtNext = true
        return this.temp
    }

    return this.temp
}

//   1,2,3
//   p
//   i
