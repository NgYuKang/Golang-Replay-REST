CREATE TABLE "replays"(
    "replayID" INT NOT NULL PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    "replayTitle" VARCHAR NOT NULL,
    "stageName" VARCHAR NOT NULL,
    "createdAt" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE "replayComments"(
    "replayID" INT NOT NULL,
    "commentID" INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    "commentContent" VARCHAR,
    PRIMARY KEY ("replayID", "commentID"),
    FOREIGN KEY ("replayID") REFERENCES replays ("replayID")
    -- can add userID
);

CREATE TABLE "replayLikes"(
    "replayID" INT NOT NULL,
    "likeID" INT NOT NULL GENERATED ALWAYS AS IDENTITY,
    PRIMARY KEY ("replayID", "likeID"),
    FOREIGN KEY ("replayID") REFERENCES replays ("replayID")
    -- can add userID, set unique key, prevent user from liking again
);