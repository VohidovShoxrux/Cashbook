package helper

import "cashbookTeam/models"

func MaxIDUser(UserArray []models.UserModel) int{
	var maxID = 0
	for i := 0; i < len (UserArray); i++ {
		if maxID < UserArray[i].Id{
			maxID=UserArray[i].Id
		}
	}
	return maxID+1
}
func MaxIDCard(UserArray []models.CardModel) int{
	var maxID = 0
	for i := 0; i < len (UserArray); i++ {
		if maxID < UserArray[i].Id{
			maxID=UserArray[i].Id
		}
	}
	return maxID+1
}
func MaxIDTransaction(UserArray []models.TransactionModel) int{
	var maxID = 0
	for i := 0; i < len (UserArray); i++ {
		if maxID < UserArray[i].Id{
			maxID=UserArray[i].Id
		}
	}
	return maxID+1
}