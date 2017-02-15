# pinky-02

This is the pinky pig #02

## Description
Small API that returns flat information about a product, based on its ID. The backing of this API is a SQLite DB, which is present in this repo.

## Notes
The project uses `glide` to manage dependencies and make for scripting

### Installation
`make install`

### Run 
`make run`

## Improvements
* Develop a caching technique to avoid doing the multiple queries to the DB
* Actually implement the locale feature
* Unit testing