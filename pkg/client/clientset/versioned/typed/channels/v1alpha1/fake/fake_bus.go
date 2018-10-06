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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/knative/eventing/pkg/apis/channels/v1alpha1"
	system "github.com/knative/eventing/pkg/system"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuses implements BusInterface
type FakeBuses struct {
	Fake *FakeChannelsV1alpha1
	ns   string
}

var busesResource = schema.GroupVersionResource{Group: "channels.knative.dev", Version: "v1alpha1", Resource: "buses"}

var busesKind = schema.GroupVersionKind{Group: "channels.knative.dev", Version: "v1alpha1", Kind: system.KindBus}

// Get takes name of the bus, and returns the corresponding bus object, and an error if there is any.
func (c *FakeBuses) Get(name string, options v1.GetOptions) (result *v1alpha1.Bus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(busesResource, c.ns, name), &v1alpha1.Bus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bus), err
}

// List takes label and field selectors, and returns the list of Buses that match those selectors.
func (c *FakeBuses) List(opts v1.ListOptions) (result *v1alpha1.BusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(busesResource, busesKind, c.ns, opts), &v1alpha1.BusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BusList{ListMeta: obj.(*v1alpha1.BusList).ListMeta}
	for _, item := range obj.(*v1alpha1.BusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buses.
func (c *FakeBuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(busesResource, c.ns, opts))

}

// Create takes the representation of a bus and creates it.  Returns the server's representation of the bus, and an error, if there is any.
func (c *FakeBuses) Create(bus *v1alpha1.Bus) (result *v1alpha1.Bus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(busesResource, c.ns, bus), &v1alpha1.Bus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bus), err
}

// Update takes the representation of a bus and updates it. Returns the server's representation of the bus, and an error, if there is any.
func (c *FakeBuses) Update(bus *v1alpha1.Bus) (result *v1alpha1.Bus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(busesResource, c.ns, bus), &v1alpha1.Bus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bus), err
}

// Delete takes name of the bus and deletes it. Returns an error if one occurs.
func (c *FakeBuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(busesResource, c.ns, name), &v1alpha1.Bus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(busesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BusList{})
	return err
}

// Patch applies the patch and returns the patched bus.
func (c *FakeBuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Bus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(busesResource, c.ns, name, data, subresources...), &v1alpha1.Bus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bus), err
}
