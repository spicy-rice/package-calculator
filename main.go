package main

import (
	"encoding/json"
	"net/http"
	"os"
	"sort"
	"strconv"
)

var packSizes []int

func main() {
    if err := loadConfig("config.json"); err != nil {
        panic(err)
    }

    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs) // Serve static files

    http.HandleFunc("/calculate-packs", handleCalculatePacks)

    println("Server is running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

func handleCalculatePacks(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    orderQuantityStr := query.Get("orderQuantity")
    orderQuantity, err := strconv.Atoi(orderQuantityStr)
    if err != nil || orderQuantity <= 0 {
        http.Error(w, "Invalid order quantity. Please provide a positive integer.", http.StatusBadRequest)
        return
    }

    packsNeeded, err := calculatePacksNeeded(orderQuantity, packSizes)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(packsNeeded)
}

func loadConfig(configPath string) error {
    fileBytes, err := os.ReadFile(configPath) // Updated from ioutil.ReadFile to os.ReadFile
    if err != nil {
        return err
    }

    var config struct {
        PackSizes []int `json:"packSizes"`
    }

    if err = json.Unmarshal(fileBytes, &config); err != nil {
        return err
    }

    packSizes = config.PackSizes
    sort.Sort(sort.Reverse(sort.IntSlice(packSizes))) // Ensure pack sizes are sorted in descending order
    return nil
}
