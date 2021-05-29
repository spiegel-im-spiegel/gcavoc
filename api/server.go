package api

import "net/url"

const (
	defaultScheme = "http"
	defaultHost   = "cavoc.org"
)

type server struct {
	*url.URL
}

func newServer() *server {
	return &server{&url.URL{Scheme: defaultScheme, Host: defaultHost}}
}

func (s *server) withPath(path string) *server {
	if s == nil {
		s = newServer()
	}
	s.Path = path
	return s
}

func (s *server) withQuery(v url.Values) *server {
	if s == nil {
		s = newServer()
	}
	s.RawQuery = v.Encode()
	return s
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
