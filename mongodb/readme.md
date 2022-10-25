# Config MongoDB 5.0

```shell
sudo -i
passwd root
sudo sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && sudo service ssh restart
cp /home/ubuntu/.ssh/authorized_keys /root/.ssh/authorized_keys && cat .ssh/authorized_keys && apt update

sudo apt-get install gnupg
wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
sudo apt-get update
sudo apt-get install -y mongodb-org
echo "mongodb-org hold" | sudo dpkg --set-selections
echo "mongodb-org-database hold" | sudo dpkg --set-selections
echo "mongodb-org-server hold" | sudo dpkg --set-selections
echo "mongodb-org-shell hold" | sudo dpkg --set-selections
echo "mongodb-org-mongos hold" | sudo dpkg --set-selections
echo "mongodb-org-tools hold" | sudo dpkg --set-selections
mkdir -p /var/lib/mongodb /var/log/mongodb
ps --no-headers -o comm 1

sudo systemctl start mongod
sudo systemctl daemon-reload
sudo systemctl status mongod
sudo systemctl enable mongod
sudo systemctl stop mongod
sudo systemctl restart mongod

mongo
#############################
use admin
db.createUser({
  user: "root",
  pwd: "test",
  roles: [
            { role: "userAdminAnyDatabase", db: "admin" },
            { role: "readWriteAnyDatabase", db: "admin" },
            { role: "dbAdminAnyDatabase",   db: "admin" }
         ]
})
use test
db.test.insertOne({"_id":"00001","hello":"world"})
db.test.find()
exit
#############################

sudo nano /etc/mongod.conf
#############################
...

# network interfaces
net:
  port: 27017
  bindIp: 0.0.0.0   #bindIp: 127.0.0.1

# security:
security:
  authorization: 'enabled'

...
#############################
sudo systemctl restart mongod

#############################
mongo <yourIpV4>/admin --username root --password
use test
db.test.find()
exit

#############################
mongo <yourIpV4>
use admin
db.auth("root","test")

use test
db.test.find()
exit
```

# Docker MongoDB

### Raspberry pi CPU set Docker image: `mongo:4.4`

- mongo:5.0
- mongo:4.4

[mongodb.yml](https://git.giaminhmedia.vn/websass/dashboard-api/-/blob/main/mongodb/mongodb.yml)

```shell
sudo apt update
sudo apt install docker.io docker-compose -y
sudo dockerd
sudo nano mongodb.yml
#############################
version: '3'
services:
  mongodb:
    restart: always
    image: mongo:5.0
    #image: mongo:4.4 #v4.4 for raspberry pi cpu
    environment:
      - "MONGO_INITDB_ROOT_USERNAME=root"
      - "MONGO_INITDB_ROOT_PASSWORD=test"
    ports:
      - "27017:27017"
    volumes:
      - "./configdb:/data/configdb"
      - "./db:/data/db"
#############################
docker-compose -f mongodb.yml up -d
```
