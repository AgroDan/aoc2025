package lights

import (
	"fmt"
	"strings"
	"utils"
)

// Now the basic rule applies before, but I'm going to instead consider
// the current state and the end state to be a pathway, and BFS will
// find the shortest path through the state space. Since I don't have
// any negative values for the buttons, any pathway that increments
// a value OVER the joltage level can be thrown out immediately. This
// should help to make this more efficient.

// This is a datatype reprepresentative of the state of a machine,
// so I can use this instead of modifying the Machine struct
type State []int

func encodeState(s State) string {
	// encodes as a string for use in map keys
	st := make([]string, len(s))
	for i, v := range s {
		st[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(st, ",")
}

func pushButton(s State, button []int) State {
	// Given a state and a button, simply returns the new state
	// after pushing said button. This is easier to make atomic
	// actions like this
	newState := make(State, len(s))

	// Make newstate the exact copy of s first
	copy(newState, s)

	for _, idx := range button {
		if idx < 0 || idx >= len(s) {
			panic("Button index out of range")
		}
		newState[idx]++
	}
	return newState
}

// Now check to see if our current state exceeds the target, if so
// we'd ordinarily throw it out
func exceedsTarget(s State, target State) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > target[i] {
			return true
		}
	}
	return false
}

func SolveJoltMachine(m *JoltMachine) (int, error) {
	// BFS for shortest path
	// note, any state over the target will be thrown out for speed
	initialState := make(State, len(m.Joltage))
	// this should be initialized to zero, no?
	// ...after testing, yes!

	startKey := encodeState(initialState)
	targetKey := encodeState(m.Joltage)

	if startKey == targetKey {
		return 0, nil // already solved, EZ!
	}

	// BFS setup
	type node struct {
		state State
		key   string
	}

	// set up my queue
	queue := utils.NewGQueue[node]()
	queue.Enqueue(node{state: initialState, key: startKey})

	// and the visited set...
	visited := utils.NewGSet[string]()
	visited.Add(startKey)

	// predecessor for path reconstruction
	type predInfo struct {
		prevKey   string
		buttonIdx int
	}

	predecessor := make(map[string]predInfo)

	// now let's rock
	for !queue.IsEmpty() {
		currentNode, _ := queue.Dequeue()

		// Try all buttons from here
		for bIdx, button := range m.Buttons {
			nextState := pushButton(currentNode.state, button)
			nextKey := encodeState(nextState)

			// check if we've exceeded target
			if exceedsTarget(nextState, m.Joltage) {
				continue // throw out
			}

			// check if we've visited
			if visited.Contains(nextKey) {
				continue // already seen
			}

			// Now let's run this new state
			visited.Add(nextKey)
			predecessor[nextKey] = predInfo{
				prevKey:   currentNode.key,
				buttonIdx: bIdx,
			}

			if nextKey == targetKey {
				// Trace back the path because we found a valid solution
				path := []int{}
				currKey := nextKey
				for currKey != startKey {
					p := predecessor[currKey]
					path = append(path, p.buttonIdx)
					currKey = p.prevKey
				}

				// reverse the path now since we built it backwards
				for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
					path[i], path[j] = path[j], path[i]
				}

				return len(path), nil
			}

			// Not solved yet, stick it in the queue
			queue.Enqueue(node{state: nextState, key: nextKey})
		}
	}

	// Unsolveable
	return 0, fmt.Errorf("no solution found")
}
