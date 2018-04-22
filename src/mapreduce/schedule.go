package mapreduce

import "fmt"

//
// schedule() starts and waits for all tasks in the given phase (Map
// or Reduce). the mapFiles argument holds the names of the files that
// are the inputs to the map phase, one per map task. nReduce is the
// number of reduce tasks. the registerChan argument yields a stream
// of registered workers; each item is the worker's RPC address,
// suitable for passing to call(). registerChan will yield all
// existing registered workers (if any) and new ones as they register.
//
func schedule(jobName string, mapFiles []string, nReduce int, phase jobPhase, registerChan chan string) {
	var ntasks int
	var n_other int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mapFiles)
		n_other = nReduce
	case reducePhase:
		ntasks = nReduce
		n_other = len(mapFiles)
	}

	fmt.Printf("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, n_other)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//type DoTaskArgs struct {
	//	JobName    string
	//	File       string   // only for map, the input file
	//	Phase      jobPhase // are we in mapPhase or reducePhase?
	//	TaskNumber int      // this task's index in the current phase
	//
	//	// NumOtherPhase is the total number of tasks in other phase; mappers
	//	// need this to compute the number of output bins, and reducers needs
	//	// this to know how many input files to collect.
	//	NumOtherPhase int
	//}
	switch phase {
	case mapPhase:
		for i := 0; i < ntasks; i++ {
			work := <- registerChan
			args := DoTaskArgs{}
			args.JobName = jobName
			args.File = mapFiles[i]
			args.Phase = mapPhase
			args.TaskNumber = i
			args.NumOtherPhase = nReduce
			ok := call(work, "Worker.DoTask", args, new(struct{}))
			if ok == false {
				fmt.Printf("DoTask: RPC %s dotask error\n", work)
			}
			if ok {
				go func() {
					registerChan <- work
				}()
				} else {
				i = i-1// if work faield the work will be abandoned
				fmt.Printf("map Task %d field\n", i)
			}

		}
	case reducePhase:
		for i := 0; i < ntasks; i++ {
			work := <- registerChan
			args := DoTaskArgs{}
			args.JobName = jobName
			args.Phase = reducePhase
			args.TaskNumber = i
			args.NumOtherPhase = len(mapFiles)
			ok := call(work, "Worker.DoTask", args, new(struct{}))
			if ok == false {
				fmt.Printf("DoTask: RPC %s dotask error\n", work)
			}
			if ok {
				go func() {
					registerChan <- work
				}()

			} else {
				i = i-1//if work faield the work will be abandoned
				fmt.Printf("reduce Task %d field\n", i)
			}

		}

	}
	fmt.Printf("Schedule: %v phase done\n", phase)
}
