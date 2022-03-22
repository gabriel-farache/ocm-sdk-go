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

// AddOnBuilder contains the data and logic needed to build 'add_on' objects.
//
// Representation of an add-on that can be installed in a cluster.
type AddOnBuilder struct {
	bitmap_              uint32
	id                   string
	href                 string
	config               *AddOnConfigBuilder
	credentialsSecret    string
	description          string
	docsLink             string
	icon                 string
	installMode          AddOnInstallMode
	label                string
	name                 string
	operatorName         string
	parameters           *AddOnParameterListBuilder
	policyPermissions    []string
	requirements         []*AddOnRequirementBuilder
	resourceCost         float64
	resourceName         string
	serviceAccount       string
	subOperators         []*AddOnSubOperatorBuilder
	targetNamespace      string
	version              *AddOnVersionBuilder
	enabled              bool
	hasExternalResources bool
	hidden               bool
	managedService       bool
}

// NewAddOn creates a new builder of 'add_on' objects.
func NewAddOn() *AddOnBuilder {
	return &AddOnBuilder{}
}

// Link sets the flag that indicates if this is a link.
func (b *AddOnBuilder) Link(value bool) *AddOnBuilder {
	b.bitmap_ |= 1
	return b
}

// ID sets the identifier of the object.
func (b *AddOnBuilder) ID(value string) *AddOnBuilder {
	b.id = value
	b.bitmap_ |= 2
	return b
}

// HREF sets the link to the object.
func (b *AddOnBuilder) HREF(value string) *AddOnBuilder {
	b.href = value
	b.bitmap_ |= 4
	return b
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *AddOnBuilder) Empty() bool {
	return b == nil || b.bitmap_&^1 == 0
}

// Config sets the value of the 'config' attribute to the given value.
//
// Representation of an add-on config.
// The attributes under it are to be used by the addon once its installed in the cluster.
func (b *AddOnBuilder) Config(value *AddOnConfigBuilder) *AddOnBuilder {
	b.config = value
	if value != nil {
		b.bitmap_ |= 8
	} else {
		b.bitmap_ &^= 8
	}
	return b
}

// CredentialsSecret sets the value of the 'credentials_secret' attribute to the given value.
//
//
func (b *AddOnBuilder) CredentialsSecret(value string) *AddOnBuilder {
	b.credentialsSecret = value
	b.bitmap_ |= 16
	return b
}

// Description sets the value of the 'description' attribute to the given value.
//
//
func (b *AddOnBuilder) Description(value string) *AddOnBuilder {
	b.description = value
	b.bitmap_ |= 32
	return b
}

// DocsLink sets the value of the 'docs_link' attribute to the given value.
//
//
func (b *AddOnBuilder) DocsLink(value string) *AddOnBuilder {
	b.docsLink = value
	b.bitmap_ |= 64
	return b
}

// Enabled sets the value of the 'enabled' attribute to the given value.
//
//
func (b *AddOnBuilder) Enabled(value bool) *AddOnBuilder {
	b.enabled = value
	b.bitmap_ |= 128
	return b
}

// HasExternalResources sets the value of the 'has_external_resources' attribute to the given value.
//
//
func (b *AddOnBuilder) HasExternalResources(value bool) *AddOnBuilder {
	b.hasExternalResources = value
	b.bitmap_ |= 256
	return b
}

// Hidden sets the value of the 'hidden' attribute to the given value.
//
//
func (b *AddOnBuilder) Hidden(value bool) *AddOnBuilder {
	b.hidden = value
	b.bitmap_ |= 512
	return b
}

// Icon sets the value of the 'icon' attribute to the given value.
//
//
func (b *AddOnBuilder) Icon(value string) *AddOnBuilder {
	b.icon = value
	b.bitmap_ |= 1024
	return b
}

// InstallMode sets the value of the 'install_mode' attribute to the given value.
//
// Representation of an add-on InstallMode field.
func (b *AddOnBuilder) InstallMode(value AddOnInstallMode) *AddOnBuilder {
	b.installMode = value
	b.bitmap_ |= 2048
	return b
}

// Label sets the value of the 'label' attribute to the given value.
//
//
func (b *AddOnBuilder) Label(value string) *AddOnBuilder {
	b.label = value
	b.bitmap_ |= 4096
	return b
}

// ManagedService sets the value of the 'managed_service' attribute to the given value.
//
//
func (b *AddOnBuilder) ManagedService(value bool) *AddOnBuilder {
	b.managedService = value
	b.bitmap_ |= 8192
	return b
}

// Name sets the value of the 'name' attribute to the given value.
//
//
func (b *AddOnBuilder) Name(value string) *AddOnBuilder {
	b.name = value
	b.bitmap_ |= 16384
	return b
}

// OperatorName sets the value of the 'operator_name' attribute to the given value.
//
//
func (b *AddOnBuilder) OperatorName(value string) *AddOnBuilder {
	b.operatorName = value
	b.bitmap_ |= 32768
	return b
}

// Parameters sets the value of the 'parameters' attribute to the given values.
//
//
func (b *AddOnBuilder) Parameters(value *AddOnParameterListBuilder) *AddOnBuilder {
	b.parameters = value
	b.bitmap_ |= 65536
	return b
}

// PolicyPermissions sets the value of the 'policy_permissions' attribute to the given values.
//
//
func (b *AddOnBuilder) PolicyPermissions(values ...string) *AddOnBuilder {
	b.policyPermissions = make([]string, len(values))
	copy(b.policyPermissions, values)
	b.bitmap_ |= 131072
	return b
}

