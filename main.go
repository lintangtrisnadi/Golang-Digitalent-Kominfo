package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type StatusResponse struct {
	Status     Status `json:"status"`
	WaterStat  string `json:"water_status"`
	WindStat   string `json:"wind_status"`
	UpdateTime string `json:"update_time"`
}

var (
	status      Status
	subscribers = make(map[chan []byte]bool)
	mutex       = make(chan struct{}, 1)
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Set up the ticker to update the JSON file every 15 seconds
	ticker := time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	// Start the ticker to broadcast status update every time it's changed
	go func() {
		for range ticker.C {
			mutex <- struct{}{}
			status.Water = rand.Intn(100) + 1
			status.Wind = rand.Intn(100) + 1
			broadcastStatus()
			<-mutex
		}
	}()

	// Set up HTTP server
	http.HandleFunc("/", statusHandler)
	http.HandleFunc("/events", eventsHandler)

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	statusResp := getStatusResponse()
	json.NewEncoder(w).Encode(statusResp)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan []byte)
	mutex <- struct{}{}
	subscribers[messageChan] = true
	<-mutex

	defer func() {
		mutex <- struct{}{}
		delete(subscribers, messageChan)
		close(messageChan)
		<-mutex
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		select {
		case message, ok := <-messageChan:
			if !ok {
				return
			}
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

func broadcastStatus() {
	statusResp := getStatusResponse()
	message, _ := json.Marshal(statusResp)
	for subscriber := range subscribers {
		subscriber <- message
	}
}

func getStatusResponse() StatusResponse {
	waterStat := getWaterStatus(status.Water)
	windStat := getWindStatus(status.Wind)
	updateTime := time.Now().Format(time.RFC3339)

	return StatusResponse{
		Status:     status,
		WaterStat:  waterStat,
		WindStat:   windStat,
		UpdateTime: updateTime,
	}
}

func getWaterStatus(water int) string {
	if water < 5 {
		return "aman"
	} else if water >= 6 && water <= 8 {
		return "siaga"
	} else {
		return "bahaya"
	}
}

func getWindStatus(wind int) string {
	if wind < 6 {
		return "aman"
	} else if wind >= 7 && wind <= 15 {
		return "siaga"
	} else {
		return "bahaya"
	}
}
