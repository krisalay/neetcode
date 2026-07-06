type Car struct {
    position int
    speed int
}

func carFleet(target int, position []int, speed []int) int {
    cars := make([]Car, len(position))
    for i, pos := range position {
        cars[i] = Car{position: pos, speed: speed[i]}
    }

    sort.Slice(cars, func(i,j int) bool {
        return cars[i].position > cars[j].position
    })
    
    stack := []float64{}
    for _, car := range cars {
        timeTaken := float64(target-car.position)/float64(car.speed)
        if len(stack) == 0 || timeTaken > stack[len(stack)-1] {
            stack = append(stack, timeTaken)
        }
    }

    return len(stack)
}
