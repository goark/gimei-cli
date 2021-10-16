package facade

import (
	"encoding/json"
	"math/rand"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gimei-cli/gimei"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/mt/mt19937"
	"github.com/spiegel-im-spiegel/mt/secure"
)

//newNameCmd returns cobra.Command instance for show sub-command
func newNameCmd(ui *rwi.RWI) *cobra.Command {
	nameCmd := &cobra.Command{
		Use:     "gimei",
		Aliases: []string{"name", "g", "n"},
		Short:   "偽名の生成",
		Long:    "偽名の生成",
		RunE: func(cmd *cobra.Command, args []string) error {
			// options
			jsonFlag, err := cmd.Flags().GetBool("json")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --json option", errs.WithCause(err)))
			}
			genderFlag, err := cmd.Flags().GetBool("gender")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --gender option", errs.WithCause(err)))
			}
			uniqueFlag, err := cmd.Flags().GetBool("unique")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --unique option", errs.WithCause(err)))
			}
			secureFlag, err := cmd.Flags().GetBool("crypt-rand")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --crypt-rand option", errs.WithCause(err)))
			}
			var src rand.Source
			if secureFlag {
				src = secure.Source{}
			} else if debugFlag {
				src = mt19937.New(19650218)
			}
			path, err := cmd.Flags().GetString("template-path")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --template-path option", errs.WithCause(err)))
			}
			count, err := cmd.Flags().GetInt("repeat")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --repeat option", errs.WithCause(err)))
			}
			if count < 1 {
				count = 1
			}

			// generate gimei
			gen := gimei.New().WithGenderOther(genderFlag).WithUnique(uniqueFlag).WithRandSource(src)
			list := []*gimei.Gimei{}
			for i := 0; i < count; i++ {
				list = append(list, gen.New())
			}

			// output
			if jsonFlag {
				if err := json.NewEncoder(ui.Writer()).Encode(list); err != nil {
					return debugPrint(ui, err)
				}
			} else if len(path) > 0 {
				if err := outputWithTemplate(ui.Writer(), path, list); err != nil {
					return debugPrint(ui, err)
				}
			} else {
				for _, g := range list {
					_ = ui.Outputln(g)
				}
			}
			return nil
		},
	}
	nameCmd.Flags().BoolP("json", "j", false, "JSON形式で表示")
	nameCmd.Flags().BoolP("gender", "g", false, "「その他」の性別を含める")
	nameCmd.Flags().BoolP("unique", "u", false, "名前の重複を避ける")
	nameCmd.Flags().BoolP("crypt-rand", "", false, "crypt/rand 乱数生成器を使う")
	nameCmd.Flags().IntP("repeat", "r", 1, "試行回数")
	nameCmd.Flags().StringP("template-path", "t", "", "テンプレートファイルへのパス")

	return nameCmd
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
