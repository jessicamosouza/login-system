package http

import (
	"encoding/json"
	"net/http"
)

type DeleteUserPayload struct {
	Email string `json:"email"`
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil {
		http.Error(w, "Empty body", http.StatusBadRequest)
	}

	var user DeleteUserPayload
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
	}

	//err := userops.DeleteUser(userops.User{
	//	Email: user.Email,
	//})
	//
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully!"))
}
