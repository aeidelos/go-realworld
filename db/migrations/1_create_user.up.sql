CREATE TABLE IF NOT EXISTS user (
    user_id     INT NOT NULL AUTO_INCREMENT,
    email       VARCHAR(30),
    token       VARCHAR(255),
    username    VARCHAR(30),
    bio         VARCHAR(255),
    image       VARCHAR(255),
    PRIMARY KEY (user_id)
)
