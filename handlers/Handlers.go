package handlers

import (
	"encoding/json"
	"net/http"
	. "newUdemy/dataloaders"
	. "newUdemy/models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// page object
	page := Page{ID: 7, Name: "Kullanıcılar", Description: "Kullanıcı Listesi", URI: "/users"}
	//dataloaders
	users := LoadUsers()
	interests := LoadInterest()
	interestMappings := LoadInterestMappings()

	// Process

	var newUsers []User

	for _, user := range users {
		for _, interestMapping := range interestMappings {
			if user.ID == interestMapping.UserID {
				for _, interest := range interests {
					if interestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
	}

	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}
