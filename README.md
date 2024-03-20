# Dcard Backend Assignment
Dcard Bcakend assignment based on https://drive.google.com/file/d/1dnDiBDen7FrzOAJdKZMDJg479IC77_zT/
## Introduction
* A backend that based on light framework of Go:Gin-gonic, along with Postresql as database.
* Used Swag for the convience of testing API.
* Used github action to maintain code consistency and validity(with go-linter and go-test).
* Containerize with docker for simplicity of deployment.
* Python for generating test data.
* Load testing with Jmeter.

## Test live now!

> https://ccns.owenowenisme.com/swagger/index.html
> 
> (Built on my server, performance might not as good as local testing due to network.)

## Built With


* Golang <img src="https://skillicons.dev/icons?i=go" alt="go" width="40" height="40"/> </a>
* Gin-gonic <img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.svg" alt="gin" width="40" height="40"/> </a>
* Postresql <img src="https://skillicons.dev/icons?i=postgres" alt="gin" width="40" height="40"/> </a>
* Swag  <img src="https://raw.githubusercontent.com/swaggo/swag/master/assets/swaggo.png" alt="swag" width="40" height="40"/> </a>
* Docker <img src="https://skillicons.dev/icons?i=docker" alt="docker" width="40" height="40"/> </a>
* Python <img src="https://skillicons.dev/icons?i=py" alt="python" width="40" height="40"/> </a>
* Github Action <img src="https://skillicons.dev/icons?i=githubactions" alt="python" width="40" height="40"/> </a>
* Jmeter <img src="https://jmeter.apache.org/images/jmeter.png" alt="python" height="40"/> </a>



## Requirement
Have docker installed on your machine.

**macOS:**

1. Download Docker Desktop for Mac from the [Docker Hub](https://hub.docker.com/editions/community/docker-ce-desktop-mac/).
2. Double-click the downloaded `.dmg` file and drag the Docker app to your Applications folder.
3. Open Docker Desktop from your Applications folder. You'll see a whale icon in the top status bar indicating that Docker is running.

**Windows:**

1. Download Docker Desktop for Windows from the [Docker Hub](https://hub.docker.com/editions/community/docker-ce-desktop-windows/).
2. Run the installer and follow the instructions.
3. Docker Desktop will start automatically once installation is complete. You'll see a whale icon in the notification area indicating that Docker is running.

**Linux (Ubuntu):**
```
 curl -fsSL https://get.docker.com -o get-docker.sh
 sudo sh get-docker.sh
```
## Usage
1. Clone this repo and enter in terminal.
2. Type in your terminal: ```docker compose up -d ```
3. Go to http://localhost:8080/swagger/index.html for API testing
4. Use ```go test -v``` for automatic testing.
## API Reference

### Retieve Ads

```
  GET /api/v1/ad
```

| Parameter  | Type     | Description                    |
| :--------  | :------- | :-------------------------     |
| `offset` | `int`    | Offset for pagination          |
| `limit` | `int`    | Limit for pagination default 5 |
| `age` | `int`    | Age to Query                   |
| `gender` | `string` | Gender                         |
| `country` | `string` | Country                        |
| `platform` | `string` | Platform                       |

### Create Ad

```
  POST /api/v1/ad
```

* Request Body
``` json
{
  "conditions": {
    "ageEnd": 0,
    "ageStart": 0,
    "country": [
      "string"
    ],
    "gender": "string",
    "platform": [
      "string"
    ]
  },
  "endAt": "string",
  "startAt": "string",
  "title": "string"
}
```
> [!NOTE]
> All field in conditions are optional
> startAt and End At should be in "yyyy-mm-ddTHH:MM:SSZ" format




## Jmeter Load Testing Result
![](https://i.imgur.com/KTG3y9y.png)

We can see that the api I wrote can't achieve 10,000 request per second,and other than by PC's performance, there are still a few improvements we can do:
*  Connection Pooling
*  Horizontal Scaling (maybe with k8s?)
*  DB Indexing
