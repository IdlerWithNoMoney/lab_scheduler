package pro_guap_scraper

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)
type User struct{
	Name 		string
	Pswd		string

	client		http.Client
}

func main(){
	jar, _ := cookiejar.New(nil)
	user := User{Name:"Nikita", Pswd:"111261", client:http.Client{Jar:jar}}
	user.Send(Init)
	user.Send(Auth)
	res, _ := user.Send(Inside_s)
	fmt.Println(res)
}


func (u *User) Send(f func(*User) (*http.Request, error))(string, error){
	if f == nil{
		err := fmt.Errorf("NPE, no function")
		return "", err
	}
	req, err := f(u)
	if err != nil{
		return "", err
	}
	res, err := u.client.Do(req)
	if err != nil{
		return "", err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil{
		return "", err
	}
	return buf.String(), nil
}

func Init(u *User)(*http.Request, error){
	res, err := http.NewRequest("GET", "https://pro.guap.ru", nil)
	if err != nil{
		err := fmt.Errorf("Can't get https://pro.guap.ru, err: %V", err)
		return nil, err
	}
	return res, nil
}

func Auth(u *User)(*http.Request, error){
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

func Inside_s(u *User)(*http.Request, error){
	req, err := http.NewRequest("GET", "https://pro.guap.ru/inside_s", nil)
	if err != nil{
		err := fmt.Errorf("Invalid credentials. err:%V", err)
		return nil, err
	}
	return req, nil
}
