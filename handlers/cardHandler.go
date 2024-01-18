package handlers

import (
	"cashbookTeam/helper"
	"cashbookTeam/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func CardHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetCard(w, r)
	case "POST":
		CreateCard(w, r)
	case "PUT":
		UpdateCard(w, r)
	case "DELETE":
		DeleteCard(w, r)

	}
}
func CreateCard(w http.ResponseWriter, r *http.Request) {
	var newCard models.CardModel
	json.NewDecoder(r.Body).Decode(&newCard)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var UserFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == newCard.UserId {
			newCard.Id = helper.MaxIDCard(userData[i].Cards)
			newCard.CreatedAt = time.Now()
			newCard.UpdatedAt = time.Now()

			if newCard.Title == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Card's Title cannot be an empty string!")
				return
			}
			userData[i].Cards = append(userData[i].Cards, newCard)
			UserFound = true
			break
		}
	}
	if !UserFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User's ID not found!")
		return
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Card created with ID: %d\n", newCard.Id)

	json.NewEncoder(w).Encode(newCard)
}
func UpdateCard(w http.ResponseWriter, r *http.Request) {
	var updateCard models.CardModel
	json.NewDecoder(r.Body).Decode(&updateCard)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var TakeCard models.CardModel

	var UserFound bool
	var CardFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == updateCard.UserId {
			for j := 0; j < len(userData[i].Cards); j++ {
				if userData[i].Cards[j].Id == updateCard.Id {
					if updateCard.Title == "" {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Card's Title cannot be an empty string!")
						return
					}
					userData[i].Cards[j].Title = updateCard.Title
					userData[i].Cards[j].UpdatedAt = time.Now()
					TakeCard = userData[i].Cards[j]
					CardFound = true
					break
				}
			}
			if !CardFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Card's ID not found!")
				return
			}
			UserFound = true
			break
		}
	}
	if !UserFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User's ID not found!")
		return
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Card updated with ID: %d\n", updateCard.Id)

	json.NewEncoder(w).Encode(TakeCard)
}
func DeleteCard(w http.ResponseWriter, r *http.Request) {
	var deleteCard models.CardModel
	json.NewDecoder(r.Body).Decode(&deleteCard)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var UserFound bool
	var CardFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == deleteCard.UserId {
			for j := 0; j < len(userData[i].Cards); j++ {
				if userData[i].Cards[j].Id == deleteCard.Id {
					userData[i].Cards = append(userData[i].Cards[:j], userData[i].Cards[j+1:]...)
					CardFound = true
					break
				}
			}
			if !CardFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Card's ID not found!")
				return
			}
			UserFound = true
			break
		}
	}
	if !UserFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User's ID not found!")
		return
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Card deleted with ID: %d\n", deleteCard.Id)
}
func GetCard(w http.ResponseWriter, r *http.Request) {
	var getCard models.CardModel
	json.NewDecoder(r.Body).Decode(&getCard)

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var TakeCard models.CardModel

	var UserFound bool
	var CardFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == getCard.UserId {
			for j := 0; j < len(userData[i].Cards); j++ {
				if userData[i].Cards[j].Id == getCard.Id {
					TakeCard = userData[i].Cards[j]
					CardFound = true
					break
				}
			}
			if !CardFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintln(w, "Card's ID not found!")
				return
			}
			UserFound = true
			break
		}
	}
	if !UserFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "User's ID not found!")
		return
	}

	res, _ := json.Marshal(userData)
	os.WriteFile("db/all.json", res, 0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Card found with ID: %d\n", getCard.Id)

	json.NewEncoder(w).Encode(TakeCard)
}
