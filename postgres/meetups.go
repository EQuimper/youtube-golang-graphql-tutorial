package postgres

import (
	"github.com/go-pg/pg/v9"

	"github.com/equimper/meetmeup/models"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Order("id").Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	return meetup, err
}

func (m *MeetupsRepo) GetByID(id string) (*models.Meetup, error) {
	var meetup models.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	return &meetup, err
}

func (m *MeetupsRepo) Update(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Update()
	return meetup, err
}

func (m *MeetupsRepo) Delete(meetup *models.Meetup) error {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Delete()
	return err
}

func (m *MeetupsRepo) GetMeetupsForUser(user *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	return meetups, err
}
