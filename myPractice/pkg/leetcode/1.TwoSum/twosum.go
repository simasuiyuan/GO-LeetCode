package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // create a hashmap
	for k, v := range nums {
		/*
			golang syntax alert!

			if idx, ok := m[target-v]; ok {
				...
			}

			if target-v has value then ok is true
			and return the value of that element (idx)
		*/
		if idx, ok := m[target-v]; ok {
			// fmt.Println(ok)
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
