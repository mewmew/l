// Code generated by "stringer -linecomment -type DwarfTag"; DO NOT EDIT.

package metadata

import "strconv"

const _DwarfTag_name = "DW_TAG_nullDW_TAG_array_typeDW_TAG_class_typeDW_TAG_entry_pointDW_TAG_enumeration_typeDW_TAG_formal_parameterDW_TAG_imported_declarationDW_TAG_labelDW_TAG_lexical_blockDW_TAG_memberDW_TAG_pointer_typeDW_TAG_reference_typeDW_TAG_compile_unitDW_TAG_string_typeDW_TAG_structure_typeDW_TAG_subroutine_typeDW_TAG_typedefDW_TAG_union_typeDW_TAG_unspecified_parametersDW_TAG_variantDW_TAG_common_blockDW_TAG_common_inclusionDW_TAG_inheritanceDW_TAG_inlined_subroutineDW_TAG_moduleDW_TAG_ptr_to_member_typeDW_TAG_set_typeDW_TAG_subrange_typeDW_TAG_with_stmtDW_TAG_access_declarationDW_TAG_base_typeDW_TAG_catch_blockDW_TAG_const_typeDW_TAG_constantDW_TAG_enumeratorDW_TAG_file_typeDW_TAG_friendDW_TAG_namelistDW_TAG_namelist_itemDW_TAG_packed_typeDW_TAG_subprogramDW_TAG_template_type_parameterDW_TAG_template_value_parameterDW_TAG_thrown_typeDW_TAG_try_blockDW_TAG_variant_partDW_TAG_variableDW_TAG_volatile_typeDW_TAG_dwarf_procedureDW_TAG_restrict_typeDW_TAG_interface_typeDW_TAG_namespaceDW_TAG_imported_moduleDW_TAG_unspecified_typeDW_TAG_partial_unitDW_TAG_imported_unitDW_TAG_conditionDW_TAG_shared_typeDW_TAG_type_unitDW_TAG_rvalue_reference_typeDW_TAG_template_aliasDW_TAG_coarray_typeDW_TAG_generic_subrangeDW_TAG_dynamic_typeDW_TAG_atomic_typeDW_TAG_call_siteDW_TAG_call_site_parameterDW_TAG_skeleton_unitDW_TAG_immutable_typeDW_TAG_MIPS_loopDW_TAG_format_labelDW_TAG_function_templateDW_TAG_class_templateDW_TAG_GNU_template_template_paramDW_TAG_GNU_template_parameter_packDW_TAG_GNU_formal_parameter_packDW_TAG_GNU_call_siteDW_TAG_GNU_call_site_parameterDW_TAG_APPLE_propertyDW_TAG_BORLAND_propertyDW_TAG_BORLAND_Delphi_stringDW_TAG_BORLAND_Delphi_dynamic_arrayDW_TAG_BORLAND_Delphi_setDW_TAG_BORLAND_Delphi_variant"

var _DwarfTag_map = map[DwarfTag]string{
	0:     _DwarfTag_name[0:11],
	1:     _DwarfTag_name[11:28],
	2:     _DwarfTag_name[28:45],
	3:     _DwarfTag_name[45:63],
	4:     _DwarfTag_name[63:86],
	5:     _DwarfTag_name[86:109],
	8:     _DwarfTag_name[109:136],
	10:    _DwarfTag_name[136:148],
	11:    _DwarfTag_name[148:168],
	13:    _DwarfTag_name[168:181],
	15:    _DwarfTag_name[181:200],
	16:    _DwarfTag_name[200:221],
	17:    _DwarfTag_name[221:240],
	18:    _DwarfTag_name[240:258],
	19:    _DwarfTag_name[258:279],
	21:    _DwarfTag_name[279:301],
	22:    _DwarfTag_name[301:315],
	23:    _DwarfTag_name[315:332],
	24:    _DwarfTag_name[332:361],
	25:    _DwarfTag_name[361:375],
	26:    _DwarfTag_name[375:394],
	27:    _DwarfTag_name[394:417],
	28:    _DwarfTag_name[417:435],
	29:    _DwarfTag_name[435:460],
	30:    _DwarfTag_name[460:473],
	31:    _DwarfTag_name[473:498],
	32:    _DwarfTag_name[498:513],
	33:    _DwarfTag_name[513:533],
	34:    _DwarfTag_name[533:549],
	35:    _DwarfTag_name[549:574],
	36:    _DwarfTag_name[574:590],
	37:    _DwarfTag_name[590:608],
	38:    _DwarfTag_name[608:625],
	39:    _DwarfTag_name[625:640],
	40:    _DwarfTag_name[640:657],
	41:    _DwarfTag_name[657:673],
	42:    _DwarfTag_name[673:686],
	43:    _DwarfTag_name[686:701],
	44:    _DwarfTag_name[701:721],
	45:    _DwarfTag_name[721:739],
	46:    _DwarfTag_name[739:756],
	47:    _DwarfTag_name[756:786],
	48:    _DwarfTag_name[786:817],
	49:    _DwarfTag_name[817:835],
	50:    _DwarfTag_name[835:851],
	51:    _DwarfTag_name[851:870],
	52:    _DwarfTag_name[870:885],
	53:    _DwarfTag_name[885:905],
	54:    _DwarfTag_name[905:927],
	55:    _DwarfTag_name[927:947],
	56:    _DwarfTag_name[947:968],
	57:    _DwarfTag_name[968:984],
	58:    _DwarfTag_name[984:1006],
	59:    _DwarfTag_name[1006:1029],
	60:    _DwarfTag_name[1029:1048],
	61:    _DwarfTag_name[1048:1068],
	63:    _DwarfTag_name[1068:1084],
	64:    _DwarfTag_name[1084:1102],
	65:    _DwarfTag_name[1102:1118],
	66:    _DwarfTag_name[1118:1146],
	67:    _DwarfTag_name[1146:1167],
	68:    _DwarfTag_name[1167:1186],
	69:    _DwarfTag_name[1186:1209],
	70:    _DwarfTag_name[1209:1228],
	71:    _DwarfTag_name[1228:1246],
	72:    _DwarfTag_name[1246:1262],
	73:    _DwarfTag_name[1262:1288],
	74:    _DwarfTag_name[1288:1308],
	75:    _DwarfTag_name[1308:1329],
	16513: _DwarfTag_name[1329:1345],
	16641: _DwarfTag_name[1345:1364],
	16642: _DwarfTag_name[1364:1388],
	16643: _DwarfTag_name[1388:1409],
	16646: _DwarfTag_name[1409:1443],
	16647: _DwarfTag_name[1443:1477],
	16648: _DwarfTag_name[1477:1509],
	16649: _DwarfTag_name[1509:1529],
	16650: _DwarfTag_name[1529:1559],
	16896: _DwarfTag_name[1559:1580],
	45056: _DwarfTag_name[1580:1603],
	45057: _DwarfTag_name[1603:1631],
	45058: _DwarfTag_name[1631:1666],
	45059: _DwarfTag_name[1666:1691],
	45060: _DwarfTag_name[1691:1720],
}

func (i DwarfTag) String() string {
	if str, ok := _DwarfTag_map[i]; ok {
		return str
	}
	return "DwarfTag(" + strconv.FormatInt(int64(i), 10) + ")"
}
