package address

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spiegel-im-spiegel/gimei-cli/name"
)

// Address is pseudo place name information.
type Address struct {
	Prefecture *name.Name
	City       *name.Name
	Town       *name.Name
}

// NewAddress function returns new Address instance.
func NewAddress(rand rand.Source) *Address {
	addr := &Address{}
	addr.Prefecture = Prefectures[int(rand.Int63())%len(Prefectures)]
	addr.City = Cities[int(rand.Int63())%len(Cities)]
	addr.Town = Towns[int(rand.Int63())%len(Towns)]
	return addr
}

func (a *Address) String() string {
	n := a.FullNameWithSep(" ")
	return fmt.Sprintf("%s (%s)", n.Name, n.Katakana)
}

// FullName method returns full name of place.
func (a *Address) FullName() *name.Name {
	return a.FullNameWithSep("")
}

// FullNameWithSep method returns full name of place with the separator.
func (a *Address) FullNameWithSep(sep string) *name.Name {
	if a == nil {
		return &name.Name{}
	}
	return &name.Name{
		Name:     strings.Join([]string{a.Prefecture.Name, a.City.Name, a.Town.Name}, sep),
		Katakana: strings.Join([]string{a.Prefecture.Katakana, a.City.Katakana, a.Town.Katakana}, sep),
		Hiragana: strings.Join([]string{a.Prefecture.Hiragana, a.City.Hiragana, a.Town.Hiragana}, sep),
		Roman:    strings.Join([]string{a.Prefecture.Roman, a.City.Roman, a.Town.Roman}, " "),
	}
}

// Prefix method returns true if matches with list of addresses.
func (a *Address) Prefix(addrs []string) bool {
	if a == nil {
		return false
	}
	if len(addrs) == 0 {
		return true
	}
	n := a.FullName().Name
	for _, a := range addrs {
		if strings.HasPrefix(n, a) {
			return true
		}
	}
	return false
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
