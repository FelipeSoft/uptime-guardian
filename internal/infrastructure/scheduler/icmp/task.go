package icmp

import (
	"context"
	// "fmt"
	"log"
	"os/exec"
	// "log"
	// "net/http"
)
 
func TestByICMP(ctx context.Context, ip string) {
	cmd := exec.CommandContext(ctx, "ping", ip)
	err := cmd.Run()
	if err != nil {
		log.Printf("Fail on ping %s; [Error] %s", ip, err.Error())
		return
	}
	log.Printf("%s is up!!!", ip)	
	// client := &http.Client{}
	// req, err := http.NewRequestWithContext(ctx, "GET", "http://"+ip, nil)
	// if err != nil {
	// 	log.Printf("Fail on req %s is down; [Error] %s", ip, err.Error())
	// 	return
	// }
	// res, err := client.Do(req)
	// if err != nil {
	// 	log.Printf("Fail on res %s is down; [Error] %s", ip, err.Error())
	// 	return
	// }
	// defer res.Body.Close()
	// log.Printf("IP %s is up!!!", ip)
}
