// Package value provides a definition of LLVM IR values.
package value

import (
	"fmt"

	"github.com/mewmew/l/ir/types"
)

// Value is an LLVM IR value.
type Value interface {
	// String returns the string representation of the value as a type-value
	// pair.
	fmt.Stringer
	// Type returns the type of the value.
	Type() types.Type
	// Ident returns the identifier associated with the value.
	Ident() string
}

// Named is a named LLVM IR value.
type Named interface {
	Value
	// Name returns the name of the value.
	GetName() string
	// SetName sets the name of the value.
	SetName(name string)
}
