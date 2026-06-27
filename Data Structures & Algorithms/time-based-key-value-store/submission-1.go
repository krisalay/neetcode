type TimedVal struct {
    Val string
    Time int
}

type TimeMap struct {
    store map[string][]TimedVal
}

func Constructor() TimeMap {
    return TimeMap{
        store: make(map[string][]TimedVal),
    }
}

func (this *TimeMap) Set(key, val string, timestamp int) {
    timedValArr, ok := this.store[key]
    if !ok {
        timedValArr = []TimedVal{TimedVal{Val: val, Time: timestamp}}
    } else {
        timedValArr = append(timedValArr, TimedVal{Val: val, Time: timestamp})
    }
    this.store[key] = timedValArr
}

func (this *TimeMap) Get(key string, timestamp int) string {
    timedValArr, ok := this.store[key]
    var res string
    if !ok {
        return res
    }
    left, right := 0, len(timedValArr) - 1
    for left <= right {
        mid := left + (right-left)/2
        if timedValArr[mid].Time == timestamp {
            return timedValArr[mid].Val
        }
        if timedValArr[mid].Time > timestamp {
            right = mid - 1
        } else {
            res = timedValArr[mid].Val
            left = mid + 1
        }
    }

    return res
}