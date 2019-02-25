# Dinner-dash

Automatic creation of food plans, based on saved recipes.

## Purpose

Are you tired of always asking yourself the question: "What should we have for dinner tonight?". Are you tired of having to go buy groceries multiple times a week, because you don't know what you are going to make for dinner throughout the week?

Dinner dash is the fix to that. The purpose of dinner dash is to provide a platform where you can save recipes, and be provided with meal plans for a week, which will specify what recipes to cook as well as what groceries to buy, in order to be able to cook the meals.

## Technology stack

### Backend technologies

* **Programming language:** [Go](http://golang.org).
* **Architecture:** Monolith exposing REST APIs to be used by the frontend.
* **Authentication:** JWT based approach (maybe auth0). [JWT introduction](https://jwt.io/introduction/)
* **Database:** PostgreSQL.

### Frontend technologies

* **Library:** [VueJS](https://vuejs.org/).
* **Typescript/JavaScript**: Currently plain JavaScript.

The frontend will be served from the Go backend. This eliminates the need for a separate web server just for serving static files. This also creates the possibility to create one single Docker image, which contains the entire application, which makes deployment much easier.

## Contributing

Contributions are more than welcome! Contact me if you find the project interesting, and we can probably figure out a way to include you =)

## Where can I use Dinner Dash?

The application is in its very early stages. Latest development builds will be put at [dinner-dash.tychsen.me](http://dinner-dash.tychsen.me).

## Name of the application

Dinner-dash is the name of a game. Thus, before this application can go live for real, it will need a new name.

At this point, there are no proposed names, but will appear in this section when received.
