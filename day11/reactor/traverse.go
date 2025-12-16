package reactor

// I'm going to start this with a Depth-First Search (DFS) traversal pattern
// to start...knowing FULL WELL this will not be optimal for part 2, whatever
// that consists of. However all it wants is the total amount of pathways
// from "you" to "out", so...I'll start there and deal with whatever happens
// next.

func TraverseDFS(current *Device, target string, deviceMap map[string]*Device, visited map[string]bool) int {

	// note there is no "out" node, we have to confirm this at the initial connection
	// if current.Name == target {
	// 	fmt.Printf("Reached target device: %s\n", target)
	// 	return 1
	// }

	// Mark the current device as visited
	visited[current.Name] = true
	totalPaths := 0

	// fmt.Printf("At device: %s, connections: %v\n", current.Name, current.Connections)
	// Explore each connection
	for _, connName := range current.Connections {
		// Check if this connection is the target (e.g., "out")
		if connName == target {
			// fmt.Printf("Reached target device: %s from device: %s\n", target, current.Name)
			totalPaths++
			continue
		}

		// Otherwise, continue traversing if not visited
		if !visited[connName] {
			if nextDevice, exists := deviceMap[connName]; exists {
				totalPaths += TraverseDFS(nextDevice, target, deviceMap, visited)
			}
		}
	}

	// Backtrack: unmark the current device as visited for other paths
	visited[current.Name] = false
	return totalPaths
}
