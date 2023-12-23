package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Collection struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey,type:uuid"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	URL         string    `json:"url" gorm:"not null"`
	Entries     []*Entry  `json:"entries,omitempty"`
}

func (c *Collection) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return err
}

func (c *Collection) FindOne(db *gorm.DB, with *Collection) (err error) {
	err = db.Model(&Collection{}).Preload(clause.Associations).First(c, with).Error
	return err
}

func (c *Collection) FindAll(db *gorm.DB) ([]Collection, error) {
	var collections []Collection
	result := db.Model(&Collection{}).Find(&collections)
	if result.Error != nil {
		return nil, result.Error
	}
	return collections, nil
}

func (c *Collection) CreateOne(tx *gorm.DB) (err error) {
	err = tx.Create(c).Error
	return err
}

func (c *Collection) DeleteOne(tx *gorm.DB) (err error) {
	err = tx.Select("Entries").Delete(c).Error
	return err
}
