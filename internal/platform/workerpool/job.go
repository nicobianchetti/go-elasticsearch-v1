package workerpool

type JobID string
type ExecutionFn func() (interface{}, error)

type JobDescriptor struct {
	ID        JobID
	StopIfErr bool
}

type Result struct {
	Data        interface{}
	Err         error
	FieldMapper string
	Descriptor  JobDescriptor
}

type Job struct {
	Descriptor  JobDescriptor
	ExecFn      ExecutionFn
	FieldMapper string
}

func CreateJob(name string, fieldMapper string, stopErr bool, execFn func() (interface{}, error)) *Job {
	return &Job{
		Descriptor: JobDescriptor{
			ID:        JobID(name),
			StopIfErr: stopErr,
		},
		ExecFn:      execFn,
		FieldMapper: fieldMapper,
	}
}

func (j *Job) execute() Result {
	data, err := j.ExecFn()

	return Result{
		Data:        data,
		Descriptor:  j.Descriptor,
		FieldMapper: j.FieldMapper,
		Err:         err,
	}
}

type Jobs []*Job

func (js Jobs) AddIfNotNull(j *Job) Jobs {
	if j != nil {
		return js.Add(*j)
	}
	return js
}

func (js Jobs) Add(j Job) Jobs {
	return append(js, &j)
}

func NewJobs(j ...*Job) Jobs {
	jobs := Jobs{}
	jobs = append(jobs, j...)
	return jobs
}
