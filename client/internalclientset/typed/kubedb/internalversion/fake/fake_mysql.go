/*
Copyright 2017 The KubeDB Authors.

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

package fake

import (
	kubedb "github.com/k8sdb/apimachinery/apis/kubedb"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMySQLs implements MySQLInterface
type FakeMySQLs struct {
	Fake *FakeKubedb
	ns   string
}

var mysqlsResource = schema.GroupVersionResource{Group: "kubedb.com", Version: "", Resource: "mysqls"}

var mysqlsKind = schema.GroupVersionKind{Group: "kubedb.com", Version: "", Kind: "MySQL"}

// Get takes name of the mySQL, and returns the corresponding mySQL object, and an error if there is any.
func (c *FakeMySQLs) Get(name string, options v1.GetOptions) (result *kubedb.MySQL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlsResource, c.ns, name), &kubedb.MySQL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubedb.MySQL), err
}

// List takes label and field selectors, and returns the list of MySQLs that match those selectors.
func (c *FakeMySQLs) List(opts v1.ListOptions) (result *kubedb.MySQLList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlsResource, mysqlsKind, c.ns, opts), &kubedb.MySQLList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubedb.MySQLList{}
	for _, item := range obj.(*kubedb.MySQLList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mySQLs.
func (c *FakeMySQLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlsResource, c.ns, opts))

}

// Create takes the representation of a mySQL and creates it.  Returns the server's representation of the mySQL, and an error, if there is any.
func (c *FakeMySQLs) Create(mySQL *kubedb.MySQL) (result *kubedb.MySQL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlsResource, c.ns, mySQL), &kubedb.MySQL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubedb.MySQL), err
}

// Update takes the representation of a mySQL and updates it. Returns the server's representation of the mySQL, and an error, if there is any.
func (c *FakeMySQLs) Update(mySQL *kubedb.MySQL) (result *kubedb.MySQL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlsResource, c.ns, mySQL), &kubedb.MySQL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubedb.MySQL), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMySQLs) UpdateStatus(mySQL *kubedb.MySQL) (*kubedb.MySQL, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mysqlsResource, "status", c.ns, mySQL), &kubedb.MySQL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubedb.MySQL), err
}

// Delete takes name of the mySQL and deletes it. Returns an error if one occurs.
func (c *FakeMySQLs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlsResource, c.ns, name), &kubedb.MySQL{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMySQLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &kubedb.MySQLList{})
	return err
}

// Patch applies the patch and returns the patched mySQL.
func (c *FakeMySQLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kubedb.MySQL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlsResource, c.ns, name, data, subresources...), &kubedb.MySQL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubedb.MySQL), err
}