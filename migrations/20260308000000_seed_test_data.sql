-- +goose Up
-- Insert test departments with 6 levels of depth
INSERT INTO "departments" ("name", "parent_id", "created_at") VALUES
-- Level 1: Root departments
('Engineering', NULL, NOW()),              -- id: 1
('Sales', NULL, NOW()),                    -- id: 2
('Human Resources', NULL, NOW()),          -- id: 3

-- Level 2: Children of Engineering (multiple siblings)
('Backend', 1, NOW()),                     -- id: 4
('Frontend', 1, NOW()),                    -- id: 5
('QA', 1, NOW()),                          -- id: 6
('DevOps', 1, NOW()),                      -- id: 7
('Security', 1, NOW()),                    -- id: 8
('Infrastructure', 1, NOW()),              -- id: 9

-- Level 3: Children of Backend
('Core Services', 4, NOW()),               -- id: 10
('API Development', 4, NOW()),             -- id: 11
('Microservices', 4, NOW()),               -- id: 12
('Batch Processing', 4, NOW()),            -- id: 13

-- Level 3: Children of DevOps
('Cloud Infrastructure', 7, NOW()),        -- id: 14
('Container Orchestration', 7, NOW()),     -- id: 15

-- Level 3: Children of Infrastructure
('Network Team', 9, NOW()),                -- id: 16
('Storage Team', 9, NOW()),                -- id: 17

-- Level 4: Children of Core Services
('Database Team', 10, NOW()),              -- id: 18
('Cache Team', 10, NOW()),                 -- id: 19
('Message Queue Team', 10, NOW()),         -- id: 20

-- Level 4: Children of API Development
('REST API', 11, NOW()),                   -- id: 21
('GraphQL Team', 11, NOW()),               -- id: 22

-- Level 4: Children of Cloud Infrastructure
('AWS Team', 14, NOW()),                   -- id: 23
('GCP Team', 14, NOW()),                   -- id: 24

-- Level 5: Children of Database Team
('PostgreSQL Specialists', 18, NOW()),     -- id: 25
('Query Optimization', 18, NOW()),         -- id: 26
('Backup & Recovery', 18, NOW()),          -- id: 27

-- Level 5: Children of Cache Team
('Redis Team', 19, NOW()),                 -- id: 28
('Memcached Team', 19, NOW()),             -- id: 29

-- Level 5: Children of REST API
('Authentication', 21, NOW()),             -- id: 30
('Data Validation', 21, NOW()),            -- id: 31

-- Level 6: Children of PostgreSQL Specialists
('PostgreSQL Performance', 25, NOW()),     -- id: 32
('PostgreSQL Security', 25, NOW()),        -- id: 33
('PostgreSQL Monitoring', 25, NOW()),      -- id: 34

-- Level 6: Children of Query Optimization
('Index Optimization', 26, NOW()),         -- id: 35
('Query Planning', 26, NOW()),             -- id: 36
('Statistics & Analysis', 26, NOW()),      -- id: 37

-- Level 6: Children of Redis Team
('Redis Performance', 28, NOW()),          -- id: 38
('Redis Replication', 28, NOW());          -- id: 39

