package array

import("fmt")
type Array struct {
	Max int
	Arr []int
}

func (re *Array) Insert(pos int, num int) {
	if pos < 0 || pos > len(re.Arr) || !(len(re.Arr)+1 <= re.Max) {
		fmt.Println("Position out of range")
		return
	}
	re.Arr = append(re.Arr, 0)
	copy(re.Arr[pos+1:], re.Arr[pos:])
	re.Arr[pos] = num
}

func (re *Array) Delete(pos int){
	if pos < 1 || pos > len(re.Arr) {
		fmt.Println("Position out of range")
		return
	}
	re.Arr = append(re.Arr[:pos] , re.Arr[pos+1:]...)
}

func (re *Array) Display(){
	fmt.Println(re.Arr)
}

func (re *Array) Find(num int) int{
	for i := 0; i < len(re.Arr); i++ {
		if re.Arr[i]==num {
			return i
		}		
	}
	return -1
}

func (re *Array) Reverse() {
    n := len(re.Arr)
    for i := 0; i < n/2; i++ {
        j := n - i - 1
        if i != j {
			// bit manipulation for fun
            re.Arr[i] = re.Arr[i] ^ re.Arr[j]
            re.Arr[j] = re.Arr[i] ^ re.Arr[j]
            re.Arr[i] = re.Arr[i] ^ re.Arr[j]
        }
    }
}

// Leetcode questions

func MergeAlternately(word1 string, word2 string) string {
	answer := ""
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		answer += string(word1[i]) + string(word2[j])
		i++
		j++
	}
	for i < len(word1) {
		answer += string(word1[i])
		i++
	}
	for j < len(word2) {
		answer += string(word2[j])
		j++
	}
	return answer
}
