package kgs

import (
	"testing"
)

func TestKGS(t *testing.T) {
	cases := []struct {
		name              string
		SaveCounterExpect []int64
		LastSavedCounter  int64
		Expected          []string
	}{
		{
			name:             "must change 'LastSavedCounter' from 0 to 1",
			LastSavedCounter: 0,
			Expected: []string{
				"1", "2",
			},
		},
		{
			name:             "must change 'LastSavedCounter' from -10000 to 1",
			LastSavedCounter: -10000,
			Expected: []string{
				"1", "2",
			},
		},
		{
			name:             "normal functionality - one key",
			LastSavedCounter: 1,
			Expected: []string{
				"1",
			},
		},
		{
			name:             "normal functionality - more key",
			LastSavedCounter: 1,
			Expected: []string{
				"1", "2",
				"3", "4",
				"5",
			},
		},
		{
			name:             "normal functionality - start from 100",
			LastSavedCounter: 100,
			Expected: []string{
				"C1", "D1",
				"E1", "F1",
				"G1",
			},
		},
		{
			name:             "normal functionality - start from 100000",
			LastSavedCounter: 100000,
			Expected: []string{
				"U0q", "V0q",
				"W0q", "X0q",
				"Y0q",
			},
		},
		{
			name:             "normal functionality - start from 9999999999999999",
			LastSavedCounter: 9999999999999999,
			Expected: []string{
				"aRgsGBBNJ", "aRgsGBBNJ",
				"aRgsGBBNJ", "cRgsGBBNJ",
				"eRgsGBBNJ",
			},
		},
		{
			name:              "normal functionality - test SaveCounter",
			SaveCounterExpect: []int64{10, 20, 30, 40, 50, 60},
			LastSavedCounter:  1,
			Expected: []string{ // 61 keys
				"1", "2", "3", "4", "5", "6", "7", "8",
				"9", "a", "b", "c", "d", "e", "f", "g",
				"h", "i", "j", "k", "l", "m", "n", "o",
				"p", "q", "r", "s", "t", "u", "v", "w",
				"x", "y", "z", "A", "B", "C", "D", "E",
				"F", "G", "H", "I", "J", "K", "L", "M",
				"N", "O", "P", "Q", "R", "S", "T", "U",
				"V", "W", "X", "Y", "Z",
			},
		},
	}

	for _, tt := range cases {
		SaveCounterInterval := 0
		kgs := New(&KGSDep{
			SaveCounter: func(c int64) {
				if tt.SaveCounterExpect == nil {
					return
				}
				// check if we have expected counter
				if tt.SaveCounterExpect[SaveCounterInterval] != c {
					t.Errorf("%s: expected %d, got %d", tt.name, tt.SaveCounterExpect[SaveCounterInterval], c)
				}
				// SaveCounterInterval is incremented after each call to SaveCounter to save the position of the next expected counter
				SaveCounterInterval++
			},
			LastSavedCounter: tt.LastSavedCounter,
		})

		for _, expected := range tt.Expected {
			shortKey := kgs.GetKey()

			if shortKey != expected {
				t.Errorf("%s: expected %s, got %s", tt.name, expected, shortKey)
			}
		}
	}
}
