// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("auditbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}