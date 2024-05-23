ALTER TABLE "replays"
ADD "replayFileName" VARCHAR NOT NULL;

ALTER TABLE "replays"
DROP COLUMN "replayURL";