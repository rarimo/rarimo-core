package types

func (m *Group) HasMember(member string) bool {
	for _, m := range m.Members {
		if m == member {
			return true
		}
	}
	return false
}
