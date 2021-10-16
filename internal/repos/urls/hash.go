package urls

import (
	"hash/fnv"
	"strconv"
)

// UrlHash is a pseudonym for url hash type
type UrlHash string

// EmptyUrlHash is an empty value of UrlHash type
const EmptyUrlHash UrlHash = UrlHash("")

// Hash calculates url hash
// Inputs:
//   url - url to calculate the hash
// Output:
//   calculated hash string
func Hash(url string) UrlHash {
	h := fnv.New32a()
	h.Write([]byte(url))

	return UrlHash(strconv.FormatUint(uint64(h.Sum32()), 10))
}
