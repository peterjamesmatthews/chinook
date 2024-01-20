-- Create "Artist" table
CREATE TABLE `Artist` (
  `ArtistId` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(120) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  PRIMARY KEY (`ArtistId`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
