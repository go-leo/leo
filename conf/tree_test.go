package conf

import (
	"context"
	"testing"
)

// Used as a workaround since we can't compare functions or their addresses
var fakeHandlerValue string

func fakeHandler(val string) HandlersChain {
	return HandlersChain{func(context.Context) {
		fakeHandlerValue = val
	}}
}

type testRequests []struct {
	path       string
	nilHandler bool
	route      string
}

func getSkippedNodes() *[]skippedNode {
	ps := make([]skippedNode, 0, 20)
	return &ps
}

func checkRequests(t *testing.T, tree *node, requests testRequests, unescapes ...bool) {
	for _, request := range requests {
		value := tree.getValue(request.path, getSkippedNodes())

		if value.handlers == nil {
			if !request.nilHandler {
				t.Errorf("handle mismatch for route '%s': Expected non-nil handle", request.path)
			}
		} else if request.nilHandler {
			t.Errorf("handle mismatch for route '%s': Expected nil handle", request.path)
		} else {
			value.handlers[0](nil)
			if fakeHandlerValue != request.route {
				t.Errorf("handle mismatch for route '%s': Wrong handle (%s != %s)", request.path, fakeHandlerValue, request.route)
			}
		}
	}
}

func checkPriorities(t *testing.T, n *node) uint32 {
	var prio uint32
	for i := range n.children {
		prio += checkPriorities(t, n.children[i])
	}

	if n.handlers != nil {
		prio++
	}

	if n.priority != prio {
		t.Errorf(
			"priority mismatch for node '%s': is %d, should be %d",
			n.path, n.priority, prio,
		)
	}

	return prio
}

func TestTreeAddAndGet(t *testing.T) {
	tree := &node{}

	routes := [...]string{
		"/hi",
		"/contact",
		"/co",
		"/c",
		"/a",
		"/ab",
		"/doc/",
		"/doc/go_faq.html",
		"/doc/go1.html",
		"/α",
		"/β",
	}
	for _, route := range routes {
		tree.addRoute(route, fakeHandler(route))
	}

	checkRequests(t, tree, testRequests{
		{"/a", false, "/a"},
		{"/", true, ""},
		{"/hi", false, "/hi"},
		{"/contact", false, "/contact"},
		{"/co", false, "/co"},
		{"/con", true, ""},  // key mismatch
		{"/cona", true, ""}, // key mismatch
		{"/no", true, ""},   // no matching child
		{"/ab", false, "/ab"},
		{"/α", false, "/α"},
		{"/β", false, "/β"},
	})

	checkPriorities(t, tree)
}
