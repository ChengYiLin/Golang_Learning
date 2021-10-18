USE website;

CREATE TABLE message (
    id bigint NOT NULL AUTO_INCREMENT,
    user_id bigint,
	content varchar(255) NOT NULL,
    time datetime DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);

-- INSERT INTO message (user_id, content)
-- VALUES (2, 'Take a Break');

-- SELECT username
-- FROM user
-- JOIN message ON user.id = message.user_id;

SELECT user.username, message.content
FROM user
JOIN message ON user.id = message.user_id AND user.username = "Martin";
