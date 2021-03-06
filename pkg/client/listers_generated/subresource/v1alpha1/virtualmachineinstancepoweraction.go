/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubrid/apiserver/pkg/apis/subresource/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualMachineInstancePoweractionLister helps list VirtualMachineInstancePoweractions.
type VirtualMachineInstancePoweractionLister interface {
	// List lists all VirtualMachineInstancePoweractions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineInstancePoweraction, err error)
	// VirtualMachineInstancePoweractions returns an object that can list and get VirtualMachineInstancePoweractions.
	VirtualMachineInstancePoweractions(namespace string) VirtualMachineInstancePoweractionNamespaceLister
	VirtualMachineInstancePoweractionListerExpansion
}

// virtualMachineInstancePoweractionLister implements the VirtualMachineInstancePoweractionLister interface.
type virtualMachineInstancePoweractionLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineInstancePoweractionLister returns a new VirtualMachineInstancePoweractionLister.
func NewVirtualMachineInstancePoweractionLister(indexer cache.Indexer) VirtualMachineInstancePoweractionLister {
	return &virtualMachineInstancePoweractionLister{indexer: indexer}
}

// List lists all VirtualMachineInstancePoweractions in the indexer.
func (s *virtualMachineInstancePoweractionLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineInstancePoweraction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineInstancePoweraction))
	})
	return ret, err
}

// VirtualMachineInstancePoweractions returns an object that can list and get VirtualMachineInstancePoweractions.
func (s *virtualMachineInstancePoweractionLister) VirtualMachineInstancePoweractions(namespace string) VirtualMachineInstancePoweractionNamespaceLister {
	return virtualMachineInstancePoweractionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualMachineInstancePoweractionNamespaceLister helps list and get VirtualMachineInstancePoweractions.
type VirtualMachineInstancePoweractionNamespaceLister interface {
	// List lists all VirtualMachineInstancePoweractions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineInstancePoweraction, err error)
	// Get retrieves the VirtualMachineInstancePoweraction from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VirtualMachineInstancePoweraction, error)
	VirtualMachineInstancePoweractionNamespaceListerExpansion
}

// virtualMachineInstancePoweractionNamespaceLister implements the VirtualMachineInstancePoweractionNamespaceLister
// interface.
type virtualMachineInstancePoweractionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualMachineInstancePoweractions in the indexer for a given namespace.
func (s virtualMachineInstancePoweractionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineInstancePoweraction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineInstancePoweraction))
	})
	return ret, err
}

// Get retrieves the VirtualMachineInstancePoweraction from the indexer for a given namespace and name.
func (s virtualMachineInstancePoweractionNamespaceLister) Get(name string) (*v1alpha1.VirtualMachineInstancePoweraction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualmachineinstancepoweraction"), name)
	}
	return obj.(*v1alpha1.VirtualMachineInstancePoweraction), nil
}
