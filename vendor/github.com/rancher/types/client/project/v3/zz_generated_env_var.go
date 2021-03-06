package client

const (
	EnvVarType           = "envVar"
	EnvVarFieldName      = "name"
	EnvVarFieldValue     = "value"
	EnvVarFieldValueFrom = "valueFrom"
)

type EnvVar struct {
	Name      string        `json:"name,omitempty"`
	Value     string        `json:"value,omitempty"`
	ValueFrom *EnvVarSource `json:"valueFrom,omitempty"`
}
