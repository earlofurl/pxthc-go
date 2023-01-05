ALTER TABLE facilities ADD CONSTRAINT UC_Name UNIQUE (name);
ALTER TABLE facilities ADD CONSTRAINT UC_License_Number UNIQUE (license_number);