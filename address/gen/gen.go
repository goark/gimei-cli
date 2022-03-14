//go:build run
// +build run

package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/goark/gimei-cli/address"
	"github.com/goark/gimei-cli/name"
	"gopkg.in/yaml.v3"
)

type item []string

type addressList struct {
	Addresses struct {
		Prefecture []item `yaml:"prefecture"`
		City       []item `yaml:"city"`
		Town       []item `yaml:"town"`
	} `yaml:"addresses"`
}

func loadPlaces(path string) (address.Places, address.Places, address.Places, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()

	var places addressList
	if err := yaml.NewDecoder(f).Decode(&places); err != nil {
		return nil, nil, nil, err
	}
	prefs := address.Places{}
	for _, i := range places.Addresses.Prefecture {
		prefs = append(prefs, name.New(i[0], i[2]))
	}
	cities := address.Places{}
	for _, i := range places.Addresses.City {
		cities = append(cities, name.New(i[0], i[2]))
	}
	towns := address.Places{}
	for _, i := range places.Addresses.Town {
		towns = append(towns, name.New(i[0], i[2]))
	}
	return prefs, cities, towns, nil
}

const template1 = `package address

import "github.com/goark/gimei-cli/name"

//Places is list of place names
type Places []*name.Name

var Prefectures = Places{
{{ range . }}	{Name: "{{ .Name }}", Katakana: "{{ .Katakana }}", Hiragana: "{{ .Hiragana }}", Roman: "{{ .Roman }}"},
{{ end }}}
`
const template2 = `
var Cities = Places{
{{ range . }}	{Name: "{{ .Name }}", Katakana: "{{ .Katakana }}", Hiragana: "{{ .Hiragana }}", Roman: "{{ .Roman }}"},
{{ end }}}
`
const template3 = `
var Towns = Places{
{{ range . }}	{Name: "{{ .Name }}", Katakana: "{{ .Katakana }}", Hiragana: "{{ .Hiragana }}", Roman: "{{ .Roman }}"},
{{ end }}}
`

func main() {
	prefs, cities, towns, err := loadPlaces("./addresses.yml")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	tpl1, err := template.New("template1").Parse(template1)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := tpl1.Execute(os.Stdout, prefs); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	tpl2, err := template.New("template2").Parse(template2)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := tpl2.Execute(os.Stdout, cities); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	tpl3, err := template.New("template3").Parse(template3)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if err := tpl3.Execute(os.Stdout, towns); err != nil {
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
