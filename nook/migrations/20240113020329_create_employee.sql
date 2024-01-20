-- Create "Employee" table
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
