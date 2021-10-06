// Copyright 2021 Praetorian Security, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package istiod

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsIstiod(t *testing.T) {
	type testcase struct {
		pod      v1.Pod
		expected bool
	}

	testcases := []testcase{
		{
			pod: v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app":   "istiod",
						"istio": "pilot",
					},
				},
			},
			expected: true,
		},
		{
			pod: v1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"k8s": "pilot",
					},
				},
			},
			expected: false,
		},
	}
	for i, tc := range testcases {
		if v := isIstiod(tc.pod); v != tc.expected {
			t.Errorf("[%d] got %t, expected %t", i, v, tc.expected)
		}
	}
}
