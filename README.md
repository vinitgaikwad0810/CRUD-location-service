# cmpe273-assignment-2

CRUD Location Service 


CMPE273-Fall15-Assignment2
Description

You will be building a location and trip planner service in Go.
Part I - CRUD Location Service [Requirement]

The location service shall have the following REST endpoints to store and retrieve locations. All the data must be persisted into MongoDB. You may want to install MongoDB locally for development testing and use free MongoLab for submission. For Go application to MongoDB, you can use driver like mgo. If you don’t like mgo, feel free to use any Go MongoDB drivers.
To lookup coordinates of a location, use Google Map Api.
Example:
Get coordinates of 1600 Amphitheatre Parkway, Mountain View, CA.
http://maps.google.com/maps/api/geocode/json?address=1600+Amphitheatre+Parkway,+Mountain+View,+CA&sensor=false
Create New Location - POST        /locations
Request:
POST /locations
{
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113"
}
Response
HTTP Response Code: 201
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113",
   "coordinate" : { 
      "lat" : 38.4220352,
     "lng" : -222.0841244
   }
}
Get a Location - GET        /locations/{location_id}
Request
GET /locations/12345
Response
HTTP Response Code: 200
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "123 Main St",
   "city" : "San Francisco",
   "state" : "CA",
   "zip" : "94113",
   "coordinate" : { 
      "lat" : 38.4220352,
     "lng" : -222.0841244
   }
}
Update a Location - PUT /locations/{location_id}
Request:
PUT /locations/12345
{
   "address" : "1600 Amphitheatre Parkway",
   "city" : "Mountain View",
   "state" : "CA",
   "zip" : "94043"
}
Response
HTTP Response Code: 201
{
   "id" : 12345,
   "name" : "John Smith",
   "address" : "1600 Amphitheatre Parkway",
   "city" : "Mountain View",
   "state" : "CA",
   "zip" : "94043"
   "coordinate" : { 
      "lat" : 37.4220352,
     "lng" : -122.0841244
   }
}
Delete a Location - DELETE /locations/{location_id}
        Request:
DELETE  /locations/12345
        Response:
HTTP Response Code: 200
Published by Google Drive–Report Abuse–Updated automatically every 5 minutes
