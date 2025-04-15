/*
Copyright © 2025 taksenov@gmail.com
*/

// Package filesystem -- utilities for working with the file system.
package filesystem

import (
	"bufio"
	"devops/app/customerrors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

func ReadFileByLines(file string) (string, error) {
	res := ""

	var err error
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("[WARNING]: reading file %s", err)
			break
		}

		res += l
	}

	return res, nil
}

func filterEmptyLinesInFile(file string) error {
	mustFiltered, err := ReadFileByLines(file)
	if err != nil {
		return err
	}

	filtered := make([]string, 0)
	b := 0
	splitted := strings.Split(mustFiltered, "\n")
	for _, v := range splitted {
		if v == "" && b == 0 {
			filtered = append(filtered, v)
			b++
		}
		if v != "" {
			filtered = append(filtered, v)
			b = 0
		}
	}

	joined := strings.Join(filtered, "\n")

	err = os.WriteFile(file, []byte(joined), 0o644) //#nosec
	if err != nil {
		return err
	}

	return nil
}

func GetManifestNames(src string) ([]string, error) {
	dirs, err := os.ReadDir(src)
	if err != nil {
		return nil, fmt.Errorf("err in 'GetManifestNames'. message: %w", err)
	}

	res := make([]string, 0)

	for _, el := range dirs {
		if el.IsDir() {
			continue
		}

		res = append(res, el.Name())
	}

	return res, nil
}

func GeneratePipe(t interface{}, dst string, tmpl string) error {
	tmplData, err := ReadFileByLines(tmpl)
	if err != nil {
		customerrors.HandleErr(err, "GeneratePipe")
	}

	tpl, err := template.New("new").Parse(tmplData)
	if err != nil {
		return err
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tpl.Execute(f, t)
	if err != nil {
		return err
	}

	err = filterEmptyLinesInFile(dst)
	if err != nil {
		return err
	}

	return nil
}
