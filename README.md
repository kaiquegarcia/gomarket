# GO MARKET

A simple project I'm creating just to spent my time with something that was on my head for a long time.

I won't use it in any project. If you found anything useful, feel free to use, but please remember to mention my name on it. It's just a gentle way to keep the sources connected.

I don't expect to earn anything from this but, just in case, if you want to encourage myself to create more open source projects like this, please send me a coffee:

- [PICPAY](https://picpay.me/kaiquegarcia.dev/10.0)
- [MERCADO PAGO](https://mpago.la/2rJb27G)

# How to use

## CLI

Just run `go run .` to start the CLI procedures. You'll see a message with more details.

## Web

Just run `go run . serve` to initialize the web application. Then, download [Postman](https://www.postman.com/) and import the [collection](./gomarket.postman_collection.json) to start sending requests.

# TO-DO List

### Product
- [P-001] Add support to update a product's materials from CLI;
- [P-002] Add user settings (including option to block access with password - the block should work for both CLI/WEB interfaces);
- [P-003] Add product's labor cost;
- [P-004] Add profit calculator.

### Technical
- [T-001] Add `util.AskPassword` function to properly request passwords on CLI (replacing the input with `*`) - required for `P-002`;
- [T-002] Add a context-oriented logging system - required for `T-003`;
- [T-003] Add RequestID middleware to automatically inject a requestID on the logging system through `*gin.Context`;
- [T-004] Add CORS middleware to validate the origin - required for `T-005`;
- [T-005] Create a web client for the web API;
- [T-006] Add `godoc comments` for HTTP handlers and install Swagger to provide compilable docs to run with the web API;
- [T-007] Add `dotenv` to properly manage environments (e.g. web API's port, etc);
- [T-008] Add more options of storage (quit from files, use MongoDB/Firebase/DynamoDB), configured by environment.