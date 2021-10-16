package pref

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spiegel-im-spiegel/gimei-cli/name"
)

// Pref is information of prefecture.
type Pref struct {
	Code string
	*name.Name
}

func (p *Pref) String() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("%s: %s (%s)", p.Code, p.Name.Name, p.Katakana)
}

// City is information of city.
type City struct {
	Code string
	Pref *Pref
	*name.Name
}

func (c *City) String() string {
	if c == nil {
		return ""
	}
	name := c.FullName()
	return fmt.Sprintf("%s: %s (%s)", c.Code, name.Name, name.Katakana)
}

// FullName method returns full name of city.
func (c *City) FullName() *name.Name {
	return c.FullNameWithSep("")
}

// FullNameWithSep method returns full name of city with the separator.
func (c *City) FullNameWithSep(sep string) *name.Name {
	if c == nil {
		return &name.Name{}
	}
	return &name.Name{
		Name:     strings.Join([]string{c.Pref.Name.Name, c.Name.Name}, sep),
		Katakana: strings.Join([]string{c.Pref.Name.Katakana, c.Name.Katakana}, sep),
		Hiragana: strings.Join([]string{c.Pref.Name.Hiragana, c.Name.Hiragana}, sep),
		Roman:    strings.Join([]string{c.Pref.Name.Roman, c.Name.Roman}, " "),
	}
}

// PrefCode function returns Pref instance by pref-code.
func PrefCode(code string) *Pref {
	return Prefs[code]
}

// PrefName function returns Pref instances by pref. name.
func PrefName(name string) []*Pref {
	prefs := []*Pref{}
	for _, v := range Prefs {
		if name == "" || strings.Contains(v.Name.Name, name) {
			prefs = append(prefs, v)
		}
	}
	return sortPrefs(prefs)
}

// CityCode function returns City instance by city-code.
func CityCode(code string) *City {
	return Cities[code]
}

// CityName function returns City instances by city name.
func CityName(name string) []*City {
	cities := []*City{}
	for _, v := range Cities {
		if name == "" || strings.Contains(v.FullName().Name, name) {
			cities = append(cities, v)
		}
	}
	return sortCities(cities)
}

func sortCities(cities []*City) []*City {
	sort.Slice(cities, func(i, j int) bool { return strings.Compare(cities[i].Code, cities[j].Code) < 0 })
	return cities
}

func sortPrefs(prefs []*Pref) []*Pref {
	sort.Slice(prefs, func(i, j int) bool { return strings.Compare(prefs[i].Code, prefs[j].Code) < 0 })
	return prefs
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
