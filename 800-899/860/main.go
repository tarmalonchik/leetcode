package main

func lemonadeChange(bills []int) bool {
	coins := []int{0, 0}
	for i := range bills {
		if bills[i] == 5 {
			addToCoins(coins, bills[i])
		} else if bills[i] == 10 {
			addToCoins(coins, 10)
			if !getFromCoins(coins, 5) {
				return false
			}
		} else if bills[i] == 20 {
			if checkCoins(coins, 10) && checkCoins(coins, 5) {
				getFromCoins(coins, 10)
				getFromCoins(coins, 5)
			} else {
				if !(getFromCoins(coins, 5) && getFromCoins(coins, 5) && getFromCoins(coins, 5)) {
					return false
				}
			}
		}
	}
	return true
}

func addToCoins(coins []int, in int) {
	switch in {
	case 5:
		coins[0]++
	case 10:
		coins[1]++
	}
}

func getFromCoins(coins []int, in int) bool {
	num := 0
	switch in {
	case 10:
		num = 1
	}
	if coins[num] > 0 {
		coins[num]--
		return true
	}
	return false
}

func checkCoins(coins []int, in int) bool {
	num := 0
	switch in {
	case 10:
		num = 1
	}
	return coins[num] > 0
}
