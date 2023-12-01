package constant

/* MajorDimension ====================================================================
 */
type MajorDimension int

const (
	MajorDimension_ROWS MajorDimension = iota
	MajorDimension_COLUMNS
)

func (m MajorDimension) String() string {
	switch m {
	case MajorDimension_ROWS:
		return "ROWS"
	case MajorDimension_COLUMNS:
		return "COLUMNS"
	}
	return "unknown"
}

/* ValueInputOption ====================================================================
 */
type ValueInputOption int

const (
	ValueInputOption_RAW ValueInputOption = iota
	ValueInputOption_USER_ENTERED
)

func (v ValueInputOption) String() string {
	switch v {
	case ValueInputOption_RAW:
		return "RAW"
	case ValueInputOption_USER_ENTERED:
		return "USER_ENTERED"
	}
	return "unknown"
}

/* LocationType ====================================================================
 */
type LocationType int

const (
	LocationType_ROW LocationType = iota
	LocationType_COLUMN
)

func (l LocationType) String() string {
	switch l {
	case LocationType_ROW:
		return "ROW"
	case LocationType_COLUMN:
		return "COLUMN"
	}
	return "unknown"
}

/* LocationType ====================================================================
 */
type InsertDataOption int

const (
	InsertDataOption_OVERWRITE InsertDataOption = iota
	InsertDataOption_INSERT_ROWS
)

func (i InsertDataOption) String() string {
	switch i {
	case InsertDataOption_OVERWRITE:
		return "OVERWRITE"
	case InsertDataOption_INSERT_ROWS:
		return "INSERT_ROWS"
	}
	return "unknown"
}

/* ValueRenderOption ====================================================================
 */
type ValueRenderOption int

const (
	ValueRenderOption_FORMATTED_VALUE ValueRenderOption = iota
	ValueRenderOption_UNFORMATTED_VALUE
	ValueRenderOption_FORMULA
)

func (v ValueRenderOption) String() string {
	switch v {
	case ValueRenderOption_FORMATTED_VALUE:
		return "FORMATTED_VALUE"
	case ValueRenderOption_UNFORMATTED_VALUE:
		return "UNFORMATTED_VALUE"
	case ValueRenderOption_FORMULA:
		return "FORMULA"
	}
	return "unknown"
}
