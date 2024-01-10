package main

import (
	"banquet/probabilistic"
	"fmt"
	"github.com/bits-and-blooms/bloom"
)

type BloomFilterParam struct {
	m uint
	k uint
}

func main() {
	bloomFilterParams := []*BloomFilterParam{
		{
			m: uint(1),
			k: uint(1),
		},
		{
			m: uint(8),
			k: uint(1),
		},
		{
			m: uint(16),
			k: uint(4),
		},
		{
			m: uint(32),
			k: uint(8),
		},
		{
			m: uint(64),
			k: uint(1),
		},
		{
			m: uint(64),
			k: uint(10),
		},
		{
			m: uint(128),
			k: uint(10),
		},
		{
			m: uint(256),
			k: uint(10),
		},
		{
			m: uint(512),
			k: uint(10),
		},
	}

	nEntries := 10

	fmt.Printf("Estimated false positive rates with %d entries: \n",
		nEntries)
	determineErrorRate(bloomFilterParams, uint(nEntries))
}

func determineErrorRateCustom(bloomFilterParams []BloomFilterParam) {
	for _, mk := range bloomFilterParams {
		// m
		size := mk.m
		// k
		numberOfHashFunc := mk.k

		errorRate := float32(0)
		bf := bloom.New(size, numberOfHashFunc)
		for i := 0; i < 100; i++ {
			errorRate += observedErrorRate(bf)
		}

		fmt.Printf("After %d tests with m=%d, k=%d there was a error rate of %f\n",
			100,
			size,
			numberOfHashFunc,
			errorRate/100)
	}
}

func observedErrorRate(bf *bloom.BloomFilter) float32 {
	for _, s := range probabilistic.SimpleInput {
		bf.AddString(s)
	}

	successes, failures := 0, 0
	var actual bool
	for _, t := range probabilistic.Tests {
		actual = bf.TestString(t.Input)
		if actual != t.Expected {
			failures++
		} else {
			successes++
		}
	}

	if failures == 0 {
		return 0
	}

	return 100.0 - (float32(100) * float32(successes) / float32(len(probabilistic.Tests)))
}

func determineErrorRate(bloomFilterParams []*BloomFilterParam, nEntries uint) {
	for _, param := range bloomFilterParams {
		filter := bloom.New(param.m, param.k)

		fmt.Printf("{m=%d, k=%d}: %f\n",
			param.m,
			param.k,
			filter.EstimateFalsePositiveRate(nEntries))
	}
}
