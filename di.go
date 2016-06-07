/*
<!--
Copyright (c) 2016 Christoph Berger. Some rights reserved.
Use of this text is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

+++
title = ""
description = ""
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2016-00-00"
publishdate = "2016-00-00"
domains = [""]
tags = ["", "", ""]
categories = ["Tutorial"]
+++

### Summary goes here

<!--more-->

## Intro goes here

## The code
*/

// ## Imports and globals
package main

import "math"

var (
	l1, l2, a1, a2, b1, b2, d float64
)

func beta1(x, y float64) float64 {
	return math.Atan2(y, x)
}

func beta2(d, l1, l2 float64) float64 {
	return (d*d + l2*l2 - l1*l1) / (2 * l2 * d)
}

func alpha2(d, l1, l2 float64) float64 {
	return (d*d - l2*l2 + l1*l1) / (2 * d * l1)
}
