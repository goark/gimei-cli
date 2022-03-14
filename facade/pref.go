package facade

import (
	"encoding/json"

	"github.com/goark/errs"
	"github.com/goark/gimei-cli/pref"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newPrefCmd returns cobra.Command instance for show sub-command
func newPrefCmd(ui *rwi.RWI) *cobra.Command {
	prefCmd := &cobra.Command{
		Use:     "prefectures",
		Aliases: []string{"prefs", "pref", "p"},
		Short:   "都道府県の表示",
		Long:    "都道府県JISコードおよび都道府県名の表示",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			jsonFlag, err := cmd.Flags().GetBool("json")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --json option", errs.WithCause(err)))
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --name option", errs.WithCause(err)))
			}
			prefs := pref.PrefName(name)
			if jsonFlag {
				if err := json.NewEncoder(ui.Writer()).Encode(prefs); err != nil {
					return debugPrint(ui, err)
				}
			} else {
				for _, p := range prefs {
					_ = ui.Outputln(p)
				}
			}
			return nil
		},
	}
	prefCmd.Flags().BoolP("json", "j", false, "JSON形式で表示")
	prefCmd.Flags().StringP("name", "n", "", "都道府県名（一部でも可）")

	return prefCmd
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
