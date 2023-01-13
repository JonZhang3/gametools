package diff

type ChangeType uint8

const (
	None ChangeType = iota
	Added
	Removed
	ModifiedOld
	ModifiedNew
)

type Difference struct {
	Type        ChangeType
	OldValue    string
	NewValue    string
	NewKeyValue string
	KeyValue    string
}
