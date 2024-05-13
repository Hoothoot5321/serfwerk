package writers

import (
	"encoding/json"
	"os"
	"serfwerk/server/misc"
	"serfwerk/server/misc/classes"
	"serfwerk/server/misc/cookie"
)

func WriteNewUser(user_req classes.UserReq) (string, classes.CustErr) {
	db_path := misc.DBPath
	b_content, err := os.ReadFile(db_path + "users.json")
	if err != nil {
		return "", classes.CreateError(500, "Fejl i at læse brugere", err, true)
	}
	var user_list []classes.User
	err2 := json.Unmarshal(b_content, &user_list)
	if err2 != nil {
		return "", classes.CreateError(500, "Fejl i at dekryptere brugere", err2, true)
	}
	for _, old_user := range user_list {
		if user_req.Username == old_user.Username {
			return "", classes.CreateError(400, "En konto har allrede dette brugernavn", nil, true)
		}
		if user_req.Email == old_user.Email {
			return "", classes.CreateError(400, "En konto brugere allerede denne Email", nil, true)
		}
	}
	cookie_val, err2 := cookie.CreateCookie(20)
	if err2 != nil {
		return "", classes.CreateError(500, "Fejl i at lave bruger cookie", err2, true)
	}
	new_user := classes.User{UserCookie: cookie_val, Username: user_req.Username, Password: user_req.Password, Email: user_req.Email, Apps: make([]classes.App, 0)}

	user_list = append(user_list, new_user)

	b_new_conent, err := json.Marshal(user_list)
	if err != nil {
		return "", classes.CreateError(500, "Fejl i at dekryptere brugere", err, true)
	}

	err3 := os.WriteFile(db_path+"users.json", b_new_conent, 0644)

	if err3 != nil {
		return "", classes.CreateError(500, "Fejl i at nedskrive bruger", err3, true)
	}
	err4 := os.Mkdir(db_path+"users/"+new_user.Username, 0744)
	if err4 != nil {
		return "", classes.CreateError(500, "Fejl i at nedskrive bruger", err4, true)
	}

	return cookie_val, classes.CreateError(200, "Succes", nil, false)
}
