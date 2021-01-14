
/*
Copyright YEAR The Kubernetes Authors.

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


package v1alpha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	. "github.com/kubrid/apiserver/pkg/apis/subresource/v1alpha1"
	. "github.com/kubrid/apiserver/pkg/client/clientset_generated/clientset/typed/subresource/v1alpha1"
)

var _ = Describe("VirtualMachineInstance", func() {
	var instance VirtualMachineInstance
	var expected VirtualMachineInstance
	var client VirtualMachineInstanceInterface

	BeforeEach(func() {
		instance = VirtualMachineInstance{}
		instance.Name = "instance-1"

		expected = instance
	})

	AfterEach(func() {
		client.Delete(context.TODO(), instance.Name, metav1.DeleteOptions{})
	})

	Describe("when sending a poweraction request", func() {
		It("should return success", func() {
			client = cs.SubresourceV1alpha1().Virtualmachineinstances("virtualmachineinstance-test-poweraction")
			_, err := client.Create(&instance)
			Expect(err).ShouldNot(HaveOccurred())

			poweraction := &VirtualMachineInstancePoweraction{}
			poweraction.Name = instance.Name
			restClient := cs.SubresourceV1alpha1().RESTClient()
			err = restClient.Post().Namespace("virtualmachineinstance-test-poweraction").
				Name(instance.Name).
				Resource("virtualmachineinstances").
				SubResource("poweraction").
				Body(poweraction).
				Do(context.TODO()).
				Error()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
