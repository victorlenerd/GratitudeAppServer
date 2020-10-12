# Gratitude App Server

Backend Server For Gratitude App. 
API for sharing notes and connecting with friends using the app.


## Features & Integrations

* Go Firebase Admin SDK Integration
    * Searching the user database with email
    * Sending push notifications using FCM
* Google Data Store
* Curate feed based on sharded notes and friends

## Usage

Set the following environment variables when running locally

```
* PORT
* DATASTORE_EMULATOR_HOST
* GOOGLE_APPLICATION_CREDENTIALS
```

## API Documentation

Endpoint          | Method | Description  
------------------| ------ | ----------- 
/                 | GET    | Welcome      
/search           | GET    | Search for a user on the app             
/tokens/{userID}  | PUT    | Create or update token for a user
/feeds/{userID}   | GET    | Get all notes shared on feed by friends of user
/friends          | PUT    | Create a friend request or update a friend request
/friends/{userID} | GET    | Get all friends for a user
/friends/{uuid}   | DELETE | Delete friend request
/notes            | GET    | Get all notes for a user
/notes            | PUT    | Upload a note or share a note to a feed
/notes/{uuid}     | DELETE | Delete a note

## License

This project is licensed under either of
 * Apache License, Version 2.0, ([LICENSE-APACHE](LICENSE-APACHE) or
   http://www.apache.org/licenses/LICENSE-2.0)
 * MIT license ([LICENSE-MIT](LICENSE-MIT) or
   http://opensource.org/licenses/MIT)

at your option.

### Contribution

Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in this project by you, as defined in the Apache-2.0 license,
shall be dual licensed as above, without any additional terms or conditions.