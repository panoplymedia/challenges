# Sales Admin

## Author: Michael Byrne

## The Application and How to Run It

The Acme admin application runs in Docker. So, with Docker installed, you can just run `docker-compose up`. This will build the frontend and backend and make the app available at localhost:3001. The interface is really simple: just upload your file where indicated and the app will display your sales total. In the background, the CSV is being uploaded to the backend, which stores the data in a Postgres table, and then it runs a pg query that computes the total. Authentication occurs via a (hardcoded) API key exchanged between the frontend and backend.

Of course, the revenue total could have been computed in any number of ways: immediately on the frontend, via some backend logic, on the frontend after a backend call that returns the CSV data. Also, the initial upload request could have returned the total itself, but that would require some coupling of different sorts of tasks. As is, there are two endpoints: one to post the file to, and a second that returns the revenue total.  

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
