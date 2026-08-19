import (
	"slices"
)

func subsets(nums []int) [][]int {
	mem := make(map[string]struct{})
	result := make([][]int, 1)
	result[0] = nums
	var i int
	for i < len(result) {
		cur := result[i]
		for i := range cur {
			newSet := make([]int, len(cur)-1)
			copy(newSet[:i], cur[:i])
			copy(newSet[i:], cur[i+1:])
			slices.Sort(newSet)
			newSetStrings := make([]string, len(cur)-1)
			for i := range newSet {
				newSetStrings[i] = strconv.Itoa(newSet[i])
			}
			memKey := strings.Join(newSetStrings, ",")
			if _, ok := mem[memKey]; !ok {
				result = append(result, newSet)
				mem[memKey] = struct{}{}
			}
		}
		i++
	}
	return result
}
