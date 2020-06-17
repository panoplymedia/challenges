# Sales Admin

## Background

Your company just acquired Acme Cult Hero Supplies. They have been using a CSV worksheet to track sales data, and you need to transform that into a web application to track revenue.

## Functional requirements

Using the web framework of your choice, deliver an application that meets the following requirements:

* Provides an interface for a user to upload the salesdata.csv file in this directory
* Parses and persists the information in the salesdata.csv file to a database
* Calculates and displays the total sales revenue to the user

## Delivery requirements

Please provide instructions for installing and running your application, including all dependencies. The simpler, the better, but we do use PostgreSQL if you want use that as a data store.

+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
I built this project using a Rails API, Postgres, and React. The theme was handled by Bootstrap.
Though it seemed simple, building out a full-stack application always poses some challenges. Some interesting twists and turns included: 
    - setting up JWT authentication - with more time, I would have explored differnet methods of storing and authenticating tokens. For now, when the user is logged in the token is saved in the browser's local storage and sent as a header to authorize requests. I would have also implemented a logout feature!
    - This was the first time I have worked with react hooks, and I had a fun time implementing them. I would definitely revisit some of the components to clean up the code and see if I could make things even more modular. 
    - Encoding form data to send to the back-end was a challenge
    - Routing and redirects could be cleaned up as well

All in all, it was a fun challenge! Thank you for your consideration!

Models 
- item - belongs_to :customer belongs_to :merchant
- customer - has_many :items
- merchant - has_many :items
- user

Features
- interface to upload CSV data to DB
- calculate sales data for user, total sales revenue
- auth

Gems added - 
- bcrypt 3.1.7
- jwt
- jbuilder
- figaro
- rack-cors

to run the app locally:
- clone this repository
- `cd sales_admin_app`
- install dependencies with `bundle install`
    - install client dependencies with `cd client && npm install`
    - remember to move back to the `sales_admin_app` directory
- setup db 
    - `rails db:setup`
    - `rails db:migrate`
    - `rails db:seed`
- `rake start`

to test: 
- `rails t`

sources: 
https://medium.com/technoetics/create-basic-login-forms-using-react-js-hooks-and-bootstrap-2ae36c15e551
https://developer.okta.com/blog/2019/10/02/jwt-react-auth


