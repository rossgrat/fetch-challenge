# Fetch

Hey Fetch Team, thanks so much for considering me for the Backend Engineer position!


# TODO
```/receipts/process```
- Validate receipt fields before creating an ID
- Check for repeat receipts by hashing receipt object


Things Unimplemented:
* Extract per handler logging to middleware
* Rate limiting
* SSL encryption


Notes:
* At present, the spec does not call for actually saving the receipts, or preventing duplicate receipts from being submitted, if we wanted to do that, we would need to save some receipt signature, and compare verify that any inbound receipt does not have an exsting signature present in our database