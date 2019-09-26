# Megaphone Code Challenges

---

## Installing the Acme Cult CSV Uploader
###**Clone Challenge Repo** 
###CD into acme-cult-hero-supplies-frontend directory
- npm install
- use command npm test to run tests
- npm start
###CD into acme-cult-hero-supplies-backend directory
- bundle install
- use command rspec to run tests
- rails s

### Functionality
- Provides an interface for a user to upload the salesdata.csv file in this directory
- Parses and persists the information in the salesdata.csv file to a database
- Calculates and displays the total sales revenue to the user

### Scaling Considerations
- Include pagination for larger CSV data sets
- validations for only CSV files
- Validation and auth in back-end
- models are created based on headers of csv not hard coded incase csv structure changes
- Depending on design requirement details may parse csv in frontend & post as json
-  Use threads to parse larger CSV
-  Utilize optimistic rendering for better UI performance

### Packages used

- [Faker][1]
- [Factory_bot_rails][2]
- [Database_cleaner][3]
- [Jest][4]

---

[1]:https://github.com/faker-ruby/faker
[2]:https://github.com/thoughtbot/factory_bot_rails
[3]:https://github.com/DatabaseCleaner/database_cleaner
[4]:https://jestjs.io/en/
