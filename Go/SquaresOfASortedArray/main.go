func sortedSquares(nums []int) []int {
  newArr := make([]int, len(nums))
  startIndex := -1
  for i, v := range nums {
    if v > 0 { 
        startIndex = i 
        break
    }
  }
  if startIndex < 0 {
      startIndex = len(nums)
  }

  negativeIndex := startIndex-1
  maxNums := len(nums)
  for i, _ := range nums {
    if startIndex >= maxNums {
        newArr[i] = nums[negativeIndex] * nums[negativeIndex]
        negativeIndex--
        continue
    }

    if negativeIndex < 0 {
        newArr[i] = nums[startIndex] * nums[startIndex]
        startIndex++
        continue
    }

    if nums[negativeIndex] * (-1) < nums[startIndex] {
        newArr[i] = nums[negativeIndex] * nums[negativeIndex]
        negativeIndex--
    } else {
        newArr[i] = nums[startIndex] * nums[startIndex]
        startIndex++ 
    }   
  }
  
  return newArr
}

