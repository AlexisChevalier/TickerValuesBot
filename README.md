# TickerValuesBot
Simple Telegram bot returning exchange ticker values coded in GO

Code is not pretty, I just wanted to try the Telegram bot API.

Uses https://www.cryptonator.com/api as data provider

Usage
=========

```` 
go run main.go <Telegram Bot Token>
```` 

Telegram Usage
==========

Send a message to the bot formatted like this:

````
BTCEUR
```` 

BTCEUR will return the price of 1 BTC in EUR

![Usage example](https://i.gyazo.com/256afe9cc781eac7716a6003d2ac858b.png)
