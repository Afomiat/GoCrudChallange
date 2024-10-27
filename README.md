# Go CRUD Challenge 🚀

This project implements a RESTful API in Go for managing a collection of persons. It allows you to create, read, update, and delete person records stored in a MongoDB database.

## Features 🌟
- **CRUD Operations**: Create, Read, Update, Delete person records.
- **Properties**:
  - `id`: Unique identifier (string, UUID)
  - `name`: Person's name (string, required)
  - `age`: Person's age (number, required)
  - `hobbies`: Array of hobbies (strings, required)

## Technologies 🛠️
- Go (Golang)
- Gin Gonic (web framework)
- MongoDB (database)
- Gorilla CORS (for cross-origin resource sharing)

## API Endpoints 🌐

- **GET /person**: Retrieve all persons.
- **GET /person/${personId}**: Retrieve a specific person by ID.
- **POST /person**: Create a new person.
- **PUT /person/${personId}**: Update an existing person.
- **DELETE /person/${personId}**: Remove a person

## CORS Support 🌍
The API is accessible by frontend apps on different domains, with CORS enabled for specified origins.