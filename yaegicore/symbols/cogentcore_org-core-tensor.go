// Code generated by 'yaegi extract cogentcore.org/core/tensor'. DO NOT EDIT.

package symbols

import (
	"cogentcore.org/core/base/metadata"
	"cogentcore.org/core/tensor"
	"reflect"
)

func init() {
	Symbols["cogentcore.org/core/tensor/tensor"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AddFunc":                   reflect.ValueOf(tensor.AddFunc),
		"AddShapes":                 reflect.ValueOf(tensor.AddShapes),
		"AlignShapes":               reflect.ValueOf(tensor.AlignShapes),
		"AnyFirstArg":               reflect.ValueOf(tensor.AnyFirstArg),
		"As1D":                      reflect.ValueOf(tensor.As1D),
		"AsFloat32Tensor":           reflect.ValueOf(tensor.AsFloat32Tensor),
		"AsFloat64Scalar":           reflect.ValueOf(tensor.AsFloat64Scalar),
		"AsFloat64Slice":            reflect.ValueOf(tensor.AsFloat64Slice),
		"AsFloat64Tensor":           reflect.ValueOf(tensor.AsFloat64Tensor),
		"AsIndexed":                 reflect.ValueOf(tensor.AsIndexed),
		"AsIntScalar":               reflect.ValueOf(tensor.AsIntScalar),
		"AsIntSlice":                reflect.ValueOf(tensor.AsIntSlice),
		"AsIntTensor":               reflect.ValueOf(tensor.AsIntTensor),
		"AsMasked":                  reflect.ValueOf(tensor.AsMasked),
		"AsReshaped":                reflect.ValueOf(tensor.AsReshaped),
		"AsRows":                    reflect.ValueOf(tensor.AsRows),
		"AsSliced":                  reflect.ValueOf(tensor.AsSliced),
		"AsStringScalar":            reflect.ValueOf(tensor.AsStringScalar),
		"AsStringSlice":             reflect.ValueOf(tensor.AsStringSlice),
		"AsStringTensor":            reflect.ValueOf(tensor.AsStringTensor),
		"Ascending":                 reflect.ValueOf(tensor.Ascending),
		"BoolToFloat64":             reflect.ValueOf(tensor.BoolToFloat64),
		"BoolToInt":                 reflect.ValueOf(tensor.BoolToInt),
		"Calc":                      reflect.ValueOf(tensor.Calc),
		"Call":                      reflect.ValueOf(tensor.Call),
		"CallAny":                   reflect.ValueOf(tensor.CallAny),
		"CallOut":                   reflect.ValueOf(tensor.CallOut),
		"CallOutMulti":              reflect.ValueOf(tensor.CallOutMulti),
		"Cells1D":                   reflect.ValueOf(tensor.Cells1D),
		"CellsSize":                 reflect.ValueOf(tensor.CellsSize),
		"Clone":                     reflect.ValueOf(tensor.Clone),
		"ColumnMajorStrides":        reflect.ValueOf(tensor.ColumnMajorStrides),
		"Comma":                     reflect.ValueOf(tensor.Comma),
		"DefaultNumThreads":         reflect.ValueOf(tensor.DefaultNumThreads),
		"DelimsN":                   reflect.ValueOf(tensor.DelimsN),
		"DelimsValues":              reflect.ValueOf(tensor.DelimsValues),
		"Descending":                reflect.ValueOf(tensor.Descending),
		"Detect":                    reflect.ValueOf(tensor.Detect),
		"Ellipsis":                  reflect.ValueOf(tensor.Ellipsis),
		"Float64ToBool":             reflect.ValueOf(tensor.Float64ToBool),
		"Float64ToString":           reflect.ValueOf(tensor.Float64ToString),
		"FullAxis":                  reflect.ValueOf(tensor.FullAxis),
		"FuncByName":                reflect.ValueOf(tensor.FuncByName),
		"Funcs":                     reflect.ValueOf(&tensor.Funcs).Elem(),
		"IntToBool":                 reflect.ValueOf(tensor.IntToBool),
		"MaxSprintLength":           reflect.ValueOf(&tensor.MaxSprintLength).Elem(),
		"MustBeSameShape":           reflect.ValueOf(tensor.MustBeSameShape),
		"MustBeValues":              reflect.ValueOf(tensor.MustBeValues),
		"NFirstLen":                 reflect.ValueOf(tensor.NFirstLen),
		"NFirstRows":                reflect.ValueOf(tensor.NFirstRows),
		"NMinLen":                   reflect.ValueOf(tensor.NMinLen),
		"NewAxis":                   reflect.ValueOf(tensor.NewAxis),
		"NewBool":                   reflect.ValueOf(tensor.NewBool),
		"NewBoolShape":              reflect.ValueOf(tensor.NewBoolShape),
		"NewByte":                   reflect.ValueOf(tensor.NewByte),
		"NewFloat32":                reflect.ValueOf(tensor.NewFloat32),
		"NewFloat64":                reflect.ValueOf(tensor.NewFloat64),
		"NewFloat64FromValues":      reflect.ValueOf(tensor.NewFloat64FromValues),
		"NewFloat64Scalar":          reflect.ValueOf(tensor.NewFloat64Scalar),
		"NewFunc":                   reflect.ValueOf(tensor.NewFunc),
		"NewIndexed":                reflect.ValueOf(tensor.NewIndexed),
		"NewInt":                    reflect.ValueOf(tensor.NewInt),
		"NewInt32":                  reflect.ValueOf(tensor.NewInt32),
		"NewIntFromValues":          reflect.ValueOf(tensor.NewIntFromValues),
		"NewIntScalar":              reflect.ValueOf(tensor.NewIntScalar),
		"NewMasked":                 reflect.ValueOf(tensor.NewMasked),
		"NewOfType":                 reflect.ValueOf(tensor.NewOfType),
		"NewReshaped":               reflect.ValueOf(tensor.NewReshaped),
		"NewRowCellsView":           reflect.ValueOf(tensor.NewRowCellsView),
		"NewRows":                   reflect.ValueOf(tensor.NewRows),
		"NewShape":                  reflect.ValueOf(tensor.NewShape),
		"NewSlice":                  reflect.ValueOf(tensor.NewSlice),
		"NewSliceInts":              reflect.ValueOf(tensor.NewSliceInts),
		"NewSliced":                 reflect.ValueOf(tensor.NewSliced),
		"NewString":                 reflect.ValueOf(tensor.NewString),
		"NewStringFromValues":       reflect.ValueOf(tensor.NewStringFromValues),
		"NewStringScalar":           reflect.ValueOf(tensor.NewStringScalar),
		"NewStringShape":            reflect.ValueOf(tensor.NewStringShape),
		"NumThreads":                reflect.ValueOf(&tensor.NumThreads).Elem(),
		"OddColumn":                 reflect.ValueOf(tensor.OddColumn),
		"OddRow":                    reflect.ValueOf(tensor.OddRow),
		"OpenCSV":                   reflect.ValueOf(tensor.OpenCSV),
		"Precision":                 reflect.ValueOf(tensor.Precision),
		"Projection2DCoords":        reflect.ValueOf(tensor.Projection2DCoords),
		"Projection2DIndex":         reflect.ValueOf(tensor.Projection2DIndex),
		"Projection2DSet":           reflect.ValueOf(tensor.Projection2DSet),
		"Projection2DSetString":     reflect.ValueOf(tensor.Projection2DSetString),
		"Projection2DShape":         reflect.ValueOf(tensor.Projection2DShape),
		"Projection2DString":        reflect.ValueOf(tensor.Projection2DString),
		"Projection2DValue":         reflect.ValueOf(tensor.Projection2DValue),
		"Range":                     reflect.ValueOf(tensor.Range),
		"ReadCSV":                   reflect.ValueOf(tensor.ReadCSV),
		"Reshape":                   reflect.ValueOf(tensor.Reshape),
		"Reslice":                   reflect.ValueOf(tensor.Reslice),
		"RowMajorStrides":           reflect.ValueOf(tensor.RowMajorStrides),
		"SaveCSV":                   reflect.ValueOf(tensor.SaveCSV),
		"SetCalcFunc":               reflect.ValueOf(tensor.SetCalcFunc),
		"SetPrecision":              reflect.ValueOf(tensor.SetPrecision),
		"SetShape":                  reflect.ValueOf(tensor.SetShape),
		"SetShapeFrom":              reflect.ValueOf(tensor.SetShapeFrom),
		"SetShapeFromMustBeValues":  reflect.ValueOf(tensor.SetShapeFromMustBeValues),
		"SetShapeNames":             reflect.ValueOf(tensor.SetShapeNames),
		"SetShapeSizesFromTensor":   reflect.ValueOf(tensor.SetShapeSizesFromTensor),
		"SetShapeSizesMustBeValues": reflect.ValueOf(tensor.SetShapeSizesMustBeValues),
		"ShapeNames":                reflect.ValueOf(tensor.ShapeNames),
		"SlicesMagicN":              reflect.ValueOf(tensor.SlicesMagicN),
		"SlicesMagicValues":         reflect.ValueOf(tensor.SlicesMagicValues),
		"Space":                     reflect.ValueOf(tensor.Space),
		"Sprint":                    reflect.ValueOf(tensor.Sprint),
		"StableSort":                reflect.ValueOf(tensor.StableSort),
		"StringToFloat64":           reflect.ValueOf(tensor.StringToFloat64),
		"Tab":                       reflect.ValueOf(tensor.Tab),
		"ThreadingThreshold":        reflect.ValueOf(&tensor.ThreadingThreshold).Elem(),
		"UnstableSort":              reflect.ValueOf(tensor.UnstableSort),
		"Vectorize":                 reflect.ValueOf(tensor.Vectorize),
		"VectorizeOnThreads":        reflect.ValueOf(tensor.VectorizeOnThreads),
		"VectorizeThreaded":         reflect.ValueOf(tensor.VectorizeThreaded),
		"WrapIndex1D":               reflect.ValueOf(tensor.WrapIndex1D),
		"WriteCSV":                  reflect.ValueOf(tensor.WriteCSV),

		// type definitions
		"Bool":          reflect.ValueOf((*tensor.Bool)(nil)),
		"Delims":        reflect.ValueOf((*tensor.Delims)(nil)),
		"FilterFunc":    reflect.ValueOf((*tensor.FilterFunc)(nil)),
		"FilterOptions": reflect.ValueOf((*tensor.FilterOptions)(nil)),
		"Func":          reflect.ValueOf((*tensor.Func)(nil)),
		"Indexed":       reflect.ValueOf((*tensor.Indexed)(nil)),
		"Masked":        reflect.ValueOf((*tensor.Masked)(nil)),
		"Reshaped":      reflect.ValueOf((*tensor.Reshaped)(nil)),
		"RowMajor":      reflect.ValueOf((*tensor.RowMajor)(nil)),
		"Rows":          reflect.ValueOf((*tensor.Rows)(nil)),
		"Shape":         reflect.ValueOf((*tensor.Shape)(nil)),
		"Slice":         reflect.ValueOf((*tensor.Slice)(nil)),
		"Sliced":        reflect.ValueOf((*tensor.Sliced)(nil)),
		"SlicesMagic":   reflect.ValueOf((*tensor.SlicesMagic)(nil)),
		"String":        reflect.ValueOf((*tensor.String)(nil)),
		"Tensor":        reflect.ValueOf((*tensor.Tensor)(nil)),
		"Values":        reflect.ValueOf((*tensor.Values)(nil)),

		// interface wrapper definitions
		"_RowMajor": reflect.ValueOf((*_cogentcore_org_core_tensor_RowMajor)(nil)),
		"_Tensor":   reflect.ValueOf((*_cogentcore_org_core_tensor_Tensor)(nil)),
		"_Values":   reflect.ValueOf((*_cogentcore_org_core_tensor_Values)(nil)),
	}
}

