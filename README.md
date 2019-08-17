# CRUD API using revel

I used revel a high-productivity web framework for creating crud API using mysql as the database layer.

### Start the web server:

revel run crud

### Go to http://localhost:9000/ and you'll see:

    "It works"

## Code Layout

The directory structure of a generated Revel application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        models/       App models go here
        views/        Templates directory
        helpers/      Helper functions can be written here

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites

## Available routes

    POST    http://localhost:9000/user      (Add User)
    GET     http://localhost:9000/user/:id  (Get User Detail)
    PUT     http://localhost:9000/user/:id  (Update User)
    DELETE  http://localhost:9000/user/:id  (Delete User)
    GET     http://localhost:9000/users     (Get List of User)

## Database config

    To change database credentials please go inside the config folder and change with your credential
    Go inside config folder and search for the below line
    db.info = root:altu123@/gomysql?charset=utf8&parseTime=True
    Where root = user, altu123 = password and gomysql = database

## Future Goals

Updating this project to use multiple database and implement usage of goroutines & channels

## Help

-   The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
-   The [Revel guides](http://revel.github.io/manual/index.html).
-   The [Revel sample apps](http://revel.github.io/examples/index.html).
-   The [API documentation](https://godoc.org/github.com/revel/revel).
