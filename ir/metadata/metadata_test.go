package metadata_test

import (
	"github.com/mewmew/l/ir/metadata"
)

// Ensure that each metadata field implements the metadata.MDField interface.
var (
	// MDField
	_ metadata.MDField = (*metadata.MDNull)(nil)
	// Metadata
	_ metadata.MDField = (*metadata.Value)(nil)
	_ metadata.MDField = (*metadata.MDString)(nil)
	_ metadata.MDField = (*metadata.MDTuple)(nil)
	_ metadata.MDField = (*metadata.MetadataDef)(nil)
	// SpecializedMDNode
	_ metadata.MDField = (*metadata.DICompileUnit)(nil)
	_ metadata.MDField = (*metadata.DIFile)(nil)
	_ metadata.MDField = (*metadata.DIBasicType)(nil)
	_ metadata.MDField = (*metadata.DISubroutineType)(nil)
	_ metadata.MDField = (*metadata.DIDerivedType)(nil)
	_ metadata.MDField = (*metadata.DICompositeType)(nil)
	_ metadata.MDField = (*metadata.DISubrange)(nil)
	_ metadata.MDField = (*metadata.DIEnumerator)(nil)
	_ metadata.MDField = (*metadata.DITemplateTypeParameter)(nil)
	_ metadata.MDField = (*metadata.DITemplateValueParameter)(nil)
	_ metadata.MDField = (*metadata.DIModule)(nil)
	_ metadata.MDField = (*metadata.DINamespace)(nil)
	_ metadata.MDField = (*metadata.DIGlobalVariable)(nil)
	_ metadata.MDField = (*metadata.DISubprogram)(nil)
	_ metadata.MDField = (*metadata.DILexicalBlock)(nil)
	_ metadata.MDField = (*metadata.DILexicalBlockFile)(nil)
	_ metadata.MDField = (*metadata.DILocation)(nil)
	_ metadata.MDField = (*metadata.DILocalVariable)(nil)
	_ metadata.MDField = (*metadata.DIExpression)(nil)
	_ metadata.MDField = (*metadata.DIGlobalVariableExpression)(nil)
	_ metadata.MDField = (*metadata.DIObjCProperty)(nil)
	_ metadata.MDField = (*metadata.DIImportedEntity)(nil)
	_ metadata.MDField = (*metadata.DIMacro)(nil)
	_ metadata.MDField = (*metadata.DIMacroFile)(nil)
	_ metadata.MDField = (*metadata.GenericDINode)(nil)
)

// Ensure that each metadata implements the metadata.Metadata interface.
var (
	// Metadata
	_ metadata.Metadata = (*metadata.Value)(nil)
	_ metadata.Metadata = (*metadata.MDString)(nil)
	_ metadata.Metadata = (*metadata.MDTuple)(nil)
	_ metadata.Metadata = (*metadata.MetadataDef)(nil)
	// SpecializedMDNode
	_ metadata.Metadata = (*metadata.DICompileUnit)(nil)
	_ metadata.Metadata = (*metadata.DIFile)(nil)
	_ metadata.Metadata = (*metadata.DIBasicType)(nil)
	_ metadata.Metadata = (*metadata.DISubroutineType)(nil)
	_ metadata.Metadata = (*metadata.DIDerivedType)(nil)
	_ metadata.Metadata = (*metadata.DICompositeType)(nil)
	_ metadata.Metadata = (*metadata.DISubrange)(nil)
	_ metadata.Metadata = (*metadata.DIEnumerator)(nil)
	_ metadata.Metadata = (*metadata.DITemplateTypeParameter)(nil)
	_ metadata.Metadata = (*metadata.DITemplateValueParameter)(nil)
	_ metadata.Metadata = (*metadata.DIModule)(nil)
	_ metadata.Metadata = (*metadata.DINamespace)(nil)
	_ metadata.Metadata = (*metadata.DIGlobalVariable)(nil)
	_ metadata.Metadata = (*metadata.DISubprogram)(nil)
	_ metadata.Metadata = (*metadata.DILexicalBlock)(nil)
	_ metadata.Metadata = (*metadata.DILexicalBlockFile)(nil)
	_ metadata.Metadata = (*metadata.DILocation)(nil)
	_ metadata.Metadata = (*metadata.DILocalVariable)(nil)
	_ metadata.Metadata = (*metadata.DIExpression)(nil)
	_ metadata.Metadata = (*metadata.DIGlobalVariableExpression)(nil)
	_ metadata.Metadata = (*metadata.DIObjCProperty)(nil)
	_ metadata.Metadata = (*metadata.DIImportedEntity)(nil)
	_ metadata.Metadata = (*metadata.DIMacro)(nil)
	_ metadata.Metadata = (*metadata.DIMacroFile)(nil)
	_ metadata.Metadata = (*metadata.GenericDINode)(nil)
)

