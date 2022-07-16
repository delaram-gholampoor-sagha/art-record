/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher` (
  `VoucherID` BINARY(16) ,
  `Type` SMALLINT ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16),
   KEY(VoucherID)
);




DELIMITER $$

CREATE PROCEDURE Save_Voucher (IN `IN_VoucherID` BINARY(16), IN `IN_Type` SMALLINT,
 IN `IN_Time` TIMESTAMP,IN `IN_RequestID` BINARY(16) , OUT `NumberOfVersion` INT) 
   BEGIN
    -- SQL EXCEPTION
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
    END;
    -- Everything between start transaction and commit is one atomic action. 
    START TRANSACTION;
        INSERT INTO Voucher (VoucherID , Type , Time , RequestID) VALUES (IN_VoucherID , IN_Type , IN_Time , IN_RequestID);
        
        SELECT COUNT(VoucherID) INTO `NumberOfVersion` FROM Voucher WHERE `VoucherID` = `VoucherID` ;
    COMMIT;

END$$

DELIMITER;




-- Return Count
SELECT COUNT(VoucherID) FROM `Voucher` WHERE `VoucherID` = ?;

-- Return Get 
SELECT *, COUNT(VoucherID) FROM Voucher WHERE `VoucherID` = ?  LIMIT = ?



