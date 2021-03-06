# Api Benchmarking Challenge

Create an api endpoint using the framework and language of choice, the fastest wins!


## Specs

- Must be a GET endpoint available on `http://localhost:5000/`
- Your application must run with a single command `docker-compose up --build`
- Must use Elasticsearch as your data store as per example
- Must create an ES record under the `events` index containing `id` and `timestamp`
- Must return a count of records under the same index

## Benchmarking 

All submissions will be run against the benchmarking tool [plow](https://github.com/six-ddc/plow)
using the following command:
```shell
plow http://localhost:5000/ -c 100 -n 100000 -d 30s
```

## Example

```shell
cd fastapi-async-es7
docker-compose up --build
```

In your browser navigate to url `http://localhost:5000/` and you should see a response like the following:

```shell
{"event_count":2}
```

Refresh the page and the `event_count` should increment.


Switch to new terminal/shell

```shell
plow http://localhost:5000/ -c 100 -n 100000 -d 30s
```

Once its completed you can `Ctrl+c` out of the terminal where docker-compose is running, and clean out the ES volume with: 

```shell
docker-compose down -v
```

## Getting started

- Clone this repo
- Get the example above running
- Create a folder of your own named appropriately to describe what is being tested
- Reuse what you can from the example
- Commit and push your changes to the `main` branch
- Good Luck!
