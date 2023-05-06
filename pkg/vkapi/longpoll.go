package vkapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/dnevsky/vk-bot-welcome/models"
)

func (v *VKApi) GetLongPollServer(group_id int) (models.LongPoolServerResponse, error) {
	var response models.LongPoolServerResponse
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s?group_id=%d&v=%s",
		VK_API_URL, "groups.getLongPollServer", group_id, v.version), nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", v.token))

	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (v *VKApi) StartLongPoll(group_id int) error {
	client := &http.Client{}

	lpserver, err := v.GetLongPollServer(group_id)
	if err != nil {
		return err
	}

	for {
		req, err := http.NewRequest("GET", fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=25",
			lpserver.Response.Server, lpserver.Response.Key, lpserver.Response.Ts), nil)

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", v.token))

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		res, err := parseResponse(resp.Body)
		if err != nil {
			return err
		}

		switch res.Failed {
		case 0:
			lpserver.Response.Ts = res.Ts
		case 1:
			lpserver.Response.Ts = res.Ts
		case 2:
			lpserver, err = v.GetLongPollServer(group_id)
			if err != nil {
				return err
			}
			continue
		case 3:
			lpserver, err = v.GetLongPollServer(group_id)
			if err != nil {
				return err
			}
			continue
		default:
			return errors.New(fmt.Sprintf("something error: %d", res.Failed))
		}

		go v.handle(res.Updates)
	}
}

func (v *VKApi) handle(obj []models.GroupEvent) {
	for _, value := range obj {
		if value.Type == "message_new" {
			var msgObj models.MessageNewObject

			err := json.Unmarshal(value.Object, &msgObj)
			if err != nil {
				log.Println(err)
			}

			go v.handleMessageNew(msgObj)

		}
	}
}

func parseResponse(reader io.Reader) (response models.Response, err error) {
	decoder := json.NewDecoder(reader)

	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return response, err
		}

		t, ok := token.(string)
		if !ok {
			continue
		}

		switch t {
		case "failed":
			raw, err := decoder.Token()
			if err != nil {
				return response, err
			}

			response.Failed = int(raw.(float64))
		case "updates":
			var updates []models.GroupEvent

			err = decoder.Decode(&updates)
			if err != nil {
				return response, err
			}

			response.Updates = updates
		case "ts":
			raw, err := decoder.Token()
			if err != nil {
				return response, err
			}

			if ts, ok := raw.(float64); ok {
				response.Ts = strconv.Itoa(int(ts))
			} else {
				response.Ts = raw.(string)
			}
		}
	}

	return response, err
}
