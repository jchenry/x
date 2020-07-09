package snowflake

import (
	"fmt"
	"hash/fnv"
	"math"
	"net"
	"sync"
	"time"
)

const (
	totalBits    = 64
	epochBits    = 32
	nodeIDBits   = 10
	sequenceBits = 12

	// Custom Epoch (January 1, 2015 Midnight UTC = 2015-01-01T00:00:00Z)
	customEpoch uint64 = 1420070400000
)

var maxNodeID uint64
var maxSequence uint64

var nodeID uint64
var lastTimestamp uint64 = 0
var sequence uint64

func init() {
	maxNodeID = uint64(math.Pow(2, nodeIDBits) - 1)
	maxSequence = uint64(math.Pow(2, sequenceBits) - 1)
	nodeID = generateNodeID()
}

func generateNodeID() uint64 {
	var nodeID uint64
	if interfaces, err := net.Interfaces(); err == nil {
		h := fnv.New32a()
		for _, i := range interfaces {
			h.Write(i.HardwareAddr)
		}
		nodeID = uint64(h.Sum32())
	} else {
		panic("interfaces not available")
	}
	nodeID = nodeID & maxNodeID
	return nodeID
}

var timestampMutex sync.Mutex
var sequenceMutex sync.Mutex

// Next returns the next logical snowflake
func Next() uint64 {
	timestampMutex.Lock()
	currentTimestamp := ts()
	timestampMutex.Unlock()

	sequenceMutex.Lock()
	if currentTimestamp == lastTimestamp {
		sequence = (sequence + 1) & maxSequence
		if sequence == 0 {
			// Sequence Exhausted, wait till next millisecond.
			currentTimestamp = waitNextMillis(currentTimestamp)
		}
	} else {
		sequence = 0
	}
	sequenceMutex.Unlock()

	lastTimestamp = currentTimestamp
	id := currentTimestamp << (totalBits - epochBits)
	fmt.Printf("%b\n", id)
	id |= (nodeID << (totalBits - epochBits - nodeIDBits))
	fmt.Printf("%b\n", id)

	id |= sequence
	fmt.Printf("%b\n", id)
	return id
}

func ts() uint64 {
	return uint64(time.Now().UnixNano()/1000000) - customEpoch
}

func waitNextMillis(currentTimestamp uint64) uint64 {
	for currentTimestamp == lastTimestamp {
		currentTimestamp = ts()
	}
	return currentTimestamp
}
