package serialization

type SerializationStrategy interface {
	Serialize(any) ([]byte, error)
	Deserialize([]byte, any) error
}
