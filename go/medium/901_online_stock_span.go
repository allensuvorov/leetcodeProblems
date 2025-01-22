type StockSpanner struct {
    decStack []quote
    curDate int
}

type quote struct {
    price int
    date int
}


func Constructor() StockSpanner {
    return StockSpanner{[]quote{}, 0}
}


func (this *StockSpanner) Next(price int) int {
    this.curDate++
    // strictly decreasing stack
    // when was the last time val was bigger?
    top := len(this.decStack) - 1
    for len(this.decStack) > 0 && this.decStack[top].price <= price {
        this.decStack = this.decStack[:top]
        top = len(this.decStack) - 1
    }
    span := this.curDate
    if len(this.decStack) > 0 {
        span = this.curDate - this.decStack[top].date
    }
    this.decStack = append(this.decStack, quote{price, this.curDate})
    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
