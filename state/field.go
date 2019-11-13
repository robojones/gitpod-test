package state

import "time"

func newField() *field {
	return &field{
		timestamp: time.Time{},
	}
}

// field is used to attach a timestamp to a field.
type field struct {
	timestamp time.Time
}

func (f *field) Timestamp() time.Time {
	return f.timestamp
}

// shouldUpdate returns true if the field shoud be affected by the update.
func (f *field) shouldUpdate(updateTime time.Time) bool {
	return updateTime.After(f.timestamp)
}

// updateTimestamp is used to update the time after a change was made to the field.
func (f *field) updateTimestamp(updateTime time.Time) {
	f.timestamp = updateTime
}
