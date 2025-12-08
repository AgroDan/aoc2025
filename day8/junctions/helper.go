package junctions

import (
	"fmt"
	"math"
	"slices"
	"utils"
)

// This will do all the "heavy lifting" for the junctions, mostly
// determining the links between all the junctions based on distance.
// sometimes the shortest distance will be whoever is linking back to
// the junction checking, so if that's the case then it is what it is

func LinkJunctions(junctions []*Junction) {
	for i, junc := range junctions {

		// Create the shortest link number as infinity
		shortestDistance := math.Inf(1)

		for k, other := range junctions {
			if i == k {
				// Can't link to itself
				continue
			}

			// calculate difference
			distance := junc.Distance(other)

			// Less than shortest, make it the shortest
			if distance < shortestDistance {
				shortestDistance = distance
				junc.Link = other
			}
		}
	}
}

// Now that all the junctions are linked, get the top 3 longest circuits
// so we have to follow the amount of linked junctions and get the top 3
func Top3LongestLinks(junctions []*Junction) (int, int, int) {
	circuits := make([]int, 0, len(junctions))
	var exists = struct{}{}

	for _, junc := range junctions {
		seen := make(map[*Junction]struct{})
		current := junc
		count := 0
		var prevLink *Junction = nil

		for {
			// If we've seen this junction before, break
			if prevLink == current.Link {
				// we're in a loop of 2
				break
			}

			if _, ok := seen[current]; ok {
				// we've seen this one before
				count++
				break
			}
			seen[current] = exists
			count++

			prevLink = current
			current = current.Link
		}
		circuits = append(circuits, count)
	}

	// I could just sort this in reverse order with the SliceFunc function,
	// but I'm just lazy so
	slices.Sort(circuits)

	fmt.Printf("Debug: %v\n", circuits)
	return circuits[len(circuits)-1], circuits[len(circuits)-2], circuits[len(circuits)-3]
}

// Going to ignore the above, I googled this and am choosing a different approach, kinda.
// Going to create the edges of each junction to the shortest distance, then using DFS
// to find the longest path of each.

type Edge struct {
	From, To *Junction
	Distance float64
}

func CreateEdges(junctions []*Junction) []Edge {
	edges := make([]Edge, 0)

	for i, junc := range junctions {
		best_edge := Edge{}
		best_distance := math.Inf(1)

		for k, other := range junctions {
			if i == k {
				continue
			}
			distance := junc.Distance(other)
			if distance < best_distance {
				best_distance = distance
				best_edge = Edge{
					From:     junc,
					To:       other,
					Distance: distance,
				}
			}
		}
		edges = append(edges, best_edge)
	}

	return edges
}

// Now to build the adjacency list from the edges, using DFS to find the longest paths

func BuildAdjacencyList(edges []Edge) map[*Junction]map[*Junction]struct{} {
	// So this is a weird one, but hear me out. It creates a _map_ of a set
	// junctions. So the idea is that if you want to see if junction A is connected
	// to junction B, you do adjacencyList[A][B] and if it exists, then they are connected.

	adjacencyList := make(map[*Junction]map[*Junction]struct{})
	var exists = struct{}{} // this is just for the filler, since unbound structs take no space

	for _, edge := range edges {
		// For every single edge determined from the above
		if _, ok := adjacencyList[edge.From]; !ok {
			adjacencyList[edge.From] = make(map[*Junction]struct{})
		}
		adjacencyList[edge.From][edge.To] = exists

		// Also, make it bidirectional because it's considered undirected
		if _, ok := adjacencyList[edge.To]; !ok {
			adjacencyList[edge.To] = make(map[*Junction]struct{})
		}
		adjacencyList[edge.To][edge.From] = exists
	}

	return adjacencyList
}

// Now to do DFS and find the longest path from each junction. And since we're looking
// for singular circuits of each junction, when we _begin_ on a node, if we've visited
// it before, skip it because it's already been counted in another circuit.

func GetCircuits(adjacencyList map[*Junction]map[*Junction]struct{}) [][]*Junction {
	visited := make(map[*Junction]struct{})
	var exists = struct{}{}
	circuits := make([][]*Junction, 0)

	for j := range adjacencyList {

		if _, ok := visited[j]; ok {
			// already visited in another circuit
			continue
		}

		// New circuit
		currentCircuit := make([]*Junction, 0)
		currentCircuit = append(currentCircuit, j)

		// mark as visited
		visited[j] = exists

		// Add this path to the stack
		pathStack := utils.NewGStack[*Junction]()
		pathStack.Push(j)

		for !pathStack.IsEmpty() {
			currentNode, _ := pathStack.Pop()
			for neighbor := range adjacencyList[currentNode] {
				if _, ok := visited[neighbor]; !ok {
					// not visited yet
					visited[neighbor] = exists
					currentCircuit = append(currentCircuit, neighbor)
					pathStack.Push(neighbor)
				}
			}
		}

		circuits = append(circuits, currentCircuit)
	}

	return circuits
}

