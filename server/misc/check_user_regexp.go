package misc

import (
	"regexp"
	"serfwerk/server/misc/classes"
)

func CheckReqxpLogin(username string, password string) classes.CustErr {
	email_regex, _ := regexp.Compile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
	username_regex, _ := regexp.Compile("[^A-Za-z0-9]")
	password_regex, _ := regexp.Compile("^[^<>{}\"/|;:.,~%^=*\\]\\\\()\\[¿§«»ω⊙¤°℃℉€¥£¢¡®©0-9_+]*$")

	if email_regex.Match([]byte(username)) {
	} else if len(username) > 6 && len(username) < 12 && !username_regex.Match([]byte(username)) {
	} else {
		return classes.CreateError(400, "Ikke valid brugernavn eller mail", nil, true)
	}

	if len(password) < 6 {
		return classes.CreateError(400, "Kodeord skal være længere end 6", nil, true)
	}
	if len(password) > 12 {
		return classes.CreateError(400, "Kodeord skal være kortere end 12", nil, true)
	}
	if !password_regex.Match([]byte(password)) {
		return classes.CreateError(400, "Kodeord skal må ikke indeholde specielle tegn", nil, true)
	}

	return classes.CreateError(200, "Succes", nil, false)
}
func CheckReqxp(username string, email string, password string) classes.CustErr {
	email_regex, _ := regexp.Compile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")
	username_regex, _ := regexp.Compile("[^A-Za-z0-9]")
	password_regex, _ := regexp.Compile("^[^<>{}\"/|;:.,~%^=*\\]\\\\()\\[¿§«»ω⊙¤°℃℉€¥£¢¡®©0-9_+]*$")

	if len(username) < 6 {
		return classes.CreateError(400, "Brugernavn skal være længere end 6", nil, true)
	}
	if len(username) > 12 {
		return classes.CreateError(400, "Brugernavn skal være kortere end 12", nil, true)
	}
	if username_regex.Match([]byte(username)) {
		return classes.CreateError(400, "Brugernavn skal må ikke indeholde specielle tegn", nil, true)
	}
	if !email_regex.Match([]byte(email)) {
		return classes.CreateError(400, "Ikke valid email", nil, true)
	}

	if len(password) < 6 {
		return classes.CreateError(400, "Kodeord skal være længere end 6", nil, true)
	}
	if len(password) > 12 {
		return classes.CreateError(400, "Kodeord skal være kortere end 12", nil, true)
	}
	if !password_regex.Match([]byte(password)) {
		return classes.CreateError(400, "Kodeord skal må ikke indeholde specielle tegn", nil, true)
	}

	return classes.CreateError(200, "Succes", nil, false)
}