// _cogentcore_org_core_tensor_RowMajor is an interface wrapper for RowMajor type
type _cogentcore_org_core_tensor_RowMajor struct {
	IValue            interface{}
	WAsValues         func() tensor.Values
	WDataType         func() reflect.Kind
	WDimSize          func(dim int) int
	WFloat            func(i ...int) float64
	WFloat1D          func(i int) float64
	WFloatRow         func(row int) float64
	WFloatRowCell     func(row int, cell int) float64
	WInt              func(i ...int) int
	WInt1D            func(i int) int
	WIntRow           func(row int) int
	WIntRowCell       func(row int, cell int) int
	WIsString         func() bool
	WLabel            func() string
	WLen              func() int
	WMetadata         func() *metadata.Data
	WNumDims          func() int
	WRowTensor        func(row int) tensor.Values
	WSetFloat         func(val float64, i ...int)
	WSetFloat1D       func(val float64, i int)
	WSetFloatRow      func(val float64, row int)
	WSetFloatRowCell  func(val float64, row int, cell int)
	WSetInt           func(val int, i ...int)
	WSetInt1D         func(val int, i int)
	WSetIntRow        func(val int, row int)
	WSetIntRowCell    func(val int, row int, cell int)
	WSetRowTensor     func(val tensor.Values, row int)
	WSetString        func(val string, i ...int)
	WSetString1D      func(val string, i int)
	WSetStringRow     func(val string, row int)
	WSetStringRowCell func(val string, row int, cell int)
	WShape            func() *tensor.Shape
	WShapeSizes       func() []int
	WString           func() string
	WString1D         func(i int) string
	WStringRow        func(row int) string
	WStringRowCell    func(row int, cell int) string
	WStringValue      func(i ...int) string
	WSubSpace         func(offs ...int) tensor.Values
}