// Finally, get the top 3 longest networks
func Top3LongestNetworks(junctions []*Junction) (int, int, int) {
	edges := CreateEdges(junctions)
	fmt.Printf("Debug: Edges: %v\n", edges)
	adjacencyList := BuildAdjacencyList(edges)
	fmt.Printf("Debug: Adjacency List: %v\n", adjacencyList)
	circuits := GetCircuits(adjacencyList)
	fmt.Printf("Debug: Circuits: %v\n", circuits)

	lengths := make([]int, 0, len(circuits))
	for _, circuit := range circuits {
		lengths = append(lengths, len(circuit))
	}

	fmt.Printf("Debug: Circuits lengths: %v\n", lengths)

	slices.Sort(lengths)

	return lengths[len(lengths)-1], lengths[len(lengths)-2], lengths[len(lengths)-3]
}

// And now that I've officially overcomplicated this, I also figured out what the puzzle
// _ACTUALLY_ wanted, which is every single pair sorted into the smallest distances I guess.

func ComputePairwiseDistances(junctions []*Junction) []Edge {
	pairs := make([]Edge, 0)

	for i, junc := range junctions {
		for j, other := range junctions {
			if i >= j {
				// avoid duplicates and self-pairing
				continue
			}
			distance := junc.Distance(other)
			pairs = append(pairs, Edge{
				From:     junc,
				To:       other,
				Distance: distance,
			})
		}
	}

	// Now sort by distance
	slices.SortFunc(pairs, func(a, b Edge) int {
		if a.Distance < b.Distance {
			return -1
		} else if a.Distance > b.Distance {
			return 1
		} else {
			return 0
		}
	})

	return pairs
}

// Now a union find given the sorted edges
func Top3LongestCircuitsWithWire(junctions []*Junction, wireAmt int) (int, int, int) {
	pairs := ComputePairwiseDistances(junctions)

	// Union-Find structure
	parent := make(map[*Junction]*Junction)
	rank := make(map[*Junction]int)

	var find func(*Junction) *Junction
	find = func(j *Junction) *Junction {
		if parent[j] != j {
			parent[j] = find(parent[j])
		}
		return parent[j]
	}

	union := func(a, b *Junction) {
		rootA := find(a)
		rootB := find(b)
		if rootA != rootB {
			if rank[rootA] < rank[rootB] {
				parent[rootA] = rootB
			} else if rank[rootA] > rank[rootB] {
				parent[rootB] = rootA
			} else {
				parent[rootB] = rootA
				rank[rootA]++
			}
		}
	}

	// Initialize union-find structure
	for _, junc := range junctions {
		parent[junc] = junc
		rank[junc] = 0
	}

	// We only care about the wireAmt number of edges
	for i := 0; i < wireAmt && i < len(pairs); i++ {
		edge := pairs[i]
		if find(edge.From) != find(edge.To) {
			union(edge.From, edge.To)
		}
	}

	// Now count the sizes of each connected component
	componentSize := make(map[*Junction]int)
	for _, junc := range junctions {
		root := find(junc)
		componentSize[root]++
	}

	// fmt.Printf("Debug: Component sizes: %v\n", componentSize)

	sizes := make([]int, 0, len(componentSize))
	for _, size := range componentSize {
		sizes = append(sizes, size)
	}

	slices.Sort(sizes)

	// fmt.Printf("Debug: Sorted sizes: %v\n", sizes)

	if len(sizes) < 3 {
		// Not enough components
		for len(sizes) < 3 {
			sizes = append(sizes, 0)
		}
	}

	return sizes[len(sizes)-1], sizes[len(sizes)-2], sizes[len(sizes)-3]
}

// Now part two...I'm going to essentially mimic the above function
// but with a different goal. This time I'm going to keep on connecting
// junctions until only one single gigantic circuit is created, and I will
// return the two junctions connected to accomplish that.

func ConnectAllJunctions(junctions []*Junction) (*Junction, *Junction, error) {
	pairs := ComputePairwiseDistances(junctions)

	// Union-Find structure
	parent := make(map[*Junction]*Junction)
	rank := make(map[*Junction]int)

	// yup yup same song and dance here
	var find func(*Junction) *Junction
	find = func(j *Junction) *Junction {
		if parent[j] != j {
			parent[j] = find(parent[j])
		}
		return parent[j]
	}

	union := func(a, b *Junction) {
		rootA := find(a)
		rootB := find(b)
		if rootA != rootB {
			if rank[rootA] < rank[rootB] {
				parent[rootA] = rootB
			} else if rank[rootA] > rank[rootB] {
				parent[rootB] = rootA
			} else {
				parent[rootB] = rootA
				rank[rootA]++
			}
		}
	}

	// Initialize union-find structure
	for _, junc := range junctions {
		parent[junc] = junc
		rank[junc] = 0
	}

	// Now go through the pairs until all junctions are connected
	totalComponents := len(junctions)
	for _, edge := range pairs {
		if find(edge.From) != find(edge.To) {
			union(edge.From, edge.To)
			totalComponents--
			if totalComponents == 1 {
				// All connected now
				return edge.From, edge.To, nil
			}
		}
	}

	return nil, nil, fmt.Errorf("Could not connect all junctions")
}
