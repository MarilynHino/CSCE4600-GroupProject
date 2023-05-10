package builtins

import (
	"errors"
	"strconv"
)

// Returned variable when the ShiftCommand function is called with an invalid number of arguments.
var (
	ErrShiftInvalidArgCount = errors.New("invalid argument count")
)

// ShiftCommand takes a slice of strings as the first argument, and an optional argument specifying how many elements to shift the slice by.
func ShiftCommand(args *[]string, shiftBy ...string) error {
	// Check if the number of arguments is valid
	if len(shiftBy) > 1 {
		return ErrShiftInvalidArgCount
	}

	shift := 1
	if len(shiftBy) == 1 {
		// If an argument was provided, parse it as an integer
		var err error
		shift, err = strconv.Atoi(shiftBy[0])
		if err != nil {
			return ErrShiftInvalidArgCount
		}
	}

	if len(*args) < shift {
		// Return an error if there are not enough elements in the slice. 
		return ErrShiftInvalidArgCount
	}

	// Remove the specified number of elements from the beginning of the slice.
	*args = (*args)[shift:]
	return nil
}