func (W _cogentcore_org_core_tensor_RowMajor) AsValues() tensor.Values  { return W.WAsValues() }
func (W _cogentcore_org_core_tensor_RowMajor) DataType() reflect.Kind   { return W.WDataType() }
func (W _cogentcore_org_core_tensor_RowMajor) DimSize(dim int) int      { return W.WDimSize(dim) }
func (W _cogentcore_org_core_tensor_RowMajor) Float(i ...int) float64   { return W.WFloat(i...) }
func (W _cogentcore_org_core_tensor_RowMajor) Float1D(i int) float64    { return W.WFloat1D(i) }
func (W _cogentcore_org_core_tensor_RowMajor) FloatRow(row int) float64 { return W.WFloatRow(row) }
func (W _cogentcore_org_core_tensor_RowMajor) FloatRowCell(row int, cell int) float64 {
	return W.WFloatRowCell(row, cell)
}
func (W _cogentcore_org_core_tensor_RowMajor) Int(i ...int) int   { return W.WInt(i...) }
func (W _cogentcore_org_core_tensor_RowMajor) Int1D(i int) int    { return W.WInt1D(i) }
func (W _cogentcore_org_core_tensor_RowMajor) IntRow(row int) int { return W.WIntRow(row) }
func (W _cogentcore_org_core_tensor_RowMajor) IntRowCell(row int, cell int) int {
	return W.WIntRowCell(row, cell)
}
func (W _cogentcore_org_core_tensor_RowMajor) IsString() bool           { return W.WIsString() }
func (W _cogentcore_org_core_tensor_RowMajor) Label() string            { return W.WLabel() }
func (W _cogentcore_org_core_tensor_RowMajor) Len() int                 { return W.WLen() }
func (W _cogentcore_org_core_tensor_RowMajor) Metadata() *metadata.Data { return W.WMetadata() }
func (W _cogentcore_org_core_tensor_RowMajor) NumDims() int             { return W.WNumDims() }
func (W _cogentcore_org_core_tensor_RowMajor) RowTensor(row int) tensor.Values {
	return W.WRowTensor(row)
}
func (W _cogentcore_org_core_tensor_RowMajor) SetFloat(val float64, i ...int) { W.WSetFloat(val, i...) }
func (W _cogentcore_org_core_tensor_RowMajor) SetFloat1D(val float64, i int)  { W.WSetFloat1D(val, i) }
func (W _cogentcore_org_core_tensor_RowMajor) SetFloatRow(val float64, row int) {
	W.WSetFloatRow(val, row)
}
func (W _cogentcore_org_core_tensor_RowMajor) SetFloatRowCell(val float64, row int, cell int) {
	W.WSetFloatRowCell(val, row, cell)
}
func (W _cogentcore_org_core_tensor_RowMajor) SetInt(val int, i ...int)   { W.WSetInt(val, i...) }
func (W _cogentcore_org_core_tensor_RowMajor) SetInt1D(val int, i int)    { W.WSetInt1D(val, i) }
func (W _cogentcore_org_core_tensor_RowMajor) SetIntRow(val int, row int) { W.WSetIntRow(val, row) }
func (W _cogentcore_org_core_tensor_RowMajor) SetIntRowCell(val int, row int, cell int) {
	W.WSetIntRowCell(val, row, cell)
}
func (W _cogentcore_org_core_tensor_RowMajor) SetRowTensor(val tensor.Values, row int) {
	W.WSetRowTensor(val, row)
}
func (W _cogentcore_org_core_tensor_RowMajor) SetString(val string, i ...int) {
	W.WSetString(val, i...)
}
func (W _cogentcore_org_core_tensor_RowMajor) SetString1D(val string, i int) { W.WSetString1D(val, i) }
func (W _cogentcore_org_core_tensor_RowMajor) SetStringRow(val string, row int) {
	W.WSetStringRow(val, row)
}
func (W _cogentcore_org_core_tensor_RowMajor) SetStringRowCell(val string, row int, cell int) {
	W.WSetStringRowCell(val, row, cell)
}
func (W _cogentcore_org_core_tensor_RowMajor) Shape() *tensor.Shape { return W.WShape() }
func (W _cogentcore_org_core_tensor_RowMajor) ShapeSizes() []int    { return W.WShapeSizes() }
func (W _cogentcore_org_core_tensor_RowMajor) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _cogentcore_org_core_tensor_RowMajor) String1D(i int) string    { return W.WString1D(i) }
func (W _cogentcore_org_core_tensor_RowMajor) StringRow(row int) string { return W.WStringRow(row) }
func (W _cogentcore_org_core_tensor_RowMajor) StringRowCell(row int, cell int) string {
	return W.WStringRowCell(row, cell)
}
func (W _cogentcore_org_core_tensor_RowMajor) StringValue(i ...int) string {
	return W.WStringValue(i...)
}
func (W _cogentcore_org_core_tensor_RowMajor) SubSpace(offs ...int) tensor.Values {
	return W.WSubSpace(offs...)
}

