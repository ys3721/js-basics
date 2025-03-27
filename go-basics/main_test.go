package main

import "testing"

func BenchmarkGoAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAdd()
	}
}

func BenchmarkGoAddInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddInt()
	}
}

func BenchmarkGoAddLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddLock()
	}
}

func BenchmarkGoAddBigLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddBigLock()
	}
}

func BenchmarkGoAddIntSerial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoAddIntSerial()
	}
}
