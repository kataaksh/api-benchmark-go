package benchmark

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Result struct {
	Duration   time.Duration
	StatusCode int
	Error      error
}

func Run(url string, totalRequests, concurrency int) error {
	var wg sync.WaitGroup
	requestsPerWorker := totalRequests / concurrency
	results := make([]Result, 0, totalRequests)
	resultsMu := sync.Mutex{}

	client := &http.Client{}
	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				start := time.Now()
				resp, err := client.Get(url)
				duration := time.Since(start)

				result := Result{
					Duration:   duration,
					StatusCode: 0,
					Error:      err,
				}

				if err == nil {
					result.StatusCode = resp.StatusCode
					resp.Body.Close()
				}

				resultsMu.Lock()
				results = append(results, result)
				resultsMu.Unlock()
			}
		}()
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	// Analyze results
	var durations []time.Duration
	statusCounts := make(map[int]int)
	errorCount := 0
	var totalDuration time.Duration
	var min, max time.Duration

	for i, r := range results {
		if r.Error != nil {
			errorCount++
			continue
		}
		durations = append(durations, r.Duration)
		totalDuration += r.Duration
		statusCounts[r.StatusCode]++

		if i == 0 || r.Duration < min {
			min = r.Duration
		}
		if r.Duration > max {
			max = r.Duration
		}
	}

	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	getPercentile := func(p float64) time.Duration {
		if len(durations) == 0 {
			return 0
		}
		index := int(float64(len(durations)) * p)
		if index >= len(durations) {
			index = len(durations) - 1
		}
		return durations[index]
	}

	avg := time.Duration(0)
	if len(durations) > 0 {
		avg = totalDuration / time.Duration(len(durations))
	}

	// Color functions
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	// Judgment
	judge := func(d time.Duration) string {
		switch {
		case d < 200*time.Millisecond:
			return green("✅ Excellent")
		case d < 500*time.Millisecond:
			return yellow("⚠️ Fair")
		default:
			return red("❌ Slow")
		}
	}

	// Output with color
	fmt.Println()
	color.Cyan("📊 Benchmark Results:")
	fmt.Printf("Total Requests:        %s\n", white(totalRequests))
	fmt.Printf("Concurrency:           %s\n", white(concurrency))
	fmt.Printf("Total Time Taken:      %s\n", white(totalTime))
	fmt.Printf("Requests per Second:   %s\n", cyan(fmt.Sprintf("%.2f", float64(len(durations))/totalTime.Seconds())))
	fmt.Printf("Average Response Time: %s %s\n", blue(avg), judge(avg))
	fmt.Printf("Min Response Time:     %s\n", blue(min))
	fmt.Printf("Max Response Time:     %s\n", blue(max))
	fmt.Printf("P95 Response Time:     %s\n", magenta(getPercentile(0.95)))
	fmt.Printf("P99 Response Time:     %s\n", magenta(getPercentile(0.99)))

	if errorCount > 0 {
		fmt.Printf("%s %d requests failed\n", red("❌ Failed Requests:"), errorCount)
	} else {
		fmt.Printf("%s %d\n", green("✅ Failed Requests:"), errorCount)
	}

	fmt.Println(magenta("Status Code Summary:"))
	for code, count := range statusCounts {
		var colorFunc func(a ...interface{}) string
		switch {
		case code >= 200 && code < 300:
			colorFunc = green
		case code >= 400 && code < 500:
			colorFunc = yellow
		case code >= 500:
			colorFunc = red
		default:
			colorFunc = white
		}
		fmt.Printf("  %d => %s responses\n", code, colorFunc(count))
	}

	return nil
}
