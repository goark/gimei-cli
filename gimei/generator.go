package gimei

import (
	"math/rand"
	"time"
)

// Generator is context for gimei generator.
type Generator struct {
	rand    rand.Source
	other   bool
	unique  bool
	history map[string]bool
}

// New function returns Generator instance.
func New() *Generator {
	return &Generator{
		rand:    rand.NewSource(time.Now().UnixNano()),
		other:   false,
		unique:  false,
		history: map[string]bool{},
	}
}

// WithRandSource method sets rand.Source.
func (gen *Generator) WithRandSource(src rand.Source) *Generator {
	if gen == nil {
		gen = New()
	}
	if src != nil {
		gen.rand = src
	}
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

// WithGenderOther methos sets other flag.
func (gen *Generator) WithGenderOther(flag bool) *Generator {
	if gen == nil {
		gen = New()
	}
	gen.other = flag
	return gen
}

// New method return new Gimei instance.
func (gen *Generator) New() *Gimei {
	for {
		gimei := NewGimei(gen.rand, gen.gender())
		if gen.other && (gen.rand.Int63()&0x0f) == 0 {
			gimei.Gender = Other
		}
		if !gen.unique {
			return gimei
		}
		fullName := gimei.FullNameWithSep(" ").Name
		if !gen.history[fullName] {
			gen.history[fullName] = true
			return gimei
		}
	}
}

func (gen *Generator) gender() Gender {
	return Gender((gen.rand.Int63() & 0x01))
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
