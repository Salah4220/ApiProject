
# ApiArtWorks - salah habibeche

## Description

the API provides general information on works of art, their year of realization, their director. 


## Installation

...

## Use

from the command line in the directory containing main.go, run the code:

$ go run .

Once the code is running, you have a running HTTP server to which you can send requests.

From a new command line window, use curl to make a request to your running web service.

$ curl http://localhost:8080/Oeuvres

From a different command line window, use curl to make a request to your running web service.

$ curl http://localhost:8080/Oeuvres\
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","Year": 1850}'


From a different command line window, use curl to make a request to your running web service.

$ curl http://localhost:8080/Oeuvres/2



...

#### Get all items

```http
  GET /api/items
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API key |

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |


## Authors

- [@Salah4220](https://github.com/Salah4220)
- Salah Habibeche

