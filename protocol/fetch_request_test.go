package protocol

import "testing"

var (
	fetchRequestNoBlocks = []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00}

	fetchRequestWithProperties = []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0xEF,
		0x00, 0x00, 0x00, 0x00}

	fetchRequestOneBlock = []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x05, 't', 'o', 'p', 'i', 'c',
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x12, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x34, 0x00, 0x00, 0x00, 0x56}
)

func TestFetchRequest(t *testing.T) {
	request := new(FetchRequest)
	testEncodable(t, "no blocks", request, fetchRequestNoBlocks)

	request.MaxWaitTime = 0x20
	request.MinBytes = 0xEF
	testEncodable(t, "with properties", request, fetchRequestWithProperties)

	request.MaxWaitTime = 0
	request.MinBytes = 0
	request.AddBlock("topic", 0x12, 0x34, 0x56)
	testEncodable(t, "one block", request, fetchRequestOneBlock)
}