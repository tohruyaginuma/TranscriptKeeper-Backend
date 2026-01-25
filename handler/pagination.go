package handler

import "strconv"

const (
	defaultOffset = 0
	defaultLimit  = 20
	maxLimit      = 100
)

func parseLimit(s string) int {
	if s == "" {
		return defaultLimit
	}

	n, err := strconv.Atoi(s)
	if err != nil || n <= 0 {
		return defaultLimit
	}
	if n > maxLimit {
		return maxLimit
	}

	return n
}

func parseOffset(s string) int {
	if s == "" {
		return defaultOffset
	}

	n, err := strconv.Atoi(s)
	if err != nil || n < 0 {
		return defaultOffset
	}

	return n
}
