package cbor2map

// Parses a CBOR array and store the result to the buffer.
//
// # Arguments
//   - (the buffer) *[]any: The buffer to store the parse result.
type CborToArray func(*[]any) error
