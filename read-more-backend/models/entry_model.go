package models

import (
	"encoding/json"
	"fmt"
	"log"

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

func (e *Entry) FindOne(db *gorm.DB, with *Entry) (err error) {
	err = db.Model(&Entry{}).Preload(clause.Associations).First(e, with).Error
	return err
}

func (e *Entry) CreateOne(db *gorm.DB, collectionTitle string) (err error) {
	return db.Transaction(func(tx *gorm.DB) (err error) {
		if e.CollectionID == uuid.Nil && len(collectionTitle) >= 1 {
			newCollection := &Collection{
				Title:       collectionTitle,
				Description: collectionTitle,
				URL:         collectionTitle,
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

func (e *Entry) UpdateOne(db *gorm.DB, with *UpdateEntryDto) (err error) {
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

type CreateEntryDto struct {
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	URL             string    `json:"url"`
	Transcription   string    `json:"transcription"`
	AudioFilename   string    `json:"audioFilename"`
	Seen            bool      `json:"seen"`
	CollectionID    uuid.UUID `json:"collectionId"`
	CollectionTitle string    `json:"collectionTitle"`
}

func (dto *CreateEntryDto) UnmarshalJSON(data []byte) error {
	var err error

	type Alias CreateEntryDto
	raw := struct {
		Alias
		CollectionID string `json:"collectionId"`
	}{}

	err = json.Unmarshal(data, &raw)
	if err != nil {
		log.Println("[Error] createEntryDto.(UnmarshalJSON) [0]:", err)
		return err
	}

	var parsedUUID uuid.UUID
	if len(raw.CollectionID) > 0 {
		parsedUUID, err = uuid.Parse(raw.CollectionID)
		if err != nil {
			log.Println("[Error] createEntryDto.(UnmarshalJSON) [1]:", err)
		}
	}

	hasCollectionId := parsedUUID != uuid.Nil
	hasCollectionTitle := len(raw.CollectionTitle) > 0
	continueBool := hasCollectionId || hasCollectionTitle
	if !continueBool {
		return fmt.Errorf("one of the fields (collectionId) or (collectionTitle) must be provided")
	}

	dto.Title = raw.Title
	dto.Description = raw.Description
	dto.URL = raw.URL
	dto.Transcription = raw.Transcription
	dto.AudioFilename = raw.AudioFilename
	dto.Seen = raw.Seen

	if hasCollectionId {
		dto.CollectionID = parsedUUID
	}

	if hasCollectionTitle {
		dto.CollectionTitle = raw.CollectionTitle
	}

	return nil
}

type UpdateEntryDto struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	URL           string    `json:"url"`
	Transcription string    `json:"transcription"`
	AudioFilename string    `json:"audioFilename"`
	Seen          bool      `json:"seen"`
	CollectionID  uuid.UUID `json:"collectionId"`
}

func (dto *UpdateEntryDto) UnmarshalJSON(data []byte) error {
	var err error

	type Alias UpdateEntryDto
	raw := struct {
		Alias
		ID           string `json:"id"`
		CollectionID string `json:"collectionId"`
	}{}

	err = json.Unmarshal(data, &raw)
	if err != nil {
		log.Println("[Error] updateEntryDto.(UnmarshalJSON) [0]:", err)
		return err
	}

	var parsedIdUUID uuid.UUID
	if len(raw.ID) > 0 {
		parsedIdUUID, err = uuid.Parse(raw.ID)
		if err != nil {
			log.Println("[Error] updateEntryDto.(UnmarshalJSON) [1]:", err)
		}
	}

	var parsedCollectionUUID uuid.UUID
	if len(raw.CollectionID) > 0 {
		parsedCollectionUUID, err = uuid.Parse(raw.CollectionID)
		if err != nil {
			log.Println("[Error] updateEntryDto.(UnmarshalJSON) [2]:", err)
		}
	}

	hasId := parsedIdUUID != uuid.Nil
	hasCollectionId := parsedCollectionUUID != uuid.Nil
	continueBool := hasId && hasCollectionId
	if !continueBool {
		return fmt.Errorf("both Entry.(ID) and Entry.(CollectionID) must be provided")
	}

	dto.Title = raw.Title
	dto.Description = raw.Description
	dto.URL = raw.URL
	dto.Transcription = raw.Transcription
	dto.AudioFilename = raw.AudioFilename
	dto.Seen = raw.Seen

	if hasId {
		dto.ID = parsedIdUUID
	}

	if hasCollectionId {
		dto.CollectionID = parsedCollectionUUID
	}

	return nil
}