// Requirements sets the value of the 'requirements' attribute to the given values.
//
//
func (b *AddOnBuilder) Requirements(values ...*AddOnRequirementBuilder) *AddOnBuilder {
	b.requirements = make([]*AddOnRequirementBuilder, len(values))
	copy(b.requirements, values)
	b.bitmap_ |= 262144
	return b
}

// ResourceCost sets the value of the 'resource_cost' attribute to the given value.
//
//
func (b *AddOnBuilder) ResourceCost(value float64) *AddOnBuilder {
	b.resourceCost = value
	b.bitmap_ |= 524288
	return b
}

// ResourceName sets the value of the 'resource_name' attribute to the given value.
//
//
func (b *AddOnBuilder) ResourceName(value string) *AddOnBuilder {
	b.resourceName = value
	b.bitmap_ |= 1048576
	return b
}

// ServiceAccount sets the value of the 'service_account' attribute to the given value.
//
//
func (b *AddOnBuilder) ServiceAccount(value string) *AddOnBuilder {
	b.serviceAccount = value
	b.bitmap_ |= 2097152
	return b
}

// SubOperators sets the value of the 'sub_operators' attribute to the given values.
//
//
func (b *AddOnBuilder) SubOperators(values ...*AddOnSubOperatorBuilder) *AddOnBuilder {
	b.subOperators = make([]*AddOnSubOperatorBuilder, len(values))
	copy(b.subOperators, values)
	b.bitmap_ |= 4194304
	return b
}

// TargetNamespace sets the value of the 'target_namespace' attribute to the given value.
//
//
func (b *AddOnBuilder) TargetNamespace(value string) *AddOnBuilder {
	b.targetNamespace = value
	b.bitmap_ |= 8388608
	return b
}

// Version sets the value of the 'version' attribute to the given value.
//
// Representation of an add-on version.
func (b *AddOnBuilder) Version(value *AddOnVersionBuilder) *AddOnBuilder {
	b.version = value
	if value != nil {
		b.bitmap_ |= 16777216
	} else {
		b.bitmap_ &^= 16777216
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *AddOnBuilder) Copy(object *AddOn) *AddOnBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.href = object.href
	if object.config != nil {
		b.config = NewAddOnConfig().Copy(object.config)
	} else {
		b.config = nil
	}
	b.credentialsSecret = object.credentialsSecret
	b.description = object.description
	b.docsLink = object.docsLink
	b.enabled = object.enabled
	b.hasExternalResources = object.hasExternalResources
	b.hidden = object.hidden
	b.icon = object.icon
	b.installMode = object.installMode
	b.label = object.label
	b.managedService = object.managedService
	b.name = object.name
	b.operatorName = object.operatorName
	if object.parameters != nil {
		b.parameters = NewAddOnParameterList().Copy(object.parameters)
	} else {
		b.parameters = nil
	}
	if object.policyPermissions != nil {
		b.policyPermissions = make([]string, len(object.policyPermissions))
		copy(b.policyPermissions, object.policyPermissions)
	} else {
		b.policyPermissions = nil
	}
	if object.requirements != nil {
		b.requirements = make([]*AddOnRequirementBuilder, len(object.requirements))
		for i, v := range object.requirements {
			b.requirements[i] = NewAddOnRequirement().Copy(v)
		}
	} else {
		b.requirements = nil
	}
	b.resourceCost = object.resourceCost
	b.resourceName = object.resourceName
	b.serviceAccount = object.serviceAccount
	if object.subOperators != nil {
		b.subOperators = make([]*AddOnSubOperatorBuilder, len(object.subOperators))
		for i, v := range object.subOperators {
			b.subOperators[i] = NewAddOnSubOperator().Copy(v)
		}
	} else {
		b.subOperators = nil
	}
	b.targetNamespace = object.targetNamespace
	if object.version != nil {
		b.version = NewAddOnVersion().Copy(object.version)
	} else {
		b.version = nil
	}
	return b
}

// Build creates a 'add_on' object using the configuration stored in the builder.
func (b *AddOnBuilder) Build() (object *AddOn, err error) {
	object = new(AddOn)
	object.id = b.id
	object.href = b.href
	object.bitmap_ = b.bitmap_
	if b.config != nil {
		object.config, err = b.config.Build()
		if err != nil {
			return
		}
	}
	object.credentialsSecret = b.credentialsSecret
	object.description = b.description
	object.docsLink = b.docsLink
	object.enabled = b.enabled
	object.hasExternalResources = b.hasExternalResources
	object.hidden = b.hidden
	object.icon = b.icon
	object.installMode = b.installMode
	object.label = b.label
	object.managedService = b.managedService
	object.name = b.name
	object.operatorName = b.operatorName
	if b.parameters != nil {
		object.parameters, err = b.parameters.Build()
		if err != nil {
			return
		}
	}
	if b.policyPermissions != nil {
		object.policyPermissions = make([]string, len(b.policyPermissions))
		copy(object.policyPermissions, b.policyPermissions)
	}
	if b.requirements != nil {
		object.requirements = make([]*AddOnRequirement, len(b.requirements))
		for i, v := range b.requirements {
			object.requirements[i], err = v.Build()
			if err != nil {
				return
			}
		}
	}
	object.resourceCost = b.resourceCost
	object.resourceName = b.resourceName
	object.serviceAccount = b.serviceAccount
	if b.subOperators != nil {
		object.subOperators = make([]*AddOnSubOperator, len(b.subOperators))
		for i, v := range b.subOperators {
			object.subOperators[i], err = v.Build()
			if err != nil {
				return
			}
		}
	}
	object.targetNamespace = b.targetNamespace
	if b.version != nil {
		object.version, err = b.version.Build()
		if err != nil {
			return
		}
	}
	return
}
