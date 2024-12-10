package icmp

import (
	"context"
	ping "github.com/andrewsjg/pro-bing"
	"time"
)

type ICMPStatistics struct {
	PacketsSent int
	PacketsRecv int
	AvgRtt      time.Duration
	IP          string
	Error       string
}

func TestByICMP(ctx context.Context, ip string) *ICMPStatistics {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return &ICMPStatistics{
			IP:          ip,
			PacketsSent: 0,
			PacketsRecv: 0,
			AvgRtt:      0,
			Error:       err.Error(),
		}
	}

	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.SetPrivileged(true)

	err = pinger.Run()
	if err != nil {
		return &ICMPStatistics{
			IP:          ip,
			PacketsSent: 0,
			PacketsRecv: 0,
			AvgRtt:      0,
			Error:       err.Error(),
		}
	}

	stats := pinger.Statistics()
	return &ICMPStatistics{
		IP:          ip,
		PacketsSent: stats.PacketsSent,
		PacketsRecv: stats.PacketsRecv,
	}
}
