package classes

type User struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	UserCookie string `json:"userCookie"`
	Apps       []App  `json:"apps"`
}

type App struct {
	AppName     string `json:"appName"`
	OwnerCookie string `json:"ownerCookie"`
	Online      bool   `json:"online"`
	BuilCmd     string `json:"buildCmd"`
	RunCmd      string `json:"runCmd"`
	PortNum     int    `json:"port"`
	AppNum      int    `json:"appNum"`
	Ram         int    `json:"ram"`
	CPU         int    `json:"cpu"`
	Lager       int    `json:"lager"`
	CPUPower    int    `json:"cpuPower"`
}

type CustErr struct {
	ErrString string `json:"errStr"`
	Code      int    `json:"code"`
	Err       error  `json:"error"`
	Exit      bool   `json:"exit"`
}

func (m *CustErr) Error() string {
	return m.ErrString
}
func CreateError(err_code int, err_string string, err error, exit bool) CustErr {
	return CustErr{Err: err, Code: err_code, ErrString: err_string, Exit: exit}
}
