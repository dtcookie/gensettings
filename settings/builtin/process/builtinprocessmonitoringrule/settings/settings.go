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

package builtinprocessmonitoringrule

import (
	"github.com/dynatrace-oss/terraform-provider-dynatrace/terraform/hcl"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Settings struct {
	HostGroupID *string `json:"-" scope:"hostGroupId"` // The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.
	_10         bool    `json:"-10"`                   // Rule id: 10
	_11         bool    `json:"-11"`                   // Rule id: 11
	_12         bool    `json:"-12"`                   // Rule id: 12
	_13         bool    `json:"-13"`                   // Rule id: 13
	_14         bool    `json:"-14"`                   // Rule id: 14
	_16         bool    `json:"-16"`                   // Rule id: 16
	_17         bool    `json:"-17"`                   // Rule id: 17
	_18         bool    `json:"-18"`                   // Rule id: 18
	_19         bool    `json:"-19"`                   // Rule id: 19
	_2          bool    `json:"-2"`                    // Rule id: 2
	_20         bool    `json:"-20"`                   // Rule id: 20
	_21         bool    `json:"-21"`                   // Rule id: 21
	_22         bool    `json:"-22"`                   // Rule id: 22
	_23         bool    `json:"-23"`                   // Rule id: 23
	_24         bool    `json:"-24"`                   // Rule id: 24
	_25         bool    `json:"-25"`                   // Rule id: 25
	_26         bool    `json:"-26"`                   // Rule id: 26
	_27         bool    `json:"-27"`                   // Rule id: 27
	_28         bool    `json:"-28"`                   // Rule id: 28
	_29         bool    `json:"-29"`                   // Rule id: 29
	_3          bool    `json:"-3"`                    // Rule id: 3
	_32         bool    `json:"-32"`                   // Rule id: 32
	_33         bool    `json:"-33"`                   // Rule id: 33
	_34         bool    `json:"-34"`                   // Rule id: 34
	_35         bool    `json:"-35"`                   // Rule id: 35
	_36         bool    `json:"-36"`                   // Rule id: 36
	_37         bool    `json:"-37"`                   // Rule id: 37
	_38         bool    `json:"-38"`                   // Rule id: 38
	_39         bool    `json:"-39"`                   // Rule id: 39
	_4          bool    `json:"-4"`                    // Rule id: 4
	_40         bool    `json:"-40"`                   // Rule id: 40
	_41         bool    `json:"-41"`                   // Rule id: 41
	_43         bool    `json:"-43"`                   // Rule id: 43
	_44         bool    `json:"-44"`                   // Rule id: 44
	_45         bool    `json:"-45"`                   // Rule id: 45
	_46         bool    `json:"-46"`                   // Rule id: 46
	_47         bool    `json:"-47"`                   // Rule id: 47
	_48         bool    `json:"-48"`                   // Rule id: 48
	_49         bool    `json:"-49"`                   // Rule id: 49
	_5          bool    `json:"-5"`                    // Rule id: 5
	_50         bool    `json:"-50"`                   // Rule id: 50
	_51         bool    `json:"-51"`                   // Rule id: 51
	_52         bool    `json:"-52"`                   // Rule id: 52
	_53         bool    `json:"-53"`                   // Rule id: 53
	_54         bool    `json:"-54"`                   // Rule id: 54
	_55         bool    `json:"-55"`                   // Rule id: 55
	_56         bool    `json:"-56"`                   // Rule id: 56
	_57         bool    `json:"-57"`                   // Rule id: 57
	_58         bool    `json:"-58"`                   // Rule id: 58
	_59         bool    `json:"-59"`                   // Rule id: 59
	_6          bool    `json:"-6"`                    // Rule id: 6
	_60         bool    `json:"-60"`                   // Rule id: 60
	_61         bool    `json:"-61"`                   // Rule id: 61
	_62         bool    `json:"-62"`                   // Rule id: 62
	_63         bool    `json:"-63"`                   // Rule id: 63
	_64         bool    `json:"-64"`                   // Rule id: 64
	_65         bool    `json:"-65"`                   // Rule id: 65
	_66         bool    `json:"-66"`                   // Rule id: 66
	_67         bool    `json:"-67"`                   // Rule id: 67
	_68         bool    `json:"-68"`                   // Rule id: 68
	_69         bool    `json:"-69"`                   // Rule id: 69
	_7          bool    `json:"-7"`                    // Rule id: 7
	_70         bool    `json:"-70"`                   // Rule id: 70
	_8          bool    `json:"-8"`                    // Rule id: 8
	_9          bool    `json:"-9"`                    // Rule id: 9
}

func (me *Settings) Schema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"host_group_id": {
			Type:        schema.TypeString,
			Description: "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope.",
			Optional:    true,
			Default:     "environment",
		},
		"_10": {
			Type:        schema.TypeBool,
			Description: "Rule id: 10",
			Required:    true,
		},
		"_11": {
			Type:        schema.TypeBool,
			Description: "Rule id: 11",
			Required:    true,
		},
		"_12": {
			Type:        schema.TypeBool,
			Description: "Rule id: 12",
			Required:    true,
		},
		"_13": {
			Type:        schema.TypeBool,
			Description: "Rule id: 13",
			Required:    true,
		},
		"_14": {
			Type:        schema.TypeBool,
			Description: "Rule id: 14",
			Required:    true,
		},
		"_16": {
			Type:        schema.TypeBool,
			Description: "Rule id: 16",
			Required:    true,
		},
		"_17": {
			Type:        schema.TypeBool,
			Description: "Rule id: 17",
			Required:    true,
		},
		"_18": {
			Type:        schema.TypeBool,
			Description: "Rule id: 18",
			Required:    true,
		},
		"_19": {
			Type:        schema.TypeBool,
			Description: "Rule id: 19",
			Required:    true,
		},
		"_2": {
			Type:        schema.TypeBool,
			Description: "Rule id: 2",
			Required:    true,
		},
		"_20": {
			Type:        schema.TypeBool,
			Description: "Rule id: 20",
			Required:    true,
		},
		"_21": {
			Type:        schema.TypeBool,
			Description: "Rule id: 21",
			Required:    true,
		},
		"_22": {
			Type:        schema.TypeBool,
			Description: "Rule id: 22",
			Required:    true,
		},
		"_23": {
			Type:        schema.TypeBool,
			Description: "Rule id: 23",
			Required:    true,
		},
		"_24": {
			Type:        schema.TypeBool,
			Description: "Rule id: 24",
			Required:    true,
		},
		"_25": {
			Type:        schema.TypeBool,
			Description: "Rule id: 25",
			Required:    true,
		},
		"_26": {
			Type:        schema.TypeBool,
			Description: "Rule id: 26",
			Required:    true,
		},
		"_27": {
			Type:        schema.TypeBool,
			Description: "Rule id: 27",
			Required:    true,
		},
		"_28": {
			Type:        schema.TypeBool,
			Description: "Rule id: 28",
			Required:    true,
		},
		"_29": {
			Type:        schema.TypeBool,
			Description: "Rule id: 29",
			Required:    true,
		},
		"_3": {
			Type:        schema.TypeBool,
			Description: "Rule id: 3",
			Required:    true,
		},
		"_32": {
			Type:        schema.TypeBool,
			Description: "Rule id: 32",
			Required:    true,
		},
		"_33": {
			Type:        schema.TypeBool,
			Description: "Rule id: 33",
			Required:    true,
		},
		"_34": {
			Type:        schema.TypeBool,
			Description: "Rule id: 34",
			Required:    true,
		},
		"_35": {
			Type:        schema.TypeBool,
			Description: "Rule id: 35",
			Required:    true,
		},
		"_36": {
			Type:        schema.TypeBool,
			Description: "Rule id: 36",
			Required:    true,
		},
		"_37": {
			Type:        schema.TypeBool,
			Description: "Rule id: 37",
			Required:    true,
		},
		"_38": {
			Type:        schema.TypeBool,
			Description: "Rule id: 38",
			Required:    true,
		},
		"_39": {
			Type:        schema.TypeBool,
			Description: "Rule id: 39",
			Required:    true,
		},
		"_4": {
			Type:        schema.TypeBool,
			Description: "Rule id: 4",
			Required:    true,
		},
		"_40": {
			Type:        schema.TypeBool,
			Description: "Rule id: 40",
			Required:    true,
		},
		"_41": {
			Type:        schema.TypeBool,
			Description: "Rule id: 41",
			Required:    true,
		},
		"_43": {
			Type:        schema.TypeBool,
			Description: "Rule id: 43",
			Required:    true,
		},
		"_44": {
			Type:        schema.TypeBool,
			Description: "Rule id: 44",
			Required:    true,
		},
		"_45": {
			Type:        schema.TypeBool,
			Description: "Rule id: 45",
			Required:    true,
		},
		"_46": {
			Type:        schema.TypeBool,
			Description: "Rule id: 46",
			Required:    true,
		},
		"_47": {
			Type:        schema.TypeBool,
			Description: "Rule id: 47",
			Required:    true,
		},
		"_48": {
			Type:        schema.TypeBool,
			Description: "Rule id: 48",
			Required:    true,
		},
		"_49": {
			Type:        schema.TypeBool,
			Description: "Rule id: 49",
			Required:    true,
		},
		"_5": {
			Type:        schema.TypeBool,
			Description: "Rule id: 5",
			Required:    true,
		},
		"_50": {
			Type:        schema.TypeBool,
			Description: "Rule id: 50",
			Required:    true,
		},
		"_51": {
			Type:        schema.TypeBool,
			Description: "Rule id: 51",
			Required:    true,
		},
		"_52": {
			Type:        schema.TypeBool,
			Description: "Rule id: 52",
			Required:    true,
		},
		"_53": {
			Type:        schema.TypeBool,
			Description: "Rule id: 53",
			Required:    true,
		},
		"_54": {
			Type:        schema.TypeBool,
			Description: "Rule id: 54",
			Required:    true,
		},
		"_55": {
			Type:        schema.TypeBool,
			Description: "Rule id: 55",
			Required:    true,
		},
		"_56": {
			Type:        schema.TypeBool,
			Description: "Rule id: 56",
			Required:    true,
		},
		"_57": {
			Type:        schema.TypeBool,
			Description: "Rule id: 57",
			Required:    true,
		},
		"_58": {
			Type:        schema.TypeBool,
			Description: "Rule id: 58",
			Required:    true,
		},
		"_59": {
			Type:        schema.TypeBool,
			Description: "Rule id: 59",
			Required:    true,
		},
		"_6": {
			Type:        schema.TypeBool,
			Description: "Rule id: 6",
			Required:    true,
		},
		"_60": {
			Type:        schema.TypeBool,
			Description: "Rule id: 60",
			Required:    true,
		},
		"_61": {
			Type:        schema.TypeBool,
			Description: "Rule id: 61",
			Required:    true,
		},
		"_62": {
			Type:        schema.TypeBool,
			Description: "Rule id: 62",
			Required:    true,
		},
		"_63": {
			Type:        schema.TypeBool,
			Description: "Rule id: 63",
			Required:    true,
		},
		"_64": {
			Type:        schema.TypeBool,
			Description: "Rule id: 64",
			Required:    true,
		},
		"_65": {
			Type:        schema.TypeBool,
			Description: "Rule id: 65",
			Required:    true,
		},
		"_66": {
			Type:        schema.TypeBool,
			Description: "Rule id: 66",
			Required:    true,
		},
		"_67": {
			Type:        schema.TypeBool,
			Description: "Rule id: 67",
			Required:    true,
		},
		"_68": {
			Type:        schema.TypeBool,
			Description: "Rule id: 68",
			Required:    true,
		},
		"_69": {
			Type:        schema.TypeBool,
			Description: "Rule id: 69",
			Required:    true,
		},
		"_7": {
			Type:        schema.TypeBool,
			Description: "Rule id: 7",
			Required:    true,
		},
		"_70": {
			Type:        schema.TypeBool,
			Description: "Rule id: 70",
			Required:    true,
		},
		"_8": {
			Type:        schema.TypeBool,
			Description: "Rule id: 8",
			Required:    true,
		},
		"_9": {
			Type:        schema.TypeBool,
			Description: "Rule id: 9",
			Required:    true,
		},
	}
}

