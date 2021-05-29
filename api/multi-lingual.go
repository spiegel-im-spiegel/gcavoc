package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const PathCVOSynonymToMultilingualCropName = "/cvo/api/CVO_MultilingualCropName.php"

//MultilingualCropName is the structure of scientific name info.
type MultilingualCropName struct {
	Ja string `json:"ja"`
	En string `json:"en"`
	Zh string `json:"zh"`
	Ko string `json:"ko"`
}

func (mn *MultilingualCropName) String() string {
	if mn == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(mn); err != nil {
		return ""
	}
	return buf.String()
}

//CVOSynonymToMultilingualCropName function returns multi lingual crop names info from CVO synonym
func CVOSynonymToMultilingualCropName(ctx context.Context, term string) (*MultilingualCropName, error) {
	params := url.Values{}
	params.Set("term", term)
	resp, err := fetch.New().Get(newServer().withPath(PathCVOSynonymToMultilingualCropName).withQuery(params).URL, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("term", term))
	}
	defer resp.Close()

	var mn MultilingualCropName
	if err := json.NewDecoder(resp.Body()).Decode(&mn); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("term", term))
	}
	return &mn, nil
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
