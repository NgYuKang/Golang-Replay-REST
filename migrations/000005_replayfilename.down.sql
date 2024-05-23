ALTER TABLE "replays"
ADD "replayURL" VARCHAR NOT NULL;

ALTER TABLE "replays"
DROP COLUMN "replayFileName";