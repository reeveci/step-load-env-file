# Reeve CI / CD - Pipeline Step: Load Env File

This is a [Reeve](https://github.com/reeveci/reeve) step for loading runtime variables from .env files.

## Configuration

See the environment variables mentioned in [Dockerfile](Dockerfile).

Params starting with `ENV_` specify which variables to load and how they should be named.
E.g. the following loads the variable `REEVE_VERSION` from the file `.env` and stores the value in the runtime variable `IMAGE_VERSION`.

```yaml
FILES: .env
ENV_REEVE_VERSION: IMAGE_VERSION
```
