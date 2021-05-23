# Amazon Wrapper API

## Description

Go api that scrapes Amazon product urls.
Plz don't hate I'm just trying to learn Go.

## Run Locally

1. `cd` into directoy.
2. Run `go run main.go`.

Congrats! API will be running on port `8080` by default.

## How To Use

From an Amazon url extract the following variables:
`https://www.amazon.ca/${item}/dp/${variant}`
and create a url like:
`http://localhost:8080/item/${item}/${variant}`

## Example

[This](https://www.amazon.ca/Mattel-GAMES-W2085-Card-Game/dp/B00CTH0A1Q) Amazon product page is [this](http://localhost:8080/item/Mattel-GAMES-W2085-Card-Game/B00CTH0A1Q) on the API. It returns the following:

```
{
    "hiResLandingImage": "https://images-na.ssl-images-amazon.com/images/I/71Jzj%2BN%2BpmS._AC_SL1500_.jpg",
    "landingImage": "https://images-na.ssl-images-amazon.com/images/I/71Jzj%2BN%2BpmS.__AC_SX300_SY300_QL70_ML2_.jpg",
    "numReviews": "5,231 ratings",
    "price": "$6.93",
    "productTitle": "UNO Card Game",
    "rating": "4.8 out of 5 stars",
    "url": "https://www.amazon.ca/Mattel-GAMES-W2085-Card-Game/dp/B00CTH0A1Q"
}

```
