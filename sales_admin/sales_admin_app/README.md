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

TODO:
- test file upload
- build client

