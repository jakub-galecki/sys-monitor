package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

//USER PID CPU MEM VSZ RSS TT STAT STARTED TIME COMMAND
type StatPerProcess struct {
	user    string
	cpu     float64
	mem     float64
	time    string
	command string
}

func getStats() string {
	cmd := "ps -exo \"user,%cpu,%mem,time, command\" | sort -nrk 2,2 | head -n 20"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

func getStatStructs(statsString string) []StatPerProcess {
	lines := strings.Split(statsString, "\n")
	var stats []StatPerProcess
	for _, line := range lines {
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		mem, err := strconv.ParseFloat(fields[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		cpu, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		s := StatPerProcess{
			user:    fields[0],
			cpu:     cpu,
			mem:     mem,
			time:    fields[3],
			command: fields[4],
		}
		stats = append(stats, s)
	}
	return stats
}

func main() {
	statsString := getStats()
	_ = getStatStructs(statsString)
}
