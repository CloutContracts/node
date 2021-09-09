![CCSLOGO](https://raw.githubusercontent.com/CloutContracts/cloutcontracts.github.io/main/assets/images/c-128x128.png)
# CloutContracts Node
*This is still actively under construction/being developed and mantained*

[CloutContracts](https://cloutcontracts.net) is a next gen blockchain development platform targeting creators and social influencers looking to quickly deploy their own network, or even professional blockchain architects. We are interplorabile and multichain compatible. Not only this, but we are the native token behind a high speed rollup layer.

## Build

```bash
$ make build
```

## Install MySQL

You can either install MySQL locally or use docker.

```bash
$ docker run --name=mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=<your-password> -d mysql
```

You might also want to install a GUI to view the database changes

```bash
$ docker run --name myadmin -d --link mysql:db -p 8080:80 phpmyadmin/phpmyadmin
```

## Initialise configration

```bash
$ make init
```

In `config.toml` you need to change the key `db_url` with your correct password

```toml
db_url="mysql://root:<your-password>@/testing?charset=utf8&parseTime=True&loc=Local"
```

In `config.toml` you also have various params for entering ethereum RPC's and contract address, do check it out!

## Run migrations

```bash
$ make migration-up
```

## Reset DB

```bash
$ make migration-down
```

## Start the node
```bash
$ make start
```
