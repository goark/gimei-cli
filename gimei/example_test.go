package gimei_test

import (
	"fmt"

	"github.com/goark/gimei-cli/gimei"
	"github.com/goark/mt/mt19937"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ExampleNewGimei() {
	fmt.Println(gimei.NewGimei(mt19937.New(19650218), gimei.Male))
	fg := gimei.NewGimei(mt19937.New(19650218), gimei.Female)
	fmt.Println(fg.FullName().Name)
	fmt.Println(fg.FullName().Katakana)
	fmt.Println(fg.FullName().Hiragana)
	fmt.Println(cases.Title(language.Und, cases.NoLower).String(fg.FullName().Roman))
	fmt.Println(fg.Email())
	// Output:
	// 菊池 貴志 (キクチ タカシ : 男性)
	// 菊池友美
	// キクチトモミ
	// きくちともみ
	// Kikuchi Tomomi
	// t.kikuchi@example.com
}

func ExampleNew() {
	gen := gimei.New().WithRandSource(mt19937.New(19650218)).WithGenderOther(true).WithUnique(true)
	for i := 0; i < 15; i++ {
		fmt.Println(gen.New())
	}
	// Output:
	// 宮原 比呂志 (ミヤハラ ヒロシ : 男性)
	// 小島 麗夏 (コジマ レイナ : 女性)
	// 中川 温彦 (ナカガワ アツヒコ : 男性)
	// 渋谷 範久 (シブヤ ノリヒサ : 男性)
	// 奥村 櫂斗 (オクムラ カイト : 男性)
	// 田島 由惠 (タジマ ヨシエ : 女性)
	// 寺田 歩武 (テラダ アユム : 男性)
	// 上野 謙一 (ウエノ ケンイチ : 男性)
	// 村瀬 殊莉 (ムラセ コトリ : 女性)
	// 荒川 望叶 (アラカワ ミカ : 女性)
	// 福永 美妙 (フクナガ ミサ : 女性)
	// 下村 智浩 (シモムラ トモヒロ : 男性)
	// 小谷 水希 (コタニ ミキ : その他)
	// 今井 氷魚 (イマイ ヒオ : 男性)
	// 桑原 草太 (クワハラ ソウタ : その他)
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
