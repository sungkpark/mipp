CREATE TABLE IF NOT EXISTS domains (
    `domainId` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `domainName` VARCHAR(255) NOT NULL,
    `companyInformation` TEXT NULL,
    `verified` BOOLEAN DEFAULT false,

    PRIMARY KEY (domainId),
    UNIQUE (domainName)
)