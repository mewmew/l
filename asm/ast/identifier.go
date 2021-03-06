package ast

import (
	"fmt"

	"github.com/mewmew/l/internal/enc"
	"github.com/mewmew/l/ir/types"
)

// === [ Identifiers ] =========================================================

// --- [ Global Identifiers ] --------------------------------------------------

// A GlobalIdent is a global identifier (e.g. @foo, @42).
type GlobalIdent struct {
	Name string
}

// String returns the string representation of the global identifier as a type-
// value pair.
func (i *GlobalIdent) String() string {
	return fmt.Sprintf("%v %v", i.Type(), i.Ident())
}

// Type returns the type of the global identifier.
func (i *GlobalIdent) Type() types.Type {
	panic(fmt.Errorf("ast.GlobalIdent.Type: unintended use of Type for global ident %q; type resolution not complete", i.Ident()))
}

// Ident returns the identifier associated with the global identifier.
func (i *GlobalIdent) Ident() string {
	// global_ident
	return enc.Global(i.Name)
}

// --- [ Local Identifiers ] ---------------------------------------------------

// A LocalIdent is a local identifier (e.g. %foo, %42).
type LocalIdent struct {
	Name string
}

// String returns the string representation of the local identifier as a type-
// value pair.
func (i *LocalIdent) String() string {
	return fmt.Sprintf("%v %v", i.Type(), i.Ident())
}

// Type returns the type of the local identifier.
func (i *LocalIdent) Type() types.Type {
	panic(fmt.Errorf("ast.LocalIdent.Type: unintended use of Type for local ident %q; type resolution not complete", i.Ident()))
}

// Ident returns the identifier associated with the local identifier.
func (i *LocalIdent) Ident() string {
	// local_ident
	return enc.Local(i.Name)
}

// --- [ Label Identifiers ] ---------------------------------------------------

// A LabelIdent is a label identifier (e.g. foo:, 42:).
type LabelIdent struct {
	Name string
}

// String returns the string representation of the label identifier.
func (i *LabelIdent) String() string {
	// label_ident
	return enc.Label(i.Name)
}

// --- [ Attribute Group Identifiers ] -----------------------------------------

// An AttrGroupID is an attribute group ID (e.g. #42).
type AttrGroupID struct {
	ID string
}

// String returns the string representation of the  attribute group ID.
func (i *AttrGroupID) String() string {
	// attr_group_id
	return enc.AttrGroup(i.ID)
}

// --- [ Comdat Identifiers ] --------------------------------------------------

// A ComdatName is a comdat name (e.g. $foo).
type ComdatName struct {
	Name string
}

// String returns the string representation of the comdat name.
func (i *ComdatName) String() string {
	// comdat_name
	return enc.Comdat(i.Name)
}

// --- [ Metadata Identifiers ] ------------------------------------------------

// A MetadataName is a metadata name (e.g. !foo).
type MetadataName struct {
	Name string
}

// String returns the string representation of the metadata name.
func (i *MetadataName) String() string {
	// metadata_name
	return enc.Metadata(i.Name)
}

// A MetadataID is a metadata ID (e.g. !42).
type MetadataID struct {
	ID string
}

// String returns the string representation of the metadata ID.
func (i *MetadataID) String() string {
	// metadata_id
	return enc.Metadata(i.ID)
}
