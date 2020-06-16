Models 
customer - has_many :items
item - belongs_to :customer belongs_to :merchant
merchant - has_many :items
user

Features
- interface to upload CSV data to DB
- calculate sales data for user, total sales revenue
- auth

Gems
- bcrypt 3.1.7
- jwt
- figaro
- rack-cors