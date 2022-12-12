package sorter

import (
	"fmt"
	"testing"

	"github.com/Bexanderthebex/flight-finder/pkg"
)

func TestTopologicalSort(t *testing.T) {
	testCases := []struct {
		name        string
		input       []Flight
		expected    Flight
		expectedErr error
	}{
		{
			name: "happy path",
			input: []Flight{
				{Source: "IND", Destination: "EWR"},
				{Source: "SFO", Destination: "ATL"},
				{Source: "GSO", Destination: "IND"},
				{Source: "ATL", Destination: "GSO"},
			},
			expected: Flight{Source: "SFO", Destination: "EWR"},
		},
		{
			name: "no possible starting path",
			input: []Flight{
				{Source: "SFO", Destination: "SFO"},
			},
			expectedErr: pkg.InvalidParameterError{Reason: "No possible starting flight path detected", Code: "no_starting_path"},
		},
		{
			name: "multiple starting paths",
			input: []Flight{
				{Source: "SFO", Destination: "ATL"},
				{Source: "IND", Destination: "GSO"},
			},
			expectedErr: pkg.InvalidParameterError{Reason: "Multiple starting flight path detected", Code: "multiple_starting_path"},
		},
		{
			name: "simultaneous paths",
			input: []Flight{
				{Source: "SFO", Destination: "ATL"},
				{Source: "ATL", Destination: "GSO"},
				{Source: "ATL", Destination: "IND"},
			},
			expectedErr: pkg.InvalidParameterError{Reason: "Simultaneous paths detected", Code: "simultaneous_path"},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			flight, err := TopologicalSort(test.input)

			if test.expectedErr != nil {
				if test.expectedErr.Error() != err.Error() {
					t.Error(fmt.Sprintf("expected %q got %q", test.expectedErr.Error(), err.Error()))
				}
				return
			}

			if test.expectedErr == nil && err != nil {
				t.Error(fmt.Sprintf("unexpected error: %s", err.Error()))
			}

			if test.expected.Source != flight.Source {
				t.Error(fmt.Sprintf("expected %s got %s", test.expected.Source, flight.Source))
			}
		})
	}
}
