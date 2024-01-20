-- Create "PlaylistTrack" table
CREATE TABLE `PlaylistTrack` (
  `PlaylistId` int NOT NULL,
  `TrackId` int NOT NULL,
  PRIMARY KEY (`PlaylistId`, `TrackId`),
  INDEX `IFK_PlaylistTrackTrackId` (`TrackId`),
  CONSTRAINT `FK_PlaylistTrackPlaylistId` FOREIGN KEY (`PlaylistId`) REFERENCES `Playlist` (`PlaylistId`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `FK_PlaylistTrackTrackId` FOREIGN KEY (`TrackId`) REFERENCES `Track` (`TrackId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
