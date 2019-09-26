
# Megaphone Code Challenges

---

## Installing the Acme Cult CSV Uploader
1. **Clone Challenge repo** Front End & Back End
2. ** Bundle Install**
3. ** Run Test **
    - npm test front-end
    - rspec back-end


### Functionality
- Provides an interface for a user to upload the salesdata.csv file in this directory
- Parses and persists the information in the salesdata.csv file to a database
- Calculates and displays the total sales revenue to the user

### Scaling Considerations
- Include pagination for larger CSV data sets
- validations for only CSV files
- Validation and auth in back-end
- models are created based on headers of csv not hard coded incase csv structure changes
- Depending on design requirement details may parse csv in frontend
-  Use threads to parse larger CSV
-  Utilize optimistic rendering for better UI performance

---