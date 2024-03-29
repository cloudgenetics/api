# CloudGenetics api

[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://raw.githubusercontent.com/cityscapelabs/cityscape/develop/LICENSE)
[![CircleCI](https://circleci.com/gh/cloudgenetics/api.svg?style=svg)](https://circleci.com/gh/cloudgenetics/api)
[![Project management](https://img.shields.io/badge/projects-view-ff69b4.svg)](https://github.com/orgs/cloudgenetics/projects/1)
[![CodeQL](https://github.com/cloudgenetics/api/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/cloudgenetics/api/actions/workflows/codeql-analysis.yml)

## Project setup
Set-up an [Auth0 API](https://auth0.com/docs/get-started/set-up-apis). Update the `.env` file:

```
API_NAME="Cloudgenetics"
AUTH0_DOMAIN="https://kks32.us.auth0.com/"
AUTH0_AUDIENCE="https://api.cloudgenetics.com"
AWS_REGION="us-east-1"
#APP_URL="https://dev-app.csgcompute.com"
APP_URL="*"
AWS_S3_BUCKET="csg-dev-data"
DB_NAME="dbname"
DB_USER="dev"
DB_HOST="dbname.randomcode.<region>.rds.amazonaws.com"
DB_PORT=5432
```

### Installing dependencies
```
go get -d
```

### Compile and run
```
GOOS=linux GOARCH=amd64 go build -o bin/application application.go
```

### Serve api
```
export GIN_MODE=release
export AWS_ACCESS_KEY_ID=<aws-id>
export AWS_SECRET_KEY=<aws-key>
./bin/application
```

### Deploy to AWS
```
eb deploy
```

