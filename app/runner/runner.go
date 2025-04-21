/*
Copyright © 2025 taksenov@gmail.com
*/

// Package runner -- ssh commands runner.
package runner

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"devops/app/customerrors"
	"devops/app/misc/conf"
)

// SshRunner -- раннер выполняющий команды по ssh.
func SshRunner() {
	config, err := conf.LoadConfig("config.yaml")
	if err != nil {
		customerrors.HandleErr(err, "cannot load config: SshRunner()")
	}

	if config.Concurency.Use {
		elapsedTime(concurrencyRun, config)
	} else {
		elapsedTime(consistentlyRun, config)
	}
}

// Последовательный запуск команд.
func consistentlyRun(config *conf.AppConfig) {
	fmt.Println("Consistently Run")

	list := config.Servers.List
	sshCommand := config.Command.Main
	sshCommandParams := strings.Join(config.Command.Params, " ")
	onServCommands := config.RemoteCommand

	for k, v := range list {
		fmt.Println("COMMAND INDEX=", k+1)

		for _, onServCommand := range onServCommands {
			fullOnservCommand := ""
			if onServCommand.IsTemplate {
				fullOnservCommand = strings.ReplaceAll(onServCommand.Command, "SERVER_NAME", setServerAddr(config, v))
			} else {
				fullOnservCommand = onServCommand.Command
			}
			fmt.Println("Remote Command:", fullOnservCommand)

			command := fmt.Sprintf(
				// Форматированная команда
				sshCommand,
				// Параметры форматирующие команду
				sshCommandParams,
				setServerAddr(config, v),
				fullOnservCommand,
			)
			fmt.Println(command)

			cmdPatch := exec.Command(
				"sh",
				"-c",
				command,
			)

			outputData, err := cmdPatch.Output()
			fmt.Println("outputData:", string(outputData))
			if err != nil {
				fmt.Println("ERR:", err)
			}
		}

		time.Sleep(50 * time.Millisecond)

	}

	fmt.Println("All requests done")
}

// Конкурентный запуск команд.
func concurrencyRun(config *conf.AppConfig) {
	fmt.Println("Concurency run. Coefficient:", config.Concurency.Coeff)

	list := config.Servers.List

	coeff := config.Concurency.Coeff
	if len(list) < coeff {
		coeff = len(list)
	}

	var countReq int32

	wg := &sync.WaitGroup{}
	semaphoreChan := make(chan struct{}, coeff)

	for _, v := range list {
		wg.Add(1)
		semaphoreChan <- struct{}{}

		go func(c *conf.AppConfig) {
			defer func() {
				<-semaphoreChan
				wg.Done()
			}()

			sshCommand := c.Command.Main
			sshCommandParams := strings.Join(c.Command.Params, " ")
			onServCommands := c.RemoteCommand

			for _, onServCommand := range onServCommands {
				fullOnservCommand := ""
				if onServCommand.IsTemplate {
					fullOnservCommand = strings.ReplaceAll(onServCommand.Command, "SERVER_NAME", setServerAddr(c, v))
				} else {
					fullOnservCommand = onServCommand.Command
				}

				command := fmt.Sprintf(
					// Форматированная команда
					sshCommand,
					// Параметры форматирующие команду
					sshCommandParams,
					setServerAddr(c, v),
					fullOnservCommand,
				)

				cmdPatch := exec.Command(
					"sh",
					"-c",
					command,
				)

				outputData, err := cmdPatch.Output()
				fmt.Println(command)
				fmt.Println("Remote Command:", fullOnservCommand)
				fmt.Println("outputData:", string(outputData))
				if err != nil {
					fmt.Println("ERR:", err)
				}
			}

			atomic.AddInt32(&countReq, 1)
		}(config)

	}

	wg.Wait()
	fmt.Println("All requests done. Count: ", countReq)
}

// elapsedTime -- функция-декоратор подсчитывающая время работы.
func elapsedTime(f func(*conf.AppConfig), c *conf.AppConfig) {
	start := time.Now()
	f(c)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed time:", elapsed)
}

// setServerAddr -- сеттер имени или IP-адреса сервера.
func setServerAddr(c *conf.AppConfig, s conf.ServerAddr) string {
	if c.Servers.Use == "ADDR" {
		return s.Addr
	}

	return s.Name
}
