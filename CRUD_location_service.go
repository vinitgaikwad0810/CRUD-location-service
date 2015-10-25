package main

import (  
    // Standard library packages
"encoding/json"
"fmt"
"net/http"
    // Third party packages
"github.com/julienschmidt/httprouter"
"gopkg.in/mgo.v2/bson"
"gopkg.in/mgo.v2"
"net/url"


)


type Request struct {
  Id bson.ObjectId `bson:"_id"`
  Name   string `json:"name" bson:"name"`
  Address string `json:"address" bson:"address"`
  City string `json:"city" bson:"city"`   
  State string `json:"state" bson:"state"`
  Zip    string    `json:"zip" bson:"zip"` 
  Latitudes float64 `json:"latitudes" bson:"latitudes"`
  Longitudes float64 `json:"longitudes" bson:"longitudes"`

}



var mgoSession *mgo.Session

var counter int

func postHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {


    fmt.Println("-----------------SERVER LOGS ARE ENABLED---------postHandler------------------")

    //session = getSession()
    userPosted:= new(Request)
    //resp := Response{}
    counter=0

    mgoSession, err := mgo.Dial("mongodb://vinitgaikwad0810:Thisisvinit0810@ds043714.mongolab.com:43714/cmpe273-assignment2-mongodb")

    // Check if connection error, is mongo running? 
    if err != nil {
        panic(err)
    }

    // Populate the user data
    json.NewDecoder(r.Body).Decode(userPosted)
    userPosted.Id = bson.NewObjectId()


    userPosted.Latitudes,userPosted.Longitudes =returnlatlng(userPosted.Address+userPosted.City+userPosted.State+userPosted.Zip)

    fmt.Println("Location co-ordinate found from Google Maps API .... \n")
    fmt.Println("lat =",userPosted.Latitudes);
    fmt.Println("lng =",userPosted.Longitudes);
    // Marshal provided interface into JSON structure
    mgoSession.DB("cmpe273-assignment2-mongodb").C("UserLocation").Insert(userPosted)
    
    userReceived :=new(Request)
    if err := mgoSession.DB("cmpe273-assignment2-mongodb").C("UserLocation").FindId(userPosted.Id).One(userReceived); err != nil {
        w.WriteHeader(404)
        return
    }

    uj, _ := json.MarshalIndent(userReceived,"","\t")
 //   uj, _ := json.MarshalIndent(userReceived,"","\t")


    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "The location data posted is as follows %s!", uj)
    counter ++

    fmt.Println("-----------------SERVER LOGS ARE ENABLED---------postHandler------------------")    
}

func returnlatlng(address string) (float64, float64){
    var Url *url.URL
    Url, err := url.Parse("http://maps.google.com")
    if err != nil {
      panic("Error Panic")
  }

  Url.Path += "/maps/api/geocode/json"
  
  parameters := url.Values{}
  parameters.Add("address", address)
  Url.RawQuery = parameters.Encode()
  Url.RawQuery += "&sensor=false"

   // fmt.Println("URL " + Url.String())

  res, err := http.Get(Url.String())
                  //
  if err != nil {
      panic("Error Panic")
  }
  //  fmt.Println(res)
  defer res.Body.Close()
  var v map[string] interface{}
  dec:= json.NewDecoder(res.Body);
  if err := dec.Decode(&v); err != nil {
      fmt.Println("ERROR: " + err.Error())
  }   

  lat := v["results"].([]interface{})[0].(map[string] interface{})["geometry"].(map[string] interface{})["location"].(map[string]interface{})["lat"].(float64)
  lng := v["results"].([]interface{})[0].(map[string] interface{})["geometry"].(map[string] interface{})["location"].(map[string]interface{})["lng"].(float64)
//fmt.Println("The latittude",lat,"and the longitude",lng)
  return lat,lng
}

func getHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

    fmt.Println("-----------------SERVER LOGS ARE ENABLED---------getHandler------------------")
                   // fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
    id := p.ByName("id")

    fmt.Println("Id given -->"+id)
    mongoIndentifier := bson.ObjectIdHex(id)

    mgoSession, err := mgo.Dial("mongodb://vinitgaikwad0810:Thisisvinit0810@ds043714.mongolab.com:43714/cmpe273-assignment2-mongodb")

    if err != nil {
        panic(err)
    }
                    // Stub user
    userReceived := new(Request)

                    // Fetch user
    if err := mgoSession.DB("cmpe273-assignment2-mongodb").C("UserLocation").FindId(mongoIndentifier).One(userReceived); err != nil {
        w.WriteHeader(404)
        fmt.Fprintf(w, "Record might be absent \n")

        return
    }
                   /* query := mgoSession.DB("go_rest_tutorial").C("users").Find(bson.M{"id":p.ByName("id")})
    fmt.Println(query)*/
                    // Marshal provided interface into JSON structure
    uj, _ := json.MarshalIndent(userReceived,"","\t")

                    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "The JSON response received from mongolab server is as follows %s", uj)

    fmt.Println("-----------------SERVER LOGS ARE ENABLED---------------getHandler---------------------")
}

