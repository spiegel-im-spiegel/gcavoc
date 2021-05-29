package facade

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gcavoc/api"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newWikipediaCmd returns cobra.Command instance for show sub-command
func newWikipediaCmd(ui *rwi.RWI) *cobra.Command {
	wikipediaCmd := &cobra.Command{
		Use:     "wikipedia [flags] <applied crop name>",
		Aliases: []string{"w"},
		Short:   "Output Wikipedia URL from applied crop name",
		Long:    "Output Wikipedia URL from applied crop name.",
		RunE: func(cmd *cobra.Command, args []string) error {
			//Options
			katakanaFlag, err := cmd.Flags().GetBool("katakana")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --katakana option", errs.WithCause(err)))
			}
			synonymFlag, err := cmd.Flags().GetBool("synonym")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --synonym option", errs.WithCause(err)))
			}
			rawFlag, err := cmd.Flags().GetBool("raw")
			if err != nil {
				return debugPrint(ui, errs.New("Error in --raw option", errs.WithCause(err)))
			}

			//Run command
			if len(args) != 1 {
				return debugPrint(ui, errs.Wrap(os.ErrInvalid))
			}
			term := args[0]
			if katakanaFlag {
				term = api.ConvertKatakana(term)
			}
			if synonymFlag {
				st, err := api.CVOSynonymToStandardTerm(context.Background(), term)
				if err != nil {
					return debugPrint(ui, errs.Wrap(err, errs.WithContext("term", term)))
				}
				if len(st.Term) > 0 {
					term = st.Term
				}
			}
			wurl, err := api.AppliedCropToWikipedia(context.Background(), term)
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("term", term)))
			}
			if rawFlag {
				return debugPrint(ui, ui.Outputln(wurl))
			}
			return debugPrint(ui, ui.Outputln(api.FixWikipediaURL(wurl.URL)))
		},
	}
	wikipediaCmd.Flags().BoolP("katakana", "k", false, "convert search term to katakana")
	wikipediaCmd.Flags().BoolP("synonym", "s", false, "input parameter as a CVO synonym")
	wikipediaCmd.Flags().BoolP("raw", "", false, "output raw data (JSON format)")

	return wikipediaCmd
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