// Ensure that each metadata node implements the metadata.MDNode interface.
var (
	// MDNode
	_ metadata.MDNode = (*metadata.MDTuple)(nil)
	_ metadata.MDNode = (*metadata.MetadataDef)(nil)
	// SpecializedMDNode
	_ metadata.MDNode = (*metadata.DICompileUnit)(nil)
	_ metadata.MDNode = (*metadata.DIFile)(nil)
	_ metadata.MDNode = (*metadata.DIBasicType)(nil)
	_ metadata.MDNode = (*metadata.DISubroutineType)(nil)
	_ metadata.MDNode = (*metadata.DIDerivedType)(nil)
	_ metadata.MDNode = (*metadata.DICompositeType)(nil)
	_ metadata.MDNode = (*metadata.DISubrange)(nil)
	_ metadata.MDNode = (*metadata.DIEnumerator)(nil)
	_ metadata.MDNode = (*metadata.DITemplateTypeParameter)(nil)
	_ metadata.MDNode = (*metadata.DITemplateValueParameter)(nil)
	_ metadata.MDNode = (*metadata.DIModule)(nil)
	_ metadata.MDNode = (*metadata.DINamespace)(nil)
	_ metadata.MDNode = (*metadata.DIGlobalVariable)(nil)
	_ metadata.MDNode = (*metadata.DISubprogram)(nil)
	_ metadata.MDNode = (*metadata.DILexicalBlock)(nil)
	_ metadata.MDNode = (*metadata.DILexicalBlockFile)(nil)
	_ metadata.MDNode = (*metadata.DILocation)(nil)
	_ metadata.MDNode = (*metadata.DILocalVariable)(nil)
	_ metadata.MDNode = (*metadata.DIExpression)(nil)
	_ metadata.MDNode = (*metadata.DIGlobalVariableExpression)(nil)
	_ metadata.MDNode = (*metadata.DIObjCProperty)(nil)
	_ metadata.MDNode = (*metadata.DIImportedEntity)(nil)
	_ metadata.MDNode = (*metadata.DIMacro)(nil)
	_ metadata.MDNode = (*metadata.DIMacroFile)(nil)
	_ metadata.MDNode = (*metadata.GenericDINode)(nil)
)

// Ensure that each DIExpression field implements the metadata.DIExpressionField
// interface.
var (
	// DIExpressionField
	_ metadata.DIExpressionField = (*metadata.MDInt)(nil)
	_ metadata.DIExpressionField = metadata.DwarfOp(0)
)

// Ensure that each specialized metadata node implements the
// metadata.SpecializedMDNode interface.
var (
	// SpecializedMDNode
	_ metadata.SpecializedMDNode = (*metadata.DICompileUnit)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIFile)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIBasicType)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DISubroutineType)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIDerivedType)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DICompositeType)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DISubrange)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIEnumerator)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DITemplateTypeParameter)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DITemplateValueParameter)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIModule)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DINamespace)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIGlobalVariable)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DISubprogram)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DILexicalBlock)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DILexicalBlockFile)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DILocation)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DILocalVariable)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIExpression)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIGlobalVariableExpression)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIObjCProperty)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIImportedEntity)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIMacro)(nil)
	_ metadata.SpecializedMDNode = (*metadata.DIMacroFile)(nil)
	_ metadata.SpecializedMDNode = (*metadata.GenericDINode)(nil)
)

// Ensure that each integer and metadata field implements the
// metadata.IntOrMDField interface.
var (
	// IntOrMDField
	_ metadata.IntOrMDField = (*metadata.MDInt)(nil)
	// MDField
	_ metadata.IntOrMDField = (*metadata.MDNull)(nil)
	// Metadata
	_ metadata.IntOrMDField = (*metadata.Value)(nil)
	_ metadata.IntOrMDField = (*metadata.MDString)(nil)
	_ metadata.IntOrMDField = (*metadata.MDTuple)(nil)
	_ metadata.IntOrMDField = (*metadata.MetadataDef)(nil)
	// SpecializedMDNode
	_ metadata.IntOrMDField = (*metadata.DICompileUnit)(nil)
	_ metadata.IntOrMDField = (*metadata.DIFile)(nil)
	_ metadata.IntOrMDField = (*metadata.DIBasicType)(nil)
	_ metadata.IntOrMDField = (*metadata.DISubroutineType)(nil)
	_ metadata.IntOrMDField = (*metadata.DIDerivedType)(nil)
	_ metadata.IntOrMDField = (*metadata.DICompositeType)(nil)
	_ metadata.IntOrMDField = (*metadata.DISubrange)(nil)
	_ metadata.IntOrMDField = (*metadata.DIEnumerator)(nil)
	_ metadata.IntOrMDField = (*metadata.DITemplateTypeParameter)(nil)
	_ metadata.IntOrMDField = (*metadata.DITemplateValueParameter)(nil)
	_ metadata.IntOrMDField = (*metadata.DIModule)(nil)
	_ metadata.IntOrMDField = (*metadata.DINamespace)(nil)
	_ metadata.IntOrMDField = (*metadata.DIGlobalVariable)(nil)
	_ metadata.IntOrMDField = (*metadata.DISubprogram)(nil)
	_ metadata.IntOrMDField = (*metadata.DILexicalBlock)(nil)
	_ metadata.IntOrMDField = (*metadata.DILexicalBlockFile)(nil)
	_ metadata.IntOrMDField = (*metadata.DILocation)(nil)
	_ metadata.IntOrMDField = (*metadata.DILocalVariable)(nil)
	_ metadata.IntOrMDField = (*metadata.DIExpression)(nil)
	_ metadata.IntOrMDField = (*metadata.DIGlobalVariableExpression)(nil)
	_ metadata.IntOrMDField = (*metadata.DIObjCProperty)(nil)
	_ metadata.IntOrMDField = (*metadata.DIImportedEntity)(nil)
	_ metadata.IntOrMDField = (*metadata.DIMacro)(nil)
	_ metadata.IntOrMDField = (*metadata.DIMacroFile)(nil)
	_ metadata.IntOrMDField = (*metadata.GenericDINode)(nil)
)
