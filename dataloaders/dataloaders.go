package dataloaders

import (
	"encoding/json"
	. "newUdemy/models"
	"newUdemy/utils"
)

func LoadUsers() []User {
	bytes, _ := utils.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterest() []Interest {
	bytes, _ := utils.ReadFile("../json/interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data
}
func LoadInterestMappings() []InterestMapping {
	bytes, _ := utils.ReadFile("../json/userInterestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data
}
