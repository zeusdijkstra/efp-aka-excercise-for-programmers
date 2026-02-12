package main

import (
	"errors"
	"testing"
)

func TestCalculateFeet(t *testing.T) {
	testCase := []struct {
		name      string
		w, h      int
		expected  int
		wantError error
	}{
		{
			name:     "BothNumber",
			w:        15,
			h:        20,
			expected: 300,
		},
		{
			name:      "NegativeNumber",
			w:         15,
			h:         -10,
			expected:  0,
			wantError: ErrNegativeValue,
		},
		{
			name:      "ZeroNumber",
			w:         0,
			h:         0,
			wantError: ErrZeroValues,
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calculateFeet(tt.w, tt.h)

			if tt.wantError != nil {
				if err == nil {
					t.Fatalf("expected error, got nothing")
				}
				if !errors.Is(err, tt.wantError) {
					t.Fatalf("expected %v, got %v instead", tt.wantError, err)
				}
			}

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestFeetToMeters(t *testing.T) {
	testCase := []struct {
		name      string
		feet      int
		want      float64
		wantError error
	}{
		{
			name: "Normal Test",
			feet: 300,
			want: 91.44,
			// wantError: ,
		},
		{
			name:      "ZeroValueTest",
			feet:      0,
			want:      0,
			wantError: ErrZeroValues,
		},
		// since this function is impossible to return negative value,
		// this section is ignored
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			result, err := feetToMeters(tc.feet)

			if tc.wantError != nil {
				if err == nil {
					t.Fatalf("expected error, got nil instead")
				}
				if !errors.Is(err, tc.wantError) {
					t.Fatalf("expected error %v, got %v instead", tc.wantError, err)
				}
			}

			if result != tc.want {
				t.Errorf("expected %v, got %v", tc.want, result)
			}
		})
	}
}
