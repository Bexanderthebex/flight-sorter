package sorter

import (
	"github.com/Bexanderthebex/flight-sorter/pkg"
)

type Flight struct {
	Source      string
	Destination string
}

func TopologicalSort(flights []Flight) (Flight, error) {
	flightDependencyCount := make(map[string]int)
	flightGraph := make(map[string][]string)

	for _, flight := range flights {
		flightGraph[flight.Source] = append(flightGraph[flight.Source], flight.Destination)
		// initialize the keys for the dependency counter first
		flightDependencyCount[flight.Source] = 0
	}

	for _, flight := range flights {
		flightDependencyCount[flight.Destination] += 1
	}

	flightEliminationQueue := make([]string, 0, len(flightDependencyCount))
	for flightSource, depedencyCount := range flightDependencyCount {
		if depedencyCount == 0 {
			flightEliminationQueue = append(flightEliminationQueue, flightSource)
		}
	}

	flightPaths := len(flightEliminationQueue)
	if flightPaths <= 0 {
		return Flight{}, pkg.InvalidParameterError{Reason: "No possible starting flight path detected", Code: "no_starting_path"}
	}

	if flightPaths > 1 {
		return Flight{}, pkg.InvalidParameterError{Reason: "Multiple starting flight path detected", Code: "multiple_starting_path"}
	}

	sortedFlights := make([]string, 0, len(flightDependencyCount))
	for len(flightEliminationQueue) > 0 {

		// a person cannot be on several flights simultaneously
		if len(flightEliminationQueue) > 1 {
			return Flight{}, pkg.InvalidParameterError{Reason: "Simultaneous paths detected", Code: "simultaneous_path"}
		}

		flight := flightEliminationQueue[0]
		flightEliminationQueue = flightEliminationQueue[1:]
		sortedFlights = append(sortedFlights, flight)

		for _, dependentFlight := range flightGraph[flight] {
			flightDependencyCount[dependentFlight] -= 1
		}

		delete(flightDependencyCount, flight)

		for flight, dependencyCount := range flightDependencyCount {
			if dependencyCount == 0 {
				flightEliminationQueue = append(flightEliminationQueue, flight)
			}
		}
	}

	return Flight{Source: sortedFlights[0], Destination: sortedFlights[len(sortedFlights)-1]}, nil
}
