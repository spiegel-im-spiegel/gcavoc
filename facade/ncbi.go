package facade

import (
	"context"
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gcavoc/api"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newNCBICmd returns cobra.Command instance for show sub-command
func newNCBICmd(ui *rwi.RWI) *cobra.Command {
	ncbiCmd := &cobra.Command{
		Use:     "ncbi [flags] <applied crop name>",
		Aliases: []string{"n"},
		Short:   "Output NCBI ID from applied crop name",
		Long:    "Output scientific name and NCBI ID from applied crop name.",
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
			nid, err := api.AppliedCropToNCBI(context.Background(), term)
			if err != nil {
				return debugPrint(ui, errs.Wrap(err, errs.WithContext("term", term)))
			}
			return debugPrint(ui, ui.Outputln(nid))
		},
	}
	ncbiCmd.Flags().BoolP("katakana", "k", false, "convert search term to katakana")
	ncbiCmd.Flags().BoolP("synonym", "s", false, "input parameter as a CVO synonym")

	return ncbiCmd
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
