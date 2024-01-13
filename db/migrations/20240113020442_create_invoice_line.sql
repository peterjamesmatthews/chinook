-- Create "InvoiceLine" table
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
