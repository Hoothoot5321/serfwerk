package readers

import (
	"encoding/json"
	"fmt"
	"os"
	"serfwerk/server/misc"
	"serfwerk/server/misc/classes"
)

func ReadCookie(cookie_val string) (classes.User, classes.CustErr) {
	db_path := misc.DBPath

	b_content, err := os.ReadFile(db_path + "users.json")
	if err != nil {
		return classes.User{}, classes.CreateError(500, "Fejl i at læse brugere", err, true)
	}
	var user_list []classes.User
	err2 := json.Unmarshal(b_content, &user_list)
	if err2 != nil {
		fmt.Println(err2)
		return classes.User{}, classes.CreateError(500, "Fejl i at dekryptere brugere", err2, true)
	}
	found_match := false
	var user classes.User
	for _, old_user := range user_list {
		if old_user.UserCookie == cookie_val {
			found_match = true
			user = old_user
		}
	}
	if !found_match {
		return classes.User{}, classes.CreateError(400, "Invalid cookie", nil, true)
	}

	return user, classes.CreateError(200, "Succes", nil, false)
}
