# MyGram


Folder Directory
- DTO = Data Structure of response 
- Entity = Model
- Repository = Interface + DB Level Query
- Service = Validate Request from user
- Handler = Get user content type request & response the result
- Pkg = Helpers & Error status

Example Login:

{
  "age": 10,
  "email": "a@a.com",
  "password": "string",
  "username": "string"
}

Example bearer token:

{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFAYS5jb20iLCJleHAiOjE2NjYyODU5MzYsImlkIjoyfQ.ahlQ3Icb2h72Aam9jWgZFM-TultwjKzymM7DGKHVVKU"
}