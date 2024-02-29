package src

const (
	StateTypeTask     = "Task"
	StateTypePass     = "Pass"
	StateTypeFail     = "Fail"
	StateTypeSucceed  = "Succeed"
	StateTypeChoice   = "Choice"
	StateTypeWait     = "Wait"
	StateTypeParallel = "Parallel"
	StateTypeMap      = "Map"
)

type Retry struct {
	ErrorEquals     []string `json:"ErrorEquals"`
	IntervalSeconds int      `json:"IntervalSeconds"`
	MaxAttempts     int      `json:"MaxAttempts"`
	BackoffRate     float64  `json:"BackoffRate"`
}

type Catch struct {
	ErrorEquals []string `json:"ErrorEquals"`
	ResultPath  string   `json:"ResultPath"`
	Next        string   `json:"Next"`
}

// State represents a state in the ASL
type State struct {
	Type           string            `json:"Type"`
	Comment        string            `json:"Comment"`
	Resource       string            `json:"Resource"`
	InputPath      string            `json:"InputPath"`
	OutputPath     string            `json:"OutputPath"`
	ResultPath     string            `json:"ResultPath"`
	Next           string            `json:"Next"`
	End            bool              `json:"End"`
	ResultSelector map[string]string `json:"ResultSelector"`
	Parameters     map[string]string `json:"Parameters"`
	Retry          []Retry           `json:"Retry"`
	Catch          []Catch           `json:"Catch"`
}

type StateMachine struct {
	Comment string           `json:"Comment"`
	StartAt string           `json:"StartAt"`
	Version string           `json:"Version"`
	States  map[string]State `json:"States"`
}
