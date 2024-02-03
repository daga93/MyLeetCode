func sequentialDigits(low int, high int) []int {
    numberOfDigitsMin, numberOfDigitsHigh := len(strconv.Itoa(low)), len(strconv.Itoa(high))

    startingDigit := int(float64(low) / math.Pow(10, float64(numberOfDigitsMin-1)))

    solution := []int{}
    DigitLoops:
    for i := numberOfDigitsMin; i <= numberOfDigitsHigh; i++ {
        for j := startingDigit; j <= 9; j++ {
        if j + i > 10 {
            startingDigit = 1
            continue DigitLoops
        }
        nb, err := createNumber(j, i)
        if nb > 123456789 {
            return solution
        }
        if err != nil {
        if nb <= high && nb >= low {
            solution = append(solution, nb)
        }
            startingDigit = 1
            continue DigitLoops
        }
        if nb > high {
            break DigitLoops
        }
        if nb < low {
            continue
        }
        solution = append(solution, nb)
        }
        
    }
    return solution
}

func createNumber(start, digits int) (int, error) {
    createdNumber := 0
    currentNumber := start
    for i := digits; i > 0; i-- {
        createdNumber = createdNumber + currentNumber * int(math.Pow(10, float64(i-1)))
        currentNumber++
        if currentNumber == 10 {
            return createdNumber, errors.New("anything")
        }
    }
    return createdNumber, nil
}
