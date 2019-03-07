package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/bstaijen/mariadb-for-microservices/vote-service/database"
	"github.com/urfave/negroni"

	"github.com/bstaijen/mariadb-for-microservices/shared/helper"
	sharedModels "github.com/bstaijen/mariadb-for-microservices/shared/models"
	"github.com/bstaijen/mariadb-for-microservices/shared/util"
)

// GetTopRatedHandler is the handler for the top rated timeline.
func GetTopRatedHandler(connection *sql.DB) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		// get offset and rows
		offset, rows := helper.PaginationFromRequest(r)

		topRated, err := db.GetTopRatedTimeline(connection, offset, rows)
		if err != nil {
			util.SendError(w, err)
			return
		}

		type Resp struct {
			Results []*sharedModels.TopRatedPhotoResponse `json:"results"`
		}
		util.SendOK(w, &Resp{Results: topRated})
	})
}

// GetHotHandler is the handler for the hot photos timeline
func GetHotHandler(connection *sql.DB) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		// get offset and rows
		offset, rows := helper.PaginationFromRequest(r)

		hot, err := db.GetHotTimeline(connection, offset, rows)
		if err != nil {
			util.SendError(w, err)
			return
		}

		type Resp struct {
			Results []*sharedModels.TopRatedPhotoResponse `json:"results"`
		}
		util.SendOK(w, &Resp{Results: hot})
	})
}

// HasVotedHandler is the handler for calculating whether the user has voted on a photo or not.
func HasVotedHandler(connection *sql.DB) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		type Collection struct {
			Objects []*sharedModels.HasVotedRequest `json:"requests"`
		}
		col := &Collection{}
		col.Objects = make([]*sharedModels.HasVotedRequest, 0)

		err := util.RequestToJSON(r, &col)
		if err != nil {
			log.Fatal(err)
		}

		counts, err := db.HasVoted(connection, col.Objects)
		if err != nil {
			util.SendError(w, err)
			return
		}

		type Resp struct {
			Results []*sharedModels.HasVotedResponse `json:"results"`
		}
		util.SendOK(w, &Resp{Results: counts})
	})
}

// GetVoteCountHandler is the handler for calculating the number of votes on a photo
func GetVoteCountHandler(connection *sql.DB) negroni.HandlerFunc {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		type Collection struct {
			Objects []*sharedModels.VoteCountRequest `json:"requests"`
		}
		col := &Collection{}
		col.Objects = make([]*sharedModels.VoteCountRequest, 0)

		err := util.RequestToJSON(r, &col)
		if err != nil {
			log.Fatal(err)
		}

		counts, err := db.VoteCount(connection, col.Objects)
		if err != nil {
			util.SendError(w, err)
			return
		}

		type Resp struct {
			Results []*sharedModels.VoteCountResponse `json:"results"`
		}
		util.SendOK(w, &Resp{Results: counts})
	})
}
