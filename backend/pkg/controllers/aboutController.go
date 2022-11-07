package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type Service struct {
	Name string `json:"name"`
	Actions []struct {
		Name string `json:"name"`
		Description string `json:"description"`
	} `json:"actions"`
	Reactions []struct {
		Name string `json:"name"`
		Description string `json:"description"`
	} `json:"reactions"`
}

type AboutJson struct {
	Client struct { 
		Host string `json:"host"`
	} `json:"client"`
	Server struct {
		CurrentTime uint `json:"current_time"`
		Services []Service `json:"services"`
	} `json:"server"`
}

func getIP(r *http.Request) (string, error) {
	ips := r.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")

	if len(splitIps) > 0 {
		netIP := net.ParseIP(splitIps[len(splitIps)-1])
		if netIP != nil {
			return netIP.String(), nil
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	netIP := net.ParseIP(ip)
	if netIP != nil {
		ip := netIP.String()
		if ip == "::1" {
			return "127.0.0.1", nil
		}
		return ip, nil
	}

	return "", errors.New("IP not found")
}

func GetAboutJson(w http.ResponseWriter, r *http.Request) {
	var aboutJson AboutJson

	ip, _ := getIP(r)	
	aboutJson.Client.Host = ip

	now := time.Now()
	aboutJson.Server.CurrentTime = uint(now.Unix())

	data, err := os.ReadFile("pkg/controllers/services.json")
	if err != nil {
        fmt.Println(err)
    }
	var services []Service
    JsonErr := json.Unmarshal([]byte(data), &services)
    if JsonErr != nil {
		panic(JsonErr)
    }
	aboutJson.Server.Services = services

	w.WriteHeader(http.StatusOK)
	js, _ := json.MarshalIndent(aboutJson, "", " ")
	w.Write(js)
}