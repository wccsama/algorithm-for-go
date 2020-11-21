package sort

func shell(nums []int) {
	for gate := len(nums); gate > 1; gate = gate / 2 {
		for i := 0; i < gate; i++ {
			for j := gate + i; j < len(nums); j = j + gate {
				current := nums[j]
				z := j
				for z > i && nums[z] > current {
					nums[z+gate] = nums[z]
					z = z - gate
				}
				nums[z] = current
			}
		}
	}
}

func shell2(nums []int) {
	for gate := len(nums) / 2; gate > 1; gate = gate / 2 {
		for i := 0; i < gate; i++ {
			for j := i + gate; i < len(nums); j++ {
				current := nums[j]
				z := j - gate
				for z > 0 && nums[z] < current {
					nums[z+gate] = nums[z]
					z -= gate
				}
				nums[z+gate] = current
			}
		}
	}
}

func shell3(a []int) {
	for gate := len(a) / 2; gate > 1; gate = gate / 2 {
		for i := 0; i <= gate; i++ {
			for j := i + gate; j < len(a); j += gate {
				current := j
				currentV := a[current]
				for k := j - gate; k > 0 && a[k] <= a[j]; k -= gate {
					a[k+gate], a[k] = a[k], a[k+gate]
					current -= gate
				}
				a[current+gate] = currentV
			}
		}
	}
}
