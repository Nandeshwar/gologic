package main

func uniqueInWindow(a []int, window int) []int {
	uniqueItemsInEachWindow := []int{}
	m := map[int]int{}
	// insert window sized value . ex. 4 is window size so take 4 values
	for _, item := range a[0:window] {
		_, ok := m[item]
		if ok {
			m[item] += m[item]
		} else {
			m[item] = 1
		}
	}
	uniqueItemsInEachWindow = append(uniqueItemsInEachWindow, len(m))

    // check first item and so on is in map or not. if yes and count is 1 delete other wise decrement count
    // Do the same thing for i th item ex. index 4 item
	for i := window; i < len(a); i++ {
		cnt, ok := m[a[i-window]]
		if ok {
			if cnt == 1 {
				delete(m, a[i-window])
			} else {
				m[a[i-window]] -= 1
			}
		}
		cnt, ok = m[a[i]]
		if ok {
			m[a[i]] += 1
		} else {
			m[a[i]] = 1
		}

		uniqueItemsInEachWindow = append(uniqueItemsInEachWindow, len(m))
	}
	return uniqueItemsInEachWindow
}
