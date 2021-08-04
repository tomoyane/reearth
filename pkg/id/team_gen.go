// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// TeamID is an ID for Team.
type TeamID ID

// NewTeamID generates a new TeamId.
func NewTeamID() TeamID {
	return TeamID(New())
}

// TeamIDFrom generates a new TeamID from a string.
func TeamIDFrom(i string) (nid TeamID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = TeamID(did)
	return
}

// MustTeamID generates a new TeamID from a string, but panics if the string cannot be parsed.
func MustTeamID(i string) TeamID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return TeamID(did)
}

// TeamIDFromRef generates a new TeamID from a string ref.
func TeamIDFromRef(i *string) *TeamID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := TeamID(*did)
	return &nid
}

// TeamIDFromRefID generates a new TeamID from a ref of a generic ID.
func TeamIDFromRefID(i *ID) *TeamID {
	if i == nil {
		return nil
	}
	nid := TeamID(*i)
	return &nid
}

// ID returns a domain ID.
func (d TeamID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d TeamID) String() string {
	return ID(d).String()
}

// GoString implements fmt.GoStringer interface.
func (d TeamID) GoString() string {
	return "id.TeamID(" + d.String() + ")"
}

// RefString returns a reference of string representation.
func (d TeamID) RefString() *string {
	id := ID(d).String()
	return &id
}

// Ref returns a reference.
func (d TeamID) Ref() *TeamID {
	d2 := d
	return &d2
}

// Contains returns whether the id is contained in the slice.
func (d TeamID) Contains(ids []TeamID) bool {
	for _, i := range ids {
		if d.ID().Equal(i.ID()) {
			return true
		}
	}
	return false
}

// CopyRef returns a copy of a reference.
func (d *TeamID) CopyRef() *TeamID {
	if d == nil {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *TeamID) IDRef() *ID {
	if d == nil {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *TeamID) StringRef() *string {
	if d == nil {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *TeamID) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *TeamID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = TeamIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *TeamID) MarshalText() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *TeamID) UnmarshalText(text []byte) (err error) {
	*d, err = TeamIDFrom(string(text))
	return
}

// Ref returns true if a ID is nil or zero-value
func (d TeamID) IsNil() bool {
	return ID(d).IsNil()
}

// TeamIDToKeys converts IDs into a string slice.
func TeamIDToKeys(ids []TeamID) []string {
	keys := make([]string, 0, len(ids))
	for _, i := range ids {
		keys = append(keys, i.String())
	}
	return keys
}

// TeamIDsFrom converts a string slice into a ID slice.
func TeamIDsFrom(ids []string) ([]TeamID, error) {
	dids := make([]TeamID, 0, len(ids))
	for _, i := range ids {
		did, err := TeamIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// TeamIDsFromID converts a generic ID slice into a ID slice.
func TeamIDsFromID(ids []ID) []TeamID {
	dids := make([]TeamID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, TeamID(i))
	}
	return dids
}

// TeamIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func TeamIDsFromIDRef(ids []*ID) []TeamID {
	dids := make([]TeamID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, TeamID(*i))
		}
	}
	return dids
}

// TeamIDsToID converts a ID slice into a generic ID slice.
func TeamIDsToID(ids []TeamID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// TeamIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func TeamIDsToIDRef(ids []*TeamID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// TeamIDSet represents a set of TeamIDs
type TeamIDSet struct {
	m map[TeamID]struct{}
	s []TeamID
}

// NewTeamIDSet creates a new TeamIDSet
func NewTeamIDSet() *TeamIDSet {
	return &TeamIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *TeamIDSet) Add(p ...TeamID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[TeamID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []TeamID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *TeamIDSet) AddRef(p *TeamID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *TeamIDSet) Has(p TeamID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *TeamIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *TeamIDSet) All() []TeamID {
	if s == nil {
		return nil
	}
	return append([]TeamID{}, s.s...)
}

// Clone returns a cloned set
func (s *TeamIDSet) Clone() *TeamIDSet {
	if s == nil {
		return NewTeamIDSet()
	}
	s2 := NewTeamIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *TeamIDSet) Merge(s2 *TeamIDSet) *TeamIDSet {
	if s == nil {
		return nil
	}
	s3 := s.Clone()
	if s2 == nil {
		return s3
	}
	s3.Add(s2.s...)
	return s3
}
