package pro_guap_scraper

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

type Guapchanin struct{
	user User

}
type Container struct {
	Subjects []struct {
		ID		string	`json:"id"`
		Subj	string	`json:"subj"`
	} 	`json:"subjects"`
}
type TaskContainer struct {
	Subject struct{
		Tasks []struct{
			Name		string	`json:"name"`
			StatusName	string	`json:"status_name"`
		}	`json:"tasks"`
	}	`json:"subject"`
}

func NewGuapchanin(name, pswd string)(* Guapchanin){
	tmp := Guapchanin{user: *NewUser(name, pswd)}
	res, _ :=tmp.user.Send(Init, "")
	res.Close()
	res, _ = tmp.user.Send(Auth, "")
	return &tmp
}

func (u *Guapchanin)GetSubjects() map[string]string{
	log.Info().Msg("Getting subjects")
	body, err := u.user.Send(Getsubjectsdictionaries, "")
	if err != nil{
		log.Info().Msg("Cookies time is over")
		u._failure()
		return u.GetSubjects()
	}else{
		defer body.Close()
	}
	var res Container
	dec := json.NewDecoder(body)
	if err := dec.Decode(&res); err != nil{
		log.Fatal().Err(err).Msg("Decoding failed")
	}
	resMap := make(map[string]string)
	for _, item := range res.Subjects{
		resMap[item.Subj] = item.ID
	}
	return resMap
}

func (u *Guapchanin)GetTasks(id string)map[string]string{
	log.Info().Msgf("Getting tasks from subject with id : %v", id)
	body, err := u.user.Send(GetSubject, id)
	if err != nil{
		log.Info().Msg("Cookies time is over")
		u._failure()
		return u.GetTasks(id)
	}else{
		defer body.Close()
	}
	var res TaskContainer
	dec := json.NewDecoder(body)
	if err := dec.Decode(&res); err != nil{
		log.Fatal().Err(err).Msg("Decoding failed")
	}
	resMap := make(map[string]string)
	for _, item := range res.Subject.Tasks{
		resMap[item.Name] = item.StatusName
	}
	return resMap
}

func (u *Guapchanin)_failure() {
	u.user.Send(Init, "")
	u.user.Send(Auth, "")
}

func Example(){
	user := NewGuapchanin("Nikita", "111261")
	ex := user.GetSubjects()
	for key, val := range ex{
		log.Debug().Msgf("key : %v, val : %v", key, val)
		tasks := user.GetTasks(val)
		for task, status := range tasks{
			log.Debug().Msgf("task : %v, status : %v" ,task, status)
		}
	}
}