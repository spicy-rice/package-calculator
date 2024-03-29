package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

var packSizes []int

func main() {
    if err := loadConfig("config.json"); err != nil {
        log.Fatal(err)
    }

    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    http.HandleFunc("/calculate-packs", handleCalculatePacks)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" 
    }

    log.Printf("Server is running on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err) // If the server fails to start, log the error and stop the application
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
    fileBytes, err := os.ReadFile(configPath)
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
    sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
    return nil
}
