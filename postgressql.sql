CREATE TABLE IF NOT EXISTS artist (
    id SERIAL PRIMARY KEY,
    name VARCHAR ( 50 ) NOT NULL,
    title VARCHAR ( 50 ) UNIQUE NOT NULL,
    price VARCHAR ( 255 ) NOT NULL
);

SELECT id,name,title,price FROM artist LIMIT 5 OFFSET 1;

UPDATE artist 
SET name='nae56',title='tle56', price='66.7667'
WHERE id=8;

SELECT * FROM artist;

SELECT * FROM artist WHERE id=8;

DELETE FROM artist WHERE id=1;

INSERT INTO artist(name,title,price) VALUES('name3','title3','price3') ON CONFLICT DO NOTHING;
