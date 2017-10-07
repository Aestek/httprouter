package httprouter

import "errors"

// RouteSet represents the set of routes used by the router
type RouteSet map[string]*node

// NewRouteSet creates a new RouteSet
func NewRouteSet() RouteSet {
	return make(RouteSet)
}

// AddRoute adds a route to the RouteSet
// An error can be returned if
// - the route does not start with a '/'
// - the route conflicts with an existing route in the set
func (s RouteSet) AddRoute(method, path string, handle Handle) error {
	if path[0] != '/' {
		return errors.New("path must begin with '/' in path '" + path + "'")
	}

	root := s[method]
	if root == nil {
		root = new(node)
		s[method] = root
	}

	return root.addRoute(path, handle)
}
