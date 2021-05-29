package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const PathCVOSynonymToScientificName = "/cvo/api/CVO_ToScientificName.php"

//ScientificName is the structure of scientific name info.
type ScientificName struct {
	Name string `json:"scientific_name"`
	En   string `json:"en_name"`
}

func (sn *ScientificName) String() string {
	if sn == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(sn); err != nil {
		return ""
	}
	return buf.String()
}

//CVOSynonymToScientificName function returns scientific name info from CVO synonym
func CVOSynonymToScientificName(ctx context.Context, term string) (*ScientificName, error) {
	params := url.Values{}
	params.Set("term", term)
	resp, err := fetch.New().Get(newServer().withPath(PathCVOSynonymToScientificName).withQuery(params).URL, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("term", term))
	}
	defer resp.Close()

	var sn ScientificName
	if err := json.NewDecoder(resp.Body()).Decode(&sn); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("term", term))
	}
	return &sn, nil
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
