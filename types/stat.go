package types

type LoadAvgInfo struct {
	Avg uint64
}
type PortStatusKind uint8

const (
	BIND PortStatusKind = iota + 1
	DIAL
)

type PortInfo struct {
	Port    uint16
	Status  PortStatusKind
	Uid     uint32
	Address string
}
type MemoryInfo struct {
	Total     uint64
	Available uint64
}
