package irx

import (
	"fmt"

	"github.com/mewmew/l/asm/internal/ast"
	"github.com/mewmew/l/ir"
	"github.com/mewmew/l/ir/metadata"
	"github.com/mewmew/l/ir/types"
)

// Translate translates the AST of the given module to an equivalent LLVM IR
// module.
func Translate(module *ast.Module) (*ir.Module, error) {
	// Create maps from identifier to definition.
	m := NewModule()
	m.indexIdents(module.Entities)

	// Resolve type definitions.
	//m.resolveTypeDefs()

	// Resolve global variables, indirect symbols and functions.
	//m.indexGlobals()
	//m.resolveGlobals()

	// Resolve metadata definitions.
	//m.indexMetadataDefs()
	//m.resolveMetadataDefs()

	return m.Module, nil
}

// indexIdents indexes top-level entity definitions.
func (m *Module) indexIdents(entities []ast.TopLevelEntity) {
	// Create maps from identifier to definition.
	for _, entity := range entities {
		switch entity := entity.(type) {
		case *ir.SourceFilename:
			// Source filename.
			m.SourceFilename = entity
		case *ir.DataLayout:
			// Data layout.
			m.DataLayout = entity
		case *ir.TargetTriple:
			// Target triple.
			m.TargetTriple = entity
		case *ir.ModuleAsm:
			// Module-level inline assembly.
			m.ModuleAsms = append(m.ModuleAsms, entity)
		case *types.NamedType:
			// Type definitions.
			m.TypeDefs = append(m.TypeDefs, entity)
			ident := entity.Name
			if prev, ok := m.typeDefs[ident]; ok {
				panic(fmt.Errorf("type identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.typeDefs[ident] = entity
		case *ir.ComdatDef:
			// Comdat definitions.
			m.ComdatDefs = append(m.ComdatDefs, entity)
			name := entity.Name
			if prev, ok := m.comdatDefs[name]; ok {
				panic(fmt.Errorf("comdat name %q already present; prev `%v`, new `%v`", name, prev, entity))
			}
			m.comdatDefs[name] = entity
		case *ir.Global:
			// Global declarations and definitions.
			m.Globals = append(m.Globals, entity)
			ident := entity.Name
			if prev, ok := m.global(ident); ok {
				panic(fmt.Errorf("global identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.globals[ident] = entity
		case *ir.IndirectSymbol:
			// Indirect symbol definitions (aliases and IFuncs).
			m.IndirectSymbols = append(m.IndirectSymbols, entity)
			ident := entity.Name
			if prev, ok := m.global(ident); ok {
				panic(fmt.Errorf("global identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.indirectSymbols[ident] = entity
		case *ir.Function:
			// Function declarations and definitions.
			m.Funcs = append(m.Funcs, entity)
			ident := entity.Name
			if prev, ok := m.global(ident); ok {
				panic(fmt.Errorf("global identifier %q already present; prev `%v`, new `%v`", ident, prev, entity))
			}
			m.funcs[ident] = entity
		case *ir.AttrGroupDef:
			// Attribute group definitions.
			m.AttrGroupDefs = append(m.AttrGroupDefs, entity)
			id := entity.ID
			if prev, ok := m.attrGroupDefs[id]; ok {
				panic(fmt.Errorf("attribute group ID %q already present; prev `%v`, new `%v`", id, prev, entity))
			}
			m.attrGroupDefs[id] = entity
		case *metadata.NamedMetadataDef:
			// Named metadata definitions.
			m.NamedMetadataDefs = append(m.NamedMetadataDefs, entity)
			name := entity.Name
			if prev, ok := m.namedMetadataDefs[name]; ok {
				panic(fmt.Errorf("metadata name %q already present; prev `%v`, new `%v`", name, prev, entity))
			}
			m.namedMetadataDefs[name] = entity
		case *metadata.MetadataDef:
			// Metadata definitions.
			m.MetadataDefs = append(m.MetadataDefs, entity)
			id := entity.ID
			if prev, ok := m.metadataDefs[id]; ok {
				panic(fmt.Errorf("metadata ID %q already present; prev `%v`, new `%v`", id, prev, entity))
			}
			m.metadataDefs[id] = entity
		case *ir.UseListOrder:
			// Use-list order directives.
			m.UseListOrders = append(m.UseListOrders, entity)
		case *ir.UseListOrderBB:
			// Basic block specific use-list order directives.
			m.UseListOrderBBs = append(m.UseListOrderBBs, entity)
		default:
			panic(fmt.Errorf("support for top-level entity %T not yet implemented", entity))
		}
	}
}
