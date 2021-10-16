package facade

import (
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gimei-cli/pref"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newCityCmd returns cobra.Command instance for show sub-command
func newCityCmd(ui *rwi.RWI) *cobra.Command {
	cityCmd := &cobra.Command{
		Use:     "cities",
		Aliases: []string{"city", "c"},
		Short:   "市区町村の表示",
		Long:    "市区町村JISコードおよび市区町村名（都道府県＋市区町村）の表示",
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
			cities := pref.CityName(name)
			if jsonFlag {
				if err := json.NewEncoder(ui.Writer()).Encode(cities); err != nil {
					return debugPrint(ui, err)
				}
			} else {
				for _, p := range cities {
					_ = ui.Outputln(p)
				}
			}
			return nil
		},
	}
	cityCmd.Flags().BoolP("json", "j", false, "JSON形式で表示")
	cityCmd.Flags().StringP("name", "n", "", "市区町村名（一部でも可）")

	return cityCmd
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
