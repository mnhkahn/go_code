package main

const primeRK = 16777619

// hashstr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashstr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		// 只有32位，超出范围的会被丢掉
		sq *= sq
	}
	return hash, pow
}

func Index(s, sep string) int {
	n := len(sep)
	// Hash sep.
	hashsep, pow := hashstr(sep)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashsep && s[:n] == sep {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashsep && s[i-n:i] == sep {
			return i - n
		}
	}
	return -1
}

func IndexFuck(s, sep string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			is := true
			j := 0
			for ; j < len(sep); j++ {
				if sep[j] != s[i] {
					is = false
					break
				}
			}
			if is {
				return i
			}
		}
	}
	return -1
}

func main() {
	a := "hello, world"
	aa := "ll"
	println(Index(a, aa))
	println(IndexFuck(a, aa))
}
