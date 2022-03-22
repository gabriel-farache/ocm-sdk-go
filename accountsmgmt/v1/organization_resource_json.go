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

package v1 // github.com/openshift-online/ocm-sdk-go/v2/accountsmgmt/v1

import (
	"io"
	"net/http"
)

func readOrganizationGetRequest(request *OrganizationGetServerRequest, r *http.Request) error {
	return nil
}
func writeOrganizationGetRequest(request *OrganizationGetRequest, writer io.Writer) error {
	return nil
}
func readOrganizationGetResponse(response *OrganizationGetResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalOrganization(reader)
	return err
}
func writeOrganizationGetResponse(response *OrganizationGetServerResponse, w http.ResponseWriter) error {
	return MarshalOrganization(response.body, w)
}
func readOrganizationUpdateRequest(request *OrganizationUpdateServerRequest, r *http.Request) error {
	var err error
	request.body, err = UnmarshalOrganization(r.Body)
	return err
}
func writeOrganizationUpdateRequest(request *OrganizationUpdateRequest, writer io.Writer) error {
	return MarshalOrganization(request.body, writer)
}
func readOrganizationUpdateResponse(response *OrganizationUpdateResponse, reader io.Reader) error {
	var err error
	response.body, err = UnmarshalOrganization(reader)
	return err
}
func writeOrganizationUpdateResponse(response *OrganizationUpdateServerResponse, w http.ResponseWriter) error {
	return MarshalOrganization(response.body, w)
}
