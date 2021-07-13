# Chef wasm

Chef wasm uses [https://github.com/autamus/chef](https://github.com/autamus/chef) in order to generate a web
interface to create a Dockerfile. Chef Autamus at your service!

![docs/img/chef.png](docs/img/chef.png)

You'll first need to generate a chef.wasm file under [docs](docs).

```bash
$ make wasm
```

You can then deploy this on GitHub pages, or run a local webserver.

```bash
$ cd docs
$ python -m http.server 9999
```
