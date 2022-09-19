package workerpool

import (
	"sync"

	"go-elasticsearch-v1/internal/platform/converter"
	"go-elasticsearch-v1/internal/platform/errors/usecase"
)

type WorkerPool struct {
	workersCount int
	jobs         chan Job
	results      chan Result
}

func New(jobs ...*Job) *WorkerPool {
	jobsCount := len(jobs)
	wpool := &WorkerPool{
		workersCount: jobsCount,
		jobs:         make(chan Job, jobsCount),
		results:      make(chan Result, jobsCount),
	}
	wpool.generateFrom(jobs)
	return wpool
}

func (wp *WorkerPool) Run() {
	var wg sync.WaitGroup
	wg.Add(wp.workersCount)

	for i := 0; i < wp.workersCount; i++ {
		go wp.worker(&wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.results)
}

func (wp *WorkerPool) GetResultChannel() <-chan Result {
	return wp.results
}

func (wp *WorkerPool) GetResult() ([]interface{}, error) {
	var results []interface{}
	for dataResult := range wp.results {
		if dataResult.Err != nil {
			if dataResult.Descriptor.StopIfErr {
				return nil, dataResult.Err
			}
		}
		results = append(results, dataResult.Data)
	}
	return results, nil
}

func (wp *WorkerPool) generateFrom(jobs []*Job) {
	for i := range jobs {
		wp.jobs <- *jobs[i]
	}
	close(wp.jobs)
}

func (wp *WorkerPool) worker(wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			//If channel is close, return
			if !ok {
				return
			}
			results <- job.execute()
		}
	}
}

func RunJobsAndMapResults(model interface{}, jobs ...*Job) error {
	resultData := RunJobs(jobs...)

	results := map[string]interface{}{}
	for dataResult := range resultData {
		if dataResult.Err != nil {
			if dataResult.Descriptor.StopIfErr {
				return dataResult.Err
			}
			continue
		}

		if dataResult.FieldMapper == "" {
			return usecase.Unknown("FieldMapper could not be empty")
		}

		results[dataResult.FieldMapper] = dataResult.Data
	}

	return converter.FillAllByFieldName(results, model)
}

func RunJobs(jobs ...*Job) <-chan Result {
	wp := New(jobs...)
	wp.Run()
	return wp.GetResultChannel()
}
