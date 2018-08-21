/*
Copyright 2016 The Kubernetes Authors.

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

package group

import (
	"k8s.io/apimachinery/pkg/util/validation/field"
	api "k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/pkg/apis/policy"
)

type nonRoot struct{}

var _ GroupStrategy = &nonRoot{}

func NewRunAsNonRoot(options *policy.RunAsGroupStrategyOptions) (GroupStrategy, error) {
	return &nonRoot{}, nil
}

// Generate creates the uid based on policy rules.  This strategy does return a GID.  It assumes
// that the user will specify a GID or the container image specifies a GID.
func (s *nonRoot) Generate(pod *api.Pod) ([]int64, error) {
	return nil, nil
}

// Generate a single value to be applied.  This is used for FSGroup.  This strategy returns nil.
func (s *nonRoot) GenerateSingle(_ *api.Pod) (*int64, error) {
	return nil, nil
}

// Validate ensures that the specified values fall within the range of the strategy.  Validation
// of this will pass if either the GID is not set, assuming that the image will provided the GID
// or if the GID is set it is not root.  Validation will fail if RunAsNonRoot is set to false.
// In order to work properly this assumes that the kubelet performs a final check on runAsGroup
// or the image GID when runAsGroup is nil.
func (s *nonRoot) Validate(scPath *field.Path, _ *api.Pod, runAsNonRoot *bool, groups []int64) field.ErrorList {
	allErrs := field.ErrorList{}

	if runAsNonRoot == nil && len(groups) == 0 {
		allErrs = append(allErrs, field.Required(scPath.Child("runAsNonRoot"), "must be true"))
		return allErrs
	}
	if runAsNonRoot != nil && *runAsNonRoot == false {
		allErrs = append(allErrs, field.Invalid(scPath.Child("runAsNonRoot"), *runAsNonRoot, "must be true"))
		return allErrs
	}
	for _, group := range groups {
		if group == 0 {
			allErrs = append(allErrs, field.Invalid(scPath.Child("runAsGroup"), group, "running with the root GID is forbidden"))
			return allErrs
		}

	}
	return allErrs
}
