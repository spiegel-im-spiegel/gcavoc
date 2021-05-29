package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const PathCVOSynonymToStandardTerm = "/cvo/api/CVO_SynonymToStandard.php"

//StandardTerm is the structure of standard term info.
type StandardTerm struct {
	Term string `json:"term"`
}

func (st *StandardTerm) String() string {
	if st == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(st); err != nil {
		return ""
	}
	return buf.String()
}

//CVOSynonymToStandardTerm function returns standard term from CVO synonym
func CVOSynonymToStandardTerm(ctx context.Context, term string) (*StandardTerm, error) {
	params := url.Values{}
	params.Set("term", term)
	resp, err := fetch.New().Get(newServer().withPath(PathCVOSynonymToStandardTerm).withQuery(params).URL, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("term", term))
	}
	defer resp.Close()

	var st StandardTerm
	if err := json.NewDecoder(resp.Body()).Decode(&st); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("term", term))
	}
	return &st, nil
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