-- Insert test employees across all levels
INSERT INTO "employees" ("full_name", "position", "hired_at", "created_at", "department_id") VALUES
-- Level 2 - Backend Department
('John Smith', 'Senior Backend Engineer', '2023-01-15', NOW(), 4),
('Marcus Johnson', 'Backend Engineer', '2023-06-01', NOW(), 4),
('Patricia Chen', 'Junior Backend Engineer', '2024-02-15', NOW(), 4),
('Daniel Lee', 'Backend Developer', '2023-11-20', NOW(), 4),
-- Level 2 - Frontend Department
('Sarah Johnson', 'Lead Frontend Engineer', '2022-06-20', NOW(), 5),
('Michael Torres', 'Senior Frontend Engineer', '2023-03-10', NOW(), 5),
('Nicole Foster', 'Frontend Developer', '2024-01-05', NOW(), 5),
('Christopher Bell', 'Junior Frontend Developer', '2024-04-12', NOW(), 5),
-- Level 2 - QA Department
('Mike Chen', 'QA Engineer', '2023-09-10', NOW(), 6),
('Lauren White', 'Senior QA Engineer', '2022-08-15', NOW(), 6),
('Brandon Harris', 'QA Automation Engineer', '2023-12-01', NOW(), 6),
-- Level 2 - DevOps Department
('Steven Garcia', 'DevOps Lead', '2022-04-12', NOW(), 7),
('Sophia Martinez', 'DevOps Engineer', '2023-05-20', NOW(), 7),
('Ryan Brown', 'Senior DevOps Engineer', '2023-01-30', NOW(), 7),
-- Level 2 - Security Department
('Jessica Rodriguez', 'Security Lead', '2022-03-15', NOW(), 8),
('Kevin Davis', 'Security Engineer', '2023-02-10', NOW(), 8),
('Megan Taylor', 'Penetration Tester', '2023-08-25', NOW(), 8),
-- Level 2 - Infrastructure Department
('Jason Miller', 'Infrastructure Lead', '2021-11-01', NOW(), 9),
('Ashley Clark', 'Infrastructure Engineer', '2022-06-14', NOW(), 9),
-- Level 3 - Core Services
('Emily Davis', 'Core Services Lead', '2022-03-01', NOW(), 10),
('Timothy Anderson', 'Senior Core Services Engineer', '2023-02-14', NOW(), 10),
('Victoria Prince', 'Core Services Developer', '2024-03-20', NOW(), 10),
-- Level 3 - API Development
('Robert Wilson', 'API Developer', '2023-11-05', NOW(), 11),
('Jennifer Lee', 'Senior API Developer', '2023-05-22', NOW(), 11),
('Nathan Phillips', 'API Engineer', '2024-01-30', NOW(), 11),
-- Level 3 - Microservices
('Grace Thompson', 'Microservices Lead', '2023-01-10', NOW(), 12),
('David Brown', 'Microservices Engineer', '2023-07-15', NOW(), 12),
-- Level 3 - Batch Processing
('Olivia Scott', 'Batch Processing Lead', '2022-11-08', NOW(), 13),
('Amanda Garcia', 'Batch Developer', '2024-03-20', NOW(), 13),
-- Level 3 - Cloud Infrastructure (DevOps child)
('Ethan Walker', 'Cloud Architect', '2023-10-25', NOW(), 14),
('Hannah Green', 'Cloud Engineer', '2024-02-05', NOW(), 14),
-- Level 3 - Container Orchestration (DevOps child)
('Thomas Jackson', 'Kubernetes Specialist', '2023-10-30', NOW(), 15),
('Rebecca Moore', 'Container Engineer', '2024-01-10', NOW(), 15),
-- Level 3 - Network Team (Infrastructure child)
('Alexander Hall', 'Network Engineer', '2023-12-20', NOW(), 16),
('Zoe Young', 'Network Specialist', '2024-04-01', NOW(), 16),
-- Level 3 - Storage Team (Infrastructure child)
('Catherine Hill', 'Storage Engineer', '2024-01-08', NOW(), 17),
('Samuel Lopez', 'Storage Specialist', '2023-09-12', NOW(), 17),
-- Level 4 - Database Team
('Lisa Anderson', 'Database Administrator', '2021-08-12', NOW(), 18),
('Kevin Chen', 'Senior DBA', '2022-04-08', NOW(), 18),
('Diana Nelson', 'Database Engineer', '2023-09-14', NOW(), 18),
('Matthew Carter', 'Junior Database Specialist', '2024-02-28', NOW(), 18),
-- Level 4 - Cache Team
('James Martinez', 'Cache Specialist', '2024-01-22', NOW(), 19),
('Elizabeth Mitchell', 'Senior Cache Engineer', '2023-07-11', NOW(), 19),
-- Level 4 - Message Queue Team
('Edward King', 'Message Queue Lead', '2023-06-05', NOW(), 20),
('Natalie Wright', 'Message Queue Engineer', '2024-03-10', NOW(), 20),
-- Level 4 - REST API
('Christopher Anderson', 'REST API Lead', '2023-02-01', NOW(), 21),
('Jessica Brown', 'REST API Developer', '2023-08-20', NOW(), 21),
-- Level 4 - GraphQL Team
('Michael Davis', 'GraphQL Specialist', '2024-01-15', NOW(), 22),
-- Level 4 - AWS Team
('Daniel Taylor', 'AWS Architect', '2022-09-10', NOW(), 23),
('Sarah Wilson', 'AWS Engineer', '2023-11-20', NOW(), 23),
-- Level 4 - GCP Team
('Lauren Martinez', 'GCP Engineer', '2023-10-15', NOW(), 24),
-- Level 5 - PostgreSQL Specialists
('Jennifer Lopez', 'PostgreSQL Engineer', '2023-07-18', NOW(), 25),
('Timothy Johnson', 'Senior PostgreSQL DBA', '2022-09-05', NOW(), 25),
('Victoria Anderson', 'PostgreSQL Developer', '2024-03-15', NOW(), 25),
-- Level 5 - Query Optimization
('Nicholas Brown', 'Query Optimization Specialist', '2024-02-01', NOW(), 26),
('Rachel Davis', 'Senior Query Analyst', '2023-08-20', NOW(), 26),
('Brandon Thompson', 'Query Developer', '2024-04-10', NOW(), 26),
-- Level 5 - Backup & Recovery
('Sophia Garcia', 'Backup Specialist', '2023-05-12', NOW(), 27),
('Lucas Martinez', 'Recovery Engineer', '2024-02-20', NOW(), 27),
-- Level 5 - Redis Team
('Amy Clark', 'Redis Specialist', '2023-09-15', NOW(), 28),
('Kevin Hall', 'Senior Redis Engineer', '2023-04-08', NOW(), 28),
-- Level 5 - Memcached Team
('Isabella White', 'Memcached Engineer', '2024-01-30', NOW(), 29),
-- Level 5 - Authentication
('Mason Rivera', 'Auth Engineer', '2023-06-22', NOW(), 30),
('Ava Rodriguez', 'Security Engineer', '2024-03-01', NOW(), 30),
-- Level 5 - Data Validation
('Ethan Phillips', 'Validation Specialist', '2023-11-14', NOW(), 31),
-- Level 6 - PostgreSQL Performance
('Oliver Jackson', 'Performance Tuning Engineer', '2023-09-15', NOW(), 32),
('Amelia Martinez', 'Senior Performance Engineer', '2022-11-08', NOW(), 32),
-- Level 6 - PostgreSQL Security
('Benjamin Garcia', 'Security Engineer', '2024-03-20', NOW(), 33),
('Charlotte Walker', 'Security Specialist', '2023-10-25', NOW(), 33),
-- Level 6 - PostgreSQL Monitoring
('Harper Lee', 'Monitoring Specialist', '2024-02-10', NOW(), 34),
-- Level 6 - Index Optimization
('Lucas Taylor', 'Index Specialist', '2023-10-30', NOW(), 35),
('Mia Johnson', 'Index Optimization Engineer', '2024-02-05', NOW(), 35),
-- Level 6 - Query Planning
('Logan Brown', 'Query Planner Expert', '2024-01-10', NOW(), 36),
('Chloe Davis', 'Query Planning Specialist', '2023-12-20', NOW(), 36),
('Noah Wilson', 'Junior Query Analyst', '2024-04-01', NOW(), 36),
-- Level 6 - Statistics & Analysis
('Ella Moore', 'Statistical Analyst', '2023-08-15', NOW(), 37),
-- Level 6 - Redis Performance
('Jackson Green', 'Redis Performance Engineer', '2024-01-20', NOW(), 38),
-- Level 6 - Redis Replication
('Avery Hill', 'Replication Specialist', '2024-02-15', NOW(), 39),
-- Level 1 - Sales
('Alice Thompson', 'VP of Sales', '2020-05-10', NOW(), 2),
('Robert Wilson', 'Senior Sales Manager', '2021-03-15', NOW(), 2),
('Nicole Foster', 'Sales Manager', '2022-07-20', NOW(), 2),
('James Martinez', 'Sales Representative', '2023-09-12', NOW(), 2),
('Jennifer Lee', 'Sales Representative', '2024-01-08', NOW(), 2),
-- Level 1 - HR
('Chris Martin', 'HR Director', '2019-12-01', NOW(), 3),
('Patricia Chen', 'Senior HR Manager', '2021-05-18', NOW(), 3),
('Michael Torres', 'HR Specialist', '2023-02-27', NOW(), 3),
('Lauren White', 'Recruiter', '2023-11-14', NOW(), 3);

-- +goose Down
DELETE FROM "employees" WHERE "department_id" IN (1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39);
DELETE FROM "departments" WHERE "name" IN ('Engineering', 'Sales', 'Human Resources', 'Backend', 'Frontend', 'QA', 'DevOps', 'Security', 'Infrastructure', 'Core Services', 'API Development', 'Microservices', 'Batch Processing', 'Cloud Infrastructure', 'Container Orchestration', 'Network Team', 'Storage Team', 'Database Team', 'Cache Team', 'Message Queue Team', 'REST API', 'GraphQL Team', 'AWS Team', 'GCP Team', 'PostgreSQL Specialists', 'Query Optimization', 'Backup & Recovery', 'Redis Team', 'Memcached Team', 'Authentication', 'Data Validation', 'PostgreSQL Performance', 'PostgreSQL Security', 'PostgreSQL Monitoring', 'Index Optimization', 'Query Planning', 'Statistics & Analysis', 'Redis Performance', 'Redis Replication');
ALTER SEQUENCE "employees_id_seq" RESTART WITH 1;
ALTER SEQUENCE "departments_id_seq" RESTART WITH 1;
