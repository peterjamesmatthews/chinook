-- Create "Customer" table
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
