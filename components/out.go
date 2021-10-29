package components

import "fmt"

type Val interface {
	GetBool() bool
	GetUint16() uint16
	GetBoolFromUint16(offset uint16) bool
	GetSel() uint8
}

type SingleChan struct {
	val bool
}

func (c *SingleChan) GetBool() bool {
	return c.val
}

func (*SingleChan) GetUint16() uint16 {
	panic("SingleChan provides only booleans")
}

func (*SingleChan) GetBoolFromUint16(uint16) bool {
	panic("SingleChan provides only booleans")
}

func (*SingleChan) GetSel() uint8 {
	panic("SingleChan provides only booleans")
}

type SelectChan struct {
	val uint8
}

func (*SelectChan) GetBool() bool {
	panic("SelectChan provides only uint8s")
}

func (*SelectChan) GetUint16() uint16 {
	panic("SelectChan provides only uint8s")
}

func (*SelectChan) GetBoolFromUint16(uint16) bool {
	panic("SelectChan provides only uint8s")
}

func (s *SelectChan) GetSel() uint8 {
	return s.val
}

type SixteenChan struct {
	val uint16
}

func (*SixteenChan) GetBool() bool {
	panic("SixteenChan provides only uint16s")
}

func (s *SixteenChan) GetUint16() uint16 {
	return s.val
}

func (s *SixteenChan) GetBoolFromUint16(offset uint16) bool {
	if offset > 15 {
		panic(fmt.Sprintf("invalid offset: %d", offset))
	}

	return s.val&(1<<offset) > 0
}

func (*SixteenChan) GetSel() uint8 {
	panic("SelectChan provides only uint16s")
}

type InvalidVal struct{}

func (*InvalidVal) GetBool() bool {
	panic("InvalidVal")
}

func (*InvalidVal) GetUint16() uint16 {
	panic("InvalidVal")
}

func (*InvalidVal) GetBoolFromUint16(uint16) bool {
	panic("InvalidVal")
}

func (*InvalidVal) GetSel() uint8 {
	panic("InvalidVal")
}

type EmptyVal struct{}

func (*EmptyVal) GetBool() bool {
	return false
}

func (*EmptyVal) GetUint16() uint16 {
	return 0
}

func (*EmptyVal) GetBoolFromUint16(uint16) bool {
	return false
}

func (*EmptyVal) GetSel() uint8 {
	return 0
}

type Out interface {
	Update(...UpdateOpts) Val
}

type UpdateOpts struct {
	target Target
	val    Val
}
