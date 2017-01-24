package topology

//RunStatus of the topology or node
type RunStatus int

const (
	//when new
	RunStatusNone RunStatus = iota

	//all elements ok
	RunStatusGreen

	//some ok,some failed
	RunStatusYellow

	//all elements failed
	RunStatusRed

	DownStatusGreen

	DownStatusYellow

	DownStatusRed
)

//Status of the topology or node
type Status struct {
	RunStatus RunStatus
	Error     string
}
