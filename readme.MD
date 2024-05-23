## Golang Replay REST API
This is a backend API written in Golang.<br>
The repo has been dockerized, so after installing docker and docker compose, it can be ran with docker compose up -d.<br>
The server will be ran on port 8080<br>
On startup, please wait until clamAV has finished setting up its database, may take up to 5 mins or more, depending on system specs
## Env
Before running, please ensure the .env is present and filled up with the appropriate values.<br>
An example env file has been given named .env.example , the only thing required is AWS_ACCESS_KEY_ID,AWS_SECRET_ACCESS_KEY,AWS_S3_REGION,AWS_S3_BUCKET and ENCRYPT_SECRET_KEY.<br>
Please ensure ENCRYPT_SECRET_KEY is a AES key of at least 16 bytes. The key can be generated in python using  os.urandom(16).hex()<br>

## Routes
It is recommended to use Postman to test out the api. Import the json file titled Replays.postman_collection.json from the repo into postman for easier testing.
### Get Top 10 replays
/api/replays GET <br>
This will list 10 replays either the latest 10 replays, or top 10 highest liked replays.<br>
The following query parameters can be used:<br>
sortBy (string): either createdAt or likes to determine which sorting to be used<br>
limit (int): Must be bigger than 0, by default is set to 10 if not provided, gives the top N results.<br>
Returns:replayID, replayTitle, stageName, createdAt, likes, commentCount<br>
### Create replay
/api/replays POST multipart/form-data<br>
Creates a replay. The Content type must be set to multipart/form-data, because we are uploading files to the API<br>
File uploaded will be encrypted and sent to AWS S3.<br>
Required fields:<br>
replayTitle: string<br>
stageTitle: string<br>
replayFile: file<br>
Returns JSON:replayID, replayTitle, stageName, createdAt, likes, commentCount<br>
### Get replay details
/api/replays/:replayID GET (example: localhost:8080/api/replays/2)<br>
Returns:replayID, replayTitle, stageName, createdAt, likes, commentCount, comments (list of comments)<br>
### Download replay file (example: localhost:8080/api/replays/file/2)
/api/replays/file/:replayID GET
This will download the replay file onto the backend server, and be decrypted and sent to the requester as application/octet-stream, with Content-Disposition set with attachment; filename={FILENAME}<br>
### Like replay
/api/replay-likes POST application/json<br>
Required fields: PostID int<br>
Returns: replayID, likeID, createdAt<br>
### Comment on replay
/api/replay-likes POST application/json<br>
Required fields: PostID int, commentContent string<br>
Returns: replayID, commentID, commentContent, createdAt<br>
## What can be improved for scaling and high performance
- The comment shown in post detail should be a different endpoint to list comments, with proper keyset/range query pagination<br>
- Ensure file size limit for replay<br>
- Use virustotal api for scanning (i originally wanted to, but for this case I opted for clamav, as i was thinking about the pricing for virustotal)<br>
- Compress the replay file before encrypting<br>
- User authentication should be implemented for access control, so comments and likes can be linked to user<br>
- (NOTE) The likes should be counted accurately because I am not updating the replay row again and again, inserts should make the count accurate.