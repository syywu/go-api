# Winelist API

## This Winelist API is deployed on Railway.app.

The Winelist API is a RESTful API that simply records all the good-tasting wines.

<br>

## Existing Routes

| Method | Path           | Additional Info | Result                      |
| ------ | -------------- | --------------- | --------------------------- |
| GET    | /winelist      |                 | all wines                   |
| GET    | /winelist/{id} |                 | get wine by a particular ID |
| POST   | /winelist      | { body }        | create a new wine           |
| PUT    | /winelist/{id} | { body }        | update wine by ID           |
| DELETE | /winelist/{id} |                 | delete wine by ID           |

<br>
