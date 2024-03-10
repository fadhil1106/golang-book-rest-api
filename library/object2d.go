package library

import "math"

type Segitiga struct {
	Alas, Tinggi int
}

type Persegi struct {
	Sisi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Lingkaran struct {
	JariJari int
}

func (s Segitiga) Luas() float64 {
	return float64(s.Alas) * float64(s.Tinggi) / float64(2)
}

func (p Persegi) Luas() int {
	return p.Sisi * p.Sisi
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func (l Lingkaran) Luas() float64 {
	return float64(math.Phi) * float64(l.JariJari) * float64(l.JariJari)
}

func (s Segitiga) Keliling() int {
	return 3 * s.Alas
}

func (p Persegi) Keliling() int {
	return 4 * p.Sisi
}

func (pp PersegiPanjang) Keliling() int {
	return 2 * (pp.Panjang + pp.Lebar)
}

func (l Lingkaran) Keliling() float64 {
	return float64(2) * float64(math.Phi) * float64(l.JariJari)
}
