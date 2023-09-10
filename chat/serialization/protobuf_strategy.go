package serialization

import "google.golang.org/protobuf/proto"

type ProtobufStrategy struct{}

func (p *ProtobufStrategy) Serialize(data interface{}) ([]byte, error) {
	return proto.Marshal(data.(proto.Message))
}

func (p *ProtobufStrategy) Deserialize(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