// _cogentcore_org_core_tensor_Tensor is an interface wrapper for Tensor type
type _cogentcore_org_core_tensor_Tensor struct {
	IValue       interface{}
	WAsValues    func() tensor.Values
	WDataType    func() reflect.Kind
	WDimSize     func(dim int) int
	WFloat       func(i ...int) float64
	WFloat1D     func(i int) float64
	WInt         func(i ...int) int
	WInt1D       func(i int) int
	WIsString    func() bool
	WLabel       func() string
	WLen         func() int
	WMetadata    func() *metadata.Data
	WNumDims     func() int
	WSetFloat    func(val float64, i ...int)
	WSetFloat1D  func(val float64, i int)
	WSetInt      func(val int, i ...int)
	WSetInt1D    func(val int, i int)
	WSetString   func(val string, i ...int)
	WSetString1D func(val string, i int)
	WShape       func() *tensor.Shape
	WShapeSizes  func() []int
	WString      func() string
	WString1D    func(i int) string
	WStringValue func(i ...int) string
}

func (W _cogentcore_org_core_tensor_Tensor) AsValues() tensor.Values        { return W.WAsValues() }
func (W _cogentcore_org_core_tensor_Tensor) DataType() reflect.Kind         { return W.WDataType() }
func (W _cogentcore_org_core_tensor_Tensor) DimSize(dim int) int            { return W.WDimSize(dim) }
func (W _cogentcore_org_core_tensor_Tensor) Float(i ...int) float64         { return W.WFloat(i...) }
func (W _cogentcore_org_core_tensor_Tensor) Float1D(i int) float64          { return W.WFloat1D(i) }
func (W _cogentcore_org_core_tensor_Tensor) Int(i ...int) int               { return W.WInt(i...) }
func (W _cogentcore_org_core_tensor_Tensor) Int1D(i int) int                { return W.WInt1D(i) }
func (W _cogentcore_org_core_tensor_Tensor) IsString() bool                 { return W.WIsString() }
func (W _cogentcore_org_core_tensor_Tensor) Label() string                  { return W.WLabel() }
func (W _cogentcore_org_core_tensor_Tensor) Len() int                       { return W.WLen() }
func (W _cogentcore_org_core_tensor_Tensor) Metadata() *metadata.Data       { return W.WMetadata() }
func (W _cogentcore_org_core_tensor_Tensor) NumDims() int                   { return W.WNumDims() }
func (W _cogentcore_org_core_tensor_Tensor) SetFloat(val float64, i ...int) { W.WSetFloat(val, i...) }
func (W _cogentcore_org_core_tensor_Tensor) SetFloat1D(val float64, i int)  { W.WSetFloat1D(val, i) }
func (W _cogentcore_org_core_tensor_Tensor) SetInt(val int, i ...int)       { W.WSetInt(val, i...) }
func (W _cogentcore_org_core_tensor_Tensor) SetInt1D(val int, i int)        { W.WSetInt1D(val, i) }
func (W _cogentcore_org_core_tensor_Tensor) SetString(val string, i ...int) { W.WSetString(val, i...) }
func (W _cogentcore_org_core_tensor_Tensor) SetString1D(val string, i int)  { W.WSetString1D(val, i) }
func (W _cogentcore_org_core_tensor_Tensor) Shape() *tensor.Shape           { return W.WShape() }
func (W _cogentcore_org_core_tensor_Tensor) ShapeSizes() []int              { return W.WShapeSizes() }
func (W _cogentcore_org_core_tensor_Tensor) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _cogentcore_org_core_tensor_Tensor) String1D(i int) string       { return W.WString1D(i) }
func (W _cogentcore_org_core_tensor_Tensor) StringValue(i ...int) string { return W.WStringValue(i...) }

