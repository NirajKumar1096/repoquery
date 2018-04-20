# repoquery
This is a Go App that allows users to filter Github data using the github.com API (api.github.com). 
 
•	It has a RESTful API in GO that filters public repos using the Github API
•	Allow users to favorite or like repos from the filtered repo list
•	Provide ability to view & edit the user’s favorite repos

## USAGE :
  go run main.go

  will run inbuilt http server on port 8088.  
  All data for Github repos as well as Users is stored in memory database as Json structures.  
  Users are stored as Json structure. 

  This can be tested using Postman, postman collection is attached as GoWebSvc.postman_collection.json.


### APIs

Environment variable 'baseUrl' should be set with "hostname:8088" e.g. localhost:8088 to use Postman REST requests.  

#### Query Repo  (GET)
  
http://{{baseurl}}/github/{query}

example: http://{{baseurl}}/github/terraform will query GitHubAPI and return repos with terraform. These will be displayed in grid on console and will be stored in memory database of GoWebApp.

#### Get Repo  (GET)
  
http://{{baseurl}}/github/repo

example: http://{{baseurl}}/github/repo will return repos stored in memory of App, it does not query GitHub API again. This is used as a mechanism to provide persistance to collection of repo that users can be set favs and likes against.

#### Get Users (GET)
  
http://{{baseurl}}/user

example: http://{{baseurl}}/user will return users from memory.

#### Get User (GET)
  
http://{{baseurl}}/user/{userid}

example: http://{{baseurl}}/user/2 will return user with id user from memory as Json.

#### Create User (POST)
  
http://{{baseurl}}/user/{userid}

example: http://{{baseurl}}/user/2 with following Json in request body will create User with Id 2 and  return updated user collection.
{"userid":"2","name":"Nic2","favorites":[2,3]}

#### Delete User (DELETE)
  
http://{{baseurl}}/user/{userid}

example: http://{{baseurl}}/user/2 will delete user with {userid} from memory as Json and return updated user collection.

#### Add User Favorite Repo (POST)
  
http://{{baseurl}}/user/{userid}/fav/{favid}

example: http://{{baseurl}}/user/1/fav/93446075 will add fav with github repo id of 93446075 to user with id of 1 and return updated user collection.

#### Remove User Favorite Repo (DELETE)
  
http://{{baseurl}}/user/{userid}/fav/{favid}

example: http://{{baseurl}}/user/1/fav/93446075 will remove fav with github repo id of 93446075, if it exists for user with id of 1 and return updated user collection.

#### User Json structure in memory

[
  {
    "userid": "1",
    "name": "Nic",
    "favorites": [
      126195231,
      93446075
    ]
  },
  {
    "userid": "2",
    "name": "Mary",
    "favorites": [
      126195231,
      120045621
    ],
    "likes": [
      93446075,
      120045621
    ]
  }
]