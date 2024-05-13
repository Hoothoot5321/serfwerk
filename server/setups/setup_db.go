package setups

import (
	"encoding/json"
	"fmt"
	"os"
	"serfwerk/server/misc"
	"serfwerk/server/misc/classes"
)

func SetupDB() {
	db_path := misc.DBPath
	f, _ := os.OpenFile(db_path+"users.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	temp := make([]byte, 100)
	f_length, err := f.Read(temp)
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir(db_path+"users", 0744)

	if f_length < 2 {
		f.Write([]byte("[]"))
	} else {

		b_content, _ := os.ReadFile(db_path + "users.json")
		var user_list []classes.User
		json.Unmarshal(b_content, &user_list)
		for _, user := range user_list {
			os.Mkdir(db_path+"users/"+user.Username, 0744)
		}
	}
}
