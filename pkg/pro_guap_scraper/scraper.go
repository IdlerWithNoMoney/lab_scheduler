package pro_guap_scraper

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)
type User struct{
	Name 		string
	Pswd		string

	client		http.Client
}


func Tutor(){
	jar, _ := cookiejar.New(nil)
	user := User{Name:"Nikita", Pswd:"111261", client:http.Client{Jar:jar}}
	user.Send(Init, "")
	user.Send(Auth, "")
	res, _ := user.Send(Getsubjectsdictionaries, "")
	fmt.Println(res)
	fmt.Println("**************************************************************************************************************")
	res, _ = user.Send(GetSubject, "2307331")
	fmt.Println(res)
}

func NewUser(name, pswd string) *User{
	jar, _ := cookiejar.New(nil)
	tmp := User{Name:name, Pswd:pswd, client:http.Client{Jar:jar}}
	return &tmp
}

func (u *User) Send(f func(*User, string) (*http.Request, error), arg string)(io.ReadCloser, error){
	if f == nil{
		err := fmt.Errorf("NPE, no function")
		return nil, err
	}
	req, err := f(u, arg)
	if err != nil{
		return nil, err
	}
	res, err := u.client.Do(req)
	if err != nil{
		return nil, err
	}
	return res.Body, nil
}

func Init(u *User, arg string)(*http.Request, error){
	res, err := http.NewRequest("GET", "https://pro.guap.ru", nil)
	if err != nil{
		err := fmt.Errorf("Can't get https://pro.guap.ru, err: %V", err)
		return nil, err
	}
	return res, nil
}

func Auth(u *User, arg string)(*http.Request, error){
	data := url.Values{}
	data.Set("_username", u.Name)
	data.Add("_password", u.Pswd)
	req, err := http.NewRequest("POST", "https://pro.guap.ru/user/login_check", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil{
		err := fmt.Errorf("Invalid credentials. err:%V", err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return req, nil
}

// useless function
func Inside_s(u *User, arg string)(*http.Request, error){
	req, err := http.NewRequest("GET", "https://pro.guap.ru/inside_s", nil)
	if err != nil{
		err := fmt.Errorf("Invalid credentials. err:%V", err)
		return nil, err
	}
	return req, nil
}

func Getsubjectsdictionaries(u *User, arg string)(*http.Request, error){
	// return json with subject and other rubish
	req, err := http.NewRequest("POST", "https://pro.guap.ru/getsubjectsdictionaries/", nil)
	if err != nil{
		err := fmt.Errorf("Invalid credentials. err:%V", err)
		return nil, err
	}
	return req, err
}
func GetSubject(u *User, arg string) (*http.Request, error){
	//curl 'https://pro.guap.ru/subjectItemStudent/' \
	//-H 'Connection: keep-alive' \
	//-H 'Accept: */*' \
	//-H 'X-Requested-With: XMLHttpRequest' \
	//-H 'User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36' \
	//-H 'Content-Type: application/x-www-form-urlencoded; charset=UTF-8' \
	//-H 'Origin: https://pro.guap.ru' \
	//-H 'Sec-Fetch-Site: same-origin' \
	//-H 'Sec-Fetch-Mode: cors' \
	//-H 'Sec-Fetch-Dest: empty' \
	//-H 'Referer: https://pro.guap.ru/inside_s' \
	//-H 'Accept-Language: en-US,en;q=0.9,ru;q=0.8' \
	//-H 'Cookie: PHPSESSID=rd43fiuk3dudgotciknmvg0pok; sharedsessioID=7c6941f2d0c98c9b9193206ee8a0a0acbee47632bb4a5afc9e4568cc4e5dfca39f9f5a829953e1ef8728d5d6d2646c3d88b14c2873b2b134da098ebe52330622' \
	//--data-raw 'id=id of subject' \
	//--compressed
	data := url.Values{}
	data.Set("id", arg)
	res, err := http.NewRequest("POST", "https://pro.guap.ru/subjectItemStudent/", bytes.NewBufferString(data.Encode()))
	if err != nil{
		err := fmt.Errorf("Error : %v", err)
		return nil, err
	}
	res.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return res, err
}

