# flight-sorter
__Story__: There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

__Goal__: To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

Required JSON structure: 
- [["SFO", "EWR"]]                                                 => ["SFO", "EWR"]
- [["ATL", "EWR"], ["SFO", "ATL"]]                                 => ["SFO", "EWR"]
- [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] => ["SFO", "EWR"]

### Solution

The chunk of the problem is a topological sorting problem. Since flights are directed (meaning going from one place to another is not commutative), the problem can then be modeled to a DAG poblem. Topological sorting is the well known solution for this and I specifically applied Kahn's algorihtm to solve the problem. However, I have identified some peculiar scenarios that could happen when a user is using the endpoint. For example, a user might enter paths that will result to the person being in 2 simultaneous flights. I added validation to catch those. Ultimately, the solution will only return an answer when it is certain that a person is in an exact source and a destination, in other words, it filters out invalid values in this case, or in a more practical sense, it imposes the person to be in a single flight at a given time.

### How to run?

First, run `go mod download`

Then run `go build .` from the root directory

Then run `./flight-sorter` from the root directory