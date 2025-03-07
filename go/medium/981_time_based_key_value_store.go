type TimeMap struct {
    data map[string][]item
}

type item struct {
    value string
    timestamp int
}


func Constructor() TimeMap {
    data := make(map[string][]item)
    return TimeMap{
        data: data,
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.data[key] = append(this.data[key], item{value,timestamp,})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    items, exists := this.data[key]

    if !exists {
        return ""
    }

    l, r := 0, len(items) - 1

    for l <= r {
        m := l + (r - l)/2
        if items[m].timestamp == timestamp {
            return items[m].value
        }

        if items[m].timestamp > timestamp {
            r = m - 1
        }
        
        if items[m].timestamp < timestamp {
            l = m + 1
        }
    }
    
    if r == -1 {
        return ""
    }

    return items[r].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
