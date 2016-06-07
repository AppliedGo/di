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

Software architecture strives to provide structure to software systems, in order to make them robust, maintainable, extendable, testable, easier to develop, and easier to document.

Many different [architecture styles and patterns](https://en.wikipedia.org/wiki/List_of_software_architecture_styles_and_patterns) have evolved over time, at different abstraction levels and with different levels of complexity. Among them, layered architectures seem to represent a fairly versatile concept, applicable to a large range of scenarios.

## The Clean Architecture

At the center of any layered software architecture is the separation of concerns. In simple words: The less each layer knows about the other layers, the better.

To achieve this, the layers form a hierarchy of abstraction levels. The Clean Architecture model describes them as (at least) four circles:

* The innermost circle contains *Entities* that implement *Enterprise Business Rules*.
* The second circle contains *Use Cases* that form the *Application Business Rules*.
* The third circle contains all *Interface Adapters*, like: *Controllers*, *Gateways*, or *Presenters*.
* The fourth circle contains the technical *Frameworks & Drivers*, like: *UI*, *Databases*, *External Interfaces*, etc.

![The Clean Architecture](cleanarchitecture.png)

Read more about The Clean Architecture [here](http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html).

The central rule of The Clean Architecture is the **Dependency Rule**.

But then, how can information flow between the layers without affecting the separation of concerns too much?





The Clean Architecture model consists of four layers (or more, if you need it) of increasing abstraction.



The layers are laid out as circles. The innermost circle


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
