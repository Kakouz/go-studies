package main

import "math"

type Forma interface {
	Area() float64
}
type Retangulo struct {
	Largura float64
	Altura  float64
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}

type Triangulo struct {
	Base float64
	Altura float64
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func Perimetro(retangulo Retangulo) float64 {
	return 2 * (retangulo.Largura + retangulo.Altura)
}
