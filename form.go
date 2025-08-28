package camunda_client_go

import (
	"fmt"
)

type ComundaFormComponentLayout struct {
	Row     string      `json:"row"`
	Columns interface{} `json:"columns"`
}

type ComundaFormComponent struct {
	Label string `json:"label"`
	Type  string `json:"type"`
	ID    string `json:"id"`
	Key   string `json:"key"`
}

type ComundaForm struct {
	ID         string                       `json:"id"`
	Type       string                       `json:"type"`
	Components []ComundaFormComponentLayout `json:"components"`
}

// userTaskApi a client for userTaskApi API
type formTaskApi struct {
	client *Client
}

// GetForm gets the form for a task
func (u *formTaskApi) GetForm(taskId string) (*ComundaForm, error) {
	res, err := u.client.doGet("/task/"+taskId+"/deployed-form", map[string]string{})
	if err != nil {
		return nil, err
	}

	var resp ComundaForm

	if err := u.client.readJsonResponse(res, &resp); err != nil {
		return nil, fmt.Errorf("can't read json response: %w", err)
	}
	return &resp, nil
}