// _cogentcore_org_core_tensor_Values is an interface wrapper for Values type
type _cogentcore_org_core_tensor_Values struct {
	IValue            interface{}
	WAppendFrom       func(from tensor.Values) error
	WAsValues         func() tensor.Values
	WBytes            func() []byte
	WClone            func() tensor.Values
	WCopyCellsFrom    func(from tensor.Values, to int, start int, n int)
	WCopyFrom         func(from tensor.Values)
	WDataType         func() reflect.Kind
	WDimSize          func(dim int) int
	WFloat            func(i ...int) float64
	WFloat1D          func(i int) float64
	WFloatRow         func(row int) float64
	WFloatRowCell     func(row int, cell int) float64
	WInt              func(i ...int) int
	WInt1D            func(i int) int
	WIntRow           func(row int) int
	WIntRowCell       func(row int, cell int) int
	WIsString         func() bool
	WLabel            func() string
	WLen              func() int
	WMetadata         func() *metadata.Data
	WNumDims          func() int
	WRowTensor        func(row int) tensor.Values
	WSetFloat         func(val float64, i ...int)
	WSetFloat1D       func(val float64, i int)
	WSetFloatRow      func(val float64, row int)
	WSetFloatRowCell  func(val float64, row int, cell int)
	WSetInt           func(val int, i ...int)
	WSetInt1D         func(val int, i int)
	WSetIntRow        func(val int, row int)
	WSetIntRowCell    func(val int, row int, cell int)
	WSetNumRows       func(rows int)
	WSetRowTensor     func(val tensor.Values, row int)
	WSetShapeSizes    func(sizes ...int)
	WSetString        func(val string, i ...int)
	WSetString1D      func(val string, i int)
	WSetStringRow     func(val string, row int)
	WSetStringRowCell func(val string, row int, cell int)
	WSetZeros         func()
	WShape            func() *tensor.Shape
	WShapeSizes       func() []int
	WSizeof           func() int64
	WString           func() string
	WString1D         func(i int) string
	WStringRow        func(row int) string
	WStringRowCell    func(row int, cell int) string
	WStringValue      func(i ...int) string
	WSubSpace         func(offs ...int) tensor.Values
}

