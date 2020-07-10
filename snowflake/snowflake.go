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
	customEpoch int64 = 1420070400000
)

var maxNodeID int64
var maxSequence int64

var nodeID int64
var lastTimestamp int64 = 0
var sequence int64

func init() {
	maxNodeID = int64(math.Pow(2, nodeIDBits) - 1)
	maxSequence = int64(math.Pow(2, sequenceBits) - 1)
	nodeID = generateNodeID()
}

func generateNodeID() int64 {
	var nodeID int64
	if interfaces, err := net.Interfaces(); err == nil {
		h := fnv.New32a()
		for _, i := range interfaces {
			h.Write(i.HardwareAddr)
		}
		nodeID = int64(h.Sum32())
	} else {
		panic("interfaces not available")
	}
	nodeID = nodeID & maxNodeID
	return nodeID
}

var timestampMutex sync.Mutex
var sequenceMutex sync.Mutex

// Next returns the next logical snowflake
func Next() int64 {
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
	// id := currentTimestamp << (totalBits - epochBits)
	// id |= (nodeID << (totalBits - epochBits - nodeIDBits))
	// id |= sequence
	var id int64 = currentTimestamp << (nodeIDBits + sequenceBits)
	id |= (nodeID << sequenceBits)
	id |= sequence

	fmt.Printf("%b\n", id)
	return id
}

func ts() int64 {
	return int64(time.Now().UnixNano()/1000000) - customEpoch
}

func waitNextMillis(currentTimestamp int64) int64 {
	for currentTimestamp == lastTimestamp {
		currentTimestamp = ts()
	}
	return currentTimestamp
}
