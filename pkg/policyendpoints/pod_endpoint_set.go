package policyendpoints

import (
	policyinfo "github.com/aws/amazon-network-policy-controller-k8s/api/v1alpha1"
	policyequality "github.com/aws/amazon-network-policy-controller-k8s/pkg/equality"
)

// podEndpointSet is a set of PodEndpoints keyed by all of their fields, including labels.
// PodEndpoint has a map field, so it isn't comparable and can't be used with sets.Set.
type podEndpointSet map[string]policyinfo.PodEndpoint

func newPodEndpointSet(endpoints ...policyinfo.PodEndpoint) podEndpointSet {
	s := make(podEndpointSet, len(endpoints))
	for _, ep := range endpoints {
		s[policyequality.PodEndpointKey(ep)] = ep
	}
	return s
}

// Has reports whether an endpoint with the same fields and labels is in the set.
func (s podEndpointSet) Has(ep policyinfo.PodEndpoint) bool {
	_, ok := s[policyequality.PodEndpointKey(ep)]
	return ok
}

// Delete removes the endpoint with the same fields and labels, if present.
func (s podEndpointSet) Delete(ep policyinfo.PodEndpoint) {
	delete(s, policyequality.PodEndpointKey(ep))
}

// UnsortedList returns the endpoints in the set in no particular order.
func (s podEndpointSet) UnsortedList() []policyinfo.PodEndpoint {
	out := make([]policyinfo.PodEndpoint, 0, len(s))
	for _, ep := range s {
		out = append(out, ep)
	}
	return out
}
