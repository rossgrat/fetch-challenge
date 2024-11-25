# Fetch Receipt Processor Challenge - Ross Grattafiori

Hey Fetch Team, thanks so much for considering me for the Backend Engineer position! I'm really excited about this opportunity, and am looking forward to your feedback on my asssessment. Hope you're having a great holiday season!

The web service that fulfills the challenge spec can be invoked in `cmd/run-receipt-processor` with `go run `, or any other flavor of building and executing a go binary.

The handlers for the endpoints defined in the API document live in `/src/receipt-process/api`, the 'database' layer is in `/src/receipt-processor/db`. `/src/mw` is some basic middleware, and `/src/logger` is a basic logger.

The `tests` directory is some basic client testing and example receipts, and the `docs` directory is the original project challenge and API definition.

# Things deferred:
* Rate limiting
* SSL encryption
* Better error configuration (error priority, (info, warning, critical), fields)
* External configuration files
