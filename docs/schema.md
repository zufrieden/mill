# Database schema

## About the format
Based on https://dbdiagram.io/ tool specific DSL format.

Table users {
  id int PK
  first_name varchar
  last_name varchar
  date_of_birth date
  email varchar
  date_start date
  date_end date
  created_at date
}

Table workingdays {
  id int PK
  day date
  working_hours
  effective_hours_am
  effective_hours_pm  
  user_id int
}

Table holidays {

}

Table flagdays {
  id int PK
  flag_date date
  flag_name varchar
}

Ref: workingweeks.user_id > users.id

/////

Table order_items {
  order_id int
  product_id int
  quantity int
}

Table products {
  id int PK
  name varchar
  merchant_id int
  price int
  status varchar
  created_at varchar
}

Table users {
  id int PK
  full_name varchar
  email varchar
  gender varchar
  date_of_birth varchar
  created_at varchar
  country_code int
}

Table merchants {
  id int PK
  merchant_name varchar
  country_code int
  created_at varchar
  admin_id int
}

Table countries {
  code int PK
  name varchar
  continent_name varchar
}

Ref: orders.user_id > users.id

Ref: order_items.order_id > orders.id

Ref: order_items.product_id > products.id

Ref: products.merchant_id > merchants.id

Ref: users.country_code > countries.code

Ref: merchants.admin_id > users.id

Ref: merchants.country_code > countries.code
