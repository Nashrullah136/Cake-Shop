# Cake Shop

## Requirement

  - Docker 20.10.17+
  - Docker Compose 3+

## Run the app

    docker-compose build
    docker-compose up

## Run the tests

    go test ./...

# REST API

- [Get All Cakes](#get-all-cakes)
- [Get Cake By ID](#get-cake-by-id)
- [Create Cake](#create-cake)
- [Update Cake](#update-cake)
- [Delete Cake](#delete-cake)


## Get All Cakes
Get all cake and sort by rating descend and title ascend. 
If page and items url query less than 0, all cake will be returned.

### Request

  Method : GET

  Endpoint : '/cakes'

  URL Query :
  - page type : int
  - items type : int

### Response
  
  Response Body :
  
    [
      {
        "id": int
        "title": str,
        "description": str,
        "rating": int
        "image": str,
        "created_at": str,
        "updated_at": str
      },
      ...
    ]
  
  Response example:
    
    [
      {
        "id": 2
        "title": "Apple Pie",
        "description": "Pie with slices of apple within",
        "rating": 10
        "image": "https://upload.wikimedia.org/wikipedia/commons/6/61/Small_apple_pie_8.jpg",
        "created_at": "2020-02-01 10:56:31",
        "updated_at": "2020-02-13 09:30:23"
      },
      {
        "id": 1,
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7.0,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2020-02-01 10:56:31",
        "updated_at": "2020-02-13 09:30:23"
      }
    ]

## Get Cake By ID

### Request

  Method : GET
  
  Endpoint : '/cakes/:id'
  
  - id type: int

### Response
  
  Response Body:
  
    {
      "id": int
      "title": str,
      "description": str,
      "rating": int
      "image": str,
      "created_at": str,
      "updated_at": str
    }
   
  Response Example:
  
    {
      "id": 1,
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7.0,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
      "created_at": "2020-02-01 10:56:31",
      "updated_at": "2020-02-13 09:30:23"
    }
        
## Create Cake

### Request

Method : POST

Endpoint : '/cakes'

Request Body:

    {
      "title": str,
      "description": str,
      "rating": int
      "image": str
    }
 
Request Body Constraint:

- title : string, required, not blank
- description : string, required, not blank
- rating : float, required, > 0
- image : string, required, not blank

Request Body Example:

    {
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7.0,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
    }

### Response
  
  Response Body:
  
    {
      "id": int
      "title": str,
      "description": str,
      "rating": int
      "image": str,
      "created_at": str,
      "updated_at": str
    }
   
  Response Example:
  
    {
      "id": 1,
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7.0,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
      "created_at": "2020-02-01 10:56:31",
      "updated_at": "2020-02-13 09:30:23"
    }
  
## Update Cake

### Request

Method : PATCH

Endpoint : '/cakes/:id'

 - id type : int

Request Body:

    {
      "title": str,
      "description": str,
      "rating": int
      "image": str
    }
 
Request Body Constraint:

- title : string, not blank
- description : string, not blank
- rating : float, > 0
- image : string, not blank

Request Body Example:

    { // Update all field
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7.0,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
    }
    
    { //update title only
      "title": "Lemon cheesecake" 
    }

### Response
  
  Response Body:
  
    {
      "id": int
      "title": str,
      "description": str,
      "rating": int
      "image": str,
      "created_at": str,
      "updated_at": str
    }
   
  Response Example:
  
    {
      "id": 1,
      "title": "Lemon cheesecake",
      "description": "A cheesecake made of lemon",
      "rating": 7.0,
      "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
      "created_at": "2020-02-01 10:56:31",
      "updated_at": "2020-02-13 09:30:23"
    }
    
    
## Delete Cake

### Request

  Method : DELETE

  Endpoint : '/cakes/:id'
  
  - id type : int

### Response
  The cake will be sort by rating descend and title ascend. 
  
  Response Body :
  
    {
      "Successful"
    }

