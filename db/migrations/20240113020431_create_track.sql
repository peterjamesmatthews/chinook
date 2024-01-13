-- Create "Track" table
CREATE TABLE `Track` (
  `TrackId` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(200) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `AlbumId` int NULL,
  `MediaTypeId` int NOT NULL,
  `GenreId` int NULL,
  `Composer` varchar(220) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Milliseconds` int NOT NULL,
  `Bytes` int NULL,
  `UnitPrice` decimal(10, 2) NOT NULL,
  PRIMARY KEY (`TrackId`),
  INDEX `IFK_TrackAlbumId` (`AlbumId`),
  INDEX `IFK_TrackGenreId` (`GenreId`),
  INDEX `IFK_TrackMediaTypeId` (`MediaTypeId`),
  CONSTRAINT `FK_TrackAlbumId` FOREIGN KEY (`AlbumId`) REFERENCES `Album` (`AlbumId`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `FK_TrackGenreId` FOREIGN KEY (`GenreId`) REFERENCES `Genre` (`GenreId`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `FK_TrackMediaTypeId` FOREIGN KEY (`MediaTypeId`) REFERENCES `MediaType` (`MediaTypeId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
