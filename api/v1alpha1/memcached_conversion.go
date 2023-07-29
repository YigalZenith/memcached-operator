package v1alpha1

import (
	"github.com/example/memcached-operator/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this Memcached to the Hub version (vbeta1).
func (src *Memcached) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Memcached)
	dst.Spec.ReplicaSize = src.Spec.Size
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.Nodes = src.Status.Nodes
	return nil
}

// ConvertFrom converts from the Hub version (vbeta1) to this version.
func (dst *Memcached) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Memcached)
	dst.Spec.Size = src.Spec.ReplicaSize
	dst.ObjectMeta = src.ObjectMeta
	dst.Status.Nodes = src.Status.Nodes
	return nil
}
