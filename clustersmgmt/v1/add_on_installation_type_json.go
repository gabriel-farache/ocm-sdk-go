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
	"io"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/v2/helpers"
)

// MarshalAddOnInstallation writes a value of the 'add_on_installation' type to the given writer.
func MarshalAddOnInstallation(object *AddOnInstallation, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeAddOnInstallation(object, stream)
	err := stream.Flush()
	if err != nil {
		return err
	}
	return stream.Error
}

// writeAddOnInstallation writes a value of the 'add_on_installation' type to the given stream.
func writeAddOnInstallation(object *AddOnInstallation, stream *jsoniter.Stream) {
	count := 0
	stream.WriteObjectStart()
	stream.WriteObjectField("kind")
	if object.bitmap_&1 != 0 {
		stream.WriteString(AddOnInstallationLinkKind)
	} else {
		stream.WriteString(AddOnInstallationKind)
	}
	count++
	if object.bitmap_&2 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("id")
		stream.WriteString(object.id)
		count++
	}
	if object.bitmap_&4 != 0 {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("href")
		stream.WriteString(object.href)
		count++
	}
	var present_ bool
	present_ = object.bitmap_&8 != 0 && object.addon != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("addon")
		writeAddOn(object.addon, stream)
		count++
	}
	present_ = object.bitmap_&16 != 0 && object.addonVersion != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("addon_version")
		writeAddOnVersion(object.addonVersion, stream)
		count++
	}
	present_ = object.bitmap_&32 != 0 && object.cluster != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("cluster")
		writeCluster(object.cluster, stream)
		count++
	}
	present_ = object.bitmap_&64 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("creation_timestamp")
		stream.WriteString((object.creationTimestamp).Format(time.RFC3339))
		count++
	}
	present_ = object.bitmap_&128 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("operator_version")
		stream.WriteString(object.operatorVersion)
		count++
	}
	present_ = object.bitmap_&256 != 0 && object.parameters != nil
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("parameters")
		stream.WriteObjectStart()
		stream.WriteObjectField("items")
		writeAddOnInstallationParameterList(object.parameters.items, stream)
		stream.WriteObjectEnd()
		count++
	}
	present_ = object.bitmap_&512 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("role_arn")
		stream.WriteString(object.roleARN)
		count++
	}
	present_ = object.bitmap_&1024 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("state")
		stream.WriteString(string(object.state))
		count++
	}
	present_ = object.bitmap_&2048 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("state_description")
		stream.WriteString(object.stateDescription)
		count++
	}
	present_ = object.bitmap_&4096 != 0
	if present_ {
		if count > 0 {
			stream.WriteMore()
		}
		stream.WriteObjectField("updated_timestamp")
		stream.WriteString((object.updatedTimestamp).Format(time.RFC3339))
	}
	stream.WriteObjectEnd()
}

// UnmarshalAddOnInstallation reads a value of the 'add_on_installation' type from the given
// source, which can be an slice of bytes, a string or a reader.
func UnmarshalAddOnInstallation(source interface{}) (object *AddOnInstallation, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	object = readAddOnInstallation(iterator)
	err = iterator.Error
	return
}

// readAddOnInstallation reads a value of the 'add_on_installation' type from the given iterator.
func readAddOnInstallation(iterator *jsoniter.Iterator) *AddOnInstallation {
	object := &AddOnInstallation{}
	for {
		field := iterator.ReadObject()
		if field == "" {
			break
		}
		switch field {
		case "kind":
			value := iterator.ReadString()
			if value == AddOnInstallationLinkKind {
				object.bitmap_ |= 1
			}
		case "id":
			object.id = iterator.ReadString()
			object.bitmap_ |= 2
		case "href":
			object.href = iterator.ReadString()
			object.bitmap_ |= 4
		case "addon":
			value := readAddOn(iterator)
			object.addon = value
			object.bitmap_ |= 8
		case "addon_version":
			value := readAddOnVersion(iterator)
			object.addonVersion = value
			object.bitmap_ |= 16
		case "cluster":
			value := readCluster(iterator)
			object.cluster = value
			object.bitmap_ |= 32
		case "creation_timestamp":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.creationTimestamp = value
			object.bitmap_ |= 64
		case "operator_version":
			value := iterator.ReadString()
			object.operatorVersion = value
			object.bitmap_ |= 128
		case "parameters":
			value := &AddOnInstallationParameterList{}
			for {
				field := iterator.ReadObject()
				if field == "" {
					break
				}
				switch field {
				case "kind":
					text := iterator.ReadString()
					value.link = text == AddOnInstallationParameterListLinkKind
				case "href":
					value.href = iterator.ReadString()
				case "items":
					value.items = readAddOnInstallationParameterList(iterator)
				default:
					iterator.ReadAny()
				}
			}
			object.parameters = value
			object.bitmap_ |= 256
		case "role_arn":
			value := iterator.ReadString()
			object.roleARN = value
			object.bitmap_ |= 512
		case "state":
			text := iterator.ReadString()
			value := AddOnInstallationState(text)
			object.state = value
			object.bitmap_ |= 1024
		case "state_description":
			value := iterator.ReadString()
			object.stateDescription = value
			object.bitmap_ |= 2048
		case "updated_timestamp":
			text := iterator.ReadString()
			value, err := time.Parse(time.RFC3339, text)
			if err != nil {
				iterator.ReportError("", err.Error())
			}
			object.updatedTimestamp = value
			object.bitmap_ |= 4096
		default:
			iterator.ReadAny()
		}
	}
	return object
}
