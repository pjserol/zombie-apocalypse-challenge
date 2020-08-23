DROP TABLE IF EXISTS `zombie_data`;

CREATE TABLE `patient_details` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(80),
    `hospital_id` int DEFAULT NULL,
    `illness_id` int DEFAULT NULL,
    `severity_level` int DEFAULT NULL
);


