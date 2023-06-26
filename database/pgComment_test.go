package database

import (
	"context"
	"github.com/dipdup-net/go-lib/mocks"
	"github.com/go-pg/pg/v10"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initMocks(t *testing.T) (*gomock.Controller, *mocks.MockSchemeCommenter, context.Context) {
	mockCtrl := gomock.NewController(t)
	mockSchemeCommenter := mocks.NewMockSchemeCommenter(mockCtrl)
	ctx := context.Background()

	return mockCtrl, mockSchemeCommenter, ctx
}

func toInterfaceSlice(origin []pg.Safe) []interface{} {
	res := make([]interface{}, len(origin))

	for i := range origin {
		res[i] = origin[i]
	}

	return res
}

func TestMakeCommentsWithTableName(t *testing.T) {
	type Ballot struct {
		//nolint
		tableName struct{} `pg:"ballots" pg-comment:"Ballot table"`
		Ballot    string   `json:"ballot"`
	}

	mockCtrl, mockSC, ctx := initMocks(t)
	defer mockCtrl.Finish()

	model := Ballot{}

	// Assert prepare
	mockSC.
		EXPECT().
		MakeTableComment(ctx, "ballots", "Ballot table").
		Times(1).
		Return(nil)

	// Act
	err := makeComments(ctx, mockSC, model)

	// Assert
	assert.Empty(t, err)
}

func TestMakeCommentsWithoutPgComment(t *testing.T) {
	type Ballot struct {
		//nolint
		tableName struct{} `pg:"ballots"`
		Ballot    string   `json:"ballot"`
	}

	mockCtrl, mockSC, ctx := initMocks(t)
	defer mockCtrl.Finish()

	model := Ballot{}

	// Assert prepare
	mockSC.
		EXPECT().
		MakeTableComment(gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	mockSC.
		EXPECT().
		MakeColumnComment(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	// Act
	err := makeComments(ctx, mockSC, model)

	// Assert
	assert.Empty(t, err)
}

func TestMakeCommentsFieldWithPgComment(t *testing.T) {
	type Ballot struct {
		//nolint
		tableName struct{} `pg:"ballots"`
		Ballot    string   `json:"ballot" pg-comment:"This is field comment"`
	}

	mockCtrl, mockSC, ctx := initMocks(t)
	defer mockCtrl.Finish()

	model := Ballot{}

	// Assert prepare
	mockSC.
		EXPECT().
		MakeColumnComment(ctx, "ballots", "ballot", "This is field comment").
		Times(1).
		Return(nil)

	// Act
	err := makeComments(ctx, mockSC, model)

	// Assert
	assert.Empty(t, err)
}

func TestMakeCommentsWithTableNameAndFieldsWithPgComment(t *testing.T) {
	type Ballot struct {
		//nolint
		tableName       struct{}    `pg:"ballots" pg-comment:"Ballot table"`
		CreatedAt       int64       `json:"-" pg-comment:"This is field comment"`
		UpdatedAt       int64       `json:"-" pg-comment:"This is field comment"`
		Network         string      `json:"network" pg:",pk" pg-comment:"This is field comment"`
		Hash            string      `json:"hash" pg:",pk" pg-comment:"This is field comment"`
		Branch          string      `json:"branch" pg-comment:"This is field comment"`
		Status          string      `json:"status" pg-comment:"This is field comment"`
		Kind            string      `json:"kind" pg-comment:"This is field comment"`
		Signature       string      `json:"signature" pg-comment:"This is field comment"`
		Protocol        string      `json:"protocol" pg-comment:"This is field comment"`
		Level           uint64      `json:"level" pg-comment:"This is field comment"`
		Errors          interface{} `json:"errors,omitempty" pg:"type:jsonb" pg-comment:"This is field comment"`
		ExpirationLevel *uint64     `json:"expiration_level" pg-comment:"This is field comment"`
		Raw             interface{} `json:"raw,omitempty" pg:"type:jsonb" pg-comment:"This is field comment"`
		Ballot          string      `json:"ballot" pg-comment:"This is field comment"`
		Period          int64       `json:"period" pg-comment:"This is field comment"`
	}

	mockCtrl, mockSC, ctx := initMocks(t)
	defer mockCtrl.Finish()

	model := Ballot{}

	// Assert prepare
	commentOnTableCall := mockSC.
		EXPECT().
		MakeTableComment(ctx, "ballots", "Ballot table").
		Times(1).
		Return(nil)

	mockSC.
		EXPECT().
		MakeColumnComment(ctx, "ballots", gomock.Any(), "This is field comment").
		Times(15).
		After(commentOnTableCall).
		Return(nil)

	// Act
	err := makeComments(ctx, mockSC, model)

	// Assert
	assert.Empty(t, err)
}

func TestMakeCommentsWithMultipleModels(t *testing.T) {
	type Ballot struct {
		//nolint
		tableName struct{} `pg:"ballots" pg-comment:"This multiple table name comment"`
		Ballot    string   `json:"ballot" pg-comment:"This is multiple field comment"`
	}

	mockCtrl, mockSC, ctx := initMocks(t)
	defer mockCtrl.Finish()

	models := []interface{}{Ballot{}, Ballot{}, Ballot{}}

	// Assert prepare
	mockSC.
		EXPECT().
		MakeTableComment(ctx, "ballots", "This multiple table name comment").
		Times(3).
		Return(nil)

	// Be aware: there is on issue with default order in checking
	// methods call: https://github.com/golang/mock/issues/653
	mockSC.
		EXPECT().
		MakeColumnComment(ctx, "ballots", "ballot", "This is multiple field comment").
		Times(3).
		Return(nil)

	// Act
	err := makeComments(ctx, mockSC, models...)

	// Assert
	assert.Empty(t, err)
}
