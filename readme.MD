## Golang Replay REST API
This is a backend API written in Golang.<br>
The repo has been dockerized, so after installing docker and docker compose, it can be ran with docker compose up -d.<br>
The server will be ran on port 8080
## Env
Before running, please ensure the .env is present and filled up with the appropriate values.<br>
An example env file has been given named .env.example , the only thing required is AWS_ACCESS_KEY_ID,AWS_SECRET_ACCESS_KEY,AWS_S3_REGION,AWS_S3_BUCKET and ENCRYPT_SECRET_KEY.<br>
Please ensure ENCRYPT_SECRET_KEY is a AES key of at least 16 bytes. The key can be generated in python using  os.urandom(16).hex()<br>

## Routes
### Get Top 10 replays
/api/replays GET <br>
This will list 10 replays either the latest 10 replays, or top 10 highest liked replays.<br>
The following query parameters can be used:<br>
sortBy (string): either createdAt or likes to determine which sorting to be used<br>
limit (int): Must be bigger than 0, by default is set to 10 if not provided, gives the top N results.<br>
Returns:<br>
replayID<br>
replayTitle<br>
stageName<br>
createdAt<br>
likes<br>
commentCount<br>
### Create replay
/api/replays POST
### Get replay details
/api/replays/:replayID GET
### Download replay file
/api/replays/file/:replayID GET
### Like a replay
/api/replay-likes
### Comment on a replay
/api/replay-likes
