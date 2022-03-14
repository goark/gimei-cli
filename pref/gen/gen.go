//go:build run
// +build run

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/goark/csvdata"
	"github.com/goark/csvdata/exceldata"
	"github.com/goark/errs"
	"github.com/goark/gimei-cli/name"
	"github.com/goark/gimei-cli/pref"
)

func importPref(path, sheet string) ([]*pref.Pref, []*pref.City, error) {
	xlsx, err := exceldata.OpenFile(path, "") // no password
	if err != nil {
		return nil, nil, errs.Wrap(err, errs.WithContext("path", path), errs.WithContext("sheet", sheet))
	}
	r, err := exceldata.New(xlsx, sheet)
	if err != nil {
		return nil, nil, errs.Wrap(err, errs.WithContext("path", path), errs.WithContext("sheetName", sheet))
	}
	rows := csvdata.NewRows(r, true)
	rows.Close() // dummy

	prefs := map[string]*pref.Pref{}
	cities := map[string]*pref.City{}
	for {
		if err := rows.Next(); err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return nil, nil, errs.Wrap(err)
		}
		code := rows.Get(0)
		prefName := rows.Get(1)
		cityName := rows.Get(2)
		prefNameKana := rows.Get(3)
		cityNameKana := rows.Get(4)
		if len(cityName) > 0 {
			p := &pref.Pref{Code: code[:2], Name: name.New(prefName, prefNameKana)}
			prefs[code[:2]] = p
			cities[code[:5]] = &pref.City{Code: code[:5], Pref: p, Name: name.New(cityName, cityNameKana)}
		}
	}
	prefList := []*pref.Pref{}
	for _, p := range prefs {
		prefList = append(prefList, p)
	}
	sort.Slice(prefList, func(i, j int) bool { return strings.Compare(prefList[i].Code, prefList[j].Code) < 0 })
	cityList := []*pref.City{}
	for _, c := range cities {
		cityList = append(cityList, c)
	}
	sort.Slice(cityList, func(i, j int) bool { return strings.Compare(cityList[i].Code, cityList[j].Code) < 0 })
	return prefList, cityList, nil
}

const template1 = `package pref

import "github.com/goark/gimei-cli/name"

var Prefs = map[string]*Pref{
{{ range . }}	"{{ .Code }}": {Code: "{{ .Code }}", Name: &name.Name{Name: "{{ .Name.Name }}", Katakana: "{{ .Name.Katakana }}", Hiragana: "{{ .Name.Hiragana }}", Roman: "{{ .Name.Roman }}"}},
{{ end }}}
`
const template2 = `
var Cities = map[string]*City{
{{ range . }}	"{{ .Code }}": {Code: "{{ .Code }}", Pref: &Pref{Code: "{{ .Pref.Code }}", Name: &name.Name{Name: "{{ .Pref.Name.Name }}", Katakana: "{{ .Pref.Name.Katakana }}", Hiragana: "{{ .Pref.Name.Hiragana }}", Roman: "{{ .Pref.Name.Roman }}"}}, Name: &name.Name{Name: "{{ .Name.Name }}", Katakana: "{{ .Name.Katakana }}", Hiragana: "{{ .Name.Hiragana }}", Roman: "{{ .Name.Roman }}"}},
{{ end }}}
`

func main() {
	prefs, cities, err := importPref("000730858.xlsx", "")
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
	// fmt.Println(prefs)
	// fmt.Println(cities)
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
