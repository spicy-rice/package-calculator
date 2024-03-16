package main

import (
	"errors"
	"sort"
)

// Modified calculatePacksNeeded to accept packSizes as a parameter
func calculatePacksNeeded(orderQuantity int, packSizes []int) (map[int]int, error) {
    if orderQuantity <= 0 {
        return nil, errors.New("order quantity must be a positive integer")
    }

    // Make a local copy of packSizes and sort it
    localPackSizes := make([]int, len(packSizes))
    copy(localPackSizes, packSizes)
    sort.Sort(sort.Reverse(sort.IntSlice(localPackSizes)))

    packsNeeded := make(map[int]int)
    remainingQuantity := orderQuantity

    for _, packSize := range localPackSizes {
        if remainingQuantity <= 0 {
            break
        }
        numberOfPacks := remainingQuantity / packSize
        if numberOfPacks > 0 {
            packsNeeded[packSize] = numberOfPacks
            remainingQuantity -= numberOfPacks * packSize
        }
    }

    if remainingQuantity > 0 {
        smallestPack := localPackSizes[len(localPackSizes)-1]
        packsNeeded[smallestPack] += 1
    }

    return packsNeeded, nil
}
