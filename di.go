/*
<!--
Copyright (c) 2016 Christoph Berger. Some rights reserved.
Use of this text is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

+++
title = "Dependency Injection in a nutshell"
description = "How to implement the Dependency Rule in Go"
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2016-06-23"
publishdate = "2016-06-23"
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

To achieve this, the layers form a hierarchy of abstraction levels. The Clean Architecture model describes them as (at least) four concentric circles that represent different abstraction levels.

![The Clean Architecture](cleanarchitecture.png)

Read more about The Clean Architecture [here](http://blog.8thlight.com/uncle-bob/2012/08/13/the-clean-architecture.html).

The central rule of The Clean Architecture is the **Dependency Rule**. The rule says that the source code of each circle can only access code in an inner circle but never any code in an outer circle. The code must not know anything about the outer circles. In short,

> Source code dependencies can only point inwards.

But then, how can information flow from the outside into the inner rings?

Imagine a common scenario:

* A user clicks a button in the UI to search for a document.
* The UI passes the action to a controller,
* the controller triggers the appropriate use case,
* the use case needs to identify the right document entity.
* The Entity needs to load the required data.
* The use case for loading that data determines that it must fetch the data from a database adapter.
* The database adapter connects the database and runs a query.
* It passes the result back to the use case,
* which delivers it to the document entity.
* The entity forwards the content to the requesting use case,
* which passes it on to a presenter,
* which finally updates the UI.

HYPE[Action flow through the layers](actionflow.html)

This action flow passes through the layers inwards and outwards.This action flow passes through the layers inwards and outwards.

How can we accomplish this if the inner rings do not know anything about the outer rings?

The solution is dependency injection.

## Dependency Injection demystified

You might already have heard about dependency injection (DI), probably in the context of (non-Go) frameworks like Spring, Guice, etc. Such frameworks might subtly give the impression that DI is not possible without the help of some opaque mechanism that injects the dependencies "for you".

One of the many good things about Go is that it is largely free of any frameworks. This allows us taking a pragmatic look at DI. And with Go, we can do that right by going through actual code.

I do not build a complete Clean Architecture example here. Rather, I want to focus on DI itself: How to pass actions or information inwards, and how to pass them outwards, without bothering the higher levels with internals of the "lesser levels".

Hence the example uses only two rings. The inner ring contains an "Announcement" entity. It provides a method for setting the message, and it knows of a "Speaker" that is able to "Speak()" the message.

In the outer ring there are two other entities, a "Messenger" that delivers a message to the Announcement entity, and three different entities that take the role of the Speaker, to take and announce a message from an Announcement.

Silly enough, but we want to focus on the DI mechanism anyway.

## The code

*/

// ## Imports and globals
package main

import (
	"fmt"
	"strings"
)

// ### The "inner ring"
//
// An Announcement contains a message and a "speaker".
type Announcement struct {
	message string
	speaker Speaker
}

// The Speaker is just an interface that defines one method, `Speak()`.
type Speaker interface {
	Speak(string)
}

// Set a new Speaker.
func (a *Announcement) SetSpeaker(s Speaker) {
	a.speaker = s
}

// Get a message delivered from the outside.
func (a *Announcement) Deliver(m string) {
	a.message = m
}

// Announce a message to the outside.
//
// Note that up to this point, there are **no dependencies on the outer layer**.
func (a *Announcement) Announce() {
	a.speaker.Speak(a.message)
}

// ## The outer ring
//
// A newscaster.
type Newscaster struct{}

// The newscaster implements the Speak method and therefore can act as a speaker for an Announcement.
func (n *Newscaster) Speak(msg string) {
	fmt.Println("Breaking news:", msg)
}

// Another speaker.
type Preacher struct{}

func (p *Preacher) Speak(msg string) {
	fmt.Println("And so the prophets say:", msg)
}

// And another.
//
// The inner ring does not have any knowledge about these speakers.
type SalesPromoter struct{}

func (s *SalesPromoter) Speak(msg string) {
	fmt.Println(strings.ToUpper(msg), "!!! BUY NOW!!!")
	fmt.Println()
}

// This is the messenger. It knows about the Announcement entity, but this
// is an inward dependency and hence perfectly fine.
type Messenger struct {
	message string
	ann     *Announcement
}

func NewMessenger(m string, a *Announcement) *Messenger {
	return &Messenger{m, a}
}

// The Messenger passes the message to the Announcement entity.
// This is how information flows inward.
func (m *Messenger) Deliver() {
	m.ann.Deliver(m.message)
}

// In main, everything is put into place. Here,
// the actual objects are created, and the dependencies get injected.
func main() {
	// Create the speakers.
	n := &Newscaster{}
	p := &Preacher{}
	s := &SalesPromoter{}

	// Create announcment and messenger.
	a := &Announcement{}
	m := NewMessenger("Gophers are really smart", a)
	m.Deliver()

	// And here we inject a dependency. Although `a` does not know
	// any types, functions, or variables from the outer ring, it now has a Speaker.
	a.SetSpeaker(n)

	// `a` can now announce its message without knowing the speaker at all.
	a.Announce()

	// Now the speaker changes, and `a` does not need to care a bit.
	a.SetSpeaker(p)
	a.Announce()

	// And another speaker.
	a.SetSpeaker(s)
	a.Announce()
}

/* Thanks to interfaces, the Announcement entity does not need to know anything about the different speaker types.

And passing information inwards is outright trivial, once the outer ring has access to the objects of the inner ring.

A question to you: Can we inject dependencies without using interfaces?


## Further reading

The excellent article [Applying The Clean Architecture to Go applications](http://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/) not only goes deep into DI but also works through all four layers of the Clean Architecture. This is a great opportunity to see how entities, use cases, interfaces, and frameworks (speaking in Clean Architecture tongue) are used to build a tiny shop system.

*/
