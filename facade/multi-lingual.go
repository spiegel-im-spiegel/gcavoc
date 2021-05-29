package facade

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gcavoc/api"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newSciCmd returns cobra.Command instance for show sub-command
func newMultiCmd(ui *rwi.RWI) *cobra.Command {
	multiCmd := &cobra.Command{
		Use:     "multi-lingual [flags] <CVO synonym>",
		Aliases: []string{"multi", "m"},
		Short:   "Output multi lingual crop names from CVO synonym",
		Long:    "Output multi lingual crop names from CVO synonym.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			katakanaFlag, err := cmd.Flags().GetBool("katakana")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --katakana option", errs.WithCause(err)))
			}

			//Run command
			if len(args) != 1 {
				return debugPrint(ui, errs.Wrap(os.ErrInvalid))
			}
			term := args[0]
			if katakanaFlag {
				term = api.ConvertKatakana(term)
			}
			mn, err := api.CVOSynonymToMultilingualCropName(context.Background(), term)
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("term", term)))
			}
			return debugPrint(ui, ui.Outputln(mn))
		},
	}
	multiCmd.Flags().BoolP("katakana", "k", false, "convert search term to katakana")

	return multiCmd
}

/* Copyright 2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
