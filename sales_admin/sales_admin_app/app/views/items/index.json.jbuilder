json.array! @items do |item|
  json.id item.id
  json.description item.description
  json.price item.price
  json.quantity item.quantity
  json.merchant item.merchant
  json.customer item.customer
end
