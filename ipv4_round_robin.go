package maxminddb_golang_example

import (
	"net"
	"sync/atomic"
)

type IPV4RoundRobin struct {
	start   uint32
	shift   uint32
	current uint32
}

func NewIPV4RoundRobin(startIP net.IP, endIP net.IP) *IPV4RoundRobin {
	start := ipV4ToUint32(startIP.To4())
	end := ipV4ToUint32(endIP.To4())

	return &IPV4RoundRobin{start: start, shift: end - start, current: start}
}

func (rr *IPV4RoundRobin) Next() net.IP {
	result := rr.start + atomic.AddUint32(&rr.current, 1)%rr.shift

	b := make([]byte, 4)
	uint32ToIPV4(b[:], result)

	return b
}

func ipV4ToUint32(b []byte) uint32 {
	_ = b[3] // bounds check hint to compiler; see golang.org/issue/14808
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

func uint32ToIPV4(b []byte, v uint32) {
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[3] = byte(v)
	b[2] = byte(v >> 8)
	b[1] = byte(v >> 16)
	b[0] = byte(v >> 24)
}
