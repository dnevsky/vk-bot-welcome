package vkapi

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/dnevsky/vk-bot-welcome/models"
	"github.com/dnevsky/vk-bot-welcome/pkg"
)

const (
	VK_API_URL = "https://api.vk.com/method/"
)

type VKApi struct {
	token   string
	version string
}

func NewVKApi(token string) *VKApi {
	return &VKApi{token: token, version: "5.131"}
}

func (v *VKApi) SendMessageWithKeyboard(text string, peer_id int, keyboard models.Keyboard) error {
	rand.Seed(time.Now().Unix())

	msg := models.MessageSend{
		UserId:      peer_id,
		PeerId:      peer_id,
		RandomId:    rand.Int(),
		Message:     text,
		Keyboard:    keyboard,
		Version:     v.version,
		AccessToken: v.token,
	}

	msgValues, err := pkg.StructToURLValues(msg)
	if err != nil {
		return err
	}

	if len(keyboard.Buttons) == 0 {
		msgValues.Del("keyboard")
	}

	resp, err := http.PostForm(VK_API_URL+"messages.send", msgValues)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "response") {
		return errors.New(fmt.Sprintf("error while send message: %s", string(body)))
	}

	return nil
}
