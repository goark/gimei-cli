package facade

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	usage = []string{ //output message of version
		Name + " " + Version,
		"repository: https://github.com/spiegel-im-spiegel/gimei-cli",
	}
)

//newVersionCmd returns cobra.Command instance for show sub-command
func newVersionCmd(ui *rwi.RWI) *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "バージョン表示",
		Long:  Name + " のバージョン表示",
		RunE: func(cmd *cobra.Command, args []string) error {
			return ui.OutputErrln(strings.Join(usage, "\n"))
		},
	}

	return versionCmd
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
