// Code generated by gocc; DO NOT EDIT.

package token

import (
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const (
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line   int
	Column int
}

func (p Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", p.Offset, p.Line, p.Column)
}

type TokenMap struct {
	typeMap []string
	idMap   map[string]Type
}

func (m TokenMap) Id(tok Type) string {
	if int(tok) < len(m.typeMap) {
		return m.typeMap[tok]
	}
	return "unknown"
}

func (m TokenMap) Type(tok string) Type {
	if typ, exist := m.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (m TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", m.Id(tok.Type), tok.Type, tok.Lit)
}

func (m TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", m.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"empty",
		"source_filename",
		"=",
		"target",
		"triple",
		"datalayout",
		"module",
		"asm",
		"type",
		"comdat",
		"any",
		"exactmatch",
		"largest",
		"noduplicates",
		"samesize",
		"thread_local",
		"(",
		")",
		"initialexec",
		"localdynamic",
		"localexec",
		"externally_initialized",
		"constant",
		"global",
		",",
		"alias",
		"ifunc",
		"declare",
		"define",
		"amdgpu_cs",
		"amdgpu_es",
		"amdgpu_gs",
		"amdgpu_hs",
		"amdgpu_kernel",
		"amdgpu_ls",
		"amdgpu_ps",
		"amdgpu_vs",
		"anyregcc",
		"arm_aapcs_vfpcc",
		"arm_aapcscc",
		"arm_apcscc",
		"avr_intrcc",
		"avr_signalcc",
		"ccc",
		"coldcc",
		"cxx_fast_tlscc",
		"fastcc",
		"ghccc",
		"hhvm_ccc",
		"hhvmcc",
		"intel_ocl_bicc",
		"msp430_intrcc",
		"preserve_allcc",
		"preserve_mostcc",
		"ptx_device",
		"ptx_kernel",
		"spir_func",
		"spir_kernel",
		"swiftcc",
		"webkit_jscc",
		"win64cc",
		"x86_64_sysvcc",
		"x86_fastcallcc",
		"x86_intrcc",
		"x86_regcallcc",
		"x86_stdcallcc",
		"x86_thiscallcc",
		"x86_vectorcallcc",
		"cc",
		"int_lit",
		"gc",
		"prefix",
		"prologue",
		"personality",
		"{",
		"}",
		"attributes",
		"!",
		"distinct",
		"uselistorder",
		"uselistorder_bb",
		"global_ident",
		"local_ident",
		"label_ident",
		"attr_group_id",
		"comdat_name",
		"metadata_name",
		"metadata_id",
		"void",
		"int_type",
		"half",
		"float",
		"double",
		"x86_fp80",
		"fp128",
		"ppc_fp128",
		"x86_mmx",
		"*",
		"<",
		"x",
		">",
		"label",
		"token",
		"metadata",
		"[",
		"]",
		"opaque",
		"sideeffect",
		"alignstack",
		"inteldialect",
		"true",
		"false",
		"float_lit",
		"null",
		"none",
		"c",
		"string_lit",
		"zeroinitializer",
		"undef",
		"blockaddress",
		"add",
		"fadd",
		"sub",
		"fsub",
		"mul",
		"fmul",
		"udiv",
		"sdiv",
		"fdiv",
		"urem",
		"srem",
		"frem",
		"shl",
		"lshr",
		"ashr",
		"and",
		"or",
		"xor",
		"extractelement",
		"insertelement",
		"shufflevector",
		"extractvalue",
		"insertvalue",
		"getelementptr",
		"inrange",
		"trunc",
		"to",
		"zext",
		"sext",
		"fptrunc",
		"fpext",
		"fptoui",
		"fptosi",
		"uitofp",
		"sitofp",
		"ptrtoint",
		"inttoptr",
		"bitcast",
		"addrspacecast",
		"icmp",
		"fcmp",
		"select",
		"alloca",
		"inalloca",
		"swifterror",
		"load",
		"atomic",
		"store",
		"fence",
		"cmpxchg",
		"weak",
		"atomicrmw",
		"max",
		"min",
		"nand",
		"umax",
		"umin",
		"xchg",
		"phi",
		"call",
		"musttail",
		"notail",
		"tail",
		"va_arg",
		"landingpad",
		"cleanup",
		"catch",
		"filter",
		"catchpad",
		"within",
		"cleanuppad",
		"ret",
		"br",
		"switch",
		"indirectbr",
		"invoke",
		"unwind",
		"resume",
		"catchswitch",
		"caller",
		"catchret",
		"from",
		"cleanupret",
		"unreachable",
		"!DICompileUnit",
		"language:",
		"producer:",
		"flags:",
		"runtimeVersion:",
		"splitDebugFilename:",
		"emissionKind:",
		"enums:",
		"retainedTypes:",
		"globals:",
		"imports:",
		"macros:",
		"dwoId:",
		"splitDebugInlining:",
		"debugInfoForProfiling:",
		"gnuPubnames:",
		"!DIFile",
		"filename:",
		"directory:",
		"checksumkind:",
		"checksum:",
		"!DIBasicType",
		"encoding:",
		"!DISubroutineType",
		"cc:",
		"types:",
		"!DIDerivedType",
		"extraData:",
		"dwarfAddressSpace:",
		"!DICompositeType",
		"elements:",
		"runtimeLang:",
		"vtableHolder:",
		"identifier:",
		"discriminator:",
		"!DISubrange",
		"count:",
		"lowerBound:",
		"!DIEnumerator",
		"value:",
		"isUnsigned:",
		"!DITemplateTypeParameter",
		"!DITemplateValueParameter",
		"!DIModule",
		"configMacros:",
		"includePath:",
		"isysroot:",
		"!DINamespace",
		"exportSymbols:",
		"!DIGlobalVariable",
		"!DISubprogram",
		"scopeLine:",
		"containingType:",
		"virtuality:",
		"virtualIndex:",
		"thisAdjustment:",
		"unit:",
		"variables:",
		"thrownTypes:",
		"!DILexicalBlock",
		"!DILexicalBlockFile",
		"!DILocation",
		"inlinedAt:",
		"!DILocalVariable",
		"arg:",
		"!DIExpression",
		"!DIGlobalVariableExpression",
		"var:",
		"expr:",
		"!DIObjCProperty",
		"setter:",
		"getter:",
		"attributes:",
		"!DIImportedEntity",
		"entity:",
		"!DIMacro",
		"!DIMacroFile",
		"nodes:",
		"!GenericDINode",
		"header:",
		"operands:",
		"dwarf_lang",
		"file:",
		"isOptimized:",
		"FullDebug",
		"LineTablesOnly",
		"NoDebug",
		"checksum_kind",
		"tag:",
		"dwarf_tag",
		"name:",
		"size:",
		"align:",
		"dwarf_att_encoding",
		"|",
		"di_flag",
		"dwarf_cc",
		"line:",
		"scope:",
		"baseType:",
		"offset:",
		"templateParams:",
		"type:",
		"linkageName:",
		"isLocal:",
		"isDefinition:",
		"declaration:",
		"dwarf_virtuality",
		"column:",
		"dwarf_op",
		"dwarf_macinfo",
		"align",
		"alwaysinline",
		"argmemonly",
		"builtin",
		"cold",
		"convergent",
		"inaccessiblemem_or_argmemonly",
		"inaccessiblememonly",
		"inlinehint",
		"jumptable",
		"minsize",
		"naked",
		"nobuiltin",
		"noduplicate",
		"noimplicitfloat",
		"noinline",
		"nonlazybind",
		"norecurse",
		"noredzone",
		"noreturn",
		"nounwind",
		"optnone",
		"optsize",
		"readnone",
		"readonly",
		"returns_twice",
		"safestack",
		"sanitize_address",
		"sanitize_hwaddress",
		"sanitize_memory",
		"sanitize_thread",
		"speculatable",
		"ssp",
		"sspreq",
		"sspstrong",
		"strictfp",
		"uwtable",
		"writeonly",
		"inreg",
		"noalias",
		"nonnull",
		"signext",
		"zeroext",
		"byval",
		"nest",
		"nocapture",
		"returned",
		"sret",
		"swiftself",
		"addrspace",
		"allocsize",
		"dereferenceable",
		"dereferenceable_or_null",
		"dllexport",
		"dllimport",
		"appending",
		"available_externally",
		"common",
		"internal",
		"linkonce",
		"linkonce_odr",
		"private",
		"weak_odr",
		"extern_weak",
		"external",
		"dso_local",
		"dso_preemptable",
		"section",
		"local_unnamed_addr",
		"unnamed_addr",
		"default",
		"hidden",
		"protected",
		"eq",
		"ne",
		"sge",
		"sgt",
		"sle",
		"slt",
		"uge",
		"ugt",
		"ule",
		"ult",
		"oeq",
		"oge",
		"ogt",
		"ole",
		"olt",
		"one",
		"ord",
		"ueq",
		"une",
		"uno",
		"nsw",
		"nuw",
		"exact",
		"afn",
		"arcp",
		"contract",
		"fast",
		"ninf",
		"nnan",
		"nsz",
		"reassoc",
		"inbounds",
		"volatile",
		"syncscope",
		"acq_rel",
		"acquire",
		"monotonic",
		"release",
		"seq_cst",
		"unordered",
		"...",
	},

	idMap: map[string]Type{
		"INVALID":                     0,
		"$":                           1,
		"empty":                       2,
		"source_filename":             3,
		"=":                           4,
		"target":                      5,
		"triple":                      6,
		"datalayout":                  7,
		"module":                      8,
		"asm":                         9,
		"type":                        10,
		"comdat":                      11,
		"any":                         12,
		"exactmatch":                  13,
		"largest":                     14,
		"noduplicates":                15,
		"samesize":                    16,
		"thread_local":                17,
		"(":                           18,
		")":                           19,
		"initialexec":                 20,
		"localdynamic":                21,
		"localexec":                   22,
		"externally_initialized":      23,
		"constant":                    24,
		"global":                      25,
		",":                           26,
		"alias":                       27,
		"ifunc":                       28,
		"declare":                     29,
		"define":                      30,
		"amdgpu_cs":                   31,
		"amdgpu_es":                   32,
		"amdgpu_gs":                   33,
		"amdgpu_hs":                   34,
		"amdgpu_kernel":               35,
		"amdgpu_ls":                   36,
		"amdgpu_ps":                   37,
		"amdgpu_vs":                   38,
		"anyregcc":                    39,
		"arm_aapcs_vfpcc":             40,
		"arm_aapcscc":                 41,
		"arm_apcscc":                  42,
		"avr_intrcc":                  43,
		"avr_signalcc":                44,
		"ccc":                         45,
		"coldcc":                      46,
		"cxx_fast_tlscc":              47,
		"fastcc":                      48,
		"ghccc":                       49,
		"hhvm_ccc":                    50,
		"hhvmcc":                      51,
		"intel_ocl_bicc":              52,
		"msp430_intrcc":               53,
		"preserve_allcc":              54,
		"preserve_mostcc":             55,
		"ptx_device":                  56,
		"ptx_kernel":                  57,
		"spir_func":                   58,
		"spir_kernel":                 59,
		"swiftcc":                     60,
		"webkit_jscc":                 61,
		"win64cc":                     62,
		"x86_64_sysvcc":               63,
		"x86_fastcallcc":              64,
		"x86_intrcc":                  65,
		"x86_regcallcc":               66,
		"x86_stdcallcc":               67,
		"x86_thiscallcc":              68,
		"x86_vectorcallcc":            69,
		"cc":                          70,
		"int_lit":                     71,
		"gc":                          72,
		"prefix":                      73,
		"prologue":                    74,
		"personality":                 75,
		"{":                           76,
		"}":                           77,
		"attributes":                  78,
		"!":                           79,
		"distinct":                    80,
		"uselistorder":                81,
		"uselistorder_bb":             82,
		"global_ident":                83,
		"local_ident":                 84,
		"label_ident":                 85,
		"attr_group_id":               86,
		"comdat_name":                 87,
		"metadata_name":               88,
		"metadata_id":                 89,
		"void":                        90,
		"int_type":                    91,
		"half":                        92,
		"float":                       93,
		"double":                      94,
		"x86_fp80":                    95,
		"fp128":                       96,
		"ppc_fp128":                   97,
		"x86_mmx":                     98,
		"*":                           99,
		"<":                           100,
		"x":                           101,
		">":                           102,
		"label":                       103,
		"token":                       104,
		"metadata":                    105,
		"[":                           106,
		"]":                           107,
		"opaque":                      108,
		"sideeffect":                  109,
		"alignstack":                  110,
		"inteldialect":                111,
		"true":                        112,
		"false":                       113,
		"float_lit":                   114,
		"null":                        115,
		"none":                        116,
		"c":                           117,
		"string_lit":                  118,
		"zeroinitializer":             119,
		"undef":                       120,
		"blockaddress":                121,
		"add":                         122,
		"fadd":                        123,
		"sub":                         124,
		"fsub":                        125,
		"mul":                         126,
		"fmul":                        127,
		"udiv":                        128,
		"sdiv":                        129,
		"fdiv":                        130,
		"urem":                        131,
		"srem":                        132,
		"frem":                        133,
		"shl":                         134,
		"lshr":                        135,
		"ashr":                        136,
		"and":                         137,
		"or":                          138,
		"xor":                         139,
		"extractelement":              140,
		"insertelement":               141,
		"shufflevector":               142,
		"extractvalue":                143,
		"insertvalue":                 144,
		"getelementptr":               145,
		"inrange":                     146,
		"trunc":                       147,
		"to":                          148,
		"zext":                        149,
		"sext":                        150,
		"fptrunc":                     151,
		"fpext":                       152,
		"fptoui":                      153,
		"fptosi":                      154,
		"uitofp":                      155,
		"sitofp":                      156,
		"ptrtoint":                    157,
		"inttoptr":                    158,
		"bitcast":                     159,
		"addrspacecast":               160,
		"icmp":                        161,
		"fcmp":                        162,
		"select":                      163,
		"alloca":                      164,
		"inalloca":                    165,
		"swifterror":                  166,
		"load":                        167,
		"atomic":                      168,
		"store":                       169,
		"fence":                       170,
		"cmpxchg":                     171,
		"weak":                        172,
		"atomicrmw":                   173,
		"max":                         174,
		"min":                         175,
		"nand":                        176,
		"umax":                        177,
		"umin":                        178,
		"xchg":                        179,
		"phi":                         180,
		"call":                        181,
		"musttail":                    182,
		"notail":                      183,
		"tail":                        184,
		"va_arg":                      185,
		"landingpad":                  186,
		"cleanup":                     187,
		"catch":                       188,
		"filter":                      189,
		"catchpad":                    190,
		"within":                      191,
		"cleanuppad":                  192,
		"ret":                         193,
		"br":                          194,
		"switch":                      195,
		"indirectbr":                  196,
		"invoke":                      197,
		"unwind":                      198,
		"resume":                      199,
		"catchswitch":                 200,
		"caller":                      201,
		"catchret":                    202,
		"from":                        203,
		"cleanupret":                  204,
		"unreachable":                 205,
		"!DICompileUnit":              206,
		"language:":                   207,
		"producer:":                   208,
		"flags:":                      209,
		"runtimeVersion:":             210,
		"splitDebugFilename:":         211,
		"emissionKind:":               212,
		"enums:":                      213,
		"retainedTypes:":              214,
		"globals:":                    215,
		"imports:":                    216,
		"macros:":                     217,
		"dwoId:":                      218,
		"splitDebugInlining:":         219,
		"debugInfoForProfiling:":      220,
		"gnuPubnames:":                221,
		"!DIFile":                     222,
		"filename:":                   223,
		"directory:":                  224,
		"checksumkind:":               225,
		"checksum:":                   226,
		"!DIBasicType":                227,
		"encoding:":                   228,
		"!DISubroutineType":           229,
		"cc:":                         230,
		"types:":                      231,
		"!DIDerivedType":              232,
		"extraData:":                  233,
		"dwarfAddressSpace:":          234,
		"!DICompositeType":            235,
		"elements:":                   236,
		"runtimeLang:":                237,
		"vtableHolder:":               238,
		"identifier:":                 239,
		"discriminator:":              240,
		"!DISubrange":                 241,
		"count:":                      242,
		"lowerBound:":                 243,
		"!DIEnumerator":               244,
		"value:":                      245,
		"isUnsigned:":                 246,
		"!DITemplateTypeParameter":    247,
		"!DITemplateValueParameter":   248,
		"!DIModule":                   249,
		"configMacros:":               250,
		"includePath:":                251,
		"isysroot:":                   252,
		"!DINamespace":                253,
		"exportSymbols:":              254,
		"!DIGlobalVariable":           255,
		"!DISubprogram":               256,
		"scopeLine:":                  257,
		"containingType:":             258,
		"virtuality:":                 259,
		"virtualIndex:":               260,
		"thisAdjustment:":             261,
		"unit:":                       262,
		"variables:":                  263,
		"thrownTypes:":                264,
		"!DILexicalBlock":             265,
		"!DILexicalBlockFile":         266,
		"!DILocation":                 267,
		"inlinedAt:":                  268,
		"!DILocalVariable":            269,
		"arg:":                        270,
		"!DIExpression":               271,
		"!DIGlobalVariableExpression": 272,
		"var:":                          273,
		"expr:":                         274,
		"!DIObjCProperty":               275,
		"setter:":                       276,
		"getter:":                       277,
		"attributes:":                   278,
		"!DIImportedEntity":             279,
		"entity:":                       280,
		"!DIMacro":                      281,
		"!DIMacroFile":                  282,
		"nodes:":                        283,
		"!GenericDINode":                284,
		"header:":                       285,
		"operands:":                     286,
		"dwarf_lang":                    287,
		"file:":                         288,
		"isOptimized:":                  289,
		"FullDebug":                     290,
		"LineTablesOnly":                291,
		"NoDebug":                       292,
		"checksum_kind":                 293,
		"tag:":                          294,
		"dwarf_tag":                     295,
		"name:":                         296,
		"size:":                         297,
		"align:":                        298,
		"dwarf_att_encoding":            299,
		"|":                             300,
		"di_flag":                       301,
		"dwarf_cc":                      302,
		"line:":                         303,
		"scope:":                        304,
		"baseType:":                     305,
		"offset:":                       306,
		"templateParams:":               307,
		"type:":                         308,
		"linkageName:":                  309,
		"isLocal:":                      310,
		"isDefinition:":                 311,
		"declaration:":                  312,
		"dwarf_virtuality":              313,
		"column:":                       314,
		"dwarf_op":                      315,
		"dwarf_macinfo":                 316,
		"align":                         317,
		"alwaysinline":                  318,
		"argmemonly":                    319,
		"builtin":                       320,
		"cold":                          321,
		"convergent":                    322,
		"inaccessiblemem_or_argmemonly": 323,
		"inaccessiblememonly":           324,
		"inlinehint":                    325,
		"jumptable":                     326,
		"minsize":                       327,
		"naked":                         328,
		"nobuiltin":                     329,
		"noduplicate":                   330,
		"noimplicitfloat":               331,
		"noinline":                      332,
		"nonlazybind":                   333,
		"norecurse":                     334,
		"noredzone":                     335,
		"noreturn":                      336,
		"nounwind":                      337,
		"optnone":                       338,
		"optsize":                       339,
		"readnone":                      340,
		"readonly":                      341,
		"returns_twice":                 342,
		"safestack":                     343,
		"sanitize_address":              344,
		"sanitize_hwaddress":            345,
		"sanitize_memory":               346,
		"sanitize_thread":               347,
		"speculatable":                  348,
		"ssp":                           349,
		"sspreq":                        350,
		"sspstrong":                     351,
		"strictfp":                      352,
		"uwtable":                       353,
		"writeonly":                     354,
		"inreg":                         355,
		"noalias":                       356,
		"nonnull":                       357,
		"signext":                       358,
		"zeroext":                       359,
		"byval":                         360,
		"nest":                          361,
		"nocapture":                     362,
		"returned":                      363,
		"sret":                          364,
		"swiftself":                     365,
		"addrspace":                     366,
		"allocsize":                     367,
		"dereferenceable":               368,
		"dereferenceable_or_null":       369,
		"dllexport":                     370,
		"dllimport":                     371,
		"appending":                     372,
		"available_externally":          373,
		"common":                        374,
		"internal":                      375,
		"linkonce":                      376,
		"linkonce_odr":                  377,
		"private":                       378,
		"weak_odr":                      379,
		"extern_weak":                   380,
		"external":                      381,
		"dso_local":                     382,
		"dso_preemptable":               383,
		"section":                       384,
		"local_unnamed_addr":            385,
		"unnamed_addr":                  386,
		"default":                       387,
		"hidden":                        388,
		"protected":                     389,
		"eq":                            390,
		"ne":                            391,
		"sge":                           392,
		"sgt":                           393,
		"sle":                           394,
		"slt":                           395,
		"uge":                           396,
		"ugt":                           397,
		"ule":                           398,
		"ult":                           399,
		"oeq":                           400,
		"oge":                           401,
		"ogt":                           402,
		"ole":                           403,
		"olt":                           404,
		"one":                           405,
		"ord":                           406,
		"ueq":                           407,
		"une":                           408,
		"uno":                           409,
		"nsw":                           410,
		"nuw":                           411,
		"exact":                         412,
		"afn":                           413,
		"arcp":                          414,
		"contract":                      415,
		"fast":                          416,
		"ninf":                          417,
		"nnan":                          418,
		"nsz":                           419,
		"reassoc":                       420,
		"inbounds":                      421,
		"volatile":                      422,
		"syncscope":                     423,
		"acq_rel":                       424,
		"acquire":                       425,
		"monotonic":                     426,
		"release":                       427,
		"seq_cst":                       428,
		"unordered":                     429,
		"...":                           430,
	},
}