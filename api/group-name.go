package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/url"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

const PathAppliedCropToGroupName = "/cvo/api/CVO_TekiyounousakumotuToCVO.php"

//GroupName is the structure of scientific name and group info.
type GroupName struct {
	Name  string `json:"scientific_name"`
	Group string `json:"group"`
}

func (gn *GroupName) String() string {
	if gn == nil {
		return ""
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(gn); err != nil {
		return ""
	}
	return buf.String()
}

//AppliedCropToGroupName function returns scientific name and group name info from applied crop name
func AppliedCropToGroupName(ctx context.Context, cropName string) (*GroupName, error) {
	params := url.Values{}
	params.Set("term", cropName)
	resp, err := fetch.New().Get(newServer().withPath(PathAppliedCropToGroupName).withQuery(params).URL, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("cropName", cropName))
	}
	defer resp.Close()

	var sn GroupName
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
