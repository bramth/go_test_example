CREATE TABLE training_user (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(20) NOT NULL
);

CREATE UNIQUE INDEX idx_unique_training_user_email ON training_user(LOWER(email));


INSERT INTO training_user(name, phone) VALUES('Wendy', 'wendy.adi@tokopedia.com', '+8201233321122');
INSERT INTO training_user(name, phone) VALUES('Yunita', 'yunita.ekawati@tokopedia.com','+82099999999999');