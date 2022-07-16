CREATE TABLE `Voucher` (
  `VoucherID` BINARY(16) PRIMARY KEY NOT NULL,
  `Type` SMALLINT NOT NULL,
  `Time` TIMESTAMP NOT NULL,
  `RequestID` BINARY(16) NOT NULL
);

CREATE INDEX Index_Voucher ON Voucher (`VoucherID`);


DELIMITER $$

CREATE PROCEDURE Save_Voucher (IN `IN_VoucherID` BINARY(16), IN `IN_Type` SMALLINT, IN `IN_Time` TIMESTAMP,IN `IN_RequestID` BINARY(16)) BEGIN
    -- SQL EXCEPTION
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
    END;
    -- Everything between start transaction and commit is one atomic action. 
    START TRANSACTION;
        INSERT INTO Voucher (VoucherID , Type , Time , RequestID) VALUES (IN_VoucherID , IN_Type , IN_Time , IN_RequestID);
        
        SELECT COUNT(VoucherID) AS Length_VoucherID FROM Voucher;
    COMMIT;

END$$

DELIMITER;




-- Return Count
SELECT COUNT(VoucherID) FROM `Voucher` WHERE `VoucherID` = ?;

-- Return Get 
-- TODO:: Not sure we had Time order or not!
SELECT *, COUNT(VoucherID) FROM Voucher WHERE `VoucherID` = ? ORDER BY Voucher.Time DESC LIMIT ?




