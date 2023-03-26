package common

// TODO: josh add all serialization methods and deserialization method
func NewSerializeGetFlightIdBySourceDest(msgId, source, dest string) []byte {
	return nil
}

func Deserialize(b []byte) map[string]string {
	// Extract request type from byte arr, and return the information in a map[string]string
	// There should be a request type and direction type, request type refers to which command, direction refers to
	// a request or response, since this method can be called by client or server
	return nil
}
