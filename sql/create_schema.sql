CREATE TABLE tracks (
    id  serial  PRIMARY KEY,
    title VARCHAR (255) UNIQUE NOT NULL,
    artist VARCHAR (255) NOT NULL,
    releaseDate TIMESTAMP DEFAULT NOW()
);
