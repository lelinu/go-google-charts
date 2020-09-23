package models

type Row struct {
	Value       *ValueWithStyle `json:"value,required"`
	ParentValue string          `json:"parent_value"`
	RowType     string          `json:"row_type,required"`
}

func NewRow(value *ValueWithStyle, parentValue string, rowType string) *Row {
	return &Row{
		Value:       value,
		ParentValue: parentValue,
		RowType:     rowType,
	}
}

type ValueWithStyle struct {
	Value     string `json:"v,required"`
	HtmlValue string `json:"f"`
}

func NewValueWithStyle(value string, htmlValue string) *ValueWithStyle {
	return &ValueWithStyle{
		Value:     value,
		HtmlValue: htmlValue,
	}
}
