# YemekSepetiClone
I made an order food application. You can create a new account, login, logout, create a new order and list your orders by using this REST Api.

The documentation of using the api:

REQUEST TYPES: 

//REGISTERS A NEW ACCOUNT RETURNS YOUR INFOS BACK <br/>
POST(localhost:8080/api/register)<br/>

   `{
     "name":"yourname",
     "surname":"yoursurname",
     "password":"yourpassword",
     "email":"youremail"<br/>
   }`

//LOGS IN AND RETURNS YOUR TOKEN<br/>
POST(localhost:8080/api/login)<br/>

   `{
     "email":"youremail",
     "password":"yourpassword"
   }`

//LOGS OUT DELETES YOUR TOKEN<br/>
Get(localhost:8080/api/logout)<br/>

//LISTS THE ORDERS BY YOUR USER ID<br/>
GET(localhost:8080/api/order)<br/>

//CREATES A NEW ORDER<br/>
POST(localhost:8080/api/order)<br/>

   `{
     "orderid":"orderid",
     "restoranid":"restoranid"
   }`
