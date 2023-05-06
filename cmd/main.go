package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dnevsky/vk-bot-welcome/pkg/vkapi"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func main() {
	godotenv.Load(".env")

	if err := init_config(); err != nil {
		log.Fatalf("error loading yml configs: %s", err.Error())
	}

	vk := vkapi.NewVKApi(os.Getenv("VK_TOKEN"))

	log.Println("Start longpoll...")
	err := vk.StartLongPoll(viper.GetInt("vk.group_id"))

	fmt.Println(err)
}

func init_config() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("default")
	return viper.ReadInConfig()
}
