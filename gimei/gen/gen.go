//go:build run
// +build run

package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/goark/gimei-cli/gimei"
	"github.com/goark/gimei-cli/name"
	"gopkg.in/yaml.v3"
)

type item []string

type gimeiList struct {
	FirstName struct {
		Male   []item `yaml:"male"`
		Female []item `yaml:"female"`
	} `yaml:"first_name"`
	LastName []item `yaml:"last_name"`
}

func loadNames(path string) (gimei.Names, gimei.Names, gimei.Names, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()

	var names gimeiList
	if err := yaml.NewDecoder(f).Decode(&names); err != nil {
		return nil, nil, nil, err
	}
	males := gimei.Names{}
	for _, i := range names.FirstName.Male {
		males = append(males, name.New(i[0], i[2]))
	}
	females := gimei.Names{}
	for _, i := range names.FirstName.Female {
		females = append(females, name.New(i[0], i[2]))
	}
	lastNames := gimei.Names{}
	for _, i := range names.LastName {
		lastNames = append(lastNames, name.New(i[0], i[2]))
	}
	return males, females, lastNames, nil
}

const template1 = `package gimei

import "github.com/goark/gimei-cli/name"

//Names is list of names
type Names []*name.Name

var Males = Names{
{{ range . }}	{Name: "{{ .Name }}", Katakana: "{{ .Katakana }}", Hiragana: "{{ .Hiragana }}", Roman: "{{ .Roman }}"},
{{ end }}}
`
const template2 = `
var Females = Names{
{{ range . }}	{Name: "{{ .Name }}", Katakana: "{{ .Katakana }}", Hiragana: "{{ .Hiragana }}", Roman: "{{ .Roman }}"},
{{ end }}}
`
const template3 = `
var LastNames = Names{
{{ range . }}	{Name: "{{ .Name }}", Katakana: "{{ .Katakana }}", Hiragana: "{{ .Hiragana }}", Roman: "{{ .Roman }}"},
{{ end }}}
`

func main() {
	males, females, lastNames, err := loadNames("./names.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	tpl1, err := template.New("template1").Parse(template1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := tpl1.Execute(os.Stdout, males); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	tpl2, err := template.New("template2").Parse(template2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := tpl2.Execute(os.Stdout, females); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	tpl3, err := template.New("template3").Parse(template3)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := tpl3.Execute(os.Stdout, lastNames); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

/* MIT License
 *
 * Copyright 2021 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
