// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transfer

const (
	SecurityPolicyName2018_11      = "TransferSecurityPolicy-2018-11"
	SecurityPolicyName2020_06      = "TransferSecurityPolicy-2020-06"
	SecurityPolicyNameFIPS_2020_06 = "TransferSecurityPolicy-FIPS-2020-06"
	SecurityPolicyName2022_03      = "TransferSecurityPolicy-2022-03"
	SecurityPolicyName2023_05      = "TransferSecurityPolicy-2023-05"
)

func SecurityPolicyName_Values() []string {
	return []string{
		SecurityPolicyName2018_11,
		SecurityPolicyName2020_06,
		SecurityPolicyNameFIPS_2020_06,
		SecurityPolicyName2022_03,
		SecurityPolicyName2023_05,
	}
}
