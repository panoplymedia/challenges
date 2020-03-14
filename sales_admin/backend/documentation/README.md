
# REST API Documentation

## Endpoints

### Upload Sales Data

**URL** : `/sales`

**Method** : `POST`

**Content-Type** : `multipart/form-data`

**Form Values** : The form should contain a csv file called "file" which is less than 10MB in size.

### Success Responses

**Status Code** : `201 CREATED`

**Body** : Response body will contain a JSON object with the newly added sales data:

```json
[
    {
        "customer_name": "Jack Burton",
        "item_description": "Premium Cowboy Boots",
        "item_price": 149.99,
        "quantity": 1,
        "merchant_name": "Carpenter Outfitters",
        "merchant_address": "99 Factory Drive"
    },
    {
        "customer_name": "Ellen Ripley",
        "item_description": "Tank Top Undershirt",
        "item_price": 9.5,
        "quantity": 2,
        "merchant_name": "Hero Outlet",
        "merchant_address": "123 Main Street"
    },
    {
        "customer_name": "Lisbeth Salander",
        "item_description": "Black Hoodie",
        "item_price": 49.99,
        "quantity": 4,
        "merchant_name": "Stockholm Supplies",
        "merchant_address": "34 Other Avenue"
    },
    {
        "customer_name": "Butch Coolidge",
        "item_description": "Tank Top Undershirt",
        "item_price": 9.5,
        "quantity": 3,
        "merchant_name": "Hero Outlet",
        "merchant_address": "123 Main Street"
    },
    {
        "customer_name": "Ellen Ripley",
        "item_description": "Stomper Shoes",
        "item_price": 129,
        "quantity": 1,
        "merchant_name": "Parker Footwear",
        "merchant_address": "77 Main Street"
    }
]
```
---

### Retrieve Sales Revenue Total

**URL** : `/sales/total`

**Method** : `GET`

### Success Responses

**Status Code** : `200 OK`

**Body** : Response body will contain a JSON object with the sales data available:

```json
[
    {
        "customer_name": "Jack Burton",
        "item_description": "Premium Cowboy Boots",
        "item_price": 149.99,
        "quantity": 1,
        "merchant_name": "Carpenter Outfitters",
        "merchant_address": "99 Factory Drive"
    },
    {
        "customer_name": "Ellen Ripley",
        "item_description": "Tank Top Undershirt",
        "item_price": 9.5,
        "quantity": 2,
        "merchant_name": "Hero Outlet",
        "merchant_address": "123 Main Street"
    },
    {
        "customer_name": "Lisbeth Salander",
        "item_description": "Black Hoodie",
        "item_price": 49.99,
        "quantity": 4,
        "merchant_name": "Stockholm Supplies",
        "merchant_address": "34 Other Avenue"
    },
    {
        "customer_name": "Butch Coolidge",
        "item_description": "Tank Top Undershirt",
        "item_price": 9.5,
        "quantity": 3,
        "merchant_name": "Hero Outlet",
        "merchant_address": "123 Main Street"
    },
    {
        "customer_name": "Ellen Ripley",
        "item_description": "Stomper Shoes",
        "item_price": 129,
        "quantity": 1,
        "merchant_name": "Parker Footwear",
        "merchant_address": "77 Main Street"
    }
]
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
