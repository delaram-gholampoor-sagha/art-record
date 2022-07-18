/* For license and copyright information please see LEGAL file in repository */

CREATE TABLE `Voucher` (
  `VoucherID` BINARY(16) ,
  `Type` SMALLINT ,
  `Time` TIMESTAMP ,
  `RequestID` BINARY(16),
   KEY(VoucherID)
);

/* https://www.javatpoint.com/mysql-procedure */

/* DELIMITER &&  
CREATE PROCEDURE procedure_name [[IN | OUT | INOUT] parameter_name datatype [, parameter datatype]) ]    
BEGIN    
    Declaration_section    
    Executable_section    
END &&  
DELIMITER ;    */


/* https://www.mysqltutorial.org/mysql-limit.aspx */


DELIMITER &&

CREATE PROCEDURE Voucher_Save (`voucherID` BINARY(16),  `type` SMALLINT,
  `time` TIMESTAMP, `requestID` BINARY(16) , OUT `NumberOfVersion` INT) 

   BEGIN

    -- Everything between start transaction and commit is one atomic action. 
  

        INSERT INTO Voucher (VoucherID , Type , Time , RequestID) VALUES (voucherID , type , time , requestID);
        
        SELECT COUNT(VoucherID) INTO `NumberOfVersion` FROM Voucher WHERE `voucherID` = `voucherID` ;


   END &&

DELIMITER;



-- count ////////////////////////

DELIMITER &&

CREATE PROCEDURE Voucher_Count ( `voucherID` BINARY(16), OUT `NumberOfVersion` INT) 
   BEGIN
  
       SELECT COUNT(VoucherID) INTO `NumberOfVersion` FROM `Voucher` WHERE `VoucherID` = `voucherID`; 

   END &&

DELIMITER;


-- Get ///////////////////////////

DELIMITER &&

CREATE PROCEDURE Voucher_Get ( `voucherID` BINARY(16), `versionOffset` INT ) 

   BEGIN

       SELECT * FROM Voucher WHERE `VoucherID` = voucherID  LIMIT = versionOffset

       SELECT COUNT(VoucherID) INTO `NumberOfVersion` FROM Voucher WHERE `voucherID` = `voucherID` ;


   END &&

DELIMITER;










