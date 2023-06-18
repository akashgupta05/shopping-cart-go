# shopping-cart-go
Shopping Cart service provides implementation of below.
- Ability to create account with two roles (admin, user) and log in

Admin should be able to
- Add items
- Suspend user

User should be able to
- List available items
- Add items to a cart (if there are items in stock)
- Remove items from their cart 


### Steps to up and running

1. Install Go on your local machine. You will also need `Postgres` running locally.

2. `git clone git@github.com:akashgupta05/shopping-cart-go.git`

3. Create a dev environment file

   `cp env.sample development.env`

   Edit values in `development.env` to match your needs

4. Run Migrations

   - Install [golang-migrate](https://github.com/golang-migrate/migrate) tool

     `brew install golang-migrate`

   -  Create database

     `createdb shopping_cart_go_development`

   - Use helper bash script to run migrations

     `./scripts/migrations.sh`

4. Run Tests

   -  Create database

     `createdb shopping_cart_go_test`

   - Use helper bash script to run migrations

     `./scripts/migrations.sh`

   - Use helper bash script to run tests

     `./scripts/run_tests.sh`

5. Install `gin`. This is needed for live-reload during development

   `GOBIN=/usr/local/bin/ go install github.com/codegangsta/gin`

6. Use helper bash script to run server locally

   `./scripts/local_run.sh`