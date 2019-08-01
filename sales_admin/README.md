# Sales Admin

## Background

Your company just acquired Acme Cult Hero Supplies. They have been using a CSV worksheet to track sales data, and you need to transform that into a web application to track revenue.

## Functional requirements

Using the web framework of your choice, deliver an application that meets the following requirements:

* Provides an interface for a user to upload the salesdata.csv file in this directory
* Parses and persists the information in the salesdata.csv file
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


<br/>
<br/>
# Brian edits

## Run the code locally
Pull the entire repo into a local directory

	git clone git@github.com:walshbm15/challenges.git

### Server code
1. Create a Python virtual environment and enter it. *This was done using Python 3.6.5*

	```
	pip install virtualenv
	virtualenv megaphone --python=python3
	source megaphone/bin/activate 
	```
2. Switch to the backend directory

	```
	cd challenges/sales_admin/backend/
	```
3. Install Python dependencies

	```
	pip install -r requirements.txt
	```
4. Create and initialize the database
	*Note for simplicity this comes packaged to work with sqlite but can be easily reconfigured to work with another database.*
	
	```
	python manage.py db upgrade
	```
5. Run the server

	```
	python run_local.py
	```

### Client code
1. Make sure npm is installed and updated. *This was done with npm 6.9.0.*
2. Switch to the frontend/sales_data directory

	```
	cd challenges/sales_admin/frontend/sales_data/
	```
3. Install dependencies

	```
	npm install
	```
4. Run client code

	```
	npm run dev
	```

### Accessing the application
Once both the Flask server and client code are running you can access the application locally by going to [localhost:8080](localhost:8080). 

This will bring you to the login screen where you can use the default user (`username=admin and password=admin`) to upload sales data and get total revenue. 

If you wish to create your own user for authentication you can do so by navigating to the Register page. For convenience there are links at the top of the page to help navigating to the corresponding pages. *To use the Upload and the Total Revenue pages successfully you must first have a user logged in.*

## Testing
The code base includes some very basic tests. With more time the tests and code coverage would be greatly improved on.

### Server tests
To run the server tests run:

	pytest challenges/sales_admin/backend/app/tests/test_api


### Client tests
To run the client tests, from the directory `challenges/sales_admin/frontend/sales_data` run:

	npm run test

