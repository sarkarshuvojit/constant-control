package domain

type CConfigRoot struct {
	SourceOfTruth SourceOfTruth `json:"sourceOfTruth" yaml:"sourceOfTruth"`
	Out           []Out         `json:"out" yaml:"out"`
}

type Out struct {
	Type string  `json:"type" yaml:"type"`
	Mode *string `json:"mode,omitempty" yaml:"mode,omitempty"`
	Opts Opts    `json:"opts" yaml:"opts"`
}

type Opts struct {
	Lang string   `json:"lang" yaml:"lang"`
	Src  []string `json:"src" yaml:"src"`
	Mode *string  `json:"mode,omitempty" yaml:"mode,omitempty"`
}

type SourceOfTruth struct {
	Type  string `json:"type" yaml:"type"`
	Conn  string `json:"conn" yaml:"conn"`
	Hooks Hooks  `json:"hooks" yaml:"hooks"`
}

type Hooks struct {
	Insert  string `json:"insert" yaml:"insert"`
	GetKeys string `json:"getKeys" yaml:"getKeys"`
	Get     string `json:"get" yaml:"get"`
	Delete  string `json:"delete" yaml:"delete"`
}
