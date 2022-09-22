package main

import (
	"fmt"
	"testing"
)

type Robot struct {
	Head string `json:"head"`
	Body string `json:"body"`
	Hand string `json:"hand"`
	Foot string `json:"foot"`
}

// builder 1
type IRobot interface {
	SetHead()
	SetBody()
	SetHand()
	SetFoot()
	Build() string
}

type GunDam struct {
	robot Robot
}

func (m *GunDam) SetHead() {
	m.robot.Head = "GunDam"
}

func (m *GunDam) SetBody() {
	m.robot.Body = "body"
}

func (m *GunDam) SetHand() {
	m.robot.Hand = "hand"
}

func (m *GunDam) SetFoot() {
	m.robot.Foot = "foot"
}
func (m *GunDam) Build() string {
	return fmt.Sprintf("%+v", m.robot)
}

type DaBanModel struct {
	robot Robot
}

func (m *DaBanModel) SetHead() {
	m.robot.Head = "daBan"
}

func (m *DaBanModel) SetBody() {
	m.robot.Body = "body"
}

func (m *DaBanModel) SetHand() {
	m.robot.Hand = "hand"
}

func (m *DaBanModel) SetFoot() {
	m.robot.Foot = "foot"
}
func (m *DaBanModel) Build() string {
	return fmt.Sprintf("%+v", m.robot)
}

func NewPlayer(builder IRobot) Player {
	return Player{
		builder: builder,
	}
}

type Player struct {
	builder IRobot
}

func (p Player) MakeFast() string {
	p.builder.SetHead()
	p.builder.SetBody()
	return p.builder.Build()
}

func (p Player) MakeSlow() string {
	p.builder.SetHead()
	p.builder.SetBody()
	p.builder.SetHand()
	p.builder.SetFoot()
	return p.builder.Build()
}

func TestBuilder1(t *testing.T) {
	var robot IRobot
	robot = new(GunDam)
	player := NewPlayer(robot)
	fmt.Println(player.MakeFast())
	robot = new(DaBanModel)
	player = NewPlayer(robot)
	fmt.Println(player.MakeSlow())
}

// builder2 functional option
type RobotOption func(*Robot)

func SetHeadOption(head string) RobotOption {
	return func(robot *Robot) {
		robot.Head = head
	}
}

func SetBodyOption(body string) RobotOption {
	return func(robot *Robot) {
		robot.Body = body
	}
}
func SetHandOption(hand string) RobotOption {
	return func(robot *Robot) {
		robot.Hand = hand
	}
}
func SetFootOption(foot string) RobotOption {
	return func(robot *Robot) {
		robot.Foot = foot
	}
}

func NewRobotBuilder2(options ...RobotOption) string {
	robot := new(Robot)
	for _, option := range options {
		option(robot)
	}
	return fmt.Sprintf("%+v", robot)
}

func TestBuilder2FunctionalOption(t *testing.T) {
	option1 := SetHeadOption("functional option")
	option2 := SetBodyOption("body")
	option3 := SetHandOption("hand")
	option4 := SetFootOption("foot")
	result := NewRobotBuilder2(option1, option2, option3, option4)
	fmt.Println(result)
}

type (
	IHeader interface {
		SetHead(string) IBody
	}
	IBody interface {
		SetBody(string) IHand
	}
	IHand interface {
		SetHand(string) IFoot
	}
	IFoot interface {
		SetFoot(string) IBuild
	}
	IBuild interface {
		Build3Robot() string
	}
	Build3 struct {
		Robot
	}
)

func (b *Build3) SetHead(p string) IBody {
	b.Hand = p
	return b
}
func (b *Build3) SetBody(p string) IHand {
	b.Body = p
	return b
}
func (b *Build3) SetHand(p string) IFoot {
	b.Hand = p
	return b
}
func (b *Build3) SetFoot(p string) IBuild {
	b.Foot = p
	return b
}
func (b *Build3) Build3Robot() string {
	return fmt.Sprintf("%+v", b)
}

func NewBuild3() IHeader {
	return &Build3{}
}

func TestBuild3(t *testing.T) {
	result := NewBuild3().SetHead("da ban").SetBody("body").SetHand("hand").SetFoot("foot").Build3Robot()
	fmt.Println(result)
}
