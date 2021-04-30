# Go lib

A mini Go library for personal use.

## Components

### csv

Parses CSV files.

### parameterize

Converts a string to snake-case.

### s3

A basic wrapper around AWS S3 SDK.

### timestamp

Set of functions to generate timestamp strings in a standart format.

## Howto

### Prerequisites

Copy the `.env.example` file and edit it as necessary.

### Run tests

```
make test
```

#### Running S3 end-to-end testing

By default these tests are disabled because they will require access to an active AWS account.

To run them:

- Add the necessary environment variables in `.env` (see `.env.example`)
- Set `RUN_S3_E2E_TESTING` environment variable to `true`
