package controllers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/fgmaia/chat/internal/entities"
	"github.com/google/uuid"
)

type ControllerHandler interface {
	Handler(ctx context.Context, buffer []byte) ([]byte, error)
}

type controllerHandler struct {
	actions map[string]func(context.Context, entities.Payload) (entities.Payload, error)
}

func NewControllerHandler() ControllerHandler {
	c := &controllerHandler{}
	c.loadActions()
	return c
}

func (c *controllerHandler) Handler(ctx context.Context, buffer []byte) ([]byte, error) {

	//get request
	var request entities.Payload
	err := json.Unmarshal(buffer, &request)

	if err != nil {
		return nil, err
	}

	//callAction
	response, err := c.callAction(ctx, request)

	//set response
	data, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *controllerHandler) send(ctx context.Context, request entities.Payload) (entities.Payload, error) {
	//TODO call use case
	//TODO call use case
	//TODO call use case
	//TODO call use case
	if request.UserID == "" {
		request.UserID = uuid.New().String()
	}
	request.Action = "receive"
	request.ReceiveMessages = []entities.Message{
		{
			ID:       uuid.New().String(),
			UserID:   uuid.New().String(),
			UserName: "teste usur 2",
			Content:  "kjfhkadhkjashdjk ashjkd",
		},
	}

	return request, nil
}

func (c *controllerHandler) loadActions() {
	c.actions = make(map[string]func(context.Context, entities.Payload) (entities.Payload, error))
	c.actions["send"] = c.send
}

func (c *controllerHandler) callAction(ctx context.Context, request entities.Payload) (entities.Payload, error) {
	if f, ok := c.actions[request.Action]; ok {
		return f(ctx, request)
	}
	return entities.Payload{}, errors.New("unknown action")
}
