func minEatingSpeed(piles []int, h int) int {
    maxPileSize := getMaxPileSize(piles)
    left, right := 1, maxPileSize
    minEatingSpeed := maxPileSize
    for left <= right {
        possibleEatingSpeed := (left+right)/2
        if isValidEatingSpeed(piles, h, possibleEatingSpeed) {
            minEatingSpeed = min(minEatingSpeed, possibleEatingSpeed)
            right = possibleEatingSpeed - 1
        } else {
            left = possibleEatingSpeed + 1
        }
    }
    return minEatingSpeed
}

func isValidEatingSpeed(piles []int, h, eatingSpeed int) bool {
    var timeTaken int
    for _, pile := range piles {
        timeTaken += int(math.Ceil(float64(pile)/float64(eatingSpeed)))
    }
    return timeTaken <= h
}

func getMaxPileSize(piles []int) int {
    maxPileSize := 0
    for _, pile := range piles {
        maxPileSize = max(maxPileSize, pile)
    }
    return maxPileSize
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
