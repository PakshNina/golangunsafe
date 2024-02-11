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

func main() {
	receivedBytes := []byte{0x00, 0x00, 0x01, 0x00, 0x09, 0x00, 0x00, 0x03, 0x06, 0x2B, 0x02, 0x64, 0x00, 0x7F, 0x00}

	header, modbusRegisters := getModbusPDU(receivedBytes)
	fmt.Printf("Header: %d", header)
	fmt.Printf("Registers: %#x", modbusRegisters)
}

func getModbusPDU(bytes []byte) (ModbusHeader, []uint16) {
	// Header
	header := *(*ModbusHeader)(unsafe.Pointer(&bytes[0]))

	// Registers
	offset := unsafe.Offsetof(header.ByteCount) + unsafe.Sizeof(header.ByteCount)
	regAddr := unsafe.Pointer(uintptr(unsafe.Pointer(&bytes[0])) + offset)
	bytesAddr := (*uint16)(regAddr)
	registers := unsafe.Slice(bytesAddr, int(header.ByteCount/2))
	return header, registers
}
