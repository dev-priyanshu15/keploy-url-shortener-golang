# Echo SQL Sample Application 🚀

## Introduction
This is a sample URL shortener app built using **Echo** and **PostgreSQL** to test Keploy integration capabilities. Buckle up, it's gonna be a fun ride! 🎢

---

## 🛠️ Platform-Specific Requirements for Keploy

| Operating System | Without Docker    | Docker Installation | Prerequisites                                                                          |
|----------------|------------------|--------------------|----------------------------------------------------------------------------------------|
| MacOS          | Not Supported     | Supported           | Docker Desktop version 4.25.2 or above                                                   |
| Windows        | Supported         | Supported           | Use WSL (`wsl --install`), Windows 10 (version 2004+) or Windows 11                      |
| Linux          | Supported         | Supported           | Linux kernel 5.15 or higher                                                              |

On **MacOS** and **Windows**, additional tools are required for Keploy due to the lack of native eBPF support.

---

## Keploy Installation

### Quick Installation Using CLI
To set up the Keploy alias, run this command:

```bash
curl --silent -O -L https://keploy.io/install.sh && source install.sh
```

### Check Installation
After installation, you should see the Keploy CLI with available commands and flags. Use the following command to check:

```bash
keploy --help
```

### Installation using Docker
Alternatively, you can install using Docker:

```bash
docker volume create --driver local --opt type=debugfs --opt device=debugfs debugfs
```

---

## Clone the Sample Application 🧪

```bash
git clone https://github.com/keploy/samples-go.git && cd samples-go/echo-sql
go mod download
```

### Using Docker Compose 🐳
Start the Postgres instance:

```bash
docker compose up
```

Build the application binary:

```bash
docker build -t echo-app:1.0 .
```

Start recording API calls:

```bash
keploy record -c "docker run -p 8082:8082 --name echoSqlApp --network keploy-network echo-app:1.0"
```

---

## Testing the Application 🧪

### Generate Testcases
Make API calls to generate testcases:

```bash
curl --request POST \
  --url http://localhost:8082/url \
  --header 'content-type: application/json' \
  --data '{ "url": "https://github.com" }'
```

### Run Captured Testcases

```bash
keploy test -c "docker run -p 8082:8082 --name echoSqlApp --network keploy-network echo-app:1.0" --delay 10
```

---

## Running App Locally on Linux/WSL 🐧
To run the app locally with Docker for Postgres:

```bash
docker-compose up -d
go build -cover
```

Start recording:

```bash
sudo -E PATH=$PATH keploy record -c "./echo-psql-url-shortener"
```

---



## Screenshots
![App Screenshot](https://i.postimg.cc/yN0GzB3f/Screenshot-from-2025-03-31-19-15-32.png)

![App Screenshot](https://i.postimg.cc/Y0t0xpYT/Screenshot-from-2025-03-31-19-16-45.png)

![App Screenshot](https://i.postimg.cc/mkhDnmpS/Screenshot-from-2025-03-31-19-18-55.png)
## Postman output
![App Screenshot](https://i.postimg.cc/C1V4Cq17/Screenshot-from-2025-03-31-19-13-08.png)




[Documentation](https://keploy.io/docs/quickstart/samples-echo/)


## Authors

- [dev-priyanshu15](https://www.github.com/dev-priyanshu15)

