# Go Slice Modification Bug

This repository demonstrates a common error when working with slices in Go.  Modifying a slice that's a view into another slice unexpectedly modifies the original slice.

The `bug.go` file contains the buggy code. The `bugSolution.go` provides a corrected version.

## Bug Description

In Go, slices are views into underlying arrays. When you create a new slice from an existing one (using slicing), the new slice shares the same underlying array.  Modifying elements of the new slice will modify the original slice.

## Solution

The solution involves creating a copy of the slice before modification to avoid unintended side effects.

## How to Reproduce

1.  Clone this repository.
2.  Navigate to the repository directory.
3.  Run the `bug.go` file. Observe that modifying `y` also modifies `x`.
4.  Run the `bugSolution.go` file. Observe that modifying the copy of `y` does not modify `x`.
