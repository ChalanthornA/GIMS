package database

import "github.com/google/uuid"

func GenerateUUID() (uint32, error){
	id, err := uuid.NewRandom()
	if err != nil {
		return 0, err
	}
	return id.ID(), nil
}