package seeders

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"reflect"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/chat/config"
	"github.com/umardev500/chat/domain/models"
	"github.com/umardev500/chat/utils"
)

type Seeder struct {
	conn *config.Connection
}

func NewSeeder(c *config.Connection) *Seeder {
	return &Seeder{
		conn: c,
	}
}

// getData reads data from a file and unmarshals it into a struct.
//
// It takes in the file path as a string and the target struct as an interface{}.
// It returns an error if there was an issue reading or unmarshalling the data.
func (s *Seeder) GetData(filePath string, result interface{}) (err error) {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	filePath = filepath.Join(dir, filePath)
	f, err := os.ReadFile(filePath)
	if err != nil {
		log.Error().Msgf("Error reading data: %v", err)
		return
	}

	err = json.Unmarshal(f, &result)
	if err != nil {
		log.Error().Msgf("Error unmarshalling data: %v", err)
		return
	}

	return
}

func (s *Seeder) Insert(ctx context.Context, table, source string, data interface{}) (err error) {
	db := s.conn.TrOrDB(ctx)

	dataType := reflect.TypeOf(data)
	if dataType == nil {
		return errors.New("data cannot be nil")
	}

	if dataType.Kind() != reflect.Slice {
		return errors.New("data must be a slice")
	}
	elementType := dataType.Elem()
	slicePtr := reflect.New(reflect.SliceOf(elementType))
	sliceInterface := slicePtr.Interface()

	err = s.GetData(source, &sliceInterface)
	if err != nil {
		return err
	}

	slice := reflect.Indirect(reflect.ValueOf(sliceInterface)).Interface()
	query, args, err := utils.BuildBulkInsertQuery(table, slice)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return
}

func (s *Seeder) Populate(ctx context.Context) (err error) {
	var base = "database/seeders/data"

	err = s.conn.WithTransaction(ctx, func(ctx context.Context) error {
		log.Info().Msgf("Seeding users")
		err = s.Insert(ctx, "users", filepath.Join(base, "users.json"), []models.UserCreate{})
		if err != nil {
			return err
		}

		log.Info().Msgf("Chat rooms")
		err = s.Insert(ctx, "chat_rooms", filepath.Join(base, "chat_rooms.json"), []models.ChatRoomCreate{})
		if err != nil {
			return err
		}

		log.Info().Msgf("User rooms")
		err = s.Insert(ctx, "user_rooms", filepath.Join(base, "user_rooms.json"), []models.UserRoomCreate{})
		if err != nil {
			return err
		}

		log.Info().Msgf("Messages")
		err = s.Insert(ctx, "messages", filepath.Join(base, "messages.json"), []models.MessageCreate{})
		if err != nil {
			return err
		}

		return nil
	})

	return
}
