package toml

import (
	"net"
	"testing"
	"time"
)

func TestMarshalString(t *testing.T) {
	type Config struct {
		Age        int
		Cats       []string
		Pi         float64
		Perfection []int
		DOB        time.Time
		Ipaddress  net.IP
	}
	
	var inputs = Config{
		Age:        13,
		Cats:       []string{},
		Pi:         3.145,
		Perfection: []int{11, 2, 3, 4},
		DOB:        time.Now(),
		Ipaddress:  net.ParseIP("192.168.59.254"),
	}
	
	t.Log(MarshalString(inputs))
}
