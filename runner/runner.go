package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner 在给定的超时时间内执行一组任务
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

//ErrTimeout 任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

//ErrInterrupt 接收到操作系统事件时返回
var ErrInterrupt = errors.New("received interrupt")

//NewRunner 返回一个新的Runner
func NewRunner(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add 将任务附加到Runner上
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {
	//接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

//run 执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		//检测系统中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		//执行任务
		task(id)
	}

	return nil
}

//gotInterrupt 验证是否收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
