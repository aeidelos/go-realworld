CREATE TABLE IF NOT EXISTS user (
    user_id     INT NOT NULL AUTO_INCREMENT,
    email       VARCHAR(30),
    token       VARCHAR(255) NULL,
    username    VARCHAR(30) NULL,
    bio         VARCHAR(255) NULL,
    image       VARCHAR(255) NULL,
    PRIMARY KEY (user_id)
)
