# Sales Admin

Implementation by @Daniel-Edminster


### Notes on scalability
In a production environment, there are a fair number of things I'd do differently. I'll break these down by category.

##### Database:
I'd split this into 4 or so tables for data organization:


| products | merchants | customer | orders |
|---|---|---|---|
| id | id | id | id |
|name|name|first_name|customer_id|
|price | address | last_name | product |
|merchant_id | | address | |


and create the same viewable table via `JOIN`. This should allow for a lot of expandability, features, and the kitchen sink. 

##### Revenue Calculation:
Presently not scalable, is done via a backend call that adds the total of each row in the database. Ideally you'd run something like this on either a copy of the database and let it take its time, or run segmented reporting like monthly, annualized, etc. Decided not to do client-side for fear of potential test cases crashing the browser.  


##### Authentication:
Quick and dirty GitHub OAuth2 implementation using express and express-session. Express-session is not scalable seemingly by design. Sessions are stored in server memory with no overflow control like an LRU Cache. For scalability's sake, probably better to use a redis store or a file-based session store in a similar fashion as Apache Server.
####Setup:
```
git clone https://github.com/Daniel-Edminster/challenges
```
##### Front-end:
```
cd challenges/sales_admin/client/salesadmin
npm install
npm run start
```
##### Back-end:
##### Postgres Setup:
```
brew install postgres
psql
```
then run the creation queries in `challenges/salesadmin/server/schema.sql`, but not as an infile.
Lastly, edit the contents of `challenges/salesadmin/server/db.example.js`,
change to your appropriate credentials, and **resave as `db.js`**. 

##### Node setup:

```
cd challenges/sales_admin/server
npm install
```
You'll need to edit `challenges/sales_admin/server/config.example.js` to get started.
This application uses GitHub for Authentication, so you'll need to setup an OAuth app [here](https://github.com/settings/applications/new), with the following credentials:
**Homepage URL**: `http://localhost:3000`
**Authorization callback URL**: `http://localhost:8080/auth/github/callback`

Modify the contents of `config.example.js`:
```
    GITHUB_CLIENT_ID: "YOUR_CLIENT_ID_HERE",
    GITHUB_CLIENT_SECRET: "YOUR_CLIENT_SECRET_HERE",
    GITHUB_CALLBACK_URL: "http://localhost:8080/auth/github/callback"
```
and **resave as `config.js`**.

That should be everything, so fire it up with:
```
node index.js
```



## Background

Your company just acquired Acme Cult Hero Supplies. They have been using a CSV worksheet to track sales data, and you need to transform that into a web application to track revenue.

## Functional requirements

Using the web framework of your choice, deliver an application that meets the following requirements:

* Provides an interface for a user to upload the salesdata.csv file in this directory
* Parses and persists the information in the salesdata.csv file to a database
* Calculates and displays the total sales revenue to the user

Bonus points if you add authentication.

_Ideally you shouldn't spend more than 4-5 hours on your solution, but take as much time as you want._

## Delivery requirements

Please provide instructions for installing and running your application, including all dependencies. The simpler, the better, but we do use PostgreSQL if you want use that as a data store.

Think about things like:

* Testing
* How to store the data
* How would your solution differ if it had to scale?

Please submit your solution as a pull request, or package it up and send it to doug.ramsay@megaphone.fm.

## Credits

Yes, this challenge is copied from the LivingSocial code challenge. I helped put that together, so hopefully nobody will mind that much.
