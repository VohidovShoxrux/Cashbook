package handlers

import (
	"cashbookTeam/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Home(w, r)
	case "POST":
	case "PUT":
	case "DELETE":

	}
}

func Home(w http.ResponseWriter, r *http.Request) {

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)
     
	fmt.Fprintln(w,"________________________________________________")
	for i := 0; i < len(userData); i++ {
		fmt.Fprintln(w, "User's ID:", userData[i].Id)
		fmt.Fprintln(w, "User's Firstname:", userData[i].Firstname)
		fmt.Fprintln(w, "User's Lastname:", userData[i].Lastname)
		fmt.Fprintln(w, "User's CreatedAt:", userData[i].CreatedAt)
		fmt.Fprintln(w, "User's UpdatedAt:", userData[i].UpdatedAt)
		fmt.Fprintln(w,"________________________________________________")
		for j := 0; j < len(userData[i].Cards); j++ {
			fmt.Fprintln(w, "  Card's ID:", userData[i].Cards[j].Id)
			fmt.Fprintln(w, "  Card's UserID:", userData[i].Cards[j].UserId)
			fmt.Fprintln(w, "  Card's Title:", userData[i].Cards[j].Title)
			fmt.Fprintln(w, "  Card's Balance:", userData[i].Cards[j].Balance)
			fmt.Fprintln(w, "  Card's CreatedAt:", userData[i].Cards[j].CreatedAt)
			fmt.Fprintln(w, "  Card's UpdatedAt:", userData[i].Cards[j].UpdatedAt)
			fmt.Fprintln(w,"   ________________________________________________")
			for l := 0; l < len(userData[i].Cards[j].Transactions); l++ {
				fmt.Fprintln(w, "    Transaction's ID:", userData[i].Cards[j].Transactions[l].Id)
				fmt.Fprintln(w, "    Transaction's UserID:", userData[i].Cards[j].Transactions[l].UserId)
				fmt.Fprintln(w, "    Transaction's CardId:", userData[i].Cards[j].Transactions[l].CardId)
				fmt.Fprintln(w, "    Transaction's Title:", userData[i].Cards[j].Transactions[l].Title)
				fmt.Fprintln(w, "    Transaction's Description:", userData[i].Cards[j].Transactions[l].Description)
				fmt.Fprintln(w, "    Transaction's TransactionType:", userData[i].Cards[j].Transactions[l].TransactionType)
				fmt.Fprintln(w, "    Transaction's CreatedAt:", userData[i].Cards[j].Transactions[l].CreatedAt)
				fmt.Fprintln(w, "    Transaction's UpdatedAt:", userData[i].Cards[j].Transactions[l].UpdatedAt)
				fmt.Fprintln(w, "    Transaction's Amount:", userData[i].Cards[j].Transactions[l].Amount)
				fmt.Fprintln(w,"     ________________________________________________")
			}
		}
	}
	
}
