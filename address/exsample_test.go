package address_test

import (
	"fmt"
	"strings"

	"github.com/spiegel-im-spiegel/gimei-cli/address"
	"github.com/spiegel-im-spiegel/mt/mt19937"
)

func ExampleNewAddress() {
	a := address.NewAddress(mt19937.New(19650218))
	fmt.Println(a.FullName().Name)
	fmt.Println(a.FullName().Katakana)
	fmt.Println(a.FullName().Hiragana)
	fmt.Println(strings.Title(a.FullName().Roman))
	// Output:
	// 奈良県富士市ダイハツ町
	// ナラケンフジシダイハツチョウ
	// ならけんふじしだいはつちょう
	// Naraken Fujishi Daihatsuchou
}

func ExampleNew() {
	gen := address.New().WithRandSource(mt19937.New(19650218)).WithUnique(true).WithFilters("松江")
	for i := 0; i < 15; i++ {
		fmt.Println(gen.New())
	}
	// Output:
	// 島根県 松江市 上磯分内 (シマネケン マツエシ カミイソブンナイ)
	// 島根県 松江市 芥田 (シマネケン マツエシ アクタダ)
	// 島根県 松江市 日ノ出 (シマネケン マツエシ ヒノデ)
	// 島根県 松江市 前沢区二十人町 (シマネケン マツエシ マエサワクニジュウニンマチ)
	// 島根県 松江市 尾崎丁 (シマネケン マツエシ オザキチョウ)
	// 島根県 松江市 北沢 (シマネケン マツエシ キタザワ)
	// 島根県 松江市 波積町本郷 (シマネケン マツエシ ハヅミチョウホンゴウ)
	// 島根県 松江市 幸町 (シマネケン マツエシ サイワイチョウ)
	// 島根県 松江市 弓越 (シマネケン マツエシ ミコシ)
	// 島根県 松江市 大蔵 (シマネケン マツエシ オオゾウ)
	// 島根県 松江市 留辺蘂町金華 (シマネケン マツエシ ルベシベチョウカネハナ)
	// 島根県 松江市 材木町 (シマネケン マツエシ ザイモクマチ)
	// 島根県 松江市 木興町 (シマネケン マツエシ キコチョウ)
	// 島根県 松江市 飯島道東 (シマネケン マツエシ イイジマミチヒガシ)
	// 島根県 松江市 江上町 (シマネケン マツエシ エガミチョウ)
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
