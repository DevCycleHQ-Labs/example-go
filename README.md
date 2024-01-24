# DevCycle Go Server SDK Example App

An example app built using the [DevCycle Go Server SDK](https://docs.devcycle.com/sdk/server-side-sdks/go/) with local bucketing

## Running the Example
### Setup

* Run `go mod download` in the project directory to install dependencies
* Create a `.env` file and set `DEVCYCLE_SERVER_SDK_KEY` to your Environment's SDK Key.\
You can find this under [Settings > Environments](https://app.devcycle.com/r/environments) on the DevCycle dashboard.
[Learn more about environments](https://docs.devcycle.com/essentials/environments).

### Development

`go run .`

Runs the app in the development mode.\
Requests may be sent to [http://localhost:8000](http://localhost:8000).

## Documentation
For more information about using the DevCycle Go Server SDK, see [the documentation](https://docs.devcycle.com/sdk/server-side-sdks/go/)
