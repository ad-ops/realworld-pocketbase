# ![RealWorld Example App](logo.png)

> ### Go with PocketBase and htmx codebase containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [RealWorld](https://github.com/gothinkster/realworld) spec and API.


### [Demo](https://demo.realworld.io/)&nbsp;&nbsp;&nbsp;&nbsp;[RealWorld](https://github.com/gothinkster/realworld)


This codebase was created to demonstrate a fully fledged fullstack application built with [PocketBase](https://pocketbase.io) and [htmx](https://htmx.org) including CRUD operations, authentication, routing, pagination, and more.

For more information on how to this works with other frontends/backends, head over to the [RealWorld](https://github.com/gothinkster/realworld) repo.


# How it works
The motivation behind this project was to make a very simple, yet powerful way of prototyping solutions. Go is a very simple language, SQLite is a simple DB and htmx is a simple way of making reactive frontends. Pocketbase is here to make it easier to use certain features available and to change the data model in an easier way.

The hope is to build a simple prototype that is self-contained, trivial to host anywhere and powerful enough for most use cases. It is also possible to write everything in Javascript and to use one of the popular SPA-frameworks, but that is for another time...

# Getting started
```sh
go mod tidy
go run main.go serve
```

```sh
docker build -t realworld-pocketbase .
docker run --rm -it -p 8080:8080 realworld-pocketbase
```

## Login to PocketBase
You can use user: `admin@realworld.com` and password: `adminadmin` in the provided db.
