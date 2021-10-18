USE website;

-- CREATE TABLE user (
--     id bigint NOT NULL AUTO_INCREMENT,
--     name varchar(255),
--     username varchar(255),
--     password varchar(255),
--     time datetime DEFAULT CURRENT_TIMESTAMP,
--     PRIMARY KEY (id)
-- );

-- INSERT INTO user (name, username, password)
-- VALUES ('PongPong', 'PongPong', 'PongPong');

-- INSERT INTO user (id, name, username, password, time)
-- VALUES (2, 'Martin', 'Martin', 'Martin', DATE '2021-10-3');

-- INSERT INTO user (id, name, username, password, time)
-- VALUES (3, 'Viola', 'Viola', 'Viola', DATE '2021-10-5');

-- INSERT INTO user (id, name, username, password, time)
-- VALUES (4, 'Wanhsuan', 'Wanhsuan', 'Wanhsuan', DATE '2021-10-7');

-- INSERT INTO user (id, name, username, password, time)
-- VALUES (5, 'test', 'test', 'test', DATE '2021-10-9');

-- SELECT COUNT(*) FROM user;

-- SELECT * FROM user ORDER BY time DESC; 

-- SELECT * FROM user
-- WHERE id >= 2 AND id <=4
-- ORDER BY time DESC;

-- SELECT username FROM user
-- WHERE username = 'test' AND password = 'test'

-- UPDATE user
-- SET name = "丁滿" WHERE username = 'test';

-- DELETE FROM user;

SELECT * FROM user;
