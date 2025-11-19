package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// 1. Persegi
type Persegi struct {
	sisi float64
}

// 2. Persegi Panjang
type PersegiPanjang struct {
	panjang, lebar float64
}

// 3. Segitiga
type Segitiga struct {
	alas, tinggi, sisi1, sisi2, sisi3 float64
}

// 4. Lingkaran
type Lingkaran struct {
	diameter float64
}

// 5. Trapesium
type Trapesium struct {
	A, B, C, D, tinggi float64
}

// Methode Luas
func (p Persegi) Area() float64 {
	return math.Pow(p.sisi, 2)
}

func (pp PersegiPanjang) Area() float64 {
	return pp.panjang * pp.lebar
}

func (s Segitiga) Area() float64 {
	return 0.5 * s.alas * s.tinggi
}

func (l Lingkaran) Area() float64 {
	pi := 3.14
	r := 0.5 * l.diameter
	return pi * math.Pow(r, 2)
}

func (t Trapesium) Area() float64 {
	return 0.5 * t.tinggi * (t.A + t.B)
}

// Methode Keliling
func (p Persegi) Perimeter() float64 {
	return 4 * p.sisi
}

func (pp PersegiPanjang) Perimeter() float64 {
	return 2 * (pp.panjang + pp.lebar)
}

func (s Segitiga) Perimeter() float64 {
	return s.sisi1 + s.sisi2 + s.sisi3
}

func (l Lingkaran) Perimeter() float64 {
	pi := 3.14
	return pi * l.diameter
}

func (t Trapesium) Perimeter() float64 {
	return t.A + t.B + t.C + t.D
}

func ProcessShape(sh Shape) {
	fmt.Println("Hasil Luas:", sh.Area())
	fmt.Println("Hasil Keliling:", sh.Perimeter())
}

func main() {
	persegi := Persegi{sisi: 15}
	persegiPanjang := PersegiPanjang{panjang: 20, lebar: 12}
	segitiga := Segitiga{
		alas:   8,
		tinggi: 6,
		sisi1:  8,
		sisi2:  10,
		sisi3:  12,
	}
	lingkaran := Lingkaran{diameter: 14}
	trapesium := Trapesium{
		tinggi: 5,
		A:      10,
		B:      18,
		C:      6,
		D:      7,
	}

	fmt.Println("======== Persegi ========")
	ProcessShape(persegi)
	fmt.Println("-------------------------")

	fmt.Println("==== Persegi Panjang ====")
	ProcessShape(persegiPanjang)
	fmt.Println("-------------------------")

	fmt.Println("======== Segitiga =======")
	ProcessShape(segitiga)
	fmt.Println("-------------------------")

	fmt.Println("======= Lingkaran =======")
	ProcessShape(lingkaran)
	fmt.Println("-------------------------")

	fmt.Println("======= Trapesium =======")
	ProcessShape(trapesium)
	fmt.Println("-------------------------")
}
