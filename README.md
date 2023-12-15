# Reeve CI / CD - Pipeline Step: Load Env File

This is a [Reeve](https://github.com/reeveci/reeve) step for loading runtime variables from a .env file.

## Configuration

See the environment variables mentioned in [Dockerfile](Dockerfile).

Params starting with `ENV_` specify which variables to load and how they should be named.
E.g. the following loads the variable `REEVE_VERSION` from the file `.env` and stores the value in the runtime variable `IMAGE_VERSION`.

```yaml
FILE: .env
ENV_REEVE_VERSION: IMAGE_VERSION
```
