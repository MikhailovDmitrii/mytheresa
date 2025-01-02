# mytheresa Promotions Test

## Description

We want you to implement a REST API endpoint that given a list of products, applies some discounts to them and can be filtered.

The full description of the task can be found [here](Task.md)

## How to run it

There is Makefile in the root directory with some handy receipts.
* "build", "run", "test" - do as them state, build a docker image, run it, or run unit rests.
* "insert" - it inserts products to the database. Those products are from the task description, just to have something in db. The input is the [products.json](products.json) file. You can swap it to a bigger file to have more variety in DB.

## Decisions

1. I tried to slice the project using DDD approach, means there are layers what communicate to each other only by entities. 
    For example, `/internal/api` package doesn't know anything about infrastructure, so we can switch our DB engine without huge refactoring on the API layer.
2. I skipped the part with configuration to reduce the scope of the task. However, the most valuable parameters are set in constants at [cmd/api/main.go#14](cmd/api/main.go#14). You can flex it if you need. 
3. I chose Sqlite as database to be able to send it to you and skip the part with migrations etc, just to reduce the scope of task
4. There is CompositePromotion which enables us to apply several promotions to a product regardless of the concrete promotion. Naming could've been better, like `AtLeastOnePromotion`. Also, we can add more complex promotions to stack them up. Looks flexible enough, but could be improved too.
5. There are basic CRUD endpoints to give you an easier way to check the DB during you testing. The whole list of endpoints you can find here [cmd/api/main.go#47](cmd/api/main.go#47)
6. I skipped tests for converter and those basic endpoint in favor of more complex parts like the `handler.go`, promotions themselves, and their ways of application.
7. There is `init_db.sql` file, it isn't used anywhere in code or scripts, it just useful sometimes.