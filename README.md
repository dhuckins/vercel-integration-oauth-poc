# vercel-integration-oauth-test

# How to run

1. Create an integration on vercel

2. put the keys in a `.env` file OR export as environment variables
```dotenv
VERCEL_CLIENT_ID="<provided in ui when registering a vercel integration>"
VERCEL_CLIENT_SECRET="<provided in ui when registering a vercel integration>"

```
3. run the app via `go run .`
4. go to the marketplace url for the app and click "Add Integration"


## Notes
- recommend to use the webhook `integration-configuration.removed|updated` to see if the integration has been uninstalled or changed