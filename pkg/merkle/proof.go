package merkle

type Proof struct {
	Hash string
	// BatchId is the proof's batch ID
	BatchId string
	// AnchorType is the proof's anchor type
	// AnchorType anchor.Anchor_Type
	// // BatchStatus is the proof's batch status
	// BatchStatus anchor.Batch_Status
	// // Format is the proof format
	// Format anchor.Proof_Format
	// // Data is the proof data in base64
	// Data string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// // Batch is the proof's batch detail
	// Batch *anchor.Batch `protobuf:"bytes,7,opt,name=batch,proto3" json:"batch,omitempty"`
}
