package models

import (
	"errors"
	"read-more-backend/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Entry struct {
	ID            uuid.UUID   `json:"id" gorm:"primaryKey,type:uuid"`
	Title         string      `json:"title" gorm:"not null"`
	Description   string      `json:"description" gorm:"not null"`
	URL           string      `json:"url" gorm:"not null"`
	Transcription string      `json:"transcription" gorm:"not null"`
	AudioFilename string      `json:"audioFilename" gorm:"not null"`
	Seen          bool        `json:"seen" gorm:"not null"`
	CollectionID  uuid.UUID   `json:"collectionId" gorm:"not null"`
	Collection    *Collection `json:"collection,omitempty"`
}

func (e *Entry) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return err
}

type CreateEntryDto struct {
	// Must be provided
	Title string `json:"title"`
	// Must be provided
	Transcription string `json:"transcription"`

	// Optional
	Description string `json:"description"`
	// Optional
	URL string `json:"url"`
	// Optional
	AudioFilename string `json:"audioFilename"`

	// CollectionID or CollectionTitle must be provided
	CollectionID string `json:"collectionId"`
	// CollectionID or CollectionTitle must be provided
	CollectionTitle string `json:"collectionTitle"`
}

func CreateEntryFromDto(dto CreateEntryDto, collectionId uuid.UUID) *Entry {
	return &Entry{
		Title:         dto.Title,
		Description:   dto.Description,
		URL:           dto.URL,
		Transcription: dto.Transcription,
		AudioFilename: dto.AudioFilename,
		Seen:          false,
		CollectionID:  collectionId,
	}
}

func (e *Entry) CreateOne(db *gorm.DB, collectionTitle string) (err error) {
	return db.Transaction(func(tx *gorm.DB) (err error) {
		if e.CollectionID == uuid.Nil {
			if utils.IsEmpty(collectionTitle) {
				return errors.New("both (collectionId) and (collectionTitle) were not provided")
			}
			newCollection := &Collection{
				Title: collectionTitle,
			}
			err = newCollection.CreateOne(tx)
			if err != nil {
				return err
			}
			e.CollectionID = newCollection.ID
		}

		err = tx.Create(e).Error

		return err
	})
}

func (e *Entry) FindOne(db *gorm.DB, with *Entry) (err error) {
	err = db.Model(&Entry{}).Preload(clause.Associations).First(e, with).Error
	return err
}

func (e *Entry) UpdateOne(db *gorm.DB, with *Entry) (err error) {
	err = db.Model(e).Updates(map[string]interface{}{
		"title":          with.Title,
		"description":    with.Description,
		"url":            with.URL,
		"transcription":  with.Transcription,
		"audio_filename": with.AudioFilename,
		"seen":           with.Seen,
		"collection_id":  with.CollectionID,
	}).Error
	return err
}

func (e *Entry) DeleteOne(tx *gorm.DB) (err error) {
	err = tx.Delete(e).Error
	return err
}
