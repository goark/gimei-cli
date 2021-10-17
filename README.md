# [gimei-cli] -- 姓名・住所データ生成ツール

[![check vulns](https://github.com/spiegel-im-spiegel/gimei-cli/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/gimei-cli/actions)
[![lint status](https://github.com/spiegel-im-spiegel/gimei-cli/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/gimei-cli/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/gimei-cli/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/gimei-cli.svg)](https://github.com/spiegel-im-spiegel/gimei-cli/releases/latest)

[willnet/gimei](https://github.com/willnet/gimei "willnet/gimei: random Japanese name and address generator") および [mattn/go-gimei](https://github.com/mattn/go-gimei) からの fork で，コマンドライン上で姓名や住所名のダミーデータを大量に生成することを目標に作成しています。

## ビルド & インストール

```
$ go install github.com/spiegel-im-spiegel/gimei-cli@latest
```

### 実行バイナリ

See [latest release](https://github.com/spiegel-im-spiegel/gimei-cli/releases/latest).

## 簡単な使い方

```
$ gimei-cli -h
姓名・住所データ生成ツール

Usage:
  gimei-cli [flags]
  gimei-cli [command]

Available Commands:
  address     住所名の生成
  cities      市区町村の表示
  completion  generate the autocompletion script for the specified shell
  gimei       姓名の生成
  help        Help about any command
  prefectures 都道府県の表示
  version     バージョン表示

Flags:
      --crypt-rand             crypt/rand 乱数生成器を使う
      --debug                  for debug
  -g, --gender                 「その他」の性別を含める
  -h, --help                   help for gimei-cli
  -j, --json                   JSON形式で表示
  -n, --name string            絞り込み市区町村名（一部でも可）
  -r, --repeat int             試行回数 (default 1)
  -t, --template-path string   テンプレートファイルへのパス
  -u, --unique                 姓名＋住所名の重複を避ける

Use "gimei-cli [command] --help" for more information about a command.

$ gimei-cli
岡村 智梨 (オカムラ サトリ : 女性) -> 長崎県 北杜市 俵柳 (ナガサキケン ホクトシ タワラヤナギ)

$ gimei-cli --json | jq .
[
  {
    "Gimei": {
      "Gender": 0,
      "FirstName": {
        "Name": "洋輔",
        "Katakana": "ヨウスケ",
        "Hiragana": "ようすけ",
        "Roman": "yousuke"
      },
      "LastName": {
        "Name": "福原",
        "Katakana": "フクハラ",
        "Hiragana": "ふくはら",
        "Roman": "fukuhara"
      }
    },
    "Address": {
      "Prefecture": {
        "Name": "大阪府",
        "Katakana": "オオサカフ",
        "Hiragana": "おおさかふ",
        "Roman": "oosakafu"
      },
      "City": {
        "Name": "安芸郡東洋町",
        "Katakana": "アキグントウヨウチョウ",
        "Hiragana": "あきぐんとうようちょう",
        "Roman": "akiguntouyouchou"
      },
      "Town": {
        "Name": "久々宇",
        "Katakana": "クグウ",
        "Hiragana": "くぐう",
        "Roman": "kuguu"
      }
    }
  }
]
```

- 姓名および住所名のデータは [willnet/gimei](https://github.com/willnet/gimei "willnet/gimei: random Japanese name and address generator") のものを利用している
- 性別は男性・女性が概ね半々になるよう調整している。また `--gender` (`-g`) オプションを付けると 1/16 の確率で「その他」の性が出現する
- 住所名は Prefecture, City, Town をランダムに組み合わせているため実在しない住所が多く出る。なお `--name` (`-n`) オプションでキーワードを指定するとキーワードを含む市区町村で絞り込むことができる

```
$ gimei-cli -r 10 -n 広島
高松 昌史 (タカマツ マサシ : 男性) -> 広島県 府中市 矢田 (ヒロシマケン フチュウシ ヤタ)
平山 颯馬 (ヒラヤマ ソウマ : 男性) -> 広島県 大竹市 木伏 (ヒロシマケン オオタケシ キップシ)
中原 成果 (ナカハラ ナルミ : 女性) -> 広島県 三次市 丸山町 (ヒロシマケン ミヨシシ マルヤマチョウ)
和田 鮎夢 (ワダ アユム : 男性) -> 北海道 北広島市 御津町豊沢樽美 (ホッカイドウ キタヒロシマシ ミトチョウトヨサワタルミ)
前川 啓文 (マエカワ ヒロフミ : 男性) -> 広島県 東広島市 相島 (ヒロシマケン ヒガシヒロシマシ アイシマ)
向井 竜次 (ムカイ リュウジ : 男性) -> 広島県 広島市安佐北区 赤城町 (ヒロシマケン ヒロシマシアサキタク アカシロチョウ)
水野 成樹 (ミズノ セイキ : 男性) -> 広島県 安芸高田市 美濃池町 (ヒロシマケン アキタカタシ ミノノイケチョウ)
平川 昌紀 (ヒラカワ マサキ : 男性) -> 広島県 廿日市市 八幡町 (ヒロシマケン ハツカイチシ ハチマンチョウ)
岡本 夏甫 (オカモト カホ : 女性) -> 広島県 広島市安佐南区 小谷美濃山町 (ヒロシマケン ヒロシマシアサミナミク オダニミノヤマチョウ)
川田 規加 (カワタ ミカ : 女性) -> 広島県 尾道市 氷上町絹山 (ヒロシマケン オノミチシ ヒカミチョウキヌヤマ)
```

- `--unique` (`-u`) オプションを付けると姓名＋住所名の組み合わせで重複するものを排除する
- `--template-path` (`-t`) オプションでテンプレートファイルを指定し，出力を任意に整形することが可能

```
$ cat testdata/template.txt 
| 名前 | 名前カタカナ | 性別 | メールアドレス | 住所 | 住所カタカナ |
| --- | --- | :---: | --- | --- | --- |{{ range . }}
| {{ .Gimei.FullName.Name }} | {{ (.Gimei.FullNameWithSep " ").Katakana }} | {{ .Gimei.Gender }} | {{ .Gimei.Email }} | {{ .Address.FullName.Name }} | {{ (.Address.FullNameWithSep " ").Katakana }} |{{ end }}

$ gimei-cli -r 10 -t testdata/template.txt
| 名前 | 名前カタカナ | 性別 | メールアドレス | 住所 | 住所カタカナ |
| --- | --- | :---: | --- | --- | --- |
| 小谷駿太 | コタニ シュンタ | 男性 | s.kotani@example.com | 青森県駿東郡小山町宍道町白石 | アオモリケン スントウグンオヤマチョウ シンジチョウハクイシ |
| 野口菜瑠 | ノグチ ナル | 女性 | n.noguchi@example.com | 岡山県稲敷郡美浦村覚王寺 | オカヤマケン イナシキグンミホムラ カクオウジ |
| 堀江凌太 | ホリエ リョウタ | 男性 | r.horie@example.com | 静岡県三沢市俵柳 | シズオカケン ミサワシ タワラヤナギ |
| 青柳縁蓮 | アオヤギ エレン | 女性 | e.aoyagi@example.com | 長野県西牟婁郡白浜町西方町金井 | ナガノケン ニシムログンシラハマチョウ ニシカタマチカナイ |
| 藤岡尚登 | フジオカ ナオト | 男性 | n.fujioka@example.com | 滋賀県青森市中畑 | シガケン アオモリシ ナカハタ |
| 小川泰誠 | オガワ タイセイ | 男性 | t.ogawa@example.com | 兵庫県日高郡日高川町酪陽 | ヒョウゴケン ヒダカグンヒダカガワチョウ ラクヨウ |
| 成田文夫 | ナリタ フミオ | 男性 | f.narita@example.com | 沖縄県厚岸郡浜中町新開町 | オキナワケン アッケシグンハマナカチョウ シンカイチョウ |
| 岩瀬恋春 | イワセ コハル | 女性 | k.iwase@example.com | 宮崎県札幌市豊平区上野新 | ミヤザキケン サッポロシトヨヒラク ウワノシン |
| 馬場椎苗 | ババ シイナ | 女性 | s.baba@example.com | 愛知県甘楽郡下仁田町大淀南 | アイチケン カンラグンシモニタマチ オオヨドミナミ |
| 松井善和 | マツイ ヨシカズ | 男性 | y.matsui@example.com | 和歌山県常滑市坪山 | ワカヤマケン トコナメシ ツボヤマ |
```

### 姓名の生成

```
$ gimei-cli gimei -h
姓名の生成

Usage:
  gimei-cli gimei [flags]

Aliases:
  gimei, name, g, n

Flags:
      --crypt-rand             crypt/rand 乱数生成器を使う
  -g, --gender                 「その他」の性別を含める
  -h, --help                   help for gimei
  -j, --json                   JSON形式で表示
  -r, --repeat int             試行回数 (default 1)
  -t, --template-path string   テンプレートファイルへのパス
  -u, --unique                 姓名の重複を避ける

Global Flags:
      --debug   for debug

$ gimei-cli gimei
多田 爽太 (タダ ソウタ : 男性)

$ gimei-cli gimei --json | jq .
[
  {
    "Gender": 1,
    "FirstName": {
      "Name": "有星",
      "Katakana": "アリセ",
      "Hiragana": "ありせ",
      "Roman": "arise"
    },
    "LastName": {
      "Name": "菅原",
      "Katakana": "スガハラ",
      "Hiragana": "すがはら",
      "Roman": "sugahara"
    }
  }
]
```

- 性別は男性・女性が概ね半々になるよう調整している。また `--gender` (`-g`) オプションを付けると 1/16 の確率で「その他」の性が出現する
- `--unique` (`-u`) オプションを付けると重複する姓名（読み仮名は無視）を排除する
- `--template-path` (`-t`) オプションでテンプレートファイルを指定し，出力を任意に整形することが可能

```
$ cat testdata/gimei-template.txt 
| 名前 | カタカナ | 性別 | メールアドレス |
| --- | --- | :---: | --- |{{ range . }}
| {{ .FullName.Name }} | {{ (.FullNameWithSep " ").Katakana }} | {{ .Gender }} | {{ .Email }} |{{ end }}

$ gimei-cli gimei -g -r 10 -t testdata/gimei-template.txt 
| 名前 | カタカナ | 性別 | メールアドレス |
| --- | --- | :---: | --- |
| 竹内秋甫 | タケウチ アキホ | 女性 | a.takeuchi@example.com |
| 清水学 | シミズ マナブ | 男性 | m.shimizu@example.com |
| 松原優佑 | マツバラ ユウスケ | 男性 | y.matsubara@example.com |
| 大久保彩郁 | オオクボ アヤカ | 女性 | a.ookubo@example.com |
| 三好慎之介 | ミヨシ シンノスケ | 男性 | s.miyoshi@example.com |
| 須田咲枝 | スダ サエ | 女性 | s.suda@example.com |
| 滝沢祐作 | タキザワ ユウサク | 男性 | y.takizawa@example.com |
| 寺田彰 | テラダ アキラ | その他 | a.terada@example.com |
| 関口昌則 | セキグチ マサノリ | 男性 | m.sekiguchi@example.com |
| 中本知保 | ナカモト チホ | 女性 | c.nakamoto@example.com |
```

### 住所名の生成

```
$ gimei-cli address -h
住所名の生成

Usage:
  gimei-cli address [flags]

Aliases:
  address, addr, a

Flags:
      --crypt-rand             crypt/rand 乱数生成器を使う
  -h, --help                   help for address
  -j, --json                   JSON形式で表示
  -n, --name string            絞り込み市区町村名（一部でも可）
  -r, --repeat int             試行回数 (default 1)
  -t, --template-path string   テンプレートファイルへのパス
  -u, --unique                 住所名の重複を避ける

Global Flags:
      --debug   for debug

$ gimei-cli address
佐賀県 新座市 原町田 (サガケン ニイザシ ハラマチダ)

$ gimei-cli address --json | jq .
[
  {
    "Prefecture": {
      "Name": "山口県",
      "Katakana": "ヤマグチケン",
      "Hiragana": "やまぐちけん",
      "Roman": "yamaguchiken"
    },
    "City": {
      "Name": "西諸県郡高原町",
      "Katakana": "ニシモロカタグンタカハルチョウ",
      "Hiragana": "にしもろかたぐんたかはるちょう",
      "Roman": "nishimorokataguntakaharuchou"
    },
    "Town": {
      "Name": "岩崎",
      "Katakana": "イワサキ",
      "Hiragana": "いわさき",
      "Roman": "iwasaki"
    }
  }
]
```

- `--name` (`-n`) オプションでキーワードを指定するとキーワードを含む市区町村で絞り込むことができる

```
$ gimei-cli address -r 10 -n 広島
広島県 大竹市 大淀南 (ヒロシマケン オオタケシ オオヨドミナミ)
広島県 竹原市 本渡町本渡 (ヒロシマケン タケハラシ ホンドマチホンド)
広島県 府中市 宇筒原 (ヒロシマケン フチュウシ ウトウバラ)
北海道 北広島市 宮丸 (ホッカイドウ キタヒロシマシ ミヤマル)
広島県 府中市 波積町本郷 (ヒロシマケン フチュウシ ハヅミチョウホンゴウ)
広島県 尾道市 加子母 (ヒロシマケン オノミチシ カシモ)
広島県 廿日市市 高坂 (ヒロシマケン ハツカイチシ タカサカ)
広島県 庄原市 吉川区泉 (ヒロシマケン ショウバラシ ヨシカワクイズミ)
広島県 広島市安佐北区 吉田町 (ヒロシマケン ヒロシマシアサキタク ヨシダマチ)
広島県 三次市 新宮町市野保 (ヒロシマケン ミヨシシ シングウチョウイチノホ)
```

- `--unique` (`-u`) オプションを付けると重複する住所名を排除する
- `--template-path` (`-t`) オプションでテンプレートファイルを指定し，出力を任意に整形することが可能

```
$ cat testdata/address-template.txt 
| 都道府県 | 市区町村 | 住所 | カタカナ |
| --- | --- | --- | --- |{{ range . }}
| {{ .Prefecture.Name }} | {{ .City.Name }} | {{ .Town.Name }} | {{ (.FullNameWithSep " ").Katakana }} |{{ end }}

$ gimei-cli address -r 10 -t testdata/address-template.txt 
| 都道府県 | 市区町村 | 住所 | カタカナ |
| --- | --- | --- | --- |
| 長崎県 | 碧南市 | 南崩山町 | ナガサキケン ヘキナンシ ミナミクエヤママチ |
| 佐賀県 | 伊都郡九度山町 | 大瀬戸町雪浦久良木郷 | サガケン イトグンクドヤマチョウ オオセトチョウユキノウラキュウラギゴウ |
| 鳥取県 | 韮崎市 | 宝町 | トットリケン ニラサキシ タカラマチ |
| 鳥取県 | 島牧郡島牧村 | 赤橋 | トットリケン シママキグンシママキムラ アカイバシ |
| 山梨県 | 神戸市北区 | 南新町 | ヤマナシケン コウベシキタク ミナミシンマチ |
| 山口県 | 東蒲原郡阿賀町 | 猫田 | ヤマグチケン ヒガシカンバラグンアガマチ ネコダ |
| 徳島県 | 伊万里市 | 朽木平良 | トクシマケン イマリシ クツキヘラ |
| 奈良県 | 加賀郡吉備中央町 | 亀崎 | ナラケン カガグンキビチュウオウチョウ カメザキ |
| 福井県 | 横浜市青葉区 | 長府新松原町 | フクイケン ヨコハマシアオバク チョウフシンマツバラチョウ |
| 香川県 | 能美市 | 美星町烏頭 | カガワケン ノミシ ビセイチョウウトウ |
```

### 都道府県の表示

```
$ gimei-cli prefectures -h
都道府県JISコードおよび都道府県名の表示

Usage:
  gimei-cli prefectures [flags]

Aliases:
  prefectures, prefs, pref, p

Flags:
  -h, --help          help for prefectures
  -j, --json          JSON形式で表示
  -n, --name string   都道府県名（一部でも可）

Global Flags:
      --debug   for debug

$ gimei-cli prefectures
01: 北海道 (ホッカイドウ)
02: 青森県 (アオモリケン)

...

46: 鹿児島県 (カゴシマケン)
47: 沖縄県 (オキナワケン)
```

- `--name` (`-n`) オプションでキーワードを指定するとキーワードを含む都道府県名で絞り込むことができる

```
$ gimei-cli prefectures -n 島
07: 福島県 (フクシマケン)
32: 島根県 (シマネケン)
34: 広島県 (ヒロシマケン)
36: 徳島県 (トクシマケン)
46: 鹿児島県 (カゴシマケン)

$ gimei-cli prefectures -n 島 --json | jq .
[
  {
    "Code": "07",
    "Name": "福島県",
    "Katakana": "フクシマケン",
    "Hiragana": "ふくしまけん",
    "Roman": "fukushimaken"
  },
  {
    "Code": "32",
    "Name": "島根県",
    "Katakana": "シマネケン",
    "Hiragana": "しまねけん",
    "Roman": "shimaneken"
  },
  {
    "Code": "34",
    "Name": "広島県",
    "Katakana": "ヒロシマケン",
    "Hiragana": "ひろしまけん",
    "Roman": "hiroshimaken"
  },
  {
    "Code": "36",
    "Name": "徳島県",
    "Katakana": "トクシマケン",
    "Hiragana": "とくしまけん",
    "Roman": "tokushimaken"
  },
  {
    "Code": "46",
    "Name": "鹿児島県",
    "Katakana": "カゴシマケン",
    "Hiragana": "かごしまけん",
    "Roman": "kagoshimaken"
  }
]
```

### 市区町村の表示

```
$ gimei-cli cities -h
市区町村JISコードおよび市区町村名（都道府県＋市区町村）の表示

Usage:
  gimei-cli cities [flags]

Aliases:
  cities, city, c

Flags:
  -h, --help          help for cities
  -j, --json          JSON形式で表示
  -n, --name string   市区町村名（一部でも可）

Global Flags:
      --debug   for debug

$ gimei-cli cities
01100: 北海道札幌市 (ホッカイドウサッポロシ)
01202: 北海道函館市 (ホッカイドウハコダテシ)

...

47381: 沖縄県竹富町 (オキナワケンタケトミチョウ)
47382: 沖縄県与那国町 (オキナワケンヨナグニチョウ)
```

- JIS都道府県コードおよびJIS市区町村コードは総務省の「[全国地方公共団体コード](https://www.soumu.go.jp/denshijiti/code.html "総務省｜地方行政のデジタル化｜全国地方公共団体コード")」で公開されているデータを利用している
- `--name` (`-n`) オプションでキーワードを指定するとキーワードを含む市区町村名で絞り込むことができる

```
$ gimei-cli cities -n 広島市
01234: 北海道北広島市 (ホッカイドウキタヒロシマシ)
34100: 広島県広島市 (ヒロシマケンヒロシマシ)
34212: 広島県東広島市 (ヒロシマケンヒガシヒロシマシ)

$ gimei-cli cities -n 広島市 --json | jq .
[
  {
    "Code": "01234",
    "Pref": {
      "Code": "01",
      "Name": "北海道",
      "Katakana": "ホッカイドウ",
      "Hiragana": "ほっかいどう",
      "Roman": "hokkaidou"
    },
    "Name": "北広島市",
    "Katakana": "キタヒロシマシ",
    "Hiragana": "きたひろしまし",
    "Roman": "kitahiroshimashi"
  },
  {
    "Code": "34100",
    "Pref": {
      "Code": "34",
      "Name": "広島県",
      "Katakana": "ヒロシマケン",
      "Hiragana": "ひろしまけん",
      "Roman": "hiroshimaken"
    },
    "Name": "広島市",
    "Katakana": "ヒロシマシ",
    "Hiragana": "ひろしまし",
    "Roman": "hiroshimashi"
  },
  {
    "Code": "34212",
    "Pref": {
      "Code": "34",
      "Name": "広島県",
      "Katakana": "ヒロシマケン",
      "Hiragana": "ひろしまけん",
      "Roman": "hiroshimaken"
    },
    "Name": "東広島市",
    "Katakana": "ヒガシヒロシマシ",
    "Hiragana": "ひがしひろしまし",
    "Roman": "higashihiroshimashi"
  }
]
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

[gimei-cli]: https://github.com/spiegel-im-spiegel/gimei-cli "spiegel-im-spiegel/gimei-cli: 姓名・住所データ生成ツール"
