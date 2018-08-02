/*
Copyright 2018 The KubeDB Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEtcdVersions implements EtcdVersionInterface
type FakeEtcdVersions struct {
	Fake *FakeKubedbV1alpha1
}

var etcdversionsResource = schema.GroupVersionResource{Group: "kubedb.com", Version: "v1alpha1", Resource: "etcdversions"}

var etcdversionsKind = schema.GroupVersionKind{Group: "kubedb.com", Version: "v1alpha1", Kind: "EtcdVersion"}

// Get takes name of the etcdVersion, and returns the corresponding etcdVersion object, and an error if there is any.
func (c *FakeEtcdVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.EtcdVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(etcdversionsResource, name), &v1alpha1.EtcdVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdVersion), err
}

// List takes label and field selectors, and returns the list of EtcdVersions that match those selectors.
func (c *FakeEtcdVersions) List(opts v1.ListOptions) (result *v1alpha1.EtcdVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(etcdversionsResource, etcdversionsKind, opts), &v1alpha1.EtcdVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EtcdVersionList{ListMeta: obj.(*v1alpha1.EtcdVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.EtcdVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested etcdVersions.
func (c *FakeEtcdVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(etcdversionsResource, opts))
}

// Create takes the representation of a etcdVersion and creates it.  Returns the server's representation of the etcdVersion, and an error, if there is any.
func (c *FakeEtcdVersions) Create(etcdVersion *v1alpha1.EtcdVersion) (result *v1alpha1.EtcdVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(etcdversionsResource, etcdVersion), &v1alpha1.EtcdVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdVersion), err
}

// Update takes the representation of a etcdVersion and updates it. Returns the server's representation of the etcdVersion, and an error, if there is any.
func (c *FakeEtcdVersions) Update(etcdVersion *v1alpha1.EtcdVersion) (result *v1alpha1.EtcdVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(etcdversionsResource, etcdVersion), &v1alpha1.EtcdVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdVersion), err
}

// Delete takes name of the etcdVersion and deletes it. Returns an error if one occurs.
func (c *FakeEtcdVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(etcdversionsResource, name), &v1alpha1.EtcdVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEtcdVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(etcdversionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EtcdVersionList{})
	return err
}

// Patch applies the patch and returns the patched etcdVersion.
func (c *FakeEtcdVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EtcdVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(etcdversionsResource, name, data, subresources...), &v1alpha1.EtcdVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EtcdVersion), err
}
