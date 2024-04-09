-- 001_create_cars_table.sql
CREATE TABLE IF NOT EXISTS cars (
                                    id SERIAL PRIMARY KEY,
                                    reg_num VARCHAR(20) NOT NULL,
                                    mark VARCHAR(50) NOT NULL,
                                    model VARCHAR(50) NOT NULL,
                                    year INTEGER NOT NULL,
                                    owner_name VARCHAR(50),
                                    owner_surname VARCHAR(50),
    ...
);

-- 002_create_people_table.sql
CREATE TABLE IF NOT EXISTS people (
                                      id SERIAL PRIMARY KEY,
                                      name VARCHAR(50) NOT NULL,
                                      surname VARCHAR(50) NOT NULL,
                                      patronymic VARCHAR(50),
);