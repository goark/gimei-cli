package all

import (
	"math/rand"

	"github.com/goark/gimei-cli/address"
	"github.com/goark/gimei-cli/gimei"
)

// Generator is context for gimei and address generator.
type Generator struct {
	gimeiGenerator   *gimei.Generator
	addressGenerator *address.Generator
	unique           bool
	history          map[string]bool
}

// New function returns Generator instance.
func New() *Generator {
	return &Generator{
		gimeiGenerator:   gimei.New(),
		addressGenerator: address.New(),
		unique:           false,
		history:          map[string]bool{},
	}
}

// WithRandSource methos sets rand.Source.
func (gen *Generator) WithRandSource(src rand.Source) *Generator {
	if gen == nil {
		gen = New()
	}
	gen.gimeiGenerator.WithRandSource(src)
	gen.addressGenerator.WithRandSource(src)
	return gen
}

// WithUnique method sets unique flag.
func (gen *Generator) WithUnique(flag bool) *Generator {
	if gen == nil {
		gen = New()
	}
	gen.unique = flag
	return gen
}

// WithFilters method sets filters.
func (gen *Generator) WithFilters(keyword string) *Generator {
	if gen == nil {
		gen = New()
	}
	gen.addressGenerator.WithFilters(keyword)
	return gen
}

// WithGenderOther methos sets other flag.
func (gen *Generator) WithGenderOther(flag bool) *Generator {
	if gen == nil {
		gen = New()
	}
	gen.gimeiGenerator.WithGenderOther(flag)
	return gen
}

// New method return new Info instance.
func (gen *Generator) New() *Info {
	for {
		info := &Info{
			Gimei:   gen.gimeiGenerator.New(),
			Address: gen.addressGenerator.New(),
		}
		if !gen.unique {
			return info
		}
		fullpath := info.serialize()
		if !gen.history[fullpath] {
			gen.history[fullpath] = true
			return info
		}
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
