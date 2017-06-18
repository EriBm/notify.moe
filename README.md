# Anime Notifier

## Installation

### Prerequisites
* Install [Go](https://golang.org/dl/) (1.9 or higher)
* Install [Aerospike](http://www.aerospike.com/download) (3.14.0 or higher)

### Database
* Remove all namespaces in `/etc/aerospike/aerospike.conf`
* Add a namespace called `arn`:
```
namespace arn {
    storage-engine device {
        file /home/YOUR_NAME/YOUR_PATH/notify.moe/db/arn-dev.dat
        filesize 60M
        data-in-memory true

        # Maximum object size. 128K is ideal for SSDs
        write-block-size 128K

        # Write block size x Post write queue = Cache size (for write block buffers)
        post-write-queue 128
    }
}
```
* Start the database using `sudo service aerospike start`
* Confirm that the status is "green": `sudo service aerospike status`

### Hosts
* Add `127.0.0.1 arn-db` to `/etc/hosts`
* Add `127.0.0.1 beta.notify.moe` to `/etc/hosts`

### Download repository
* `go get github.com/animenotifier/notify.moe`

### HTTPS
* Create directory `notify.moe/security`
* Create the certificate `notify.moe/security/fullchain.pem` (domain: `beta.notify.moe`)
* Create the private key `notify.moe/security/privkey.pem`
* Add `notify.moe/security/fullchain.pem`

### API keys
* Get a Google OAuth 2.0 client key & secret from [console.developers.google.com](https://console.developers.google.com)
* Create the file `notify.moe/security/api-keys.json`:
```json
{
	"google": {
		"id": "YOUR_KEY",
		"secret": "YOUR_SECRET"
	}
}
```

### Build all
* Run `make all`

### Fetch data
* Run `jobs/sync-anime/sync-anime` from this repository to fetch anime data

### Install pack & run
* `go install github.com/aerogo/pack`
* `go install github.com/aerogo/run`

### Run
* Start the web server in notify.moe directory: `run`