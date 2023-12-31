# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "aws_vpc" "main" {
  cidr_block = cidrsubnet(var.base_cidr_block, 4, var.region_numbers[var.region])
}

resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id
}
