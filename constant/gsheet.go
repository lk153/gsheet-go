package constant

/* MajorDimension ====================================================================
 */
type MajorDimension int

const (
	Unknown = "unknown"

	MajorDimensionRows MajorDimension = iota
	MajorDimensionColumns
)

func (m MajorDimension) String() string {
	switch m {
	case MajorDimensionRows:
		return "ROWS"
	case MajorDimensionColumns:
		return "COLUMNS"
	}
	return Unknown
}

/* ValueInputOption ====================================================================
 */
type ValueInputOption int

const (
	ValueInputOptionRaw ValueInputOption = iota
	ValueInputOptionUserEntered
)

func (v ValueInputOption) String() string {
	switch v {
	case ValueInputOptionRaw:
		return "RAW"
	case ValueInputOptionUserEntered:
		return "USER_ENTERED"
	}
	return Unknown
}

/* LocationType ====================================================================
 */
type LocationType int

const (
	LocationTypeRow LocationType = iota
	LocationTypeColumn
)

func (l LocationType) String() string {
	switch l {
	case LocationTypeRow:
		return "ROW"
	case LocationTypeColumn:
		return "COLUMN"
	}
	return Unknown
}

/* LocationType ====================================================================
 */
type InsertDataOption int

const (
	InsertDataOptionOverwrite InsertDataOption = iota
	InsertDataOptionInsertRows
)

func (i InsertDataOption) String() string {
	switch i {
	case InsertDataOptionOverwrite:
		return "OVERWRITE"
	case InsertDataOptionInsertRows:
		return "INSERT_ROWS"
	}
	return Unknown
}

/* ValueRenderOption ====================================================================
 */
type ValueRenderOption int

const (
	ValueRenderOptionFormattedValue ValueRenderOption = iota
	ValueRenderOptionUnFormattedValue
	ValueRenderOptionFormula
)

func (v ValueRenderOption) String() string {
	switch v {
	case ValueRenderOptionFormattedValue:
		return "FORMATTED_VALUE"
	case ValueRenderOptionUnFormattedValue:
		return "UNFORMATTED_VALUE"
	case ValueRenderOptionFormula:
		return "FORMULA"
	}
	return Unknown
}
