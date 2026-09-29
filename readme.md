# Teepee Marketplace API Documentation

## Overview

Frontend uses HTML, CSS, and Vanilla JavaScript to communicate with the
Go backend API using fetch().

------------------------------------------------------------------------

# How to run

1.install docker
2.pull the repository
3.run docker compose -f {path_to_repo}/docker_compose.yaml up
4.the service is accessable at the localhost with 1323 port but changable inside the docker env
5.inside docker if the env seeddb is set to false it won't automatically filled the db with seed data


------------------------------------------------------------------------

# API Response Format

All APIs return:

``` json
{
    "status": true,
    "message_error": "",
    "data": {}
}
```

The frontend reads the payload from `data`.

If `status` is false, the frontend displays `message_error`.

------------------------------------------------------------------------

# Running Application

Start backend:

``` bash
go run ./cmd/webservice
```

Serve static frontend:

``` go
e.Static("/static", "static")
```

Access:

    GET /web

------------------------------------------------------------------------

# Shop API

## Add Shop

Endpoint:

    POST /shops/addshop

Request type:

    multipart/form-data

Fields:

    name
    location

------------------------------------------------------------------------

## View Shop List

Endpoint:

    GET /shops/viewshoplist

Response:

``` json
[
    {
        "id":"shop-id",
        "name":"Shop A",
        "location":"Bali"
    }
]
```

------------------------------------------------------------------------

## View Shop Items

Endpoint:

    GET /shops/:shop_id/viewitem

Example:

    GET /shops/123/viewitem

Response:

``` json
{
    "name":"Shop A",
    "location":"Bali",
    "id":"123",
    "items":[
        {
            "id":"item1",
            "name":"Keyboard",
            "stock":20,
            "available":"available"
        }
    ]
}
```

------------------------------------------------------------------------

# Product API

## Add Product

Endpoint:

    POST /products/additem

Request:

    name
    price
    stock
    available
    shop_id

Uses:

    multipart/form-data

------------------------------------------------------------------------

## View Product Detail

Endpoint:

    GET /products/viewitemdetail/:item_id

------------------------------------------------------------------------

## Add Product Stock

Endpoint:

    PUT /products/additemstock

Request:

    item_id
    stock

------------------------------------------------------------------------

## Update Availability

Endpoint:

    PUT /products/updateavailabilty

Request:

    id
    availabilty

Values:

    available
    not_available

------------------------------------------------------------------------

## Checkout Product

Endpoint:

    POST /products/:id/checkout

Body:

``` json
{
    "quantity":3
}
```

Rules:

-   Product must exist
-   Product must be available
-   Stock must be sufficient
-   Stock cannot become negative

------------------------------------------------------------------------

# Frontend Functions

Load shop list:

    loadShop()

Load shop products:

    loadItems(shopId)

Add stock:

    addStock(itemId)

Checkout:

    checkout(itemId)

Change availability:

    changeAvailability(itemId, availability)

------------------------------------------------------------------------

# Request Types

Use FormData for:

-   Add Shop
-   Add Product
-   Add Stock
-   Update Availability

Do not manually set multipart Content-Type because the browser creates
the boundary.

Use JSON for checkout requests.

