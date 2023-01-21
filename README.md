[![License](https://img.shields.io/github/license/tarakaprabhuchinta/github-merged-branch-remover)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/tarakaprabhuchinta/github-merged-branch-remover)](https://goreportcard.com/report/github.com/tarakaprabhuchinta/github-merged-branch-remover) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=tarakaprabhuchinta_github-merged-branch-remover&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=tarakaprabhuchinta_github-merged-branch-remover)

# github-merged-branch-remover
Go code to check and remove merged branches in Gihub repositories in which merged branches has their commits more than 10 days ago.

### Github Token
A Github token is needed to run this code. The token needs to have read and write permissions to the repo.

```

 export GITHUB_TOKEN=<your-token>

 ```

## Run the code

```

go run main.go json-decode.go <owner/organization-name> <repo-name>

```

## Get this project as go module

```

go get github.com/tarakaprabhuchinta/github-merged-branch-remover@v1.0.3

```

## Build the Docker file 

```

docker build . -t github-merged-branch-remover:1.0.3

```

To view the output when building Dockerfile

```

docker build . --progress plain --no-cache

```

Export Github token as environment variable

```

export GITHUB_TOKEN=<github-token>

```

## Run the docker image

```
docker run -e GITHUB_TOKEN=$GITHUB_TOKEN github-merged-branch-remover:1.0.3 <owner-name> <repo-name>

````