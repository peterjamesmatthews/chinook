-- -- Create "Artist" table
CREATE TABLE `Artist` (
  `ArtistId` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(120) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  PRIMARY KEY (`ArtistId`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "Album" table
CREATE TABLE `Album` (
  `AlbumId` int NOT NULL AUTO_INCREMENT,
  `Title` varchar(160) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `ArtistId` int NOT NULL,
  PRIMARY KEY (`AlbumId`),
  INDEX `IFK_AlbumArtistId` (`ArtistId`),
  CONSTRAINT `FK_AlbumArtistId` FOREIGN KEY (`ArtistId`) REFERENCES `Artist` (`ArtistId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "Employee" table
CREATE TABLE `Employee` (
  `EmployeeId` int NOT NULL AUTO_INCREMENT,
  `LastName` varchar(20) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `FirstName` varchar(20) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `Title` varchar(30) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `ReportsTo` int NULL,
  `BirthDate` datetime NULL,
  `HireDate` datetime NULL,
  `Address` varchar(70) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `City` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `State` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Country` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `PostalCode` varchar(10) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Phone` varchar(24) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Fax` varchar(24) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Email` varchar(60) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  PRIMARY KEY (`EmployeeId`),
  INDEX `IFK_EmployeeReportsTo` (`ReportsTo`),
  CONSTRAINT `FK_EmployeeReportsTo` FOREIGN KEY (`ReportsTo`) REFERENCES `Employee` (`EmployeeId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "Customer" table
CREATE TABLE `Customer` (
  `CustomerId` int NOT NULL AUTO_INCREMENT,
  `FirstName` varchar(40) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `LastName` varchar(20) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `Company` varchar(80) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Address` varchar(70) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `City` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `State` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Country` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `PostalCode` varchar(10) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Phone` varchar(24) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Fax` varchar(24) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Email` varchar(60) CHARSET utf8mb3 NOT NULL COLLATE utf8mb3_general_ci,
  `SupportRepId` int NULL,
  PRIMARY KEY (`CustomerId`),
  INDEX `IFK_CustomerSupportRepId` (`SupportRepId`),
  CONSTRAINT `FK_CustomerSupportRepId` FOREIGN KEY (`SupportRepId`) REFERENCES `Employee` (`EmployeeId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "Invoice" table
CREATE TABLE `Invoice` (
  `InvoiceId` int NOT NULL AUTO_INCREMENT,
  `CustomerId` int NOT NULL,
  `InvoiceDate` datetime NOT NULL,
  `BillingAddress` varchar(70) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingCity` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingState` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingCountry` varchar(40) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `BillingPostalCode` varchar(10) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  `Total` decimal(10, 2) NOT NULL,
  PRIMARY KEY (`InvoiceId`),
  INDEX `IFK_InvoiceCustomerId` (`CustomerId`),
  CONSTRAINT `FK_InvoiceCustomerId` FOREIGN KEY (`CustomerId`) REFERENCES `Customer` (`CustomerId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "Genre" table
CREATE TABLE `Genre` (
  `GenreId` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(120) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  PRIMARY KEY (`GenreId`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "MediaType" table
CREATE TABLE `MediaType` (
  `MediaTypeId` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(120) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  PRIMARY KEY (`MediaTypeId`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "Track" table
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

-- -- Create "InvoiceLine" table
CREATE TABLE `InvoiceLine` (
  `InvoiceLineId` int NOT NULL AUTO_INCREMENT,
  `InvoiceId` int NOT NULL,
  `TrackId` int NOT NULL,
  `UnitPrice` decimal(10, 2) NOT NULL,
  `Quantity` int NOT NULL,
  PRIMARY KEY (`InvoiceLineId`),
  INDEX `IFK_InvoiceLineInvoiceId` (`InvoiceId`),
  INDEX `IFK_InvoiceLineTrackId` (`TrackId`),
  CONSTRAINT `FK_InvoiceLineInvoiceId` FOREIGN KEY (`InvoiceId`) REFERENCES `Invoice` (`InvoiceId`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `FK_InvoiceLineTrackId` FOREIGN KEY (`TrackId`) REFERENCES `Track` (`TrackId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "Playlist" table
CREATE TABLE `Playlist` (
  `PlaylistId` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(120) CHARSET utf8mb3 NULL COLLATE utf8mb3_general_ci,
  PRIMARY KEY (`PlaylistId`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

-- -- Create "PlaylistTrack" table
CREATE TABLE `PlaylistTrack` (
  `PlaylistId` int NOT NULL,
  `TrackId` int NOT NULL,
  PRIMARY KEY (`PlaylistId`, `TrackId`),
  INDEX `IFK_PlaylistTrackTrackId` (`TrackId`),
  CONSTRAINT `FK_PlaylistTrackPlaylistId` FOREIGN KEY (`PlaylistId`) REFERENCES `Playlist` (`PlaylistId`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `FK_PlaylistTrackTrackId` FOREIGN KEY (`TrackId`) REFERENCES `Track` (`TrackId`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
