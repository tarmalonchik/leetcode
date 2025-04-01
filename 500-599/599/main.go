package main

func findRestaurant(list1 []string, list2 []string) []string {
	mp := make(map[string]int, len(list1))
	for i := range list1 {
		if _, ok := mp[list1[i]]; !ok {
			mp[list1[i]] = i
		}
	}

	sum := -1
	out := make([]string, 0)

	for i := range list2 {
		if val, ok := mp[list2[i]]; ok {
			if sum == -1 {
				sum = val + i
				out = append(out, list2[i])
			} else if val+i < sum {
				sum = val + i
				out = make([]string, 0)
				out = append(out, list2[i])
			} else if val+i == sum {
				out = append(out, list2[i])
			}
		}
	}
	return out
}
