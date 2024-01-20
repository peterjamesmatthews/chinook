-- Create "Album" table
CREATE TABLE `Album` (
  `AlbumId` int NOT NULL AUTO_INCREMENT,
  `Title` varchar(160) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `ArtistId` int NOT NULL,
  PRIMARY KEY (`AlbumId`),
  INDEX `IFK_AlbumArtistId` (`ArtistId`),
  CONSTRAINT `FK_AlbumArtistId` FOREIGN KEY (`ArtistId`) REFERENCES `Artist` (`ArtistId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
