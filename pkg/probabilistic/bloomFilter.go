package probabilistic

import "encoding/binary"

type BloomFilter struct {
	// Length of the bloom filter
	m uint
	// Number of hash functions
	k uint
	// Bits in this set get turned to 1 if a hash functions returns that index
	set []bool
}

func New(m uint, k uint) *BloomFilter {
	return &BloomFilter{k: max(1, k), m: max(1, m), set: make([]bool, m)}
}

func (bf *BloomFilter) Add(data []byte) *BloomFilter {
	hashes := baseHashes(data)
	for i := uint(0); i < bf.k; i++ {
		bf.set[(bf.location(hashes, i))] = true
	}
	return bf
}

func (bf *BloomFilter) AddString(data string) *BloomFilter {
	return bf.Add([]byte(data))
}

func (bf *BloomFilter) Test(data []byte) bool {
	hashes := baseHashes(data)
	for i := uint(0); i < bf.k; i++ {
		if !bf.set[bf.location(hashes, i)] {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) TestString(data string) bool {
	return bf.Test([]byte(data))
}

// location returns the ith hashed location using the four base hash values
func location(h [4]uint64, i uint) uint64 {
	ii := uint64(i)
	return h[ii%2] + ii*h[2+(((ii+(ii%2))%4)/2)]
}

// bf.location calculates the final index within the Bloom filter's bit array
func (bf *BloomFilter) location(h [4]uint64, i uint) uint {
	hashedLocation := location(h, i)
	filterSize := uint64(bf.m)

	// Modulo to make the hashed location a valid array index
	return uint(hashedLocation % filterSize)
}

func (bf *BloomFilter) EstimateFalsePositiveRate(n uint) (fpRate float64) {
	rounds := uint32(100000)
	bf.set = make([]bool, bf.m)
	n1 := make([]byte, 4)
	for i := uint32(0); i < uint32(n); i++ {
		binary.BigEndian.PutUint32(n1, i)
		bf.Add(n1)
	}
	fp := 0
	// test for number of rounds
	for i := uint32(0); i < rounds; i++ {
		binary.BigEndian.PutUint32(n1, i+uint32(n)+1)
		if bf.Test(n1) {
			//fmt.Printf("%v failed.\n", i+uint32(n)+1)
			fp++
		}
	}
	fpRate = float64(fp) / (float64(rounds))
	bf.set = make([]bool, bf.m)
	return
}
