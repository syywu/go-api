# Winelist API

## This [Winelist API](https://go-api-production-ad21.up.railway.app/winelist) is deployed on Railway.app.

The Winelist API is a RESTful API that simply record all the good-tasting wines.

<br>

## Existing Routes

| Method | Path        | Additional Info | Result                        |
| ------ | ----------- | --------------- | ----------------------------- |
| GET    | /winelist   |                 | all wines                     |
| GET    | /books/{id} |                 | get a wine by a particular ID |
| POST   | /books      | { body }        | create a new wine             |
| PUT    | /books/{id} | { body }        | update wine by ID             |
| DELETE | /books/{id} |                 | delete wine by ID             |

<br>
