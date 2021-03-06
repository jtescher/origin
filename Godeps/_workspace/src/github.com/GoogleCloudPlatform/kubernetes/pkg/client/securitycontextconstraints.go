/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package client

import (
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/watch"
)

type SecurityContextConstraintsInterface interface {
	SecurityContextConstraints() SecurityContextConstraintInterface
}

type SecurityContextConstraintInterface interface {
	Get(name string) (result *api.SecurityContextConstraints, err error)
	Create(scc *api.SecurityContextConstraints) (*api.SecurityContextConstraints, error)
	List(label labels.Selector, field fields.Selector) (*api.SecurityContextConstraintsList, error)
	Delete(name string) error
	Update(*api.SecurityContextConstraints) (*api.SecurityContextConstraints, error)
	Watch(label labels.Selector, field fields.Selector, resourceVersion string) (watch.Interface, error)
}

// securityContextConstraints implements SecurityContextConstraintInterface
type securityContextConstraints struct {
	client *Client
}

// newSecurityContextConstraints returns a securityContextConstraints object.
func newSecurityContextConstraints(c *Client) *securityContextConstraints {
	return &securityContextConstraints{c}
}

func (s *securityContextConstraints) Create(scc *api.SecurityContextConstraints) (*api.SecurityContextConstraints, error) {
	result := &api.SecurityContextConstraints{}
	err := s.client.Post().
		Resource("securityContextConstraints").
		Body(scc).
		Do().
		Into(result)

	return result, err
}

// List returns a list of SecurityContextConstraints matching the selectors.
func (s *securityContextConstraints) List(label labels.Selector, field fields.Selector) (*api.SecurityContextConstraintsList, error) {
	result := &api.SecurityContextConstraintsList{}

	err := s.client.Get().
		Resource("securityContextConstraints").
		LabelsSelectorParam(label).
		FieldsSelectorParam(field).
		Do().
		Into(result)

	return result, err
}

// Get returns the given SecurityContextConstraints, or an error.
func (s *securityContextConstraints) Get(name string) (*api.SecurityContextConstraints, error) {
	result := &api.SecurityContextConstraints{}
	err := s.client.Get().
		Resource("securityContextConstraints").
		Name(name).
		Do().
		Into(result)

	return result, err
}

// Watch starts watching for SecurityContextConstraints matching the given selectors.
func (s *securityContextConstraints) Watch(label labels.Selector, field fields.Selector, resourceVersion string) (watch.Interface, error) {
	return s.client.Get().
		Prefix("watch").
		Resource("securityContextConstraints").
		Param("resourceVersion", resourceVersion).
		LabelsSelectorParam(label).
		FieldsSelectorParam(field).
		Watch()
}

func (s *securityContextConstraints) Delete(name string) error {
	return s.client.Delete().
		Resource("securityContextConstraints").
		Name(name).
		Do().
		Error()
}

func (s *securityContextConstraints) Update(scc *api.SecurityContextConstraints) (result *api.SecurityContextConstraints, err error) {
	result = &api.SecurityContextConstraints{}
	err = s.client.Put().
		Resource("securityContextConstraints").
		Name(scc.Name).
		Body(scc).
		Do().
		Into(result)

	return
}
