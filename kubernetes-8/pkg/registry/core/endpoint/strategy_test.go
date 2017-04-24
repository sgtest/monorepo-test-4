/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package endpoint

import (
	"testing"

	"github.com/sourcegraph/monorepo-test-1/kubernetes-8/pkg/api"
	apitesting "github.com/sourcegraph/monorepo-test-1/kubernetes-8/pkg/api/testing"
)

func TestSelectableFieldLabelConversions(t *testing.T) {
	apitesting.TestSelectableFieldLabelConversionsOfKind(t,
		api.Registry.GroupOrDie(api.GroupName).GroupVersion.String(),
		"Endpoints",
		EndpointsToSelectableFields(&api.Endpoints{}),
		nil,
	)
}