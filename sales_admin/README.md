# Sales Admin

![Sales App Home](home.png)
![Sales App Dashboard](dashboard.png)

## Requirements
- docker
- docker-compose
- open ports on 3000 and 5432

## To Run
```bash
docker-compose up
```
Once running, visit localhost:3000 on your browser

## Stack
- API: endpoints and templates built using GoBuffalo
- DB: PostgreSQL

## API

#### Backend
Uses the buffalo app as to route requests, with a pop.Connection middleware
used to access the db. I chose buffalo for this challenge because it allows you
to generate much of the repetitive boilerplate code in a basic API

#### Frontend
I used Bootsrap 4 to style components in the two routes visible on the app's UI. Plush templating was used to create the html

## DB
- PostgreSQL used as persistent storage
- I used Pop to manage the database schema and migrations
- Collections:
    - Merchants
    - Products
        - products belong to one merchant
    - Customers
    - Orders
        - orders have one merchant, customer, and product

## Scale Considerations
- Total revenue calculation is a heavy burden on the database because it is
  being calculated on the server using all rows in the orders table. A sql
  query would be cleaner and faster
- If this application would scale across geographic regions, adding additional
  db nodes and having some sort of data replication among all database nodes
  would improve query time and spread queries across nodes
- As number of records increases, indexing the database would improve query
  times

## Other Considerations and Improvements
- Implement Unit testing
- Better UX on the upload button. Currently it returns an unhandled error if
  the button is pressed before selecting a file. The button should be inactive
  until that point, and flash messages should be used if unsupported filetypes
  are uploaded to the app
- Add validation to objects upon import so that merchant, customer, and product
  rows are not duplicated
- Auth, with roles enabling object editing and deletion 
