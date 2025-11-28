# Advent of Code 2025 - AI Coding Agent Instructions

## Project Overview
This is an Advent of Code 2025 workspace using **Go 1.23.4** with a shared `utils` package for common algorithms and data structures. Each daily challenge lives in a separate `dayN/` directory with its own Go module.

## Architecture

### Module Structure
- **Root**: Contains prep scripts and the shared `utils/` package
- **Daily challenges**: Each in `dayN/` directory with independent Go module
- **Go workspaces**: Uses `go work use dayN` to link daily modules to shared utils (no go.work file exists yet until first day is prepped)

### Coordinate System Convention
**Critical**: This codebase uses a **top-left origin** coordinate system where:
- `Y` increases going **DOWN** (reading order)
- `Y` decreases going **UP**
- `X` increases going **RIGHT**
- `X` decreases going **LEFT**

Cardinal directions in `utils/coords.go`:
```go
const (
    N = 0  // Y--
    E = 1  // X++
    S = 2  // Y++
    W = 3  // X--
)
```

## Daily Workflow

### Starting a New Day
Use the prep script for your platform to bootstrap a new day:
```powershell
# PowerShell (Windows)
$Env:AOC_SESSION="<your_session_cookie>"
.\prep.ps1 5

# Bash
export AOC_SESSION="<your_session_cookie>"
./prep.sh 5

# Fish
set -x AOC_SESSION "<your_session_cookie>"
source ./prep.fish 5
```

This automatically:
1. Creates `dayN/` directory
2. Downloads puzzle input to `dayN/input`
3. Generates `dayN/dayN.go` stub with standard imports
4. Runs `go mod init dayN`
5. Runs `go work use dayN` to link to utils

### Standard Day Stub Structure
Every generated `dayN.go` includes:
```go
package main

import (
    "fmt"
    "flag"
    "time"
    "utils"
)

func main() {
    t := time.Now()
    filePtr := flag.String("f", "input", "Input file if not 'input'")
    flag.Parse()

    // Choose input parser:
    // lines, err := utils.GetFileLines(*filePtr)  // Line-by-line
    // challengeText, err := utils.GetTextBlob(*filePtr)  // Whole file

    // Solution code here

    fmt.Printf("Total time elapsed: %s\n", time.Since(t))
}
```

### Running Solutions
```powershell
cd dayN
go run . -f input        # Use default input
go run . -f sample.txt   # Use alternate input file
```

## Utils Package Reference

### File Parsing (`fileparse.go`)
- `GetFileLines(filename string) ([]string, error)` - Returns slice of lines
- `GetTextBlob(filename string) (string, error)` - Returns entire file as string

### Coordinate & Grid Navigation (`coords.go`)
- `Coord{X, Y int}` struct with methods:
  - `Move(dir int)` - Modifies coord in-place
  - `Peek(dir int) Coord` - Returns new coord without modifying
  - `AllAvailable() []Coord` - Returns 4 cardinal neighbors (N/E/S/W)
  - `Neighbors() []Coord` - Returns 8 neighbors including diagonals
  - `Parallel(dir int) []Coord` - Returns coords in parallel direction
- Direction helpers: `TurnRight(dir)`, `TurnLeft(dir)`, `Opposite(dir)`
- `ManhattanDistance(c1, c2 Coord) int`
- `ManhattanRadius(c Coord, radius int) []Coord` - All points within Manhattan distance
- `IsInSquare(loc, topLeft, botRight Coord) bool`

### Runemap (`runemap.go`)
For 2D character grids:
- `NewRunemap(in []string) Runemap` - Create from string slice
- `IsInBounds(c Coord) bool`
- `Get(c Coord) (rune, error)` - Retrieve character at coord
- `Set(c Coord, setRune rune) error` - Modify character at coord
- `Find(items ...rune) (Coord, error)` - Find first occurrence
- `FindAll(item rune) []Coord` - Find all occurrences
- `DeepCopy() *Runemap` - Creates independent copy
- `Print()` - Display map to stdout
- `Width()`, `Height()` - Dimensions
- `GetRaw() [][]rune` - Access underlying slice

### Data Structures

**Generic (preferred for new code)**:
- `GQueue[T]` (`gQueue.go`) - Generic FIFO queue with `Enqueue`, `Dequeue`, `Peek`, `IsEmpty`
- `GStack[T]` (`gStack.go`) - Generic LIFO stack with `Push`, `Pop`, `Peek`, `IsEmpty`

**Legacy (interface{} based)**:
- `Queue` (`queue.go`) - Use generics instead for new code
- `Stack` (`stack.go`) - Use generics instead for new code

**Other**:
- `set` (`sets.go`) - String-based set with `Add`, `Remove`, `Contains`
- `Breadcrumb` (`breadcrumbs.go`) - Tracks visited coords with direction:
  - `Add(c Coord, dir int)`, `Contains(c Coord)`, `GetDir(c Coord)`, `DeepCopy()`

### Algorithms & Math
- `Cache` (`cache.go`) - Memoization wrapper:
  ```go
  cache := NewCache()
  result := cache.Get("key", func() interface{} {
      return expensiveComputation()
  }).(ExpectedType)
  ```
- `EuclideanModulus(x, y int)` (`math.go`) - Proper modulo for wrapping (e.g., `-1 % 10 = 9`)
- `Heap(arr []int, n int, result *[][]int)` (`heap.go`) - Generate permutations
- `CartesianProduct[T any](slice1, slice2 []T) [][]T` (`heap.go`)

### Utilities
- `ReverseString(s string)` (`misc.go`)
- `ZFill(s string, width int)` (`misc.go`) - Zero-pad string
- `FlushScreen(s string)` (`debug.go`) - Update single line in terminal with `\r`

## Common Patterns

### Grid-Based Puzzles
```go
lines, _ := utils.GetFileLines(*filePtr)
grid := utils.NewRunemap(lines)
start, _ := grid.Find('S')

for _, neighbor := range start.AllAvailable() {
    if grid.IsInBounds(neighbor) {
        char, _ := grid.Get(neighbor)
        // Process neighbor
    }
}
```

### BFS with Visited Tracking
```go
queue := utils.NewGQueue[utils.Coord]()
visited := utils.NewBreadcrumb()
queue.Enqueue(start)
visited.Add(start, 0)

for !queue.IsEmpty() {
    current, _ := queue.Dequeue()
    for _, next := range current.AllAvailable() {
        if !visited.Contains(next) && isValid(next) {
            queue.Enqueue(next)
            visited.Add(next, directionToNext)
        }
    }
}
```

### Memoized Recursion
```go
cache := utils.NewCache()
var solve func(state string) int
solve = func(state string) int {
    return cache.Get(state, func() interface{} {
        // Base case
        if isBaseCase(state) { return result }
        // Recursive computation
        return solve(nextState)
    }).(int)
}
```

## Notes
- Prefer **generic data structures** (`GQueue`, `GStack`) over legacy `interface{}` versions
- Always check `IsInBounds()` before `Get()`/`Set()` on Runemap
- Use `EuclideanModulus()` for wrapping coordinates on toroidal grids
- Add timing with `time.Since(t)` pattern already in stub
- Input files are named `input` by default; use `-f` flag for samples
