package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewTeapodSrv(name string, repos repository.Repos) *TeapodSrv {
	return &TeapodSrv{
		Name: name,
		repos: repos,
	}
}

type TeapodSrv struct {
	// deprecated
	Name string
	repos repository.Repos
}

type Teapod struct {
	Name string `json:"name"`
}
func (srv *TeapodSrv) List() ([]Teapod, error) {
	list := make([]Teapod, 0)
	if err := srv.repos.DB.FindAll("teapods", bson.D{}, &list); err != nil {
		return list, err
	}
	return list, nil
}

func (srv *TeapodSrv) GetProvider() (plug.ProviderInterface, error) {
	command := fmt.Sprintf("teapod-%s", srv.Name)
	client := plug.NewClient(command)
	// defer client.Kill()

	rpcc, err := client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcc.Dispense("main")
	if err != nil {
		return nil, err
	}

	return raw.(plug.ProviderInterface), nil
}

func (srv *TeapodSrv) GetInfo() (plug.Info, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Info{}, err
	}
	return provider.Info()
}

func (srv *TeapodSrv) GetTeabox(name string) (plug.Teabox, error) {
	info, err := srv.GetInfo()
	if err != nil {
		return plug.Teabox{}, fmt.Errorf("teabox not found.")
	}
	for _, teabox := range info.Teaboxes {
		if name == teabox.Name {
			return teabox, nil
		}
	}
	return plug.Teabox{}, fmt.Errorf("teabox not found.")
}

func (srv *TeapodSrv) ListTeas(teaboxName string) ([]plug.Tea, error) {
	return make([]plug.Tea, 0), nil
}

func (srv *TeapodSrv) Act(name string, vals []plug.Val) (string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return "", err
	}

	message, err := provider.Act(plug.ActProps{
		Name: name,
		Vals: vals,
	})
	if err != nil {
		return "", err
	}
	return message, nil
}

// func (srv *TeapodSrv) ValidateTeaboxVals(teabox plug.Teabox, vals plug.Vals) error {
// 	for key := range teabox.Vals {
// 		if _, ok := vals[key]; !ok {
// 			return fmt.Errorf("key `%s` is required.", key)
// 		}
// 	}

// 	for key := range vals {
// 		if _, ok := teabox.Vals[key]; !ok {
// 			return fmt.Errorf("additional key `%s` is not allowd.", key)
// 		}
// 	}
// 	return nil
// }
