# CloudGenetics api

[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://raw.githubusercontent.com/cityscapelabs/cityscape/develop/LICENSE)
[![CircleCI](https://circleci.com/gh/cloudgenetics/api.svg?style=svg)](https://circleci.com/gh/cloudgenetics/api)
[![Project management](https://img.shields.io/badge/projects-view-ff69b4.svg)](https://github.com/cloudgenetics/api/projects/)

## Project setup
Set-up an [Auth0 API](https://auth0.com/docs/get-started/set-up-apis). Update the `.env` file with `AUTH0_DOMAIN` and `AUTH0_AUDIENCE` variables.

### Installing dependencies
```
go get -d
```

### Compile and run
```
go build -o api
```

### Serve api
```
export GIN_MODE=release
./api
```