package main

import (
	"internal/cpu"
)

func archAvailableIEEE() bool {
	return cpu.X86.HasPCLMULQDQ && cpu.X86.HasSSE41
}

