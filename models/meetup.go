package models

type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}

func (m *Meetup) IsOwner(user *User) bool {
	return m.UserID == user.ID
}
