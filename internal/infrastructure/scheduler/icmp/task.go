package icmp

import (
	"context"
	"log"
	"time"
	"github.com/go-ping/ping"
)

func TestByICMP(ctx context.Context, ip string) {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		log.Printf("Error creating pinger for %s: %v", ip, err)
		return
	}

	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		log.Printf("Fail on ping %s; [Error] %s", ip, err.Error())
		return
	}

	stats := pinger.Statistics()
	log.Printf("%s Sent = %d, Received = %d, Lost = %d, Latency = %v", ip, stats.PacketsSent, stats.PacketsRecv, stats.PacketsSent-stats.PacketsRecv, stats.AvgRtt)
}
