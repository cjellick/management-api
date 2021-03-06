package client

const (
	IDRangeType     = "idRange"
	IDRangeFieldMax = "max"
	IDRangeFieldMin = "min"
)

type IDRange struct {
	Max *int64 `json:"max,omitempty"`
	Min *int64 `json:"min,omitempty"`
}
