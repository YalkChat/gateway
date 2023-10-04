package serialization

// The chat package provides only the interface, the implementation
// is in the main package. The rest of the code should do the same
// providing interfaces and leaving implementations outside

type SerializationStrategy interface {
	Serialize(any) ([]byte, error)
	Deserialize([]byte, any) error
}
