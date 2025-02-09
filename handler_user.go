package main

import (
	"encoding/json"
	"fmt"
	"my_rss_proj/internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (ApiCfg *apiConfig) handlerCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	user, err := ApiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %s", err))
		return
	}


	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}