package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	tests := []struct {
		name  string
		in1   int
		in2   int
		saida int
	}{
		{
			name:  "positives",
			in1:   2,
			in2:   5,
			saida: 7,
		},
		{
			name:  "negatives",
			in1:   -2,
			in2:   -5,
			saida: -7,
		},
		{
			name:  "negatives_and_positives",
			in1:   -2,
			in2:   5,
			saida: 3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Soma(test.in1, test.in2)
			if got != test.saida {
				t.Error("want: ", test.saida, " got: ", got)
			}
		})
	}
}

func TestFactorial(t *testing.T) {
	cases := []struct {
		caso     string
		valor    int64
		saida int64
	}{
		{
			caso:		"zero",
			valor:  	0,
			saida:	 1,
		},
	}

	for _, casos := range cases {
		t.Run(casos.caso, func(t *testing.T) {
			got := Fatorial(casos.valor)
			if got != casos.saida {
				t.Error("want: ", casos.saida, " got: ", got)
			}
		})
	}
}

func TestSearchBinaryIndex(t *testing.T) {
	tests := []struct {
		style  string
		in     int
		listIn []int
		out    int
	}{
		{
			style:  " => ordenado/crescente",
			in:     3,
			listIn: []int{3, 4, 5, 6, 7, 9},
			out:    0,
		},
		{
			style:  " => ordenado/decrescente",
			in:     2,
			listIn: []int{7, 6, 5, 4, 3, 2, 1},
			out:    1,
		},
		 {
			style: " => aleatorio",
			in: 8,
			listIn: []int{3,6,8,2,1,9,7},
			out: 5,
		},
		{
			style: " => valor inexistente",
			in: 8,
			listIn: []int{7, 6, 5, 4, 3, 2, 1},
			out: -1,
		},
	}
	
	for _, caso := range tests {
		t.Run(caso.style, func(t *testing.T) {
			got := SearchBinaryIndex(caso.listIn, caso.in)
			if got != caso.out {
				t.Error("want", caso.out, "got", got)
			}
		})
	}
}
