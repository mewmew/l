// Code generated by "stringer -linecomment -type DwarfOp"; DO NOT EDIT.

package metadata

import "strconv"

const (
	_DwarfOp_name_0 = "DW_OP_addr"
	_DwarfOp_name_1 = "DW_OP_deref"
	_DwarfOp_name_2 = "DW_OP_const1uDW_OP_const1sDW_OP_const2uDW_OP_const2sDW_OP_const4uDW_OP_const4sDW_OP_const8uDW_OP_const8sDW_OP_constuDW_OP_constsDW_OP_dupDW_OP_dropDW_OP_overDW_OP_pickDW_OP_swapDW_OP_rotDW_OP_xderefDW_OP_absDW_OP_andDW_OP_divDW_OP_minusDW_OP_modDW_OP_mulDW_OP_negDW_OP_notDW_OP_orDW_OP_plusDW_OP_plus_uconstDW_OP_shlDW_OP_shrDW_OP_shraDW_OP_xorDW_OP_braDW_OP_eqDW_OP_geDW_OP_gtDW_OP_leDW_OP_ltDW_OP_neDW_OP_skipDW_OP_lit0DW_OP_lit1DW_OP_lit2DW_OP_lit3DW_OP_lit4DW_OP_lit5DW_OP_lit6DW_OP_lit7DW_OP_lit8DW_OP_lit9DW_OP_lit10DW_OP_lit11DW_OP_lit12DW_OP_lit13DW_OP_lit14DW_OP_lit15DW_OP_lit16DW_OP_lit17DW_OP_lit18DW_OP_lit19DW_OP_lit20DW_OP_lit21DW_OP_lit22DW_OP_lit23DW_OP_lit24DW_OP_lit25DW_OP_lit26DW_OP_lit27DW_OP_lit28DW_OP_lit29DW_OP_lit30DW_OP_lit31DW_OP_reg0DW_OP_reg1DW_OP_reg2DW_OP_reg3DW_OP_reg4DW_OP_reg5DW_OP_reg6DW_OP_reg7DW_OP_reg8DW_OP_reg9DW_OP_reg10DW_OP_reg11DW_OP_reg12DW_OP_reg13DW_OP_reg14DW_OP_reg15DW_OP_reg16DW_OP_reg17DW_OP_reg18DW_OP_reg19DW_OP_reg20DW_OP_reg21DW_OP_reg22DW_OP_reg23DW_OP_reg24DW_OP_reg25DW_OP_reg26DW_OP_reg27DW_OP_reg28DW_OP_reg29DW_OP_reg30DW_OP_reg31DW_OP_breg0DW_OP_breg1DW_OP_breg2DW_OP_breg3DW_OP_breg4DW_OP_breg5DW_OP_breg6DW_OP_breg7DW_OP_breg8DW_OP_breg9DW_OP_breg10DW_OP_breg11DW_OP_breg12DW_OP_breg13DW_OP_breg14DW_OP_breg15DW_OP_breg16DW_OP_breg17DW_OP_breg18DW_OP_breg19DW_OP_breg20DW_OP_breg21DW_OP_breg22DW_OP_breg23DW_OP_breg24DW_OP_breg25DW_OP_breg26DW_OP_breg27DW_OP_breg28DW_OP_breg29DW_OP_breg30DW_OP_breg31DW_OP_regxDW_OP_fbregDW_OP_bregxDW_OP_pieceDW_OP_deref_sizeDW_OP_xderef_sizeDW_OP_nopDW_OP_push_object_addressDW_OP_call2DW_OP_call4DW_OP_call_refDW_OP_form_tls_addressDW_OP_call_frame_cfaDW_OP_bit_pieceDW_OP_implicit_valueDW_OP_stack_valueDW_OP_implicit_pointerDW_OP_addrxDW_OP_constxDW_OP_entry_valueDW_OP_const_typeDW_OP_regval_typeDW_OP_deref_typeDW_OP_xderef_typeDW_OP_convertDW_OP_reinterpret"
	_DwarfOp_name_3 = "DW_OP_GNU_push_tls_address"
	_DwarfOp_name_4 = "DW_OP_GNU_addr_indexDW_OP_GNU_const_index"
	_DwarfOp_name_5 = "DW_OP_LLVM_fragment"
)

var (
	_DwarfOp_index_2 = [...]uint16{0, 13, 26, 39, 52, 65, 78, 91, 104, 116, 128, 137, 147, 157, 167, 177, 186, 198, 207, 216, 225, 236, 245, 254, 263, 272, 280, 290, 307, 316, 325, 335, 344, 353, 361, 369, 377, 385, 393, 401, 411, 421, 431, 441, 451, 461, 471, 481, 491, 501, 511, 522, 533, 544, 555, 566, 577, 588, 599, 610, 621, 632, 643, 654, 665, 676, 687, 698, 709, 720, 731, 742, 753, 763, 773, 783, 793, 803, 813, 823, 833, 843, 853, 864, 875, 886, 897, 908, 919, 930, 941, 952, 963, 974, 985, 996, 1007, 1018, 1029, 1040, 1051, 1062, 1073, 1084, 1095, 1106, 1117, 1128, 1139, 1150, 1161, 1172, 1183, 1194, 1205, 1217, 1229, 1241, 1253, 1265, 1277, 1289, 1301, 1313, 1325, 1337, 1349, 1361, 1373, 1385, 1397, 1409, 1421, 1433, 1445, 1457, 1469, 1479, 1490, 1501, 1512, 1528, 1545, 1554, 1579, 1590, 1601, 1615, 1637, 1657, 1672, 1692, 1709, 1731, 1742, 1754, 1771, 1787, 1804, 1820, 1837, 1850, 1867}
	_DwarfOp_index_4 = [...]uint8{0, 20, 41}
)

func (i DwarfOp) String() string {
	switch {
	case i == 3:
		return _DwarfOp_name_0
	case i == 6:
		return _DwarfOp_name_1
	case 8 <= i && i <= 169:
		i -= 8
		return _DwarfOp_name_2[_DwarfOp_index_2[i]:_DwarfOp_index_2[i+1]]
	case i == 224:
		return _DwarfOp_name_3
	case 251 <= i && i <= 252:
		i -= 251
		return _DwarfOp_name_4[_DwarfOp_index_4[i]:_DwarfOp_index_4[i+1]]
	case i == 4096:
		return _DwarfOp_name_5
	default:
		return "DwarfOp(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
