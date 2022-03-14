package all_test

import (
	"fmt"

	"github.com/goark/gimei-cli/all"
	"github.com/goark/mt/mt19937"
)

func ExampleNew() {
	gen := all.New().WithRandSource(mt19937.New(19650218)).WithUnique(true).WithFilters("松江").WithGenderOther(true)
	for i := 0; i < 15; i++ {
		fmt.Println(gen.New())
	}
	// Output:
	// 宮原 比呂志 (ミヤハラ ヒロシ : 男性) -> 島根県 松江市 三木里町 (シマネケン マツエシ ミキサトチョウ)
	// 梅田 鈴望 (ウメダ レミ : 女性) -> 島根県 松江市 古市場 (シマネケン マツエシ フルイチバ)
	// 庄司 洋裕 (ショウジ ヒロヤス : 男性) -> 島根県 松江市 芥田 (シマネケン マツエシ アクタダ)
	// 稲垣 美晴 (イナガキ ミハル : 女性) -> 島根県 松江市 日置野田 (シマネケン マツエシ ヘキノダ)
	// 菊池 鼓太郎 (キクチ コタロウ : その他) -> 島根県 松江市 三輪 (シマネケン マツエシ ミワ)
	// 竹内 乃聖 (タケウチ ノア : 女性) -> 島根県 松江市 尾崎丁 (シマネケン マツエシ オザキチョウ)
	// 宮田 芽 (ミヤタ モエ : 女性) -> 島根県 松江市 本宮北ノ内 (シマネケン マツエシ モトミヤキタノウチ)
	// 藤原 十実 (フジワラ トミ : 女性) -> 島根県 松江市 原町田 (シマネケン マツエシ ハラマチダ)
	// 青木 光彦 (アオキ ミツヒコ : 男性) -> 島根県 松江市 波積町本郷 (シマネケン マツエシ ハヅミチョウホンゴウ)
	// 堀内 美風 (ホリウチ ミカ : 女性) -> 島根県 松江市 衣川区桧山沢 (シマネケン マツエシ コロモガワクヒヤマサワ)
	// 内田 麗流 (ウチダ レナ : 女性) -> 島根県 松江市 浦川原区上柿野 (シマネケン マツエシ ウラガワラクカミガキノ)
	// 園田 逸平 (ソノダ イッペイ : 男性) -> 島根県 松江市 留辺蘂町金華 (シマネケン マツエシ ルベシベチョウカネハナ)
	// 中川 洸 (ナカガワ コウ : 男性) -> 島根県 松江市 上熊本 (シマネケン マツエシ カミクマモト)
	// 森山 照高 (モリヤマ テルタカ : 男性) -> 島根県 松江市 井堀中郷町 (シマネケン マツエシ イボリナカゴウチョウ)
	// 柴田 満夫 (シバタ ミツオ : 男性) -> 島根県 松江市 飯島道東 (シマネケン マツエシ イイジマミチヒガシ)
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
