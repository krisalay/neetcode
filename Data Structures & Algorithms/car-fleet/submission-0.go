type Car struct {
    Position int
    Speed int
}

func carFleet(target int, position []int, speed []int) int {
    cars := []Car{}
    for i, pos := range position {
        cars = append(cars, Car{Position: pos, Speed: speed[i]})
    }
    sort.Slice(cars, func(i, j int) bool {
        return cars[i].Position > cars[j].Position
    })

    fmt.Println(cars)

    stack := []float64{}
    for _, car := range cars {
        timeTaken := float64(target-car.Position)/float64(car.Speed)
        if len(stack) == 0 || timeTaken > stack[len(stack)-1] {
            stack = append(stack, timeTaken)
        }
    }
    return len(stack)
}
