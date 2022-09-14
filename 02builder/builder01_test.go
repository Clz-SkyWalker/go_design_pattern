package main

import (
	"fmt"
	"testing"
)

type Profile struct {
	id   int
	name string
	age  int
}

// builder 1
func NewProfileBuilder() ProfileBuilder {
	return ProfileBuilder{profile: Profile{}}
}

type ProfileBuilder struct {
	profile Profile
}

func (b ProfileBuilder) WithId(id int) ProfileBuilder {
	b.profile.id = id
	return b
}
func (b ProfileBuilder) WithName(name string) ProfileBuilder {
	b.profile.name = name
	return b
}
func (b ProfileBuilder) WithAge(age int) ProfileBuilder {
	b.profile.age = age
	return b
}

func (b ProfileBuilder) build() Profile {
	return b.profile
}

func TestBuilder1(t *testing.T) {
	builder := NewProfileBuilder()
	fmt.Println(builder.WithId(1).WithAge(2).WithName("t").build().age)
}

// builder 2
type ProfileOption func(*Profile)

func NewBuilder2(options ...ProfileOption) Profile {
	p := &Profile{}
	for _, option := range options {
		option(p)
	}
	return *p
}

func SetId(id int) ProfileOption {
	return func(p *Profile) {
		p.id = id
	}
}

func SetName(name string) ProfileOption {
	return func(p *Profile) {
		p.name = name
	}
}

func TestBuilder2(t *testing.T) {
	fmt.Println(NewBuilder2(SetId(1), SetName("t")).name)
}

// builder 3 fluent api
type (
	IdBuilder interface {
		SetId(id int) NameBuilder
	}
	NameBuilder interface {
		SetName(name string) EndBuilder
	}
	EndBuilder interface {
		Build() Profile
	}

	BuilderStruct struct {
		p Profile
	}
)

func (b BuilderStruct) SetId(id int) NameBuilder {
	b.p.id = id
	return b
}

func (b BuilderStruct) SetName(name string) EndBuilder {
	b.p.name = name
	return b
}

func (b BuilderStruct) Build() Profile {
	return b.p
}

func NewBuilder3() IdBuilder {
	return BuilderStruct{}
}

func TestBuilder3(t *testing.T) {
	fmt.Println(NewBuilder3().SetId(1).SetName("t").Build().name)
}
