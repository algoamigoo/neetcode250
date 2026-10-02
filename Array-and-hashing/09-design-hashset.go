type MyHashSet struct {
	myarr []int
}

func Constructor() MyHashSet {
	return MyHashSet{
		myarr: []int{},
	}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.myarr = append(this.myarr, key)
	}
}

func (this *MyHashSet) Remove(key int) {
	for idx, val := range this.myarr {
		if val == key {
			this.myarr = append(this.myarr[:idx], this.myarr[idx+1:]...)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, val := range this.myarr {
		if val == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */