package models_test

import (
	"encoding/json"
	"read-more-backend/models"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func marshalUnmarshal[T any](toJson T, fromJson *T) error {
	j, err := json.Marshal(toJson)
	if err != nil {
		return err
	}
	err = json.Unmarshal(j, &fromJson)
	if err != nil {
		return err
	}
	return nil
}

func TestCreateEntryDtoUnmarshal(t *testing.T) {
	t.Parallel()

	want := models.CreateEntryDto{
		Title:           "EntryTitle",
		Description:     "EntryDescription",
		URL:             "EntryURL",
		Transcription:   "EntryTranscription",
		AudioFilename:   "EntryAudioFilename",
		Seen:            false,
		CollectionID:    uuid.Nil,
		CollectionTitle: "EntryCollectionTitle",
	}

	var got models.CreateEntryDto
	if err := marshalUnmarshal(want, &got); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %+v, got: %+v", got, want)
	}
}

func TestCreateEntryDtoUnmarshalInvalid(t *testing.T) {
	t.Parallel()

	base := models.CreateEntryDto{
		Title:         "EntryTitle",
		Description:   "EntryDescription",
		URL:           "EntryURL",
		Transcription: "EntryTranscription",
		AudioFilename: "EntryAudioFilename",
		Seen:          false,
	}

	// with both collectionTitle and collectionId set to their zero values
	base.CollectionID = uuid.Nil
	base.CollectionTitle = ""

	var got models.CreateEntryDto
	if err := marshalUnmarshal(base, &got); err == nil {
		t.Fatal("want error got nil")
	}

	// without providing a collectionID
	// TODO: TEST

	// without providing a collectionTitle
	// TODO: TEST

	// without providing both
	// TODO: TEST
}
