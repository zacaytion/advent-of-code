package utils

import (
	"bytes"
	"flag"
	"fmt"
	"log/slog"
	"os/exec"
	"strconv"
)

var part int

func init() {
	flag.IntVar(&part, "part", 1, "Part 1 or 2")
	flag.Parse()
}

func Run[A any](p1 func() A, p2 func() A) {
	slog.Info("Running", "part", part)

	var answer A
	switch part {
	case 1:
		answer = p1()
	case 2:
		answer = p2()
	default:
		slog.Error("Unexpected part", "part", part)
	}

	copyToClipboard(answer)
	slog.Info("Finished", "part", part, "answer", answer)
}

// tip: https://github.com/alexchao26/advent-of-code-go/blob/main/util/CopyToClipboard.go

// CopyToClipboard is for macOS
func copyToClipboard(i any) error {
	var txt string
	switch v := i.(type) {
	case string:
		txt = v
	case int, int32:
		txt = fmt.Sprintf("%d", v)
	}
	command := exec.Command("pbcopy")
	command.Stdin = bytes.NewReader([]byte(txt))

	if err := command.Start(); err != nil {
		return fmt.Errorf("error starting pbcopy command: %w", err)
	}

	err := command.Wait()
	if err != nil {
		return fmt.Errorf("error running pbcopy %w", err)
	}

	return nil
}

func RunesToInt(f, l rune) int {
				s := fmt.Sprintf("%c%c", f, l)
				i, err := strconv.Atoi(s)
				if err != nil {
								slog.Error("error converting runes to integer", "in", s ,"err", err)
				}
				return i
}
