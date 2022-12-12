package api

import (
	"fmt"
	"net/http"

	"github.com/Bexanderthebex/flight-sorter/internal/sorter"
	"github.com/labstack/echo/v4"
)

func SortFlights(ctx echo.Context) error {
	var flightInput flights
	if err := ctx.Bind(&flightInput); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid input format. Should be an array of 2 element strings")
	}

	flightsToSort := make([]sorter.Flight, 0, len(flightInput))
	for _, f := range flightInput {
		inputLength := len(f)
		if inputLength != 2 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input format, expecting 2 array elements but got %d", inputLength))
		}

		flightsToSort = append(flightsToSort, sorter.Flight{Source: f[0], Destination: f[1]})
	}

	sortedFlightPath, err := sorter.TopologicalSort(flightsToSort)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(http.StatusOK, []string{sortedFlightPath.Source, sortedFlightPath.Destination})
}
