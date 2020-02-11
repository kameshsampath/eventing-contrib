/*
Copyright 2020 The Knative Authors

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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "knative.dev/eventing/pkg/apis/messaging/v1alpha1"
)

// FakeSequences implements SequenceInterface
type FakeSequences struct {
	Fake *FakeMessagingV1alpha1
	ns   string
}

var sequencesResource = schema.GroupVersionResource{Group: "messaging.knative.dev", Version: "v1alpha1", Resource: "sequences"}

var sequencesKind = schema.GroupVersionKind{Group: "messaging.knative.dev", Version: "v1alpha1", Kind: "Sequence"}

// Get takes name of the sequence, and returns the corresponding sequence object, and an error if there is any.
func (c *FakeSequences) Get(name string, options v1.GetOptions) (result *v1alpha1.Sequence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sequencesResource, c.ns, name), &v1alpha1.Sequence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sequence), err
}

// List takes label and field selectors, and returns the list of Sequences that match those selectors.
func (c *FakeSequences) List(opts v1.ListOptions) (result *v1alpha1.SequenceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sequencesResource, sequencesKind, c.ns, opts), &v1alpha1.SequenceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SequenceList{ListMeta: obj.(*v1alpha1.SequenceList).ListMeta}
	for _, item := range obj.(*v1alpha1.SequenceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sequences.
func (c *FakeSequences) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sequencesResource, c.ns, opts))

}

// Create takes the representation of a sequence and creates it.  Returns the server's representation of the sequence, and an error, if there is any.
func (c *FakeSequences) Create(sequence *v1alpha1.Sequence) (result *v1alpha1.Sequence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sequencesResource, c.ns, sequence), &v1alpha1.Sequence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sequence), err
}

// Update takes the representation of a sequence and updates it. Returns the server's representation of the sequence, and an error, if there is any.
func (c *FakeSequences) Update(sequence *v1alpha1.Sequence) (result *v1alpha1.Sequence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sequencesResource, c.ns, sequence), &v1alpha1.Sequence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sequence), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSequences) UpdateStatus(sequence *v1alpha1.Sequence) (*v1alpha1.Sequence, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sequencesResource, "status", c.ns, sequence), &v1alpha1.Sequence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sequence), err
}

// Delete takes name of the sequence and deletes it. Returns an error if one occurs.
func (c *FakeSequences) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sequencesResource, c.ns, name), &v1alpha1.Sequence{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSequences) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sequencesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SequenceList{})
	return err
}

// Patch applies the patch and returns the patched sequence.
func (c *FakeSequences) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Sequence, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sequencesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Sequence{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sequence), err
}