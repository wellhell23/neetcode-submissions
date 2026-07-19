func maxArea(heights []int) int {
	low, high:= 0, len(heights)-1
	area:= 0
	for low < high {
		tempArea:= 0
		if heights[low] < heights[high] {
			tempArea = heights[low]* (high-low)
			low++
		} else {
			tempArea = heights[high]*(high-low)
			high--
		}
		if area	< tempArea{
			area = tempArea
		}
	}
	return area
}
