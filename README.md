# Movie-Application-Microservice-Project

## Application for movie lovers

### Features

1. Get the movie metadata (such as title, year, description, director) and aggregated movie rating

2. Rate a movie

### Movie metadata

- ID
- Title
- Year
- Description
- Director
- List of actors

Such information about movies doesn’t generally change unless somebody wants to update the description, but for simplicity, we may assume that we are dealing with a static dataset. We would retrieve the records based on their IDs, so we could use any key-value or document database to store and access the metadata.

### Rating

rating operations:

1. Store a movie rating

2. Get the aggregated movie rating

At some point, we may want to extend the rating functionality to other types of movie-related records. A user may be able to do the following:

1. Rate the actor’s performance in some movies

2. Rate the movie soundtrack

3. Rate the movie’s costume design

Let’s define the API for such a rating component:

1. Store the rating record, including the following:

- ID of the user who gave the rating
- Type of record
- ID of the record
- Rating value

2. Get the aggregated rating for a record by its ID and type


the data models for both components are also quite different. The movie metadata component stores static data, which is going to be retrieved by ID, while the rating component stores dynamic data, which requires aggregation

Both components seem to be relatively independent of each other. This is a perfect example of a situation where we may benefit from splitting the application into separate services:
- Logic is loosely coupled
- Data models are different
- Data is generally independent

Let’s list the services we would split the application into:

1. **Movie metadata service:** Store and retrieve the movie metadata records by movie IDs.
2. **Rating service:** Store ratings for different types of records and retrieve aggregated ratings for records.
3. **Movie service:** Provide complete information to the callers about a movie or a set of movies, including the movie metadata and its rating.

Metadata service <-------movie service-----> Rating service

application components:
- **controller:** Business logic
- **gateway:** Logic for interacting with other services
- **handler:** API handlers• repository: Database logic

### Movie Metadata Service

Logic of metadata service

- **API:** Get metadata for a movie
- **Database:** Movie metadata database
- **Interacts with services:** None
- **Data model type:** Movie metadata

This logic would translate into the following packages:
- **cmd:** Contains the main function for starting the service
- **controller:** Our service logic (read the movie metadata)
- **handler:** API handler for a service
- **repository:** Logic for accessing the movie metadata database


- metadata/cmd
- metadata/internal/controller
- metadata/internal/handler
- metadata/internal/repository
- metadata/pkg

### Rating Service

the logic of the rating service:

- **API:** Get the aggregated rating for a record and write a rating.
- **Database:** Rating database.
- **Interacts with services:** None.
- **Data model type:** Rating.
This logic would translate into the following packages:

- **cmd:** Contains the main function for starting the service
- **controller:** Our service logic (read and write ratings)
- **handler:** API handler for a service
- **repository:** Logic for accessing the movie metadata database

- rating/cmd
- rating/internal/controller
- rating/internal/handler
- rating/internal/repository
- rating/pkg

### Movie Service

logic of the movie service:

- **API:** Get the details for a movie, including the aggregated movie rating and movie metadata.
- **Database:** None.
- **Interacts with services:** Movie metadata and rating.
- **Data model type:** Movie details

This logic would translate into the following packages:

- cmd: Contains the main function for starting the service
- controller: Our service logic (read rating and metadata)
- gateway: Logic for calling the other services
- handler: API handler for a service

The directory structure is as follows:

- movie/cmd
- movie/internal/controller
- movie/internal/gateway
- movie/internal/handler
- movie/pkg
