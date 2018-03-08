package irx

import (
	"fmt"

	"github.com/mewmew/l/asm/internal/ast"
	"github.com/mewmew/l/ir"
	"github.com/mewmew/l/ir/metadata"
	"github.com/mewmew/l/ll"
	"github.com/mewmew/l/ll/types"
)

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	// Create maps from identifier to definition.
	m := NewModule()
	m.indexIdents(module.Entities)

	// Resolve type definitions.
	m.resolveTypeDefs()

	// Resolve global variables, indirect symbols and functions.
	m.indexGlobals()
	m.resolveGlobals()

	// Resolve metadata definitions.
	m.indexMetadataDefs()
	m.resolveMetadataDefs()

	return m.Module, nil
}

// indexIdents indexes top-level entity definitions.
func (m *Module) indexIdents(entities []ast.TopLevelEntity) {
	// Create maps from identifier to definition.
	for _, entity := range entities {
		switch entity := entity.(type) {
		case *types.NamedType:
			// Type definitions.
			ident := entity.Name
			if prev, ok := m.localIdent[ident]; ok {
				panic(fmt.Errorf("type identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.localIdent[ident] = entity
		case *ll.ComdatDef:
			// Comdat definitions.
			name := entity.Name
			if prev, ok := m.comdatName[name]; ok {
				panic(fmt.Errorf("comdat name %q already present; prev `%v`, new `%v`", name, prev, entity))
			}
			m.comdatName[name] = entity
		case *ir.Global:
			// Global declarations and definitions.
			ident := entity.Name
			if prev, ok := m.globalIdent[ident]; ok {
				panic(fmt.Errorf("global identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.globalIdent[ident] = entity
		case *ast.IndirectSymbol:
			// Indirect symbol definitions.
			ident := entity.Name
			if prev, ok := m.globalIdent[ident]; ok {
				panic(fmt.Errorf("global identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.globalIdent[ident] = entity
		case *ast.Function:
			ident := entity.Header.Name
			if prev, ok := m.globalIdent[ident]; ok {
				panic(fmt.Errorf("global identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.globalIdent[ident] = entity
		case *ir.AttrGroupDef:
			// Attribute group definitions.
			id := entity.ID
			if prev, ok := m.attrGroupID[id]; ok {
				panic(fmt.Errorf("attribute group ID %q already present; prev `%v`, new `%v`", id, prev, entity))
			}
			m.attrGroupID[id] = entity
		case *metadata.NamedMetadataDef:
			// Named metadata definitions.
			name := entity.Name
			if prev, ok := m.metadataName[name]; ok {
				panic(fmt.Errorf("metadata name %q already present; prev `%v`, new `%v`", name, prev, entity))
			}
			m.metadataName[name] = entity
		case *metadata.MetadataDef:
			// Metadata definitions.
			id := entity.ID
			if prev, ok := m.metadataID[id]; ok {
				panic(fmt.Errorf("metadata ID %q already present; prev `%v`, new `%v`", id, prev, entity))
			}
			m.metadataID[id] = entity
		default:
			panic(fmt.Errorf("support for top-level entity %T not yet implemented", entity))
		}
	}
}