func (me *Settings) MarshalHCL(properties hcl.Properties) error {
	return properties.EncodeAll(map[string]any{
		"host_group_id": me.HostGroupID,
		"_10":           me._10,
		"_11":           me._11,
		"_12":           me._12,
		"_13":           me._13,
		"_14":           me._14,
		"_16":           me._16,
		"_17":           me._17,
		"_18":           me._18,
		"_19":           me._19,
		"_2":            me._2,
		"_20":           me._20,
		"_21":           me._21,
		"_22":           me._22,
		"_23":           me._23,
		"_24":           me._24,
		"_25":           me._25,
		"_26":           me._26,
		"_27":           me._27,
		"_28":           me._28,
		"_29":           me._29,
		"_3":            me._3,
		"_32":           me._32,
		"_33":           me._33,
		"_34":           me._34,
		"_35":           me._35,
		"_36":           me._36,
		"_37":           me._37,
		"_38":           me._38,
		"_39":           me._39,
		"_4":            me._4,
		"_40":           me._40,
		"_41":           me._41,
		"_43":           me._43,
		"_44":           me._44,
		"_45":           me._45,
		"_46":           me._46,
		"_47":           me._47,
		"_48":           me._48,
		"_49":           me._49,
		"_5":            me._5,
		"_50":           me._50,
		"_51":           me._51,
		"_52":           me._52,
		"_53":           me._53,
		"_54":           me._54,
		"_55":           me._55,
		"_56":           me._56,
		"_57":           me._57,
		"_58":           me._58,
		"_59":           me._59,
		"_6":            me._6,
		"_60":           me._60,
		"_61":           me._61,
		"_62":           me._62,
		"_63":           me._63,
		"_64":           me._64,
		"_65":           me._65,
		"_66":           me._66,
		"_67":           me._67,
		"_68":           me._68,
		"_69":           me._69,
		"_7":            me._7,
		"_70":           me._70,
		"_8":            me._8,
		"_9":            me._9,
	})
}

