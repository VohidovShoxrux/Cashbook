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

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetTransaction(w, r)
	case "POST":
		CreateTransaction(w, r)
	case "PUT":
		UpdateTransaction(w, r)
	case "DELETE":
		DeleteTransaction(w, r)
	}
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var newTransaction models.TransactionModel
	json.NewDecoder(r.Body).Decode(&newTransaction)

	//---------------------------------------------

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var UserFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == newTransaction.UserId {
			var CardFound bool
			for j := 0; j < len(userData[i].Cards); j++ {
				if userData[i].Cards[j].Id == newTransaction.CardId {
					newTransaction.Id = helper.MaxIDTransaction(userData[i].Cards[j].Transactions)
					newTransaction.CreatedAt = time.Now()
					newTransaction.UpdatedAt = time.Now()
					if newTransaction.Title == "" {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Transaction's Title cannot be an empty string!")
						return
					} else if newTransaction.Description == "" {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Transaction's Title can not be an empty string!")
						return
					} else if newTransaction.TransactionType ==""{
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Transaction's transactionType cannot be an empty string!")
						return
					} else if newTransaction.Amount == 0 {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Transaction's amount cannot be an empty string!")
						return
					}
					userData[i].Cards[j].Transactions = append(userData[i].Cards[j].Transactions, newTransaction)
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
	fmt.Fprintf(w, "Transaction created with ID: %d\n", newTransaction.Id)

	json.NewEncoder(w).Encode(newTransaction)
}
func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	
	var updateTransaction models.TransactionModel
	json.NewDecoder(r.Body).Decode(&updateTransaction)

	//---------------------------------------------

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var TakeTransaction models.TransactionModel

	var UserFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == updateTransaction.UserId {
			var CardFound bool
			for j := 0; j < len(userData[i].Cards); j++ {
				if userData[i].Cards[j].Id == updateTransaction.CardId {
					var Transactionfound bool
					for l := 0; l < len(userData[i].Cards[j].Transactions); l++ {
						if userData[i].Cards[j].Transactions[l].Id == updateTransaction.Id {
							if updateTransaction.Title == "" {
								w.WriteHeader(http.StatusBadRequest)
								fmt.Fprintln(w, "Transaction's Title cannot be an empty string!")
								return
							} else if updateTransaction.Description == "" {
								w.WriteHeader(http.StatusBadRequest)
								fmt.Fprintln(w, "Transaction's Title cannot be an empty string!")
								return
							} else if updateTransaction.TransactionType ==""{
								w.WriteHeader(http.StatusBadRequest)
								fmt.Fprintln(w, "Transaction's transactionType cannot be an empty string!")
								return
							} else if updateTransaction.Amount == 0 {
								w.WriteHeader(http.StatusBadRequest)
								fmt.Fprintln(w, "Transaction's amount cannot be an empty string!")
								return
							}
							userData[i].Cards[j].Transactions[l].Title = updateTransaction.Title
							userData[i].Cards[j].Transactions[l].Description = updateTransaction.Description
							userData[i].Cards[j].Transactions[l].TransactionType = updateTransaction.TransactionType
							userData[i].Cards[j].Transactions[l].Amount =updateTransaction.Amount
							userData[i].Cards[j].Transactions[l].UpdatedAt = time.Now()
							TakeTransaction = userData[i].Cards[j].Transactions[l]
							Transactionfound = true
							break
						}
					}
					if !Transactionfound {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Transaction's ID not found!")
						return
					}
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
	fmt.Fprintf(w, "Transaction created with ID: %d\n", updateTransaction.Id)

	json.NewEncoder(w).Encode(TakeTransaction)
}
func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var deleteTransaction models.TransactionModel
	json.NewDecoder(r.Body).Decode(&deleteTransaction)

	//---------------------------------------------

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var UserFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == deleteTransaction.UserId {
			var CardFound bool
			for j := 0; j < len(userData[i].Cards); j++ {
				if userData[i].Cards[j].Id == deleteTransaction.CardId {
					var Transactionfound bool
					for l := 0; l < len(userData[i].Cards[j].Transactions); l++ {
						if userData[i].Cards[j].Transactions[l].Id == deleteTransaction.Id {
							userData[i].Cards[j].Transactions= append(userData[i].Cards[j].Transactions[:l], userData[i].Cards[j].Transactions[l+1:]...)
							Transactionfound = true
							break
						}
					}
					if !Transactionfound {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Transaction's ID not found!")
						return
					}
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
	fmt.Fprintf(w, "Transaction deleted with ID: %d\n", deleteTransaction.Id)
}
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var getTransaction models.TransactionModel
	json.NewDecoder(r.Body).Decode(&getTransaction)

	//---------------------------------------------

	var userData []models.UserModel
	userByte, _ := os.ReadFile("db/all.json")
	json.Unmarshal(userByte, &userData)

	var TakeTransaction models.TransactionModel

	var UserFound bool
	for i := 0; i < len(userData); i++ {
		if userData[i].Id == getTransaction.UserId {
			var CardFound bool
			for j := 0; j < len(userData[i].Cards); j++ {
				if userData[i].Cards[j].Id == getTransaction.CardId {
					var Transactionfound bool
					for l := 0; l < len(userData[i].Cards[j].Transactions); l++ {
						if userData[i].Cards[j].Transactions[l].Id == getTransaction.Id {
							TakeTransaction = userData[i].Cards[j].Transactions[l]
							Transactionfound = true
							break
						}
					}
					if !Transactionfound {
						w.WriteHeader(http.StatusBadRequest)
						fmt.Fprintln(w, "Transaction's ID not found!")
						return
					}
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
	fmt.Fprintf(w, "Transaction found with ID: %d\n", getTransaction.Id)
	
	json.NewEncoder(w).Encode(TakeTransaction)
}
