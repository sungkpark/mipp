CREATE TABLE IF NOT EXISTS ideas (
    `ideaId` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `username` VARCHAR(255) NOT NULL,
    `capturedUrl` VARCHAR(2083) NOT NULL,
    `domainId` INT UNSIGNED NOT NULL,
    `createdAt` TIMESTAMP NOT NULL,
    `updatedAt` TIMESTAMP,

    PRIMARY KEY (ideaId),
    FOREIGN KEY (domainId) REFERENCES domains(domainId)
)