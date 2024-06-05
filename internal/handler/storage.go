package handler

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Nyarum/diho_gpkov2/internal/actor"
	"github.com/Nyarum/diho_gpkov2/internal/packets"
	"go.etcd.io/bbolt"
)

type StorageReturnType struct {
	DB  *bbolt.DB
	Err error
}

type SaveAccount struct {
	Name string
	Data packets.Auth
}

type GetAccount struct {
	Name string
}

type SaveCharacter struct {
	Login string
	Data  packets.Character
}

type GetCharacters struct {
	Login string
}

type Storage struct {
	Accounts   map[string]packets.Auth
	Characters map[string][]packets.Character
}

func SaveToLocalStorage(storage Storage, db *bbolt.DB) {
	db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("storage"))
		if err != nil {
			return err
		}

		buf, err := json.Marshal(storage)
		if err != nil {
			return err
		}

		return bucket.Put([]byte("storage"), buf)
	})
}

func GetFromLocalStorage(db *bbolt.DB) (Storage, error) {
	var storage Storage

	db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("storage"))

		if bucket == nil {
			return nil
		}

		buf := bucket.Get([]byte("storage"))

		err := json.Unmarshal(buf, &storage)
		if err != nil {
			return err
		}

		return nil
	})

	if storage.Accounts == nil {
		storage.Accounts = make(map[string]packets.Auth)
	}

	if storage.Characters == nil {
		storage.Characters = make(map[string][]packets.Character)
	}

	return storage, nil
}

func NewStorageActor(ctx context.Context) (actor.ActorHandle, StorageReturnType) {
	db, err := bbolt.Open("database.db", 0600, nil)
	if err != nil {
		return nil, StorageReturnType{
			DB:  nil,
			Err: err,
		}
	}

	storage, err := GetFromLocalStorage(db)
	if err != nil {
		return nil, StorageReturnType{
			DB:  nil,
			Err: err,
		}
	}

	return func(me actor.ActorInterface, message any) any {
			switch v := message.(type) {
			case SaveAccount:
				fmt.Println("save account", v.Name)

				storage.Accounts[v.Name] = v.Data
			case GetAccount:
				fmt.Println("Get account", v.Name)

				return storage.Accounts[v.Name]
			case SaveCharacter:
				characters := append(storage.Characters[v.Login], v.Data)
				storage.Characters[v.Login] = characters

			case GetCharacters:
				return storage.Characters[v.Login]

			case actor.ActorReady:
				fmt.Println("storage actor is ready")
			}

			SaveToLocalStorage(storage, db)

			return nil
		}, StorageReturnType{
			DB:  db,
			Err: err,
		}
}
