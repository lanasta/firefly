// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apiserver

import (
	"net/http"

	"github.com/hyperledger/firefly/internal/batch"
	"github.com/hyperledger/firefly/internal/coremsgs"
	"github.com/hyperledger/firefly/internal/oapispec"
)

var getStatusBatchManager = &oapispec.Route{
	Name:            "getStatusBatchManager",
	Path:            "status/batchmanager",
	Method:          http.MethodGet,
	PathParams:      nil,
	QueryParams:     nil,
	FilterFactory:   nil,
	Description:     coremsgs.APIEndpointsGetStatusBatchManager,
	JSONInputValue:  nil,
	JSONOutputValue: func() interface{} { return &batch.ManagerStatus{} },
	JSONOutputCodes: []int{http.StatusOK},
	JSONHandler: func(r *oapispec.APIRequest) (output interface{}, err error) {
		output = getOr(r.Ctx).BatchManager().Status()
		return output, nil
	},
}