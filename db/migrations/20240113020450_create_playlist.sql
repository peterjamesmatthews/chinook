-- Create "Playlist" table
CREATE TABLE `Playlist` (
  `PlaylistId` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(120) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  PRIMARY KEY (`PlaylistId`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
