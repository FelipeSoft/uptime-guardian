package worker

import (
	probing "github.com/andrewsjg/pro-bing"
)

type PingStatistics struct {
	PacketsSent int     `json:"PackageSent"`
	PacketLoss  float64 `json:"sentPackages"`
	PacketsRecv int     `json:"receivedPackages"`
}

func TestByICMP(ip string) (*PingStatistics, error) {
	pinger, err := probing.NewPinger(ip)
	if err != nil {
		panic(err)
	}
	pinger.SetPrivileged(true)
	pinger.Count = 3
	pinger.Timeout = 10 * 1e9
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics()
	return &PingStatistics{
		PacketsSent: stats.PacketsSent,
		PacketLoss:  stats.PacketLoss,
		PacketsRecv: stats.PacketsRecv,
	}, nil
}