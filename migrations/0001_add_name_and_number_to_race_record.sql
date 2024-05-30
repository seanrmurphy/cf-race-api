-- Migration number: 0001 	 2024-05-23
ALTER TABLE race_results ADD COLUMN first_name TEXT;
ALTER TABLE race_results ADD COLUMN last_name TEXT;
ALTER TABLE race_results ADD COLUMN race_number INTEGER ;
