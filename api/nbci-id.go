package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const PathAppliedCropToNBCI = "/cvo/api/CVO_TekiyounousakumotuToNCBI.php"

//GroupName is the structure of scientific name and group info.
type NBCI struct {
	ID   string `json:"NCBI_id"`
	Name string `json:"scientific_name"`
}

func (n *NBCI) String() string {
	if n == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(n); err != nil {
		return ""
	}
	return buf.String()
}

//AppliedCropToNBCI function returns scientific name and NBCI ID info from applied crop name
func AppliedCropToNBCI(ctx context.Context, cropName string) (*NBCI, error) {
	params := url.Values{}
	params.Set("term", cropName)
	resp, err := fetch.New().Get(newServer().withPath(PathAppliedCropToNBCI).withQuery(params).URL, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("cropName", cropName))
	}
	defer resp.Close()

	var sn NBCI
	if err := json.NewDecoder(resp.Body()).Decode(&sn); err != nil {
		return nil, errs.Wrap(err, errs.WithContext("cropName", cropName))
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
