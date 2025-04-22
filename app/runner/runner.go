/*
Copyright © 2025 taksenov@gmail.com
*/

// Package runner -- ssh commands runner.
package runner

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"devops/app/customerrors"
	"devops/app/logger"
	"devops/app/misc/conf"
)

// SshRunner -- раннер выполняющий команды по ssh.
func SshRunner(ctx context.Context) {
	// Инициализация логгера в файл
	log, err := logger.InitFileLogger(
		"logs/log-"+time.Now().Format("2006-01-02--15-04-05")+".log", // путь к файлу лога
		"",               // префикс (не используется в нашей реализации)
		logger.LevelInfo, // уровень логирования
		true,             // включать информацию о месте вызова
	)
	if err != nil {
		panic(err)
	}
	defer log.Close()

	// Добавляем логгер в контекст
	ctx = logger.WithContext(ctx, log)

	config, err := conf.LoadConfig("config.yaml")
	if err != nil {
		log.Error("cannot load config: %s", err)
		customerrors.HandleErr(err, "cannot load config: SshRunner()")

	}

	if config.Concurency.Use {
		elapsedTime(concurrencyRun, ctx, config)
	} else {
		elapsedTime(consistentlyRun, ctx, config)
	}
}

// Последовательный запуск команд.
func consistentlyRun(ctx context.Context, config *conf.AppConfig) {
	logInfo(ctx, "Consistently Run")
	fmt.Println("Consistently Run")

	list := config.Servers.List
	sshCommand := config.Command.Main
	sshCommandParams := strings.Join(config.Command.Params, " ")
	onServCommands := config.RemoteCommand

	for k, v := range list {
		logInfo(ctx, "COMMAND INDEX= %d", k+1)
		fmt.Println("COMMAND INDEX=", k+1)

		for _, onServCommand := range onServCommands {
			fullOnservCommand := ""
			if onServCommand.IsTemplate {
				fullOnservCommand = strings.ReplaceAll(onServCommand.Command, "SERVER_NAME", setServerAddr(config, v))
			} else {
				fullOnservCommand = onServCommand.Command
			}
			logInfo(ctx, "Remote Command: %s", blueMsg(fullOnservCommand))
			fmt.Println("Remote Command:", blueMsg(fullOnservCommand))

			command := fmt.Sprintf(
				// Форматированная команда
				sshCommand,
				// Параметры форматирующие команду
				sshCommandParams,
				setServerAddr(config, v),
				fullOnservCommand,
			)
			logInfo(ctx, "%s", greenMsg(command))
			fmt.Println(greenMsg(command))

			cmdRun := exec.Command(
				"sh",
				"-c",
				command,
			)

			outputData, err := cmdRun.Output()
			logInfo(ctx, "Output Data: \n %s", string(outputData))
			fmt.Println("Output Data: \n", string(outputData))
			if err != nil {
				logInfo(ctx, "ERR: %s", err)
				fmt.Println("ERR:", err)
			}
		}

		time.Sleep(50 * time.Millisecond)

	}

	logInfo(ctx, "All requests done")
	fmt.Println("All requests done")
}

// Конкурентный запуск команд.
func concurrencyRun(ctx context.Context, config *conf.AppConfig) {
	list := config.Servers.List
	coeff := config.Concurency.Coeff
	if len(list) < coeff {
		coeff = len(list)
	}

	logInfo(ctx, "Concurency run. Coefficient: %d", coeff)
	fmt.Println("Concurency run. Coefficient:", coeff)

	var countReq int32

	wg := &sync.WaitGroup{}
	semaphoreChan := make(chan struct{}, coeff)
	var outputMutex sync.Mutex

	for _, v := range list {
		wg.Add(1)
		semaphoreChan <- struct{}{}

		contextLocal := ctx

		go func(ctx context.Context, c *conf.AppConfig) {
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

				cmdRun := exec.Command(
					"sh",
					"-c",
					command,
				)

				outputData, err := cmdRun.Output()

				outputMutex.Lock()
				logInfo(ctx, "%s", greenMsg(command))
				logInfo(ctx, "Remote Command: %s", blueMsg(fullOnservCommand))
				logInfo(ctx, "Output Data: \n %s", string(outputData))
				fmt.Println(greenMsg(command))
				fmt.Println("Remote Command:", blueMsg(fullOnservCommand))
				fmt.Println("Output Data: \n", string(outputData))
				if err != nil {
					logError(ctx, "ERR: %s", err)
					fmt.Println("ERR:", err)
				}
				outputMutex.Unlock()
			}

			atomic.AddInt32(&countReq, 1)
		}(contextLocal, config)

	}

	wg.Wait()
	logInfo(ctx, "All requests done. Count: %d", countReq)
	fmt.Println("All requests done. Count:", countReq)
}

// elapsedTime -- функция-декоратор подсчитывающая время работы.
func elapsedTime(f func(context.Context, *conf.AppConfig), ctx context.Context, c *conf.AppConfig) {
	start := time.Now()
	f(ctx, c)
	t := time.Now()
	elapsed := t.Sub(start)
	logInfo(ctx, "Elapsed time: %s", elapsed)
	fmt.Println("Elapsed time:", elapsed)
}

// setServerAddr -- сеттер имени или IP-адреса сервера.
func setServerAddr(c *conf.AppConfig, s conf.ServerAddr) string {
	if c.Servers.Use == "ADDR" {
		return s.Addr
	}

	return s.Name
}

func greenMsg(s string) string {
	return string("\033[92m") + s + string("\033[0m")
}

func blueMsg(s string) string {
	return string("\033[36m") + s + string("\033[0m")
}

func logInfo(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Info(format, args...)
}

func logError(ctx context.Context, format string, args ...interface{}) {
	logger.FromContext(ctx).Error(format, args...)
}
