package main

import (
	"fmt"
	"unsafe"
)

type ModbusHeader struct {
	TransactionID uint16
	ProtocolID    uint16
	Length        uint16
	Address       uint8
	FunctionCode  uint8
	ByteCount     uint8
}

var (
	modbusFrame = ModbusHeader{
		TransactionID: 0,
		ProtocolID:    0,
		Length:        5,
		Address:       5,
		FunctionCode:  3,
		ByteCount:     3,
	}
	registers = []uint16{255, 255, 0}
)

func getBytes(pdu ModbusHeader, registers []uint16) []byte {
	if len(registers) < 1 {
		return nil
	}
	headerSize := unsafe.Offsetof(pdu.ByteCount) + unsafe.Sizeof(pdu.ByteCount)
	addrH := (*byte)(unsafe.Pointer(&pdu))
	bytesHeader := unsafe.Slice(addrH, headerSize)

	addrR := (*byte)(unsafe.Pointer(&registers[0]))
	bytesRegisters := unsafe.Slice(addrR, pdu.ByteCount*uint8(unsafe.Sizeof(registers[0])))
	return append(bytesHeader, bytesRegisters...)
}

func main() {
	bytes := getBytes(modbusFrame, registers)
	fmt.Printf("%d", bytes)
}
