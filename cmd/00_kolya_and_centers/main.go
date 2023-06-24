package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type DataCenter struct {
	ID, Restarts, Active int

	ServersActivity []bool
}

func (n *DataCenter) Greater(other *DataCenter) bool {
	a := n.Value()
	b := other.Value()
	if a == b {
		return n.ID < other.ID
	}

	return a > b
}

func (n *DataCenter) Less(other *DataCenter) bool {
	a := n.Value()
	b := other.Value()
	if a == b {
		return n.ID < other.ID
	}

	return a < b
}

func (n *DataCenter) Disable(serverID int) {
	if n.ServersActivity[serverID-1] == false {
		n.ServersActivity[serverID-1] = true
		n.Active--
	}
}

func (n *DataCenter) Reset(serversC int) {
	n.Restarts++
	n.Active = serversC
	n.ServersActivity = make([]bool, serversC)
}

func (n *DataCenter) Value() int {
	return n.Restarts * n.Active
}

type Answer struct {
	Node  *DataCenter
	Dirty bool
}

func main() {
	var nodeC, serversC, eventsC int
	_, _ = fmt.Fscanf(reader, "%d %d %d\n", &nodeC, &serversC, &eventsC)

	// Use two balanced binary search trees to access min and max DataCenter by Value() and ID.
	nodes := make([]*DataCenter, 0, 1000000)

	for i := 1; i <= nodeC; i++ {
		nodes = append(nodes, &DataCenter{
			ID:              i,
			Restarts:        0,
			Active:          serversC,
			ServersActivity: make([]bool, serversC), // false for active
		})
	}

	var (
		maxAns = Answer{Dirty: true}
		minAns = Answer{Dirty: true}
	)

	var kind string
	for ; eventsC > 0; eventsC-- {
		_, _ = fmt.Fscanf(reader, "%s", &kind)
		switch kind {
		case "DISABLE":
			nodeID, serverID := 0, 0
			_, _ = fmt.Fscanf(reader, "%d %d\n", &nodeID, &serverID)

			node := nodes[nodeID-1]
			node.Disable(serverID)

			if !maxAns.Dirty && maxAns.Node.ID == node.ID {
				maxAns.Dirty = true
			}
			minAns.Dirty = true
		case "RESET":
			var nodeID int
			_, _ = fmt.Fscanf(reader, "%d\n", &nodeID)

			node := nodes[nodeID-1]
			node.Reset(serversC)

			if !minAns.Dirty && minAns.Node.ID == node.ID {
				minAns.Dirty = true
			}
			maxAns.Dirty = true
		case "GETMAX":
			if maxAns.Dirty {
				// update
				maxAns.Node = nodes[0]
				for _, n := range nodes {
					if n.Greater(maxAns.Node) {
						maxAns.Node = n
					}
				}
				maxAns.Dirty = false
			}

			_, _ = fmt.Fscanf(reader, "\n")
			_, _ = fmt.Fprintln(writer, maxAns.Node.ID)
		case "GETMIN":
			if minAns.Dirty {
				// update
				minAns.Node = nodes[0]
				for _, n := range nodes {
					if n.Less(minAns.Node) {
						minAns.Node = n
					}
				}
				minAns.Dirty = false
			}

			_, _ = fmt.Fscanf(reader, "\n")
			_, _ = fmt.Fprintln(writer, minAns.Node.ID)
		default:
			panic(fmt.Sprintf("unknown event kind: %s", kind))
		}
	}

	_ = writer.Flush()
}
