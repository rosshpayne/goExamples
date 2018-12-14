package main

import (
        "testing"
)
func BenchmarkAI_1000(b *testing.B) { benchmarkAI(b,1000) }
func BenchmarkAI_10000(b *testing.B) { benchmarkAI(b,10000) }
func BenchmarkAI_100000(b *testing.B) { benchmarkAI(b,100000) }
func BenchmarkAI_1000000(b *testing.B) { benchmarkAI(b,1000000) }

func benchmarkAI(b *testing.B,size int32) {
	var aa []int32

	aa=make([]int32,size,size)

	b.ResetTimer()
        for i:=0; i<b.N; i++ {
		arrayIndex(aa,size)
	}
}

func BenchmarkSA_1000(b *testing.B) { benchmarkSA(b,1000) }
func BenchmarkSA_10000(b *testing.B) { benchmarkSA(b,10000) }
func BenchmarkSA_100000(b *testing.B) { benchmarkSA(b,100000) }
func BenchmarkSA_1000000(b *testing.B) { benchmarkSA(b,1000000) }

func benchmarkSA(b *testing.B, size int32) {
	var aa []int32

	aa=make([]int32,0,size)
	b.ResetTimer()
        for i:=0; i<b.N; i++ {
		sliceAppend(aa,size)
	}
}

func BenchmarkCopy_1000(b *testing.B) { benchmarkCopy(b,1000) }
func BenchmarkCopy_10000(b *testing.B) { benchmarkCopy(b,10000) }
func BenchmarkCopy_100000(b *testing.B) { benchmarkCopy(b,100000) }
func BenchmarkCopy_1000000(b *testing.B) { benchmarkCopy(b,1000000) }

func benchmarkCopy(b *testing.B,size int32) {
	var aa []int32
	var bb []int32

	aa=make([]int32,size,size)
        for i:=int32(0); i<size;i++	{
		aa[i]=5*i
	}
	bb=make([]int32,size,size)

	b.ResetTimer()
        for i:=0; i<b.N; i++ {
		sliceCopy(aa,bb)
	}
}

func sliceCopy(aa,bb []int32) 	{
		copy(aa,bb)
}

func sliceAppend(aa []int32, size int32) {

	for i:=int32(0);i<size;i++ {
		aa=append(aa,i*5)
	}
}

func arrayIndex(aa []int32, size int32) {

	for i:=int32(0);i<size;i++ {
		aa[i]=i*5
	}
}

