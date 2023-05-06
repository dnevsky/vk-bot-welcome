package vkapi

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/dnevsky/vk-bot-welcome/models"
)

func (v *VKApi) handleMessageNew(obj models.MessageNewObject) {
	log.Println(obj)

	payload := make(map[string]interface{})

	if len(obj.Message.Payload) != 0 {
		err := json.Unmarshal([]byte(obj.Message.Payload), &payload)
		if err != nil {
			log.Println(err)
			return
		}

		if value, ok := payload["button"]; ok {
			switch value {
			case "main":
				kboard := mainKeyboard()

				err = v.SendMessageWithKeyboard("Возвращаю, что ещё?", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "1":
				kboard := models.Keyboard{
					OneTime: false,
					Buttons: [][]models.Button{},
				}

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"11"}`,
						Label:   `Идём вабанк`,
					},
					Color: "positive",
				})

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"main"}`,
						Label:   `Верни меня обратно`,
					},
					Color: "primary",
				})

				err = v.SendMessageWithKeyboard("Ты хочешь чтобы я передал всему миру привет? Тебе это обойдётся в копечку", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "2":
				kboard := models.Keyboard{
					OneTime: false,
					Buttons: [][]models.Button{},
				}

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"21"}`,
						Label:   `Круто, а github можно?`,
					},
					Color: "positive",
				})

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"main"}`,
						Label:   `Верни меня обратно`,
					},
					Color: "primary",
				})

				err = v.SendMessageWithKeyboard("ЁЛКИ ПАЛКИ!!!1", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "3":
				kboard := models.Keyboard{
					OneTime: false,
					Buttons: [][]models.Button{},
				}

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"31"}`,
						Label:   `Круто, а какие ещё фразочки?`,
					},
					Color: "positive",
				})

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"main"}`,
						Label:   `Верни меня обратно`,
					},
					Color: "primary",
				})

				err = v.SendMessageWithKeyboard("— Скажи моё имя.\n— Хайзенберг.\n— Ты чертовски прав.", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "4":
				kboard := models.Keyboard{
					OneTime: false,
					Buttons: [][]models.Button{},
				}

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"41"}`,
						Label:   `Блин(`,
					},
					Color: "positive",
				})

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"main"}`,
						Label:   `Верни меня обратно`,
					},
					Color: "primary",
				})

				err = v.SendMessageWithKeyboard("Ахах, самоуничтожение, ты что? ОТСТАВИТь САМОУНИЧТОЖЕНИЕ!!!.\nHAPPY LIVES MATTER!!!", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "11":
				kboard := mainKeyboard()

				err = v.SendMessageWithKeyboard("Ну с тебя тогда сто мильён денег.\n\nВсем привет!!!!!", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "21":
				kboard := mainKeyboard()

				err = v.SendMessageWithKeyboard("Можно.\nhttps://github.com/dnevsky", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "31":
				walterWhite := []string{
					"Моя жена ждет моей смерти. Теперь у меня остался только бизнес, только он, а ты хочешь его отобрать.",
					"За чистые машины и чистые деньги!",
					"— У меня идея получше.\n— Ну, слава Богу. Давай, мистер Уайт, выкладывай.\n— Бобы.\n— Бобы?!\n— Бобы клещевины.\n— И чё с ними будем делать? Вырастим волшебный стебель до неба, да? Заберёмся и убежим?\n— Мы сделаем из них рицин.\n— Ты сказал кретин?\n— Рицин!",
					"— Пока с мухой не разберёмся, варить ме...ин не будем.",
					"Мой бизнес — не ме...ин. Мой бизнес — империя.",
					"— Идти можешь?\n— Да.\n— Так иди отсюда нахер и никогда не возвращайся.",
					"Мне не грозит опасность, Скайлер, я сам опасность! Кто-то откроет дверь и схватит пулю. Думаешь, им буду я?! Нет. Это я постучу в дверь.",
					"— ... У меня планы.\n— Курить мар.х.а.у, жрать «Читос» и дрочить не подходит под определение «планы».",
					"— Джесси, ты теперь миллионер, и до сих пор ноешь? Да в каком мире ты живешь?\n— В мире, где тех, кто реально горбатится, не трахают кулаком в жопу.",
					"Позитив — это хорошо. Я за позитив, но позитив не меняет реальность.",
					"Можно вопрос? Куда ушли чистые бизнес-отношения? Почему только я один продолжаю вести себя профессионально?",
					"— Как ты ее так быстро усыпляешь? Ты такой скучный?\n— Я успокаивающий.",
					"Я один должен ощутить последствия моего выбора, и никто другой. А те последствия всё ближе. Хватит оттягивать неизбежное.",
					"Меня греет мысль, что бриллиант и женщина, надевшая его на палец, состоят из одного и того же элемента.",
					"О Боже... Я столько лгу, что начинаю путаться в своей лжи.",
					"Если мы делаем что-то во благо, то нам не о чем переживать... и нет блага выше, чем семья.",
					"— Пошёл хоронить трупы?\n— Грабить поезд.",
					"— Это какой химический элемент?\n— Проволока!",
					"Я спас тебе жизнь, Джесси. Спасешь ли ты мою?",
					// всё хватит
				}

				kboard := models.Keyboard{
					OneTime: false,
					Buttons: [][]models.Button{},
				}

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"31"}`,
						Label:   `Круто, а какие ещё фразочки?`,
					},
					Color: "positive",
				})

				kboard.AddButton(0, models.Button{
					Action: models.ButtonText{
						Type:    "text",
						Payload: `{"button":"main"}`,
						Label:   `Верни меня обратно`,
					},
					Color: "primary",
				})

				rand.Seed(time.Now().Unix())

				quote := walterWhite[rand.Intn(len(walterWhite))]

				err = v.SendMessageWithKeyboard(quote, obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			case "41":
				kboard := mainKeyboard()

				err = v.SendMessageWithKeyboard("Шо блин, плохо о таком думать, надо больше позитива.\n\n\nhttps://youtu.be/dQw4w9WgXcQ", obj.Message.PeerID, kboard)
				if err != nil {
					log.Println(err)
					return
				}

				return
			}
		}
	}

	kboard := mainKeyboard()

	err := v.SendMessageWithKeyboard("Привет, мой уважаемый друг", obj.Message.PeerID, kboard)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func mainKeyboard() models.Keyboard {
	kboard := models.Keyboard{
		OneTime: false,
		Buttons: [][]models.Button{},
	}

	kboard.AddButton(0, models.Button{
		Action: models.ButtonText{
			Type:    "text",
			Payload: `{"button":"1"}`,
			Label:   `Написать в чат "Всем ку!"`,
		},
		Color: "primary",
	})

	kboard.AddButton(0, models.Button{
		Action: models.ButtonText{
			Type:    "text",
			Payload: `{"button":"2"}`,
			Label:   `Написать в чат "Ёлки палки!"`,
		},
		Color: "primary",
	})

	kboard.AddButton(1, models.Button{
		Action: models.ButtonText{
			Type:    "text",
			Payload: `{"button":"3"}`,
			Label:   `Передать привет Walter White`,
		},
		Color: "primary",
	})

	kboard.AddButton(1, models.Button{
		Action: models.ButtonText{
			Type:    "text",
			Payload: `{"button":"4"}`,
			Label:   `Режим самоликвидации`,
		},
		Color: "negative",
	})

	return kboard
}
