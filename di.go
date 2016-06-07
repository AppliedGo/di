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
domains = ["Architecture"]
tags = ["Dependency Injection", "Separation Of Concerns", "Interface"]
categories = ["Tutorial"]
+++

Layered software architectures adhere to the Dependency Rule: Source code dependencies can only point from lower-level layers to higher-level layers, and never vice versa. But at some point you need to connect the lower layers to the higher layers without violating the rule. Here is how.

<!--more-->

## Software Architecture

The nightmare of every software maintainer: After fixing a bug in a small utility function from the PDF formatter module, database queries start returning wrong results. After long and laborious search, it turns out that the database code uses the PDF utility function for formatting some strings.

Sounds familiar?

Bugs of this kind have inspired software engineers to think up ways of strucring software that effectively prohibits uncontrolled proliferation of mutual dependencies between the entities of a software system. The result is called a software architecture.

Many different [architecture styles and patterns](https://en.wikipedia.org/wiki/List_of_software_architecture_styles_and_patterns) have evolved over time, at different abstraction levels and with different levels of complexity. Among them, layered architectures seem to represent a fairly versatile concept, applicable to a large range of scenarios.


## The Clean Architecture

An excellent summary of the principles of good layered architectures is [The Clean Architecture](http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html), a blog article about the architecture of the same name. The Clean Architecture model consists of four layers (or more, if you need it) of increasing abstraction.

![The Clean Architecture](cleanarchitecture.png)

The layers are laid out as circles. The innermost circle

At the center of any software architecture is the separation of concerns.

## The code

*/

// ## Imports and globals
package main

import (
	"fmt"
	"strings"
)

type Speaker interface {
	Speak(string)
}

type Announcement struct {
	message string
	speaker Speaker
}

func NewAnnouncement(m string, s Speaker) *Announcement {
	return &Announcement{m, s}
}

func (a *Announcement) SetSpeaker(s Speaker) {
	a.speaker = s
}
func (a *Announcement) Announce() {
	a.speaker.Speak(a.message)
}

type Newscaster struct{}

func (n *Newscaster) Speak(msg string) {
	fmt.Println("Breaking news:")
	fmt.Println(msg)
	fmt.Println()
}

type Preacher struct{}

func (p *Preacher) Speak(msg string) {
	fmt.Println("And so the prophets say:")
	fmt.Println(msg)
	fmt.Println()
}

type SalesPromoter struct{}

func (s *SalesPromoter) Speak(msg string) {
	fmt.Println(strings.ToUpper(msg), "!!! BUY NOW!!!")
	fmt.Println()
}

func main() {
	n := &Newscaster{}
	p := &Preacher{}
	s := &SalesPromoter{}

	a := NewAnnouncement("Gophers are incredibly smart", n)
	a.Announce()

	a.SetSpeaker(p)
	a.Announce()

	a.SetSpeaker(s)
	a.Announce()
}
