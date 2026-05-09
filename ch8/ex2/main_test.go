package main

import "testing"

func BenchmarkPrintableIntF_String(b *testing.B) {
	var pif PrintableIntF
	pif = 80

	b.ResetTimer()

	for b.Loop() {
		_ = pif.String()
	}
}

func BenchmarkPrintableInt_String(b *testing.B) {
	var pi PrintableInt
	pi = 80

	b.ResetTimer()

	for b.Loop() {
		_ = pi.String()
	}
}
