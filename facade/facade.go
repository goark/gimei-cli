package facade

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"runtime"

	"github.com/goark/errs"
	"github.com/goark/gimei-cli/all"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/goark/mt/mt19937"
	"github.com/goark/mt/secure"
	"github.com/spf13/cobra"
)

var (
	//Name is applicatin name
	Name = "gimei-cli"
	//Version is version for applicatin
	Version = "developer version"
)
var (
	debugFlag bool //debug flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "姓名・住所データ生成ツール",
		Long:  "姓名・住所データ生成ツール",
		RunE: func(cmd *cobra.Command, args []string) error {
			// options
			jsonFlag, err := cmd.Flags().GetBool("json")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --json option", errs.WithCause(err)))
			}
			uniqueFlag, err := cmd.Flags().GetBool("unique")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --unique option", errs.WithCause(err)))
			}
			genderFlag, err := cmd.Flags().GetBool("gender")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --gender option", errs.WithCause(err)))
			}
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --name option", errs.WithCause(err)))
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

			// generate gimei and address
			gen := all.New().WithGenderOther(genderFlag).WithUnique(uniqueFlag).WithFilters(name).WithRandSource(src)
			list := []*all.Info{}
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
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)            //arguments of command-line
	rootCmd.SetIn(ui.Reader())       //Stdin
	rootCmd.SetOut(ui.ErrorWriter()) //Stdout -> Stderr
	rootCmd.SetErr(ui.ErrorWriter()) //Stderr
	rootCmd.AddCommand(
		newVersionCmd(ui),
		newPrefCmd(ui),
		newCityCmd(ui),
		newNameCmd(ui),
		newAddressCmd(ui),
	)

	//options
	rootCmd.Flags().BoolP("json", "j", false, "JSON形式で表示")
	rootCmd.Flags().BoolP("unique", "u", false, "姓名＋住所名の重複を避ける")
	rootCmd.Flags().StringP("name", "n", "", "絞り込み市区町村名（一部でも可）")
	rootCmd.Flags().BoolP("gender", "g", false, "「その他」の性別を含める")
	rootCmd.Flags().BoolP("crypt-rand", "", false, "crypt/rand 乱数生成器を使う")
	rootCmd.Flags().IntP("repeat", "r", 1, "試行回数")
	rootCmd.Flags().StringP("template-path", "t", "", "テンプレートファイルへのパス")
	//global options
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")

	return rootCmd
}

func debugPrint(ui *rwi.RWI, err error) error {
	if debugFlag && err != nil {
		fmt.Fprintf(ui.Writer(), "%+v\n", err)
	}
	return err
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