//curl -XPOST -H 'CONTENT-Type: application/json' -d '{"name":"Vinit","Address":"475 West San Carlos Street","City":"San Jose", "State":"CA","Zip":"95110"}' http://localhost:8080/locations

func deleteHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {  
    // TODO: only write status for now

 fmt.Println("-----------------SERVER LOGS ARE ENABLED--------------------deleteHandler-------")

 id := p.ByName("id")
 fmt.Println("Id given -->"+id)
 mongoIndentifier := bson.ObjectIdHex(id)



 mgoSession, err := mgo.Dial("mongodb://vinitgaikwad0810:Thisisvinit0810@ds043714.mongolab.com:43714/cmpe273-assignment2-mongodb")
    // Check if connection error, is mongo running? 
 if err != nil {
    panic(err)
}
    // Stub user
   // u := Response{}

if err := mgoSession.DB("cmpe273-assignment2-mongodb").C("UserLocation").RemoveId(mongoIndentifier); err != nil {
    w.WriteHeader(404)
    fmt.Fprintf(w, "Deletion failed as Record might be absent \n")
    return
}

    // Write status
w.WriteHeader(200)
fmt.Fprintf(w, "The data is removed. Please check with firing a get request \n")



fmt.Println("-----------------SERVER LOGS ARE ENABLED--------------------deleteHandler-------")

}


func putHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params){
    fmt.Println("-----------------SERVER LOGS ARE ENABLED--------------------putHandler-------")
    id := p.ByName("id")
    fmt.Println("Id given -->"+id)
    userPut := new(Request)
    json.NewDecoder(r.Body).Decode(userPut)
   // fmt.Println("Address -->"+userPut.Address)   
    //fmt.Println()

    fmt.Println("Id given -->"+id)
    mongoIndentifier := bson.ObjectIdHex(id)

    userPut.Latitudes,userPut.Longitudes =returnlatlng(userPut.Address+userPut.City+userPut.State+userPut.Zip)

    mgoSession, err := mgo.Dial("mongodb://vinitgaikwad0810:Thisisvinit0810@ds043714.mongolab.com:43714/cmpe273-assignment2-mongodb")

    if err != nil {
        panic(err)
    }

    if err := mgoSession.DB("cmpe273-assignment2-mongodb").C("UserLocation").UpdateId(mongoIndentifier,bson.M{"$set":bson.M{"address": userPut.Address,"city":userPut.City, "state":userPut.State,"zip":userPut.Zip, "longitudes":userPut.Longitudes,"latitudes":userPut.Latitudes}}); err != nil {
        w.WriteHeader(404)
        return
    }

    fmt.Println("The record is updated.\n\n")

    userReceived := new(Request)

    if err := mgoSession.DB("cmpe273-assignment2-mongodb").C("UserLocation").FindId(mongoIndentifier).One(userReceived); err != nil {
        w.WriteHeader(404)
        return
    }
                   /* query := mgoSession.DB("go_rest_tutorial").C("users").Find(bson.M{"id":p.ByName("id")})
    fmt.Println(query)*/
                    // Marshal provided interface into JSON structure
    uj, _ := json.MarshalIndent(userReceived,"","\t")

                    // Write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "The updated record is as follows :-> %s", uj)


    fmt.Println("-----------------SERVER LOGS ARE ENABLED--------------------putHandler-------")
}

//curl -X "DELETE" http://www.url.com/page

func main() {  
    fmt.Println("-----------------CRUD LOCATION SERVICE LISTENING---------------------------")


    r := httprouter.New()

    r.POST("/locations", postHandler)
    r.GET("/locations/:id", getHandler)
    r.DELETE("/locations/:id", deleteHandler)
    r.PUT("/locations/:id",putHandler)
//r.DELETE("/locations/:id", delete)

    server := http.Server{
        Addr:        "0.0.0.0:8080",
        Handler: r,
    }
    server.ListenAndServe()

}


//curl -XPOST -H 'CONTENT-Type: application/json' -d '{"name":"Vinit","Address":"475 West San Carlos Street","City":"San Jose", "State":"CA","Zip":"95110"}' http://localhost:8080/locations

//curl -X "DELETE" http://localhost:8080/locations/562c2616c2be631f33c4b4c5

//curl -XPUT -H 'CONTENT-Type: application/json' -d '{"address":"1600 Amphitheatre Parkway","city":"Mountain View", "state":"CA","zip":"94043"}' http://localhost:8080/locations/562c6239c2be632ae01b4a7a

//curl -XPUT -H 'CONTENT-Type: application/json' -d '{"address":"Geeta Nagar 123","city":"Pune", "state":"MH","zip":"410038"}' http://localhost:8080/locations/562c67f3c2be632b940788f1
