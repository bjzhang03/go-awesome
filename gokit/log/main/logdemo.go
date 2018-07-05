package main

import (
	"github.com/go-kit/kit/log"
	"os"
	"fmt"
	"time"
	"sync"
	"math/rand"
)

type Logger interface {
	Log(keyvals ...interface{}) error
}
type Task struct {
	ID int
}

func RunTask(task Task, logger log.Logger) string {
	logger.Log("taskID", task.ID, "event", "starting task")
	logger.Log("taskID", task.ID, "event", "task complete")
	return ""
}
func basic() {
	logger := log.NewLogfmtLogger(os.Stdout)

	RunTask := func(task Task, logger log.Logger) {
		logger.Log("taskID", task.ID, "event", "starting task")

		logger.Log("taskID", task.ID, "event", "task complete")
	}

	RunTask(Task{ID: 1}, logger)
}

func contextual() {
	logger := log.NewLogfmtLogger(os.Stdout)

	// 重写了上面定义的数据结构
	type Task struct {
		ID  int
		Cmd string
	}

	taskHelper := func(cmd string, logger log.Logger) {
		// execute(cmd)
		logger.Log("cmd", cmd, "dur", 42*time.Millisecond)
	}

	RunTask := func(task Task, logger log.Logger) {
		logger = log.With(logger, "taskID", task.ID)
		logger.Log("event", "starting task")

		taskHelper(task.Cmd, logger)

		logger.Log("event", "task complete")
	}

	RunTask(Task{ID: 1, Cmd: "echo Hello, world!"}, logger)
}

func debuginfo() {
	logger := log.NewLogfmtLogger(os.Stdout)
	// make time predictable for this test
	baseTime := time.Date(2015, time.February, 3, 10, 0, 0, 0, time.UTC)
	mockTime := func() time.Time {
		baseTime = baseTime.Add(time.Second)
		return baseTime
	}
	logger = log.With(logger, "time", log.Timestamp(mockTime), "caller", log.DefaultCaller)
	logger.Log("call", "first")
	logger.Log("call", "second")
	// ...
	logger.Log("call", "third")
}

func syncwriter() {
	w := log.NewSyncWriter(os.Stdout)
	logger := log.NewLogfmtLogger(w)

	type Task struct {
		ID int
	}
	var wg sync.WaitGroup
	RunTask := func(task Task, logger log.Logger) {
		logger.Log("taskID", task.ID, "event", "starting task")
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		logger.Log("taskID", task.ID, "event", "task complete")
		wg.Done()
	}
	wg.Add(2)
	go RunTask(Task{ID: 1}, logger)
	go RunTask(Task{ID: 2}, logger)
	wg.Wait()
}

func valuer() {
	logger := log.NewLogfmtLogger(os.Stdout)
	count := 0
	counter := func() interface{} {
		count++
		return count
	}
	logger = log.With(logger, "count", log.Valuer(counter))
	logger.Log("call", "first")
	logger.Log("call", "second")
}

func main() {
	fmt.Println("basic:")
	basic();
	fmt.Println("contextual:")
	contextual();
	fmt.Println("debuginfo:")
	debuginfo()
	fmt.Println("syncwriter:")
	syncwriter()
	fmt.Println("valuer:")
	valuer()
}
