// Code generated by 'yaegi extract debug/dwarf'. DO NOT EDIT.

//go:build go1.22
// +build go1.22

package stdlib

import (
	"debug/dwarf"
	"reflect"
)

func init() {
	Symbols["debug/dwarf/dwarf"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AttrAbstractOrigin":        reflect.ValueOf(dwarf.AttrAbstractOrigin),
		"AttrAccessibility":         reflect.ValueOf(dwarf.AttrAccessibility),
		"AttrAddrBase":              reflect.ValueOf(dwarf.AttrAddrBase),
		"AttrAddrClass":             reflect.ValueOf(dwarf.AttrAddrClass),
		"AttrAlignment":             reflect.ValueOf(dwarf.AttrAlignment),
		"AttrAllocated":             reflect.ValueOf(dwarf.AttrAllocated),
		"AttrArtificial":            reflect.ValueOf(dwarf.AttrArtificial),
		"AttrAssociated":            reflect.ValueOf(dwarf.AttrAssociated),
		"AttrBaseTypes":             reflect.ValueOf(dwarf.AttrBaseTypes),
		"AttrBinaryScale":           reflect.ValueOf(dwarf.AttrBinaryScale),
		"AttrBitOffset":             reflect.ValueOf(dwarf.AttrBitOffset),
		"AttrBitSize":               reflect.ValueOf(dwarf.AttrBitSize),
		"AttrByteSize":              reflect.ValueOf(dwarf.AttrByteSize),
		"AttrCallAllCalls":          reflect.ValueOf(dwarf.AttrCallAllCalls),
		"AttrCallAllSourceCalls":    reflect.ValueOf(dwarf.AttrCallAllSourceCalls),
		"AttrCallAllTailCalls":      reflect.ValueOf(dwarf.AttrCallAllTailCalls),
		"AttrCallColumn":            reflect.ValueOf(dwarf.AttrCallColumn),
		"AttrCallDataLocation":      reflect.ValueOf(dwarf.AttrCallDataLocation),
		"AttrCallDataValue":         reflect.ValueOf(dwarf.AttrCallDataValue),
		"AttrCallFile":              reflect.ValueOf(dwarf.AttrCallFile),
		"AttrCallLine":              reflect.ValueOf(dwarf.AttrCallLine),
		"AttrCallOrigin":            reflect.ValueOf(dwarf.AttrCallOrigin),
		"AttrCallPC":                reflect.ValueOf(dwarf.AttrCallPC),
		"AttrCallParameter":         reflect.ValueOf(dwarf.AttrCallParameter),
		"AttrCallReturnPC":          reflect.ValueOf(dwarf.AttrCallReturnPC),
		"AttrCallTailCall":          reflect.ValueOf(dwarf.AttrCallTailCall),
		"AttrCallTarget":            reflect.ValueOf(dwarf.AttrCallTarget),
		"AttrCallTargetClobbered":   reflect.ValueOf(dwarf.AttrCallTargetClobbered),
		"AttrCallValue":             reflect.ValueOf(dwarf.AttrCallValue),
		"AttrCalling":               reflect.ValueOf(dwarf.AttrCalling),
		"AttrCommonRef":             reflect.ValueOf(dwarf.AttrCommonRef),
		"AttrCompDir":               reflect.ValueOf(dwarf.AttrCompDir),
		"AttrConstExpr":             reflect.ValueOf(dwarf.AttrConstExpr),
		"AttrConstValue":            reflect.ValueOf(dwarf.AttrConstValue),
		"AttrContainingType":        reflect.ValueOf(dwarf.AttrContainingType),
		"AttrCount":                 reflect.ValueOf(dwarf.AttrCount),
		"AttrDataBitOffset":         reflect.ValueOf(dwarf.AttrDataBitOffset),
		"AttrDataLocation":          reflect.ValueOf(dwarf.AttrDataLocation),
		"AttrDataMemberLoc":         reflect.ValueOf(dwarf.AttrDataMemberLoc),
		"AttrDecimalScale":          reflect.ValueOf(dwarf.AttrDecimalScale),
		"AttrDecimalSign":           reflect.ValueOf(dwarf.AttrDecimalSign),
		"AttrDeclColumn":            reflect.ValueOf(dwarf.AttrDeclColumn),
		"AttrDeclFile":              reflect.ValueOf(dwarf.AttrDeclFile),
		"AttrDeclLine":              reflect.ValueOf(dwarf.AttrDeclLine),
		"AttrDeclaration":           reflect.ValueOf(dwarf.AttrDeclaration),
		"AttrDefaultValue":          reflect.ValueOf(dwarf.AttrDefaultValue),
		"AttrDefaulted":             reflect.ValueOf(dwarf.AttrDefaulted),
		"AttrDeleted":               reflect.ValueOf(dwarf.AttrDeleted),
		"AttrDescription":           reflect.ValueOf(dwarf.AttrDescription),
		"AttrDigitCount":            reflect.ValueOf(dwarf.AttrDigitCount),
		"AttrDiscr":                 reflect.ValueOf(dwarf.AttrDiscr),
		"AttrDiscrList":             reflect.ValueOf(dwarf.AttrDiscrList),
		"AttrDiscrValue":            reflect.ValueOf(dwarf.AttrDiscrValue),
		"AttrDwoName":               reflect.ValueOf(dwarf.AttrDwoName),
		"AttrElemental":             reflect.ValueOf(dwarf.AttrElemental),
		"AttrEncoding":              reflect.ValueOf(dwarf.AttrEncoding),
		"AttrEndianity":             reflect.ValueOf(dwarf.AttrEndianity),
		"AttrEntrypc":               reflect.ValueOf(dwarf.AttrEntrypc),
		"AttrEnumClass":             reflect.ValueOf(dwarf.AttrEnumClass),
		"AttrExplicit":              reflect.ValueOf(dwarf.AttrExplicit),
		"AttrExportSymbols":         reflect.ValueOf(dwarf.AttrExportSymbols),
		"AttrExtension":             reflect.ValueOf(dwarf.AttrExtension),
		"AttrExternal":              reflect.ValueOf(dwarf.AttrExternal),
		"AttrFrameBase":             reflect.ValueOf(dwarf.AttrFrameBase),
		"AttrFriend":                reflect.ValueOf(dwarf.AttrFriend),
		"AttrHighpc":                reflect.ValueOf(dwarf.AttrHighpc),
		"AttrIdentifierCase":        reflect.ValueOf(dwarf.AttrIdentifierCase),
		"AttrImport":                reflect.ValueOf(dwarf.AttrImport),
		"AttrInline":                reflect.ValueOf(dwarf.AttrInline),
		"AttrIsOptional":            reflect.ValueOf(dwarf.AttrIsOptional),
		"AttrLanguage":              reflect.ValueOf(dwarf.AttrLanguage),
		"AttrLinkageName":           reflect.ValueOf(dwarf.AttrLinkageName),
		"AttrLocation":              reflect.ValueOf(dwarf.AttrLocation),
		"AttrLoclistsBase":          reflect.ValueOf(dwarf.AttrLoclistsBase),
		"AttrLowerBound":            reflect.ValueOf(dwarf.AttrLowerBound),
		"AttrLowpc":                 reflect.ValueOf(dwarf.AttrLowpc),
		"AttrMacroInfo":             reflect.ValueOf(dwarf.AttrMacroInfo),
		"AttrMacros":                reflect.ValueOf(dwarf.AttrMacros),
		"AttrMainSubprogram":        reflect.ValueOf(dwarf.AttrMainSubprogram),
		"AttrMutable":               reflect.ValueOf(dwarf.AttrMutable),
		"AttrName":                  reflect.ValueOf(dwarf.AttrName),
		"AttrNamelistItem":          reflect.ValueOf(dwarf.AttrNamelistItem),
		"AttrNoreturn":              reflect.ValueOf(dwarf.AttrNoreturn),
		"AttrObjectPointer":         reflect.ValueOf(dwarf.AttrObjectPointer),
		"AttrOrdering":              reflect.ValueOf(dwarf.AttrOrdering),
		"AttrPictureString":         reflect.ValueOf(dwarf.AttrPictureString),
		"AttrPriority":              reflect.ValueOf(dwarf.AttrPriority),
		"AttrProducer":              reflect.ValueOf(dwarf.AttrProducer),
		"AttrPrototyped":            reflect.ValueOf(dwarf.AttrPrototyped),
		"AttrPure":                  reflect.ValueOf(dwarf.AttrPure),
		"AttrRanges":                reflect.ValueOf(dwarf.AttrRanges),
		"AttrRank":                  reflect.ValueOf(dwarf.AttrRank),
		"AttrRecursive":             reflect.ValueOf(dwarf.AttrRecursive),
		"AttrReference":             reflect.ValueOf(dwarf.AttrReference),
		"AttrReturnAddr":            reflect.ValueOf(dwarf.AttrReturnAddr),
		"AttrRnglistsBase":          reflect.ValueOf(dwarf.AttrRnglistsBase),
		"AttrRvalueReference":       reflect.ValueOf(dwarf.AttrRvalueReference),
		"AttrSegment":               reflect.ValueOf(dwarf.AttrSegment),
		"AttrSibling":               reflect.ValueOf(dwarf.AttrSibling),
		"AttrSignature":             reflect.ValueOf(dwarf.AttrSignature),
		"AttrSmall":                 reflect.ValueOf(dwarf.AttrSmall),
		"AttrSpecification":         reflect.ValueOf(dwarf.AttrSpecification),
		"AttrStartScope":            reflect.ValueOf(dwarf.AttrStartScope),
		"AttrStaticLink":            reflect.ValueOf(dwarf.AttrStaticLink),
		"AttrStmtList":              reflect.ValueOf(dwarf.AttrStmtList),
		"AttrStrOffsetsBase":        reflect.ValueOf(dwarf.AttrStrOffsetsBase),
		"AttrStride":                reflect.ValueOf(dwarf.AttrStride),
		"AttrStrideSize":            reflect.ValueOf(dwarf.AttrStrideSize),
		"AttrStringLength":          reflect.ValueOf(dwarf.AttrStringLength),
		"AttrStringLengthBitSize":   reflect.ValueOf(dwarf.AttrStringLengthBitSize),
		"AttrStringLengthByteSize":  reflect.ValueOf(dwarf.AttrStringLengthByteSize),
		"AttrThreadsScaled":         reflect.ValueOf(dwarf.AttrThreadsScaled),
		"AttrTrampoline":            reflect.ValueOf(dwarf.AttrTrampoline),
		"AttrType":                  reflect.ValueOf(dwarf.AttrType),
		"AttrUpperBound":            reflect.ValueOf(dwarf.AttrUpperBound),
		"AttrUseLocation":           reflect.ValueOf(dwarf.AttrUseLocation),
		"AttrUseUTF8":               reflect.ValueOf(dwarf.AttrUseUTF8),
		"AttrVarParam":              reflect.ValueOf(dwarf.AttrVarParam),
		"AttrVirtuality":            reflect.ValueOf(dwarf.AttrVirtuality),
		"AttrVisibility":            reflect.ValueOf(dwarf.AttrVisibility),
		"AttrVtableElemLoc":         reflect.ValueOf(dwarf.AttrVtableElemLoc),
		"ClassAddrPtr":              reflect.ValueOf(dwarf.ClassAddrPtr),
		"ClassAddress":              reflect.ValueOf(dwarf.ClassAddress),
		"ClassBlock":                reflect.ValueOf(dwarf.ClassBlock),
		"ClassConstant":             reflect.ValueOf(dwarf.ClassConstant),
		"ClassExprLoc":              reflect.ValueOf(dwarf.ClassExprLoc),
		"ClassFlag":                 reflect.ValueOf(dwarf.ClassFlag),
		"ClassLinePtr":              reflect.ValueOf(dwarf.ClassLinePtr),
		"ClassLocList":              reflect.ValueOf(dwarf.ClassLocList),
		"ClassLocListPtr":           reflect.ValueOf(dwarf.ClassLocListPtr),
		"ClassMacPtr":               reflect.ValueOf(dwarf.ClassMacPtr),
		"ClassRangeListPtr":         reflect.ValueOf(dwarf.ClassRangeListPtr),
		"ClassReference":            reflect.ValueOf(dwarf.ClassReference),
		"ClassReferenceAlt":         reflect.ValueOf(dwarf.ClassReferenceAlt),
		"ClassReferenceSig":         reflect.ValueOf(dwarf.ClassReferenceSig),
		"ClassRngList":              reflect.ValueOf(dwarf.ClassRngList),
		"ClassRngListsPtr":          reflect.ValueOf(dwarf.ClassRngListsPtr),
		"ClassStrOffsetsPtr":        reflect.ValueOf(dwarf.ClassStrOffsetsPtr),
		"ClassString":               reflect.ValueOf(dwarf.ClassString),
		"ClassStringAlt":            reflect.ValueOf(dwarf.ClassStringAlt),
		"ClassUnknown":              reflect.ValueOf(dwarf.ClassUnknown),
		"ErrUnknownPC":              reflect.ValueOf(&dwarf.ErrUnknownPC).Elem(),
		"New":                       reflect.ValueOf(dwarf.New),
		"TagAccessDeclaration":      reflect.ValueOf(dwarf.TagAccessDeclaration),
		"TagArrayType":              reflect.ValueOf(dwarf.TagArrayType),
		"TagAtomicType":             reflect.ValueOf(dwarf.TagAtomicType),
		"TagBaseType":               reflect.ValueOf(dwarf.TagBaseType),
		"TagCallSite":               reflect.ValueOf(dwarf.TagCallSite),
		"TagCallSiteParameter":      reflect.ValueOf(dwarf.TagCallSiteParameter),
		"TagCatchDwarfBlock":        reflect.ValueOf(dwarf.TagCatchDwarfBlock),
		"TagClassType":              reflect.ValueOf(dwarf.TagClassType),
		"TagCoarrayType":            reflect.ValueOf(dwarf.TagCoarrayType),
		"TagCommonDwarfBlock":       reflect.ValueOf(dwarf.TagCommonDwarfBlock),
		"TagCommonInclusion":        reflect.ValueOf(dwarf.TagCommonInclusion),
		"TagCompileUnit":            reflect.ValueOf(dwarf.TagCompileUnit),
		"TagCondition":              reflect.ValueOf(dwarf.TagCondition),
		"TagConstType":              reflect.ValueOf(dwarf.TagConstType),
		"TagConstant":               reflect.ValueOf(dwarf.TagConstant),
		"TagDwarfProcedure":         reflect.ValueOf(dwarf.TagDwarfProcedure),
		"TagDynamicType":            reflect.ValueOf(dwarf.TagDynamicType),
		"TagEntryPoint":             reflect.ValueOf(dwarf.TagEntryPoint),
		"TagEnumerationType":        reflect.ValueOf(dwarf.TagEnumerationType),
		"TagEnumerator":             reflect.ValueOf(dwarf.TagEnumerator),
		"TagFileType":               reflect.ValueOf(dwarf.TagFileType),
		"TagFormalParameter":        reflect.ValueOf(dwarf.TagFormalParameter),
		"TagFriend":                 reflect.ValueOf(dwarf.TagFriend),
		"TagGenericSubrange":        reflect.ValueOf(dwarf.TagGenericSubrange),
		"TagImmutableType":          reflect.ValueOf(dwarf.TagImmutableType),
		"TagImportedDeclaration":    reflect.ValueOf(dwarf.TagImportedDeclaration),
		"TagImportedModule":         reflect.ValueOf(dwarf.TagImportedModule),
		"TagImportedUnit":           reflect.ValueOf(dwarf.TagImportedUnit),
		"TagInheritance":            reflect.ValueOf(dwarf.TagInheritance),
		"TagInlinedSubroutine":      reflect.ValueOf(dwarf.TagInlinedSubroutine),
		"TagInterfaceType":          reflect.ValueOf(dwarf.TagInterfaceType),
		"TagLabel":                  reflect.ValueOf(dwarf.TagLabel),
		"TagLexDwarfBlock":          reflect.ValueOf(dwarf.TagLexDwarfBlock),
		"TagMember":                 reflect.ValueOf(dwarf.TagMember),
		"TagModule":                 reflect.ValueOf(dwarf.TagModule),
		"TagMutableType":            reflect.ValueOf(dwarf.TagMutableType),
		"TagNamelist":               reflect.ValueOf(dwarf.TagNamelist),
		"TagNamelistItem":           reflect.ValueOf(dwarf.TagNamelistItem),
		"TagNamespace":              reflect.ValueOf(dwarf.TagNamespace),
		"TagPackedType":             reflect.ValueOf(dwarf.TagPackedType),
		"TagPartialUnit":            reflect.ValueOf(dwarf.TagPartialUnit),
		"TagPointerType":            reflect.ValueOf(dwarf.TagPointerType),
		"TagPtrToMemberType":        reflect.ValueOf(dwarf.TagPtrToMemberType),
		"TagReferenceType":          reflect.ValueOf(dwarf.TagReferenceType),
		"TagRestrictType":           reflect.ValueOf(dwarf.TagRestrictType),
		"TagRvalueReferenceType":    reflect.ValueOf(dwarf.TagRvalueReferenceType),
		"TagSetType":                reflect.ValueOf(dwarf.TagSetType),
		"TagSharedType":             reflect.ValueOf(dwarf.TagSharedType),
		"TagSkeletonUnit":           reflect.ValueOf(dwarf.TagSkeletonUnit),
		"TagStringType":             reflect.ValueOf(dwarf.TagStringType),
		"TagStructType":             reflect.ValueOf(dwarf.TagStructType),
		"TagSubprogram":             reflect.ValueOf(dwarf.TagSubprogram),
		"TagSubrangeType":           reflect.ValueOf(dwarf.TagSubrangeType),
		"TagSubroutineType":         reflect.ValueOf(dwarf.TagSubroutineType),
		"TagTemplateAlias":          reflect.ValueOf(dwarf.TagTemplateAlias),
		"TagTemplateTypeParameter":  reflect.ValueOf(dwarf.TagTemplateTypeParameter),
		"TagTemplateValueParameter": reflect.ValueOf(dwarf.TagTemplateValueParameter),
		"TagThrownType":             reflect.ValueOf(dwarf.TagThrownType),
		"TagTryDwarfBlock":          reflect.ValueOf(dwarf.TagTryDwarfBlock),
		"TagTypeUnit":               reflect.ValueOf(dwarf.TagTypeUnit),
		"TagTypedef":                reflect.ValueOf(dwarf.TagTypedef),
		"TagUnionType":              reflect.ValueOf(dwarf.TagUnionType),
		"TagUnspecifiedParameters":  reflect.ValueOf(dwarf.TagUnspecifiedParameters),
		"TagUnspecifiedType":        reflect.ValueOf(dwarf.TagUnspecifiedType),
		"TagVariable":               reflect.ValueOf(dwarf.TagVariable),
		"TagVariant":                reflect.ValueOf(dwarf.TagVariant),
		"TagVariantPart":            reflect.ValueOf(dwarf.TagVariantPart),
		"TagVolatileType":           reflect.ValueOf(dwarf.TagVolatileType),
		"TagWithStmt":               reflect.ValueOf(dwarf.TagWithStmt),

		// type definitions
		"AddrType":        reflect.ValueOf((*dwarf.AddrType)(nil)),
		"ArrayType":       reflect.ValueOf((*dwarf.ArrayType)(nil)),
		"Attr":            reflect.ValueOf((*dwarf.Attr)(nil)),
		"BasicType":       reflect.ValueOf((*dwarf.BasicType)(nil)),
		"BoolType":        reflect.ValueOf((*dwarf.BoolType)(nil)),
		"CharType":        reflect.ValueOf((*dwarf.CharType)(nil)),
		"Class":           reflect.ValueOf((*dwarf.Class)(nil)),
		"CommonType":      reflect.ValueOf((*dwarf.CommonType)(nil)),
		"ComplexType":     reflect.ValueOf((*dwarf.ComplexType)(nil)),
		"Data":            reflect.ValueOf((*dwarf.Data)(nil)),
		"DecodeError":     reflect.ValueOf((*dwarf.DecodeError)(nil)),
		"DotDotDotType":   reflect.ValueOf((*dwarf.DotDotDotType)(nil)),
		"Entry":           reflect.ValueOf((*dwarf.Entry)(nil)),
		"EnumType":        reflect.ValueOf((*dwarf.EnumType)(nil)),
		"EnumValue":       reflect.ValueOf((*dwarf.EnumValue)(nil)),
		"Field":           reflect.ValueOf((*dwarf.Field)(nil)),
		"FloatType":       reflect.ValueOf((*dwarf.FloatType)(nil)),
		"FuncType":        reflect.ValueOf((*dwarf.FuncType)(nil)),
		"IntType":         reflect.ValueOf((*dwarf.IntType)(nil)),
		"LineEntry":       reflect.ValueOf((*dwarf.LineEntry)(nil)),
		"LineFile":        reflect.ValueOf((*dwarf.LineFile)(nil)),
		"LineReader":      reflect.ValueOf((*dwarf.LineReader)(nil)),
		"LineReaderPos":   reflect.ValueOf((*dwarf.LineReaderPos)(nil)),
		"Offset":          reflect.ValueOf((*dwarf.Offset)(nil)),
		"PtrType":         reflect.ValueOf((*dwarf.PtrType)(nil)),
		"QualType":        reflect.ValueOf((*dwarf.QualType)(nil)),
		"Reader":          reflect.ValueOf((*dwarf.Reader)(nil)),
		"StructField":     reflect.ValueOf((*dwarf.StructField)(nil)),
		"StructType":      reflect.ValueOf((*dwarf.StructType)(nil)),
		"Tag":             reflect.ValueOf((*dwarf.Tag)(nil)),
		"Type":            reflect.ValueOf((*dwarf.Type)(nil)),
		"TypedefType":     reflect.ValueOf((*dwarf.TypedefType)(nil)),
		"UcharType":       reflect.ValueOf((*dwarf.UcharType)(nil)),
		"UintType":        reflect.ValueOf((*dwarf.UintType)(nil)),
		"UnspecifiedType": reflect.ValueOf((*dwarf.UnspecifiedType)(nil)),
		"UnsupportedType": reflect.ValueOf((*dwarf.UnsupportedType)(nil)),
		"VoidType":        reflect.ValueOf((*dwarf.VoidType)(nil)),

		// interface wrapper definitions
		"_Type": reflect.ValueOf((*_debug_dwarf_Type)(nil)),
	}
}

// _debug_dwarf_Type is an interface wrapper for Type type
type _debug_dwarf_Type struct {
	IValue  interface{}
	WCommon func() *dwarf.CommonType
	WSize   func() int64
	WString func() string
}

func (W _debug_dwarf_Type) Common() *dwarf.CommonType { return W.WCommon() }
func (W _debug_dwarf_Type) Size() int64               { return W.WSize() }
func (W _debug_dwarf_Type) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
