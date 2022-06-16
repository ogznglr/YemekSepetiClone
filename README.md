# YemekSepetiClone
Golang ile yemeksepeti uygulamasını yazdım. Yeni hesap oluşturma, login olma, logout olma, yeni sipariş oluşturma, mevcut siparişleri listleme gibi özellikler mevcut.


I made an order food application. You can create new account, login, logout, create a new order and list your orders by using this REST Api.

The documentation of using the api:

REQUEST TYPES: 

//REGISTERS A NEW ACCOUNT RETURNS YOUR INFOS BACK
POST(localhost:8080/api/register)
   
   {
     "name":"yourname",
     "surname":"yoursurname",
     "password":"yourpassword",
     "email":"youremail"
   }

//LOGS IN AND RETURNS YOUR TOKEN
POST(localhost:8080/api/login)

   {
     "email":"youremail",
     "password":"yourpassword"
   }

//LOGS OUT DELETES YOUR TOKEN
Get(localhost:8080/api/logout)

//LISTS THE ORDERS BY YOUR USER ID
GET(localhost:8080/api/order)

//CREATES A NEW ORDER
POST(localhost:8080/api/order)

   {
     "orderid":"orderid",
     "restoranid":"restoranid"
   }