func (me *Settings) UnmarshalHCL(decoder hcl.Decoder) error {
	return decoder.DecodeAll(map[string]any{
		"host_group_id": &me.HostGroupID,
		"_10":           &me._10,
		"_11":           &me._11,
		"_12":           &me._12,
		"_13":           &me._13,
		"_14":           &me._14,
		"_16":           &me._16,
		"_17":           &me._17,
		"_18":           &me._18,
		"_19":           &me._19,
		"_2":            &me._2,
		"_20":           &me._20,
		"_21":           &me._21,
		"_22":           &me._22,
		"_23":           &me._23,
		"_24":           &me._24,
		"_25":           &me._25,
		"_26":           &me._26,
		"_27":           &me._27,
		"_28":           &me._28,
		"_29":           &me._29,
		"_3":            &me._3,
		"_32":           &me._32,
		"_33":           &me._33,
		"_34":           &me._34,
		"_35":           &me._35,
		"_36":           &me._36,
		"_37":           &me._37,
		"_38":           &me._38,
		"_39":           &me._39,
		"_4":            &me._4,
		"_40":           &me._40,
		"_41":           &me._41,
		"_43":           &me._43,
		"_44":           &me._44,
		"_45":           &me._45,
		"_46":           &me._46,
		"_47":           &me._47,
		"_48":           &me._48,
		"_49":           &me._49,
		"_5":            &me._5,
		"_50":           &me._50,
		"_51":           &me._51,
		"_52":           &me._52,
		"_53":           &me._53,
		"_54":           &me._54,
		"_55":           &me._55,
		"_56":           &me._56,
		"_57":           &me._57,
		"_58":           &me._58,
		"_59":           &me._59,
		"_6":            &me._6,
		"_60":           &me._60,
		"_61":           &me._61,
		"_62":           &me._62,
		"_63":           &me._63,
		"_64":           &me._64,
		"_65":           &me._65,
		"_66":           &me._66,
		"_67":           &me._67,
		"_68":           &me._68,
		"_69":           &me._69,
		"_7":            &me._7,
		"_70":           &me._70,
		"_8":            &me._8,
		"_9":            &me._9,
	})
}
