package dojo

import (
	"reflect"
	"testing"
)

type In struct {
	person  []Person
	address []Address
}

var tests = []struct {
	name string
	in   In
	out  []PersonAddress
}{
	{
		name: "simple case",
		in: In{
			person: []Person{
				{
					ID:        1,
					FirstName: "Allen",
					LastName:  "Wang",
				},
				{
					ID:        2,
					FirstName: "Bob",
					LastName:  "Alice",
				},
			},
			address: []Address{
				{
					ID:       1,
					City:     "New York City",
					State:    "New York",
					PersonID: 2,
				},
				{
					ID:       2,
					City:     "Leetcode",
					State:    "California",
					PersonID: 3,
				},
			},
		},
		out: []PersonAddress{
			{
				FirstName: "Allen",
				LastName:  "Wang",
				City:      "Null",
				State:     "Null",
			},
			{
				FirstName: "Bob",
				LastName:  "Alice",
				City:      "New York City",
				State:     "New York",
			},
		},
	},
	{
		name: "complex case",
		in: In{
			person: []Person{
				{
					ID:        1,
					FirstName: "Allen",
					LastName:  "Wang",
				},
				{
					ID:        2,
					FirstName: "Bob",
					LastName:  "Alice",
				},
				{
					ID:        3,
					FirstName: "Gildson",
					LastName:  "Bezerra",
				},
				{
					ID:        4,
					FirstName: "Djair",
					LastName:  "Alves",
				},
			},
			address: []Address{
				{
					ID:       1,
					City:     "New York City",
					State:    "New York",
					PersonID: 2,
				},
				{
					ID:       2,
					City:     "Leetcode",
					State:    "California",
					PersonID: 5,
				},
				{
					ID:       3,
					City:     "Macau",
					State:    "RN",
					PersonID: 4,
				},
			},
		},
		out: []PersonAddress{
			{
				FirstName: "Allen",
				LastName:  "Wang",
				City:      "Null",
				State:     "Null",
			},
			{
				FirstName: "Bob",
				LastName:  "Alice",
				City:      "New York City",
				State:     "New York",
			},
			{
				FirstName: "Gildson",
				LastName:  "Bezerra",
				City:      "Null",
				State:     "Null",
			},
			{
				FirstName: "Djair",
				LastName:  "Alves",
				City:      "Macau",
				State:     "RN",
			},
		},
	},
}

func TestCombineTwoTables1(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := CombineTwoTables1(test.in.person, test.in.address)
			if !reflect.DeepEqual(test.out, res) {
				t.Errorf("want: %#v, got: %#v", test.out, res)
			}
		})
	}
}

func TestCombineTwoTables2(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := CombineTwoTables2(test.in.person, test.in.address)
			if !reflect.DeepEqual(test.out, res) {
				t.Errorf("want: %#v, got: %#v", test.out, res)
			}
		})
	}
}

func BenchmarkCombineTwoTables1(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombineTwoTables1(test.in.person, test.in.address)
			}
		})
	}
}

func BenchmarkCombineTwoTables2(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombineTwoTables2(test.in.person, test.in.address)
			}
		})
	}
}
