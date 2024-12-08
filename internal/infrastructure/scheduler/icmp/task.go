package icmp

import (
	"context"
	"log"
	"net/http"
)

func TestByICMP(ctx context.Context, ip string) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", "http://"+ip, nil)
	if err != nil {
		log.Printf("Fail on req %s is down; [Error] %s", ip, err.Error())
		return
	}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Fail on res %s is down; [Error] %s", ip, err.Error())
		return
	}
	defer res.Body.Close()
	log.Printf("IP %s is up!!!", ip)
}
