package fixture

import (
	"path/filepath"

	"github.com/acmota2/musmgr/backend/internal/model"
	"github.com/acmota2/musmgr/backend/internal/policies"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var MusmgrEvents = []model.CreateEventParams{
	{
		ID:          uuid.New(),
		HappenedAt:  "2019-03-02",
		Name:        "The big game event 2019",
		Description: pgtype.Text{String: "A big game event", Valid: true},
		EventType:   model.MusmgrEventTypeOther,
	},
	{
		ID:          uuid.New(),
		HappenedAt:  "2022-05-17",
		Name:        "The Spring Festival 2022",
		Description: pgtype.Text{Valid: false},
		EventType:   model.MusmgrEventTypeFestival,
	},
}

var Pieces = []model.CreatePieceParams{
	{
		ID:              uuid.New(),
		ComposedAt:      "2018",
		Description:     "A game sounding piece",
		Instrumentation: model.MusmgrInstrumentationNameEnsemble,
		Title:           "Level I",
	},
	{
		ID:              uuid.New(),
		ComposedAt:      "2021",
		Description:     "A game sounding piece",
		Instrumentation: model.MusmgrInstrumentationNameEnsemble,
		Title:           "Level II",
	},
}

var Files = map[string]model.CreateFileParams{
	filepath.Join("testdata", "Level I.pdf"): {
		ID:             uuid.New(),
		ContentType:    "application/pdf",
		Classification: int16(policies.ClassProtected),
		FileType:       model.MusmgrFileTypeScoreFull,
		Name:           "Level I.pdf",
		Origin:         model.MusmgrFileOriginUser,
		ParentID:       pgtype.UUID{Valid: false},
		PieceID:        Pieces[0].ID,
	},
	filepath.Join("testdata", "Level II.pdf"): {
		ID:             uuid.New(),
		ContentType:    "application/pdf",
		Classification: int16(policies.ClassProtected),
		FileType:       model.MusmgrFileTypeScoreFull,
		Name:           "Level II.pdf",
		Origin:         model.MusmgrFileOriginUser,
		ParentID:       pgtype.UUID{Valid: false},
		PieceID:        Pieces[1].ID,
	},
	filepath.Join("testdata", "Level II.mp3"): {
		ID:             uuid.New(),
		ContentType:    "audio/mpeg",
		Classification: int16(policies.ClassPublic),
		FileType:       model.MusmgrFileTypeAudioRecording,
		Name:           "Level I.mp3",
		Origin:         model.MusmgrFileOriginUser,
		ParentID:       pgtype.UUID{Valid: false},
		PieceID:        Pieces[0].ID,
	},
}
