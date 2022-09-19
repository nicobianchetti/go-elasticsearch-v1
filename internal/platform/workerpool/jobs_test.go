package workerpool

import (
	"errors"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Job_Ok(t *testing.T) {
	//Setup
	job := Job{
		Descriptor: JobDescriptor{
			ID:        "Job-Test",
			StopIfErr: false,
		},
		ExecFn: func() (interface{}, error) {
			return 1, nil
		},
	}

	//Exercise
	result := job.execute()
	data := result.Data.(int)

	//Verify
	assert.NoError(t, result.Err)
	assert.Equal(t, 1, data)
	assert.Equal(t, result.Descriptor, job.Descriptor)
}

func Test_Job_Err(t *testing.T) {
	//Setup
	job := Job{
		Descriptor: JobDescriptor{
			ID:        "Job-Test",
			StopIfErr: false,
		},
		ExecFn: func() (interface{}, error) {
			return nil, errors.New("mock error")
		},
	}

	//Exercise
	result := job.execute()

	//Verify
	assert.Error(t, result.Err)
	assert.Nil(t, result.Data)
	assert.Equal(t, result.Descriptor, job.Descriptor)
}

func Test_CreateJob(t *testing.T) {
	//Setup
	name := "User"
	fieldMapper := "Field"
	stopErr := true
	execFn := func() (interface{}, error) { return nil, nil }

	//Exercise
	result := CreateJob(name, fieldMapper, stopErr, execFn)

	//Verify
	assert.Equal(t, JobID(name), result.Descriptor.ID)
	assert.Equal(t, stopErr, result.Descriptor.StopIfErr)
	assert.Equal(t, fieldMapper, result.FieldMapper)
	funcExpected := runtime.FuncForPC(reflect.ValueOf(execFn).Pointer()).Name()
	funcResult := runtime.FuncForPC(reflect.ValueOf(result.ExecFn).Pointer()).Name()
	assert.Equal(t, funcExpected, funcResult)
}
