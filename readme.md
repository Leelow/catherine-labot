# catherine-labot

> A tiny slack bot which gives you weather information.

## Philosophy

`catherine-labot` is just a proof of concept to show how it can be easy to play with API to have a bot which is enough "clever" to understand natural langage and give you weather information (five days forecast due to the API limitation).
The name `catherine-labot` is just a tribute of the famous french weather presenter [Catherine Laborde](https://fr.wikipedia.org/wiki/Catherine_Laborde) retired in 2017, but don't worry she speaks english ;)

## Requirements

Please make sur Golang is installed correctly. The bot was tested on Golang 1.8.

## Config

The config file `config/config.json` contains api keys. You must edit it before running the bot.

```json
    {
      "slack_token": "SLACK_TOKEN",
      "openweathermap_api_key": "OPENWEATHERMAP_API_KEY",
      "meaningcloud_api_key": "MEANINGCLOUD_API_KEY"
    }

```

The free plan of [OpenWeatherMap](https://openweathermap.org/) and [MeaningCloud](https://www.meaningcloud.com/) are enough to play and have fun discussing with the bot !

## Usage

After setting your bot in Slack, you can discuss with it mentioning it every time with `@<your_bot_name>`.
For instance :

    @catherine-labot What is the weather at Rennes today at 5pm ?
    > The estimated temperature at Rennes the 07/04 at 05 pm is 30°C. A clear sky is also planned.

    @catherine-labot Is it raining next friday at Paris ?
    > The estimated temperature at Paris the 07/07 is 24°C. A light rain is also planned.
    
    @catherine-labot Do you like baguette ?
    > Sorry, I don't understand, I am just a weather bot !

## Docker

A `Dockerfile` is present if you want to build `catherine-labot` and tun it as a Docker container.

## Credits

Thanks [@rapidloop](https://github.com/rapidloop) for the [slack library](https://github.com/rapidloop/mybot).

## License

[MIT](LICENSE) © [Léo Lozach](https://github.com/Leelow)
