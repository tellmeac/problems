package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var n int

	_, _ = fmt.Fscanf(reader, "%d\n", &n)

	var rawTs string
	stamps := make([]int, 0, n)
	for ; n > 0; n-- {
		_, _ = fmt.Fscanf(reader, "%s ", &rawTs)
		st, _ := time.Parse("15:04", rawTs)
		stamps = append(stamps, st.Minute()+st.Hour()*60)
	}

	rotate := 24 * 60
	if len(stamps) == 1 {
		_, _ = fmt.Fprintf(writer, "%d\n", rotate)
		return
	}

	sort.Slice(stamps, func(i, j int) bool {
		return stamps[i] < stamps[j]
	})

	minDiff := math.Abs(float64(stamps[1] - stamps[0]))
	for ind := 1; ind < len(stamps); ind++ {
		minDiff = min(minDiff, math.Abs(float64(stamps[ind]-stamps[ind-1])))
	}

	// 00 00 to 23 59 case
	cycleDiff := math.Abs(float64(stamps[len(stamps)-1] - stamps[0] - rotate))
	minDiff = min(minDiff, cycleDiff)

	_, _ = fmt.Fprintf(writer, "%d\n", int(minDiff))

	_ = writer.Flush()
}
