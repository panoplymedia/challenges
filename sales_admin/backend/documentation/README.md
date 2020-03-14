
# REST API Documentation

## Endpoints

### Upload Sales Data

**URL** : `/sales`

**Method** : `POST`

**Content-Type** : `multipart/form-data`

**Form Values** : The form should contain a csv file called "file" which is less than 10MB in size.

### Success Responses

**Status Code** : `201 CREATED`

**Body** : Response body will be empty.

### Retrieve Sales Revenue Total

**URL** : `/sales/total`

**Method** : `GET`

### Success Responses

**Status Code** : `200 OK`

**Body** : Response body will contain a JSON object with a "total" property describing the sum of sales revenue from the data available:

```json
{
	"total": 435.14
}
```
---
## Error Responses

### Bad Request

**Condition** : Parameters are invalid.

**Status Code** : `400 BAD REQUEST`

**Body example** :

```json
{
    "message": "Missing Parameters: file"
}
```
---
### Not Found

**Condition** : The requested entity could not be found.

**Status Code** : `404 NOT FOUND`

---
### Internal Server Error

**Condition** : The server encountered an internal error.

**Status Code** : `500 Internal Server Error`
