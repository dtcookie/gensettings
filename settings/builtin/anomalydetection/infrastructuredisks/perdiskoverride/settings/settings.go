/**
* @license
* Copyright 2020 Dynatrace LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package perdiskoverride

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	DiskID                              string                           `json:"-" scope:"diskId"` // The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	DiskLowInodesDetection              *DiskLowInodesDetection          `json:"diskLowInodesDetection,omitempty"`
	DiskLowSpaceDetection               *DiskLowSpaceDetection           `json:"diskLowSpaceDetection,omitempty"`
	DiskSlowWritesAndReadsDetection     *DiskSlowWritesAndReadsDetection `json:"diskSlowWritesAndReadsDetection,omitempty"`
	OverrideDiskLowSpaceDetection       bool                             `json:"overrideDiskLowSpaceDetection"`       // Override low disk space detection settings
	OverrideLowInodesDetection          bool                             `json:"overrideLowInodesDetection"`          // Override low inodes detection settings
	OverrideSlowWritesAndReadsDetection bool                             `json:"overrideSlowWritesAndReadsDetection"` // Override slow writes and reads detection settings
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"disk_id": {
			Type:        schema.TypeString,
			Description: "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.",
			Required:    true,
		},
		"disk_low_inodes_detection": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(DiskLowInodesDetection).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"disk_low_space_detection": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(DiskLowSpaceDetection).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"disk_slow_writes_and_reads_detection": {
			Type:        schema.TypeList,
			Description: "no documentation available",
			Optional:    true,

			Elem:     &schema.Resource{Schema: new(DiskSlowWritesAndReadsDetection).Schema()},
			MinItems: 1,
			MaxItems: 1,
		},
		"override_disk_low_space_detection": {
			Type:        schema.TypeBool,
			Description: "Override low disk space detection settings",
			Required:    true,
		},
		"override_low_inodes_detection": {
			Type:        schema.TypeBool,
			Description: "Override low inodes detection settings",
			Required:    true,
		},
		"override_slow_writes_and_reads_detection": {
			Type:        schema.TypeBool,
			Description: "Override slow writes and reads detection settings",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"disk_id":                                  me.DiskID,
		"disk_low_inodes_detection":                me.DiskLowInodesDetection,
		"disk_low_space_detection":                 me.DiskLowSpaceDetection,
		"disk_slow_writes_and_reads_detection":     me.DiskSlowWritesAndReadsDetection,
		"override_disk_low_space_detection":        me.OverrideDiskLowSpaceDetection,
		"override_low_inodes_detection":            me.OverrideLowInodesDetection,
		"override_slow_writes_and_reads_detection": me.OverrideSlowWritesAndReadsDetection,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"disk_id":                                  &me.DiskID,
		"disk_low_inodes_detection":                &me.DiskLowInodesDetection,
		"disk_low_space_detection":                 &me.DiskLowSpaceDetection,
		"disk_slow_writes_and_reads_detection":     &me.DiskSlowWritesAndReadsDetection,
		"override_disk_low_space_detection":        &me.OverrideDiskLowSpaceDetection,
		"override_low_inodes_detection":            &me.OverrideLowInodesDetection,
		"override_slow_writes_and_reads_detection": &me.OverrideSlowWritesAndReadsDetection,
	})
}
