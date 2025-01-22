package lgraph

type node uint

type edge struct {
	destination node
	label       rune
}

// LGraph is a function representing a directed labeled graph. If the node exists
// in the graph, the function returns true along with the set of outgoing edges
// from that node, otherwise false and nil.
type LGraph func(node) ([]edge, bool)

// FindSequence returns (S, true) if there is a sequence S of length k from node
// s to node t in graph g1 and S is not a sequence from s to t in graph g2; else
// it returns (nil, false).
//func FindSequence(g1, g2 LGraph, s, t node, k uint) ([]rune, bool) {
// TODO: Complete the function.
//	panic("TODO: implement this!")
//}
func FindSequence(g1, g2 LGraph, s, t node, k uint) ([]rune, bool) {
	// Handle the special case where k = 0
	if k == 0 {
		if s == t {
			// Empty sequence is valid if g1 recognizes s as a valid node
			_, existsInG1 := g1(s)
			_, existsInG2 := g2(s)
			if existsInG1 && !existsInG2 {
				return []rune{}, true
			}
		}
		return nil, false // k=0 but s != t, or invalid sequence
	}

	// Helper function to perform DFS and find sequences of length k
	var dfs func(current node, path []rune, length uint) ([]rune, bool)
	dfs = func(current node, path []rune, length uint) ([]rune, bool) {
		// Base case: if we reach length k
		if length == k {
			if current == t {
				// Ensure this sequence does not exist in g2
				if existsInGraph(g2, s, t, path) {
					return nil, false // Sequence exists in both graphs
				}
				return path, true // Sequence exists only in g1
			}
			return nil, false // Incorrect path length
		}

		// Get neighbors from graph g1
		neighbors, exists := g1(current)
		if !exists {
			return nil, false
		}

		// Explore each outgoing edge
		for _, e := range neighbors {
			newPath := append(path, e.label)
			if res, found := dfs(e.destination, newPath, length+1); found {
				return res, true
			}
		}
		return nil, false
	}

	// Start DFS from node s
	return dfs(s, []rune{}, 0)
}

// Helper function to check if a sequence exists in graph g
func existsInGraph(g LGraph, start, end node, sequence []rune) bool {
	current := start
	for _, label := range sequence {
		neighbors, exists := g(current)
		if !exists {
			return false
		}
		found := false
		for _, e := range neighbors {
			if e.label == label {
				current = e.destination
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return current == end
}