func (W _cogentcore_org_core_tensor_Values) AppendFrom(from tensor.Values) error {
	return W.WAppendFrom(from)
}
func (W _cogentcore_org_core_tensor_Values) AsValues() tensor.Values { return W.WAsValues() }
func (W _cogentcore_org_core_tensor_Values) Bytes() []byte           { return W.WBytes() }
func (W _cogentcore_org_core_tensor_Values) Clone() tensor.Values    { return W.WClone() }
func (W _cogentcore_org_core_tensor_Values) CopyCellsFrom(from tensor.Values, to int, start int, n int) {
	W.WCopyCellsFrom(from, to, start, n)
}
func (W _cogentcore_org_core_tensor_Values) CopyFrom(from tensor.Values) { W.WCopyFrom(from) }
func (W _cogentcore_org_core_tensor_Values) DataType() reflect.Kind      { return W.WDataType() }
func (W _cogentcore_org_core_tensor_Values) DimSize(dim int) int         { return W.WDimSize(dim) }
func (W _cogentcore_org_core_tensor_Values) Float(i ...int) float64      { return W.WFloat(i...) }
func (W _cogentcore_org_core_tensor_Values) Float1D(i int) float64       { return W.WFloat1D(i) }
func (W _cogentcore_org_core_tensor_Values) FloatRow(row int) float64    { return W.WFloatRow(row) }
func (W _cogentcore_org_core_tensor_Values) FloatRowCell(row int, cell int) float64 {
	return W.WFloatRowCell(row, cell)
}
func (W _cogentcore_org_core_tensor_Values) Int(i ...int) int   { return W.WInt(i...) }
func (W _cogentcore_org_core_tensor_Values) Int1D(i int) int    { return W.WInt1D(i) }
func (W _cogentcore_org_core_tensor_Values) IntRow(row int) int { return W.WIntRow(row) }
func (W _cogentcore_org_core_tensor_Values) IntRowCell(row int, cell int) int {
	return W.WIntRowCell(row, cell)
}
func (W _cogentcore_org_core_tensor_Values) IsString() bool           { return W.WIsString() }
func (W _cogentcore_org_core_tensor_Values) Label() string            { return W.WLabel() }
func (W _cogentcore_org_core_tensor_Values) Len() int                 { return W.WLen() }
func (W _cogentcore_org_core_tensor_Values) Metadata() *metadata.Data { return W.WMetadata() }
func (W _cogentcore_org_core_tensor_Values) NumDims() int             { return W.WNumDims() }
func (W _cogentcore_org_core_tensor_Values) RowTensor(row int) tensor.Values {
	return W.WRowTensor(row)
}
func (W _cogentcore_org_core_tensor_Values) SetFloat(val float64, i ...int) { W.WSetFloat(val, i...) }
func (W _cogentcore_org_core_tensor_Values) SetFloat1D(val float64, i int)  { W.WSetFloat1D(val, i) }
func (W _cogentcore_org_core_tensor_Values) SetFloatRow(val float64, row int) {
	W.WSetFloatRow(val, row)
}
func (W _cogentcore_org_core_tensor_Values) SetFloatRowCell(val float64, row int, cell int) {
	W.WSetFloatRowCell(val, row, cell)
}
func (W _cogentcore_org_core_tensor_Values) SetInt(val int, i ...int)   { W.WSetInt(val, i...) }
func (W _cogentcore_org_core_tensor_Values) SetInt1D(val int, i int)    { W.WSetInt1D(val, i) }
func (W _cogentcore_org_core_tensor_Values) SetIntRow(val int, row int) { W.WSetIntRow(val, row) }
func (W _cogentcore_org_core_tensor_Values) SetIntRowCell(val int, row int, cell int) {
	W.WSetIntRowCell(val, row, cell)
}
func (W _cogentcore_org_core_tensor_Values) SetNumRows(rows int) { W.WSetNumRows(rows) }
func (W _cogentcore_org_core_tensor_Values) SetRowTensor(val tensor.Values, row int) {
	W.WSetRowTensor(val, row)
}
func (W _cogentcore_org_core_tensor_Values) SetShapeSizes(sizes ...int)     { W.WSetShapeSizes(sizes...) }
func (W _cogentcore_org_core_tensor_Values) SetString(val string, i ...int) { W.WSetString(val, i...) }
func (W _cogentcore_org_core_tensor_Values) SetString1D(val string, i int)  { W.WSetString1D(val, i) }
func (W _cogentcore_org_core_tensor_Values) SetStringRow(val string, row int) {
	W.WSetStringRow(val, row)
}
func (W _cogentcore_org_core_tensor_Values) SetStringRowCell(val string, row int, cell int) {
	W.WSetStringRowCell(val, row, cell)
}
func (W _cogentcore_org_core_tensor_Values) SetZeros()            { W.WSetZeros() }
func (W _cogentcore_org_core_tensor_Values) Shape() *tensor.Shape { return W.WShape() }
func (W _cogentcore_org_core_tensor_Values) ShapeSizes() []int    { return W.WShapeSizes() }
func (W _cogentcore_org_core_tensor_Values) Sizeof() int64        { return W.WSizeof() }
func (W _cogentcore_org_core_tensor_Values) String() string {
	if W.WString == nil {
		return ""
	}
	return W.WString()
}
func (W _cogentcore_org_core_tensor_Values) String1D(i int) string    { return W.WString1D(i) }
func (W _cogentcore_org_core_tensor_Values) StringRow(row int) string { return W.WStringRow(row) }
func (W _cogentcore_org_core_tensor_Values) StringRowCell(row int, cell int) string {
	return W.WStringRowCell(row, cell)
}
func (W _cogentcore_org_core_tensor_Values) StringValue(i ...int) string { return W.WStringValue(i...) }
func (W _cogentcore_org_core_tensor_Values) SubSpace(offs ...int) tensor.Values {
	return W.WSubSpace(offs...)
}
