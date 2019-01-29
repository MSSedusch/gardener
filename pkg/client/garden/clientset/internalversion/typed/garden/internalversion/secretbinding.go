// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	garden "github.com/gardener/gardener/pkg/apis/garden"
	scheme "github.com/gardener/gardener/pkg/client/garden/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SecretBindingsGetter has a method to return a SecretBindingInterface.
// A group's client should implement this interface.
type SecretBindingsGetter interface {
	SecretBindings(namespace string) SecretBindingInterface
}

// SecretBindingInterface has methods to work with SecretBinding resources.
type SecretBindingInterface interface {
	Create(*garden.SecretBinding) (*garden.SecretBinding, error)
	Update(*garden.SecretBinding) (*garden.SecretBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*garden.SecretBinding, error)
	List(opts v1.ListOptions) (*garden.SecretBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.SecretBinding, err error)
	SecretBindingExpansion
}

// secretBindings implements SecretBindingInterface
type secretBindings struct {
	client rest.Interface
	ns     string
}

// newSecretBindings returns a SecretBindings
func newSecretBindings(c *GardenClient, namespace string) *secretBindings {
	return &secretBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secretBinding, and returns the corresponding secretBinding object, and an error if there is any.
func (c *secretBindings) Get(name string, options v1.GetOptions) (result *garden.SecretBinding, err error) {
	result = &garden.SecretBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecretBindings that match those selectors.
func (c *secretBindings) List(opts v1.ListOptions) (result *garden.SecretBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &garden.SecretBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secretBindings.
func (c *secretBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secretbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a secretBinding and creates it.  Returns the server's representation of the secretBinding, and an error, if there is any.
func (c *secretBindings) Create(secretBinding *garden.SecretBinding) (result *garden.SecretBinding, err error) {
	result = &garden.SecretBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secretbindings").
		Body(secretBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a secretBinding and updates it. Returns the server's representation of the secretBinding, and an error, if there is any.
func (c *secretBindings) Update(secretBinding *garden.SecretBinding) (result *garden.SecretBinding, err error) {
	result = &garden.SecretBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretbindings").
		Name(secretBinding.Name).
		Body(secretBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the secretBinding and deletes it. Returns an error if one occurs.
func (c *secretBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretbindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secretBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretbindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched secretBinding.
func (c *secretBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.SecretBinding, err error) {
	result = &garden.SecretBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secretbindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
