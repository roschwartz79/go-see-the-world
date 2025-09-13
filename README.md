# go-see-the-world
Learning the world of Go and GCP

This is a simple application written in GoLang.

### Key Technologies

- Go
- GIN
- Docker
- Google Cloud Platform
  - Cloud Build
  - Artifactory
  - Cloud Run

### Getting Started

#### Run the Application

To start the server, open your terminal, navigate to the project directory, and run the following command. The application will start on `http://localhost:8080`.

```bash
go run .
```

### API Endpoints

Use the following `cURL` commands to interact with the API endpoints. The Cloud Run URL is

```
https://go-see-the-world-894054645138.us-central1.run.app
```

#### Add a New Player

This command sends a **POST** request with a JSON body to the `/players` endpoint to add a new player to the team.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Wayne Gretzky", "position": "Center"}' http://localhost:8080/players/<id>
```

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Wayne Gretzky", "position": "Center"}' https://go-see-the-world-894054645138.us-central1.run.app/players/<id>
```

#### Add a New Coach

This command sends a **POST** request with a JSON body to the /coaches endpoint to add a new coach.

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Scotty Bowman", "position": "Head Coach"}' http://localhost:8080/coaches/<id>
```

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "Scotty Bowman", "position": "Head Coach"}' https://go-see-the-world-894054645138.us-central1.run.app/coaches/<id>
```

#### Create a Team

This command sends a Post request to the /teams endpoint to create a Team.

_Note: The Players and Coaches slice doesn't need to be passed in since Go will initialize the slice as empty to start off with._

```bash
curl -X POST http://localhost:8080/teams \
-H "Content-Type: application/json" \
-d '{
  "Id": 12345,
  "Name": "Go Gophers"
}'
```

```bash
curl -X POST https://go-see-the-world-894054645138.us-central1.run.app/teams \
-H "Content-Type: application/json" \
-d '{
  "Id": 12345,
  "Name": "Go Gophers"
}'
```

#### View the Full Team Roster

This command sends a GET request to the /teams endpoint to retrieve the complete team roster, including all players and coaches.

```bash
curl http://localhost:8080/teams
```

```bash
curl https://go-see-the-world-894054645138.us-central1.run.app/teams
```

### Running in Google Cloud

Google Cloud is what I'm using to build, deploy and run this application.

#### Building

To kickoff a new pipeline and push an image to Artifactory, just create a Tag. Cloud Build will be triggered off of 
a new Tag and will start a new pipeline when needed.

#### Deploying

TO deploy

