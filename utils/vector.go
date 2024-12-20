package utils

type Vector [2]int

var (
	DirUp        = Vector{0, -1}
	DirDown      = Vector{0, 1}
	DirLeft      = Vector{-1, 0}
	DirRight     = Vector{1, 0}
	DirUpLeft    = Vector{-1, -1}
	DirUpRight   = Vector{1, -1}
	DirDownLeft  = Vector{-1, 1}
	DirDownRight = Vector{1, 1}
)

var AllDirections = []Vector{
	DirUp,
	DirDown,
	DirLeft,
	DirRight,
	DirUpLeft,
	DirUpRight,
	DirDownLeft,
	DirDownRight,
}

var FourDirections = []Vector{
	DirUp,
	DirDown,
	DirLeft,
	DirRight,
}

// VectorCompare - True if both vectors have the same position
func VectorCompare(a, b Vector) bool {
	return a[0] == b[0] && a[1] == b[1]
}

// VectorAdd - Simply adds to 2 vectors and return it
func VectorAdd(a, b Vector) Vector {
	return Vector{a[0] + b[0], a[1] + b[1]}
}

// VectorSub - Subtracts vector b from vector a
func VectorSub(a, b Vector) Vector {
	return Vector{a[0] - b[0], a[1] - b[1]}
}

// VectorTurnRight - Turn a vector 90 degrees to the right
func VectorTurnRight(a Vector) Vector {
	return Vector{-a[1], a[0]}
}
