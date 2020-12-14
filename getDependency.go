package main

import (
	"fmt"
	"os/exec"
	"sync"
)

var pkgs = [...]string{
	"github.com/easierway/concurrent_map",
	"github.com/stretchr/testify/assert",
	"github.com/smartystreets/goconvey/convey",
}

func main() {
	wg := sync.WaitGroup{}
	for _, pkg := range pkgs {
		wg.Add(1)
		go func(pkg string) {
			defer func() {
				wg.Done()
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			exec.Command("go get", "-u", pkg)
		}(pkg)
	}
	wg.Wait()
}
