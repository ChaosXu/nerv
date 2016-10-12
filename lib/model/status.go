package model

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
)

//JobStatus of the operation executing on the topology or node
type JobStatus int

const (
	//nonr jon is running
	JobStatusNone JobStatus = iota

	//job is doing
	JobStatusDoing

	//job has been done.if error the RunStatus is RunStatusRed and information in the Error
	JobStatusDone
)

//Status of the topology or node
type Status struct {
	RunStatus RunStatus
	JobStatus JobStatus
	Error     string
}
