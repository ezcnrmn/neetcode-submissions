type timeMark struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]timeMark
}

func Constructor() TimeMap {
	tm := TimeMap{
		m: make(map[string][]timeMark),
	}
	return tm
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	index := this.binarySearch(this.m[key], timestamp)
	if index != -1 && this.m[key][index].timestamp == timestamp {
		this.m[key][index].value = value
		return
	}

	this.m[key] = append(this.m[key], timeMark{})
	for i := len(this.m[key])-1; i > index+1; i-- {
		this.m[key][i] = this.m[key][i-1]
	}
	this.m[key][index+1] = timeMark{timestamp: timestamp, value: value}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	index := this.binarySearch(this.m[key], timestamp)
	if index == -1 {
		return ""
	}

	return this.m[key][index].value
}

func (this *TimeMap) binarySearch(marks []timeMark, timestamp int) int {
	left, right := 0, len(marks)-1
	for left <= right {
		mid := left + (right-left)/2
		if mid == len(marks)-1 || (marks[mid].timestamp <= timestamp && marks[mid+1].timestamp > timestamp) {
			return mid
		}

		if marks[mid].timestamp > timestamp {
			right = mid-1
		} else {
			left = mid+1
		}
	}

	return -1
}