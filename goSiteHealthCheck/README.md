# Go Site Health Checker
Command line tool used to check if a given web address is running or currently down.

All data saved locally using structs and slices.

## Featues
- Two API endpoints
- Five Routes
- Five API Methods

## API Notes
|Route |Functions |Endpoints | Methods
|--|--|--|--|
|`GET ALL`|getAllFilms|`/films/`|GET|
|`GET BY ID`|getFilm|`/films/id`|GET|
|`CREATE`|createFilm|`/films`|POST|
|`UPDATE`|updateFilm|`/films/id`|PUT|
|`DELETE`|deleteFilm|`/films/id`|DELETE|

## Installation and Usage
```bash
# Clone the repo
$  git clone git@github.com:lokeam/go-micro-projects.git

# Navigate to the goEmailVerify directory
$  cd goSiteHealthCheck

# Install all app dependencies
$  go mody tidy

# Boot up the app, including a web domain (exclude https://) to check:
# Example:
$  go run . --domain go.dev/
```
Example responses:
```bash

# go.dev
[UP] go.dev is reachable,
 From: 192.168.50.209:56338
 To: 216.239.36.21:80

 # possiblybrokensite.com
 [DOWN] possiblybrokensite.com is unreachable,
 Error: dial tcp: lookup possiblybrokensite.com: no such host
```
