package pro_guap_scraper

import (
	"testing"
)

func TestUser_Send(t *testing.T) {
	user := User{Name: "ACAB", Pswd:"11111111"}
	_, err := user.Send(nil)
	if err == nil{
		t.Error("Function mustn't work with func(nil) ")
	}else{
		t.Log("Idiot test passed")
	}
	res, _ := user.Send(Init)
	if res == ""{
		t.Error("Look at func 'Send', it feel bad, or you haven't internet connection")
	}else{
		t.Log("Another test passed")
	}
}

func TestInit(t *testing.T) {
	user := User{Name:"asdf", Pswd: "fdsa"}
	_, err := Init(&user)
	if err != nil{
		t.Errorf("(Init)Some problems, err %v:", err)
		t.Fail()
	}else{
		t.Log("Init request generated Success!!")
	}
}

func TestAuth(t *testing.T) {
	user := User{Name:"asdf", Pswd: "fdsa"}
	_, err := Auth(&user)
	if err != nil{
		t.Errorf("(Auth)Some problems, err %v:", err)
		t.Fail()
	}else{
		t.Log("Auth request generated Success!!")
	}
}

func TestInside_s(t *testing.T) {
	user := User{Name:"asdf", Pswd: "fdsa"}
	_, err := Inside_s(&user)
	if err != nil{
		t.Errorf("(Inside_s)Some problems, err %v:", err)
		t.Fail()
	}else{
		t.Log("Inside_s request generated Success!!")
	}
}