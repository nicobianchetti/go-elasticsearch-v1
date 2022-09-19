package workerpool

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	jobsCount = 2
)

func Test_Worker_Pool_Result_Channel_Ok(t *testing.T) {
	//Setup
	wpool := New(getJobs(false)...)

	//Exercise
	wpool.Run()
	resultDataChannel := wpool.GetResultChannel()
	data, _ := getData(resultDataChannel, t)

	//Verify
	assert.Equal(t, 0, data[0])
	assert.Equal(t, 2, data[1])
}

func Test_Worker_Pool_Result_Ok(t *testing.T) {
	//Setup
	wpool := New(getJobs(false)...)

	//Exercise
	wpool.Run()
	resultData, err := wpool.GetResult()
	resultFirst := resultData[0].(int)
	resultSecond := resultData[1].(int)

	//Verify
	assert.NoError(t, err)
	assert.Equal(t, 0, resultFirst)
	assert.Equal(t, 2, resultSecond)
}

func Test_Worker_Pool_Result_Err(t *testing.T) {
	//Setup
	wpool := New(getJobs(true)...)

	//Exercise
	wpool.Run()
	resultData, err := wpool.GetResult()

	//Verify
	assert.Error(t, err)
	assert.Nil(t, resultData)
}

func getData(resultData <-chan Result, t *testing.T) ([]int, error) {
	result := []int{}

	for dataResult := range resultData {
		if dataResult.Err != nil {
			t.Error(dataResult.Err)
		}

		data := dataResult.Data.(int)
		result = append(result, data)
	}

	return result, nil
}

func getJobs(isErr bool) []*Job {
	jobs := []*Job{}

	for i := 0; i < jobsCount; i++ {
		jobs = append(jobs, &Job{
			Descriptor: JobDescriptor{ID: "Test Job", StopIfErr: isErr},
			ExecFn:     getFn(i, isErr),
		})
	}

	return jobs
}

func getFn(i int, isErr bool) func() (interface{}, error) {
	if !isErr {
		return getFnWithourErr(i)
	}
	return getFnWithErr(i)
}

func getFnWithourErr(i int) func() (interface{}, error) {
	return func() (interface{}, error) {
		return i * 2, nil
	}
}

func getFnWithErr(i int) func() (interface{}, error) {
	return func() (interface{}, error) {
		return nil, errors.New("mock error")
	}
}
