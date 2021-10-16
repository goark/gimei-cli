package gimei

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spiegel-im-spiegel/gimei-cli/name"
)

// DomainExample is "example.com" domain name.
const DomainExample = "example.com"

// Gimei is pseudo personal name (gimei) information.
type Gimei struct {
	Gender    Gender
	FirstName *name.Name
	LastName  *name.Name
}

// NewGimei function returns new Gimei instance.
func NewGimei(rand rand.Source, g Gender) *Gimei {
	gimei := &Gimei{Gender: g, FirstName: nil, LastName: nil}
	if g == Female {
		gimei.FirstName = Females[int(rand.Int63())%len(Females)]
	} else {
		gimei.FirstName = Males[int(rand.Int63())%len(Males)]
	}
	gimei.LastName = LastNames[int(rand.Int63())%len(LastNames)]
	return gimei
}

func (g *Gimei) String() string {
	n := g.FullNameWithSep(" ")
	return fmt.Sprintf("%s (%s : %v)", n.Name, n.Katakana, g.Gender)
}

// FullName method returns full name of gimei.
func (g *Gimei) FullName() *name.Name {
	return g.FullNameWithSep("")
}

// FullNameWithSep method returns full name of gimei with the separator.
func (g *Gimei) FullNameWithSep(sep string) *name.Name {
	if g == nil {
		return &name.Name{}
	}
	return &name.Name{
		Name:     strings.Join([]string{g.LastName.Name, g.FirstName.Name}, sep),
		Katakana: strings.Join([]string{g.LastName.Katakana, g.FirstName.Katakana}, sep),
		Hiragana: strings.Join([]string{g.LastName.Hiragana, g.FirstName.Hiragana}, sep),
		Roman:    strings.Join([]string{g.LastName.Roman, g.FirstName.Roman}, " "),
	}
}

// Email method returns email address.
func (g *Gimei) Email() string {
	return g.EmailWithDomain(DomainExample)
}

// EmailWithDomain method returns email address with the domain name.
func (g *Gimei) EmailWithDomain(domain string) string {
	if g == nil {
		return ""
	}
	return string([]rune(g.FirstName.Roman)[0:1]) + "." + g.LastName.Roman + "@" + domain
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
