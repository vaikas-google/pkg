/*
Copyright 2018 The Knative Authors

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
	v1alpha1 "github.com/knative/pkg/apis/duck/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SubscribableLister helps list Subscribables.
type SubscribableLister interface {
	// List lists all Subscribables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Subscribable, err error)
	// Subscribables returns an object that can list and get Subscribables.
	Subscribables(namespace string) SubscribableNamespaceLister
	SubscribableListerExpansion
}

// subscribableLister implements the SubscribableLister interface.
type subscribableLister struct {
	indexer cache.Indexer
}

// NewSubscribableLister returns a new SubscribableLister.
func NewSubscribableLister(indexer cache.Indexer) SubscribableLister {
	return &subscribableLister{indexer: indexer}
}

// List lists all Subscribables in the indexer.
func (s *subscribableLister) List(selector labels.Selector) (ret []*v1alpha1.Subscribable, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Subscribable))
	})
	return ret, err
}

// Subscribables returns an object that can list and get Subscribables.
func (s *subscribableLister) Subscribables(namespace string) SubscribableNamespaceLister {
	return subscribableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubscribableNamespaceLister helps list and get Subscribables.
type SubscribableNamespaceLister interface {
	// List lists all Subscribables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Subscribable, err error)
	// Get retrieves the Subscribable from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Subscribable, error)
	SubscribableNamespaceListerExpansion
}

// subscribableNamespaceLister implements the SubscribableNamespaceLister
// interface.
type subscribableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Subscribables in the indexer for a given namespace.
func (s subscribableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Subscribable, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Subscribable))
	})
	return ret, err
}

// Get retrieves the Subscribable from the indexer for a given namespace and name.
func (s subscribableNamespaceLister) Get(name string) (*v1alpha1.Subscribable, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subscribable"), name)
	}
	return obj.(*v1alpha1.Subscribable), nil
}
