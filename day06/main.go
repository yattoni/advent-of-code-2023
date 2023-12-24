package main

import (
	"log/slog"
	"os"
	"strings"

	util "github.com/yattoni/advent-of-code-2023"
)

func solve(times []int, distances []int) int {
	answer := 1
	for i := 0; i < len(times); i++ {
		slog.Debug("solve", "race:", i)
		waysToWin := 0
		for speed := 1; speed < times[i]; speed++ {
			// speed * time left, where speed == wait time at the start
			if speed*(times[i]-speed) > distances[i] {
				waysToWin++
				slog.Debug("solve", "speed", speed)
			}
		}
		slog.Debug("solve", "waysToWin", waysToWin)
		answer *= waysToWin
	}
	return answer
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo})))

	lines := util.ReadFileWithSpacesToLines("prompt-input")
	timeStrs := strings.Fields(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	times := make([]int, len(timeStrs))
	for i, s := range timeStrs {
		times[i] = util.MustAtoi(s)
	}
	distanceStrs := strings.Fields(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	distances := make([]int, len(distanceStrs))
	for i, s := range distanceStrs {
		distances[i] = util.MustAtoi(s)
	}

	slog.Info("parsed input", slog.Group("", "times", times, "distances", distances))

	answer := solve(times, distances)
	slog.Info("part1", "answer", answer) // 128700

	combinedTime := util.MustAtoi(strings.Join(timeStrs, ""))
	combinedDistance := util.MustAtoi(strings.Join(distanceStrs, ""))

	answer2 := solve([]int{combinedTime}, []int{combinedDistance})
	slog.Info("part2", "answer", answer2) // 39594072
}
