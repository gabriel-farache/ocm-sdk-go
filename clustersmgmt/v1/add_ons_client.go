/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/v2/clustersmgmt/v1

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/openshift-online/ocm-sdk-go/v2/errors"
	"github.com/openshift-online/ocm-sdk-go/v2/helpers"
)

// AddOnsClient is the client of the 'add_ons' resource.
//
// Manages the collection of add-ons.
type AddOnsClient struct {
	transport http.RoundTripper
	path      string
}

// NewAddOnsClient creates a new client for the 'add_ons'
// resource using the given transport to send the requests and receive the
// responses.
func NewAddOnsClient(transport http.RoundTripper, path string) *AddOnsClient {
	return &AddOnsClient{
		transport: transport,
		path:      path,
	}
}

// Add creates a request for the 'add' method.
//
// Create a new add-on and add it to the collection of add-ons.
func (c *AddOnsClient) Add() *AddOnsAddRequest {
	return &AddOnsAddRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// List creates a request for the 'list' method.
//
// Retrieves the list of add-ons.
func (c *AddOnsClient) List() *AddOnsListRequest {
	return &AddOnsListRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// Addon returns the target 'add_on' resource for the given identifier.
//
// Returns a reference to the service that manages a specific add-on.
func (c *AddOnsClient) Addon(id string) *AddOnClient {
	return NewAddOnClient(
		c.transport,
		path.Join(c.path, id),
	)
}

// AddOnsAddRequest is the request for the 'add' method.
type AddOnsAddRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	body      *AddOn
}

// Parameter adds a query parameter.
func (r *AddOnsAddRequest) Parameter(name string, value interface{}) *AddOnsAddRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *AddOnsAddRequest) Header(name string, value interface{}) *AddOnsAddRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *AddOnsAddRequest) Impersonate(user string) *AddOnsAddRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Body sets the value of the 'body' parameter.
//
// Description of the add-on.
func (r *AddOnsAddRequest) Body(value *AddOn) *AddOnsAddRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *AddOnsAddRequest) Send() (result *AddOnsAddResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *AddOnsAddRequest) SendContext(ctx context.Context) (result *AddOnsAddResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeAddOnsAddRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &AddOnsAddResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readAddOnsAddResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// AddOnsAddResponse is the response for the 'add' method.
type AddOnsAddResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *AddOn
}

// Status returns the response status code.
func (r *AddOnsAddResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *AddOnsAddResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *AddOnsAddResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
// Description of the add-on.
func (r *AddOnsAddResponse) Body() *AddOn {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
// Description of the add-on.
func (r *AddOnsAddResponse) GetBody() (value *AddOn, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}

// AddOnsListRequest is the request for the 'list' method.
type AddOnsListRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	order     *string
	page      *int
	search    *string
	size      *int
}

// Parameter adds a query parameter.
func (r *AddOnsListRequest) Parameter(name string, value interface{}) *AddOnsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *AddOnsListRequest) Header(name string, value interface{}) *AddOnsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Impersonate wraps requests on behalf of another user.
// Note: Services that do not support this feature may silently ignore this call.
func (r *AddOnsListRequest) Impersonate(user string) *AddOnsListRequest {
	helpers.AddImpersonationHeader(&r.header, user)
	return r
}

// Order sets the value of the 'order' parameter.
//
// Order criteria.
//
// The syntax of this parameter is similar to the syntax of the _order by_ clause of
// a SQL statement, but using the names of the attributes of the add-on instead of
// the names of the columns of a table. For example, in order to sort the add-ons
// descending by name the value should be:
//
// ```sql
// name desc
// ```
//
// If the parameter isn't provided, or if the value is empty, then the order of the
// results is undefined.
func (r *AddOnsListRequest) Order(value string) *AddOnsListRequest {
	r.order = &value
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *AddOnsListRequest) Page(value int) *AddOnsListRequest {
	r.page = &value
	return r
}

// Search sets the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the _where_ clause of an
// SQL statement, but using the names of the attributes of the add-on instead of
// the names of the columns of a table. For example, in order to retrieve all the
// add-ons with a name starting with `my` the value should be:
//
// ```sql
// name like 'my%'
// ```
//
// If the parameter isn't provided, or if the value is empty, then all the add-ons
// that the user has permission to see will be returned.
func (r *AddOnsListRequest) Search(value string) *AddOnsListRequest {
	r.search = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *AddOnsListRequest) Size(value int) *AddOnsListRequest {
	r.size = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *AddOnsListRequest) Send() (result *AddOnsListResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *AddOnsListRequest) SendContext(ctx context.Context) (result *AddOnsListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.order != nil {
		helpers.AddValue(&query, "order", *r.order)
	}
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.search != nil {
		helpers.AddValue(&query, "search", *r.search)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &AddOnsListResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	reader := bufio.NewReader(response.Body)
	_, err = reader.Peek(1)
	if err == io.EOF {
		err = nil
		return
	}
	if result.status >= 400 {
		result.err, err = errors.UnmarshalErrorStatus(reader, result.status)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readAddOnsListResponse(result, reader)
	if err != nil {
		return
	}
	return
}

// AddOnsListResponse is the response for the 'list' method.
type AddOnsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	items  *AddOnList
	page   *int
	size   *int
	total  *int
}

// Status returns the response status code.
func (r *AddOnsListResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *AddOnsListResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *AddOnsListResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of add-ons.
func (r *AddOnsListResponse) Items() *AddOnList {
	if r == nil {
		return nil
	}
	return r.items
}

// GetItems returns the value of the 'items' parameter and
// a flag indicating if the parameter has a value.
//
// Retrieved list of add-ons.
func (r *AddOnsListResponse) GetItems() (value *AddOnList, ok bool) {
	ok = r != nil && r.items != nil
	if ok {
		value = r.items
	}
	return
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *AddOnsListResponse) Page() int {
	if r != nil && r.page != nil {
		return *r.page
	}
	return 0
}

// GetPage returns the value of the 'page' parameter and
// a flag indicating if the parameter has a value.
//
// Index of the requested page, where one corresponds to the first page.
func (r *AddOnsListResponse) GetPage() (value int, ok bool) {
	ok = r != nil && r.page != nil
	if ok {
		value = *r.page
	}
	return
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
func (r *AddOnsListResponse) Size() int {
	if r != nil && r.size != nil {
		return *r.size
	}
	return 0
}

// GetSize returns the value of the 'size' parameter and
// a flag indicating if the parameter has a value.
//
// Maximum number of items that will be contained in the returned page.
func (r *AddOnsListResponse) GetSize() (value int, ok bool) {
	ok = r != nil && r.size != nil
	if ok {
		value = *r.size
	}
	return
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *AddOnsListResponse) Total() int {
	if r != nil && r.total != nil {
		return *r.total
	}
	return 0
}

// GetTotal returns the value of the 'total' parameter and
// a flag indicating if the parameter has a value.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *AddOnsListResponse) GetTotal() (value int, ok bool) {
	ok = r != nil && r.total != nil
	if ok {
		value = *r.total
	}
	return
}
