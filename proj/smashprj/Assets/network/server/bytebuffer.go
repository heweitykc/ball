package main

import(
	"encoding/binary"
)

const (
	presize = 0
	initsize = 1024
)

type ByteBuffer struct{
	_buffer		[]byte
	_prependSize int
	_readerIndex   int
	_writerIndex   int
}

func NewByteBuffer() *ByteBuffer{
	return &ByteBuffer{
		_buffer:	make([]byte, initsize),
		_prependSize: presize,
		_readerIndex: presize,
		_writerIndex: presize,
	}
}

func (self *ByteBuffer) Append(buf ...byte){
	size := len(buf)
	if size == 0{
		return
	}
	if self.WrSize() < size {
		self.WrReserve(size)
	}
	copy(self._buffer[self._writerIndex:], buf)
	self.WrFlip(size)
}

func (self *ByteBuffer) WrBuffer() []byte {
	if self._writerIndex >= len(self._buffer){
		return nil
	}
	return self._buffer[self._writerIndex:]
}

func (self *ByteBuffer) WrSize() int{
	return len(self._buffer) - self._writerIndex
}

func (self *ByteBuffer) WrFlip(size int){
	self._writerIndex += size
}

func (self *ByteBuffer) RdBuffer() []byte{
	if self._readerIndex >= len(self._buffer){
		return nil
	}	
	return self._buffer[self._readerIndex:]
}

func (self *ByteBuffer) RdReady() bool{
	return self._writerIndex > self._readerIndex
}

func (self *ByteBuffer) RdSize() int{
	return self._writerIndex - self._readerIndex
}

func (self *ByteBuffer) RdFlip(size int){
	if size < self.RdSize(){
		self._readerIndex += size
	} else {
		self.Reset()
	}
}

func (self *ByteBuffer) Reset(){
	self._readerIndex = self._prependSize
	self._writerIndex = self._prependSize
}

func (self *ByteBuffer) MaxSize() int{
	return len(self._buffer)
}
func (self *ByteBuffer) WrReserve(size int) {
	if self.WrSize()+self._readerIndex < size+self._prependSize {
		tmpbuf := make([]byte, self._writerIndex+size)
		copy(tmpbuf, self._buffer)
		self._buffer = tmpbuf
	} else {
		readable := self.RdSize()
		copy(self._buffer[self._prependSize:], self._buffer[self._readerIndex:self._writerIndex])
		self._readerIndex = self._prependSize
		self._writerIndex = self._readerIndex + readable
	}
}

func (self *ByteBuffer) Prepend(buf ...byte) bool{
	size := len(buf)
	if size == 0 {
		return false
	}
	if self._readerIndex < size{
		return false
	}
	self._readerIndex -= size
	copy(self._buffer[self._readerIndex:], buf)
	return true
}

func (self *ByteBuffer) Slice(n int) []byte{
	if n > self.RdSize(){
		return make([]byte, n)
	}
	r := self.RdBuffer()[:n]
	self.RdFlip(n)
	return r
}

func (self *ByteBuffer) ReadUint8() uint8{
	return uint8(self.Slice(1)[0])
}

func (self *ByteBuffer) ReadUint16LE() uint16{
	return binary.LittleEndian.Uint16(self.Slice(2))
}

func (self *ByteBuffer) ReadUint16BE() uint16{
	return binary.LittleEndian.Uint16(self.Slice(2))
}

func (self *ByteBuffer) ReadUint32LE() uint32{
	return binary.LittleEndian.Uint32(self.Slice(4))
}

func (self *ByteBuffer) ReadUint32BE() uint32{
	return binary.BigEndian.Uint32(self.Slice(4))
}

func (self *ByteBuffer) ReadUint64LE() uint64{
	return binary.LittleEndian.Uint64(self.Slice(8))
}

func (self *ByteBuffer) ReadUint64BE() uint64{
	return binary.BigEndian.Uint64(self.Slice(8))
}

func (self *ByteBuffer) WriteUint8(v uint8) {
	self.Append(byte(v))
}

func (self *ByteBuffer) WriteUint16LE(v uint16) {
	self.Append(byte(v), byte(v>>8))
}

func (self *ByteBuffer) WriteUint16BE(v uint16) {
	self.Append(byte(v>>8), byte(v))
}

func (self *ByteBuffer) WriteUint32LE(v uint32) {
	self.Append(byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

func (self *ByteBuffer) WriteUint32BE(v uint32) {
	self.Append(byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func (self *ByteBuffer) WriteUint64LE(v uint64) {
	self.Append(byte(v), byte(v>>8), byte(v>>16), byte(v>>24), byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
}

func (self *ByteBuffer) WriteUint64BE(v uint64) {
	self.Append(byte(v>>56), byte(v>>48), byte(v>>40), byte(v>>32), byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}




