/*
Copyright 2025.

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

package functional_test

import (
	. "github.com/onsi/ginkgo/v2" //revive:disable:dot-imports
	. "github.com/onsi/gomega"    //revive:disable:dot-imports
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("Instanceha Controller", func() {
	var instancehaName types.NamespacedName

	When("a default Instanceha gets created", func() {
		BeforeEach(func() {
			instanceha := CreateInstancehaConfig(namespace, GetDefaultInstancehaSpec())
			instancehaName.Name = instanceha.GetName()
			instancehaName.Namespace = instanceha.GetNamespace()
			DeferCleanup(th.DeleteInstance, instanceha)
		})

		It("should have created a Instanceha", func() {
			Eventually(func(_ Gomega) {
				GetInstanceha(instancehaName)
			}, timeout, interval).Should(Succeed())
		})
	})

})
