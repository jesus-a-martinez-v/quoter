# Quoter

Simple API to manage your favorite quotes by the influential people you admire the most! 

![Quote](https://i.pinimg.com/originals/b7/91/d8/b791d82bb7fcc62ff09b753259264b54.jpg)

# Dependencies

In order for Quoter to work in your local environment, you'll need:

- [Go](https://golang.org/).
- [Gingonic](https://gin-gonic.github.io/gin/).
- [PQ](https://github.com/lib/pq).
- [Docker Compose](https://docs.docker.com/compose/install/).

# Data

Quoter comes with a set of quotes that'll be loaded into the system the first time you run the application. If you want to review or modify this data set,
it is located at [/resources/quotes_all.csv](https://github.com/jesus-a-martinez-v/quoter/blob/master/resources/quotes_all.csv)

# Run

### Step 1: Run database

`docker-compose up`

NOTE: If you want to stop the container that runs the database and delete the attached volume, run: `docker-compose down -v`

### Step 2: Run the service

From the source of the project, execute: `go run src/api/main.go`

After loading the data, the server should start listening on port 8000.

# API

### Create a quote

```$xslt
POST localhost:8000/quotes

Body:

{
	"quote": "Trust the Process",
	"author": "John Sonmez"
}

Response: 

{
    "id": 42,
    "quote": "Trust the Process",
    "author": "John Sonmez"
}
```

# Get a quote by its ID.
   
   ```$xslt
   GET localhost:8000/quotes/42
   
   Response: 
   
   {
       "id": 42,
       "quote": "Trust the Process",
       "author": "John Sonmez"
   }
   ```
   
# Delete a quote by its ID.

```$xslt
DELETE localhost:8000/quotes/42

Response: 

{
    "id": 42,
    "quote": "Trust the Process",
    "author": "John Sonmez"
}
```

# Get quotes

```
GET localhost:8000/quotes

Optional query parameters:

- author: The name of the author of the quote. Example: Mark Twain.
- genre: The genre of the quote. Example: best.
- random: Flag. If true, it will retrieve a random quote based on the filters supplied above. If none is supplied, then it will consider the whole database.

Example: http://localhost:8000/quotes?random=true&author=Henry+Ford&genre=happiness

Response: 
[
    {"id": 32041,
     "author": "Henry Ford","genre":"happiness", 
     "quote":"There is joy in work. There is no happiness except in the realization that we have accomplished something."
     }
]
```
