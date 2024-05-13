package readers

import (
	"encoding/json"
	"fmt"
	"os"
	"serfwerk/server/misc"
	"serfwerk/server/misc/classes"
)

func ReadLogin(login_info classes.LoginReq) (string, classes.CustErr) {
	db_path := misc.DBPath

	b_content, err := os.ReadFile(db_path + "users.json")
	if err != nil {
		return "", classes.CreateError(500, "Fejl i at læse brugere", err, true)
	}
	var user_list []classes.User
	err2 := json.Unmarshal(b_content, &user_list)
	if err2 != nil {
		fmt.Println(err2)
		return "", classes.CreateError(500, "Fejl i at dekryptere brugere", err2, true)
	}
	found_match := false
	cookie_val := ""
	for _, old_user := range user_list {
		if login_info.Username == old_user.Username || login_info.Username == old_user.Email {
			if login_info.Password == old_user.Password {
				found_match = true
				cookie_val = old_user.UserCookie
			}
		}
	}
	if !found_match {
		return "", classes.CreateError(400, "Forkert brugernavn eller adganskode", err2, true)
	}

	return cookie_val, classes.CreateError(200, "Succes", nil, false)
}
