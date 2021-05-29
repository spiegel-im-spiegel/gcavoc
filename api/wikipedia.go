package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const PathAppliedCropToWikipedia = "/cvo/api/CVO_TekiyounousakumotuToWIKIPEDIA.php"

//GroupName is the structure of scientific name and group info.
type Wikipedia struct {
	URL string `json:"WIKIPEDIA_url"`
}

func (w *Wikipedia) String() string {
	if w == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(w); err != nil {
		return ""
	}
	return buf.String()
}

//AppliedCropToWikipedia function returns Wikipedia URL from applied crop name
func AppliedCropToWikipedia(ctx context.Context, cropName string) (*Wikipedia, error) {
	params := url.Values{}
	params.Set("term", cropName)
	resp, err := fetch.New().Get(newServer().withPath(PathAppliedCropToWikipedia).withQuery(params).URL, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("cropName", cropName))
	}
	defer resp.Close()

	var w Wikipedia
	if err := json.NewDecoder(resp.Body()).Decode(&w); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("cropName", cropName))
	}
	return &w, nil
}

const prefixWikipediaJaURL = "https://ja.wikipedia.org/wiki/"

//FixWikipediaURL function returns fixed wikipedia URL.
func FixWikipediaURL(s string) string {
	if len(s) == 0 {
		return ""
	}
	p := strings.TrimPrefix(s, prefixWikipediaJaURL)
	return prefixWikipediaJaURL + url.PathEscape(p)
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
