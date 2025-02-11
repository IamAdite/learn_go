package main

import (
	"encoding/json"
	"fmt"
	"my_rss_proj/internal/database"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error creating feed follow: %s", err))
		return
	}


	respondWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollow, err := apiCfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error getting feed follows: %s", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollow))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFoollowIDStr := chi.URLParam(r, "feedFoollowID")
	feedFollowID, err := uuid.Parse(feedFoollowIDStr)
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error parsing followId to uuid: %s", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: feedFollowID,
		UserID: user.ID,
	})
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("Error deleting feed follow: %s", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}