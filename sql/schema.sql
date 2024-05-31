-- CREATE TYPE CarAdsStatus as ENUM('New','Second');
-- CREATE TYPE BahanBakar as ENUM('Bensin','Hybrid','Electric');
-- CREATE TYPE TransType as ENUM('Automatic','Manual');
-- CREATE TYPE KapasitasMesin as ENUM('<1000 cc','>1000 - 1500 cc','>1500 - 2000 cc','>2000 - 3000 cc');

CREATE TABLE car_advertisements (
  id int NOT NULL AUTO_INCREMENT ,
  car_id int NOT NULL,
  dealer_id int NOT NULL,
  title varchar(255) NOT NULL,
  description varchar(255) NOT NULL,
  is_active tinyint(1) NOT NULL,
  posting_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  year_production int NOT NULL,
  color varchar(50) NOT NULL,
  mileage int NOT NULL,
  price int NOT NULL,
  status ENUM('New','Second') DEFAULT NULL,
  location varchar(255) NOT NULL,
  bahan_bakar ENUM('Bensin','Hybrid','Electric') NOT NULL,
  trans_type ENUM('Automatic','Manual') NOT NULL,
  kapasitas_mesin ENUM('<1000 cc','>1000 - 1500 cc','>1500 - 2000 cc','>2000 - 3000 cc') NOT NULL,
  stok int NOT NULL,
  thumbnailImage varchar(255) DEFAULT NULL,
  -- PRIMARY KEY (id),
  -- KEY fk_advertisements_cars (car_id),
  -- KEY fk_advertisements_dealers (dealer_id),
  CONSTRAINT fk_advertisements_cars FOREIGN KEY (car_id) REFERENCES cars (id),
  CONSTRAINT fk_advertisements_dealers FOREIGN KEY (dealer_id) REFERENCES dealers (id)
);


-- CREATE TYPE CarType AS ENUM('MPV','SUV','Hatchback','Sedan','Compact','Van','Minibus','Pick-Up','Truk','Double Cabin','Wagon','Coupe','Jeep','Convertible','Offroad','Sports','Classic','Bus');

CREATE TABLE cars (
  id int NOT NULL AUTO_INCREMENT ,
  brand varchar(255) NOT NULL,
  model varchar(255) NOT NULL,
  car_type ENUM('MPV','SUV','Hatchback','Sedan','Compact','Van','Minibus','Pick-Up','Truk','Double Cabin','Wagon','Coupe','Jeep','Convertible','Offroad','Sports','Classic','Bus') NOT NULL,
  PRIMARY KEY (id)
);



CREATE TABLE dealers (
  id int NOT NULL AUTO_INCREMENT ,
  dealer_name varchar(255) NOT NULL,
  no_telp varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  location varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  PRIMARY KEY (id)
);




CREATE TABLE messages (
  id int NOT NULL AUTO_INCREMENT ,
  user_id int NOT NULL,
  dealer_id int NOT NULL,
  message varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  -- PRIMARY KEY (id),
  -- KEY fk_messages_sender (user_id),
  -- KEY fk_messages_recipient (dealer_id),
  CONSTRAINT fk_messages_recipient FOREIGN KEY (dealer_id) REFERENCES dealers (id),
  CONSTRAINT fk_messages_sender FOREIGN KEY (user_id) REFERENCES users (id)
) ;







CREATE TABLE order_items (
  id int NOT NULL AUTO_INCREMENT ,
  order_id int NOT NULL,
  advertisement_id int NOT NULL,
  quantity int NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  price int DEFAULT NULL,
  -- PRIMARY KEY (id),
  -- KEY fk_order_items_orders (order_id),
  -- KEY fk_order_items_ads (advertisement_id),
  CONSTRAINT fk_order_items_ads FOREIGN KEY (advertisement_id) REFERENCES car_advertisements (id),
  CONSTRAINT fk_order_items_orders FOREIGN KEY (order_id) REFERENCES orders (id)
) ;


CREATE TABLE orders (
  id int NOT NULL AUTO_INCREMENT PRIMARY KEY  ,
  user_id int NOT NULL,
  payment_id int NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  dealer_id int DEFAULT NULL,
  price int DEFAULT NULL,
  nama_pembeli varchar(255) DEFAULT NULL,
  nomor_telepon_pembeli varchar(255) DEFAULT NULL,
  emailPembeli varchar(255) DEFAULT NULL,
  alamat_pembeli varchar(255) DEFAULT NULL,
  -- PRIMARY KEY (id),
  -- KEY fk_orders_users (user_id),
  -- KEY fk_orders_payments (payment_id),
  -- KEY fk_orders_dealers (dealer_id),
  CONSTRAINT fk_orders_dealers FOREIGN KEY (dealer_id) REFERENCES dealers (id),
  CONSTRAINT fk_orders_payments FOREIGN KEY (payment_id) REFERENCES payment_details (id),
  CONSTRAINT fk_orders_users FOREIGN KEY (user_id) REFERENCES users (id)
) ;



-- CREATE TYPE PaymentMethod as ENUM('Paypal','Credit Card','Cash','Loan','Debit Card');
-- CREATE TYPE PaymentStatus as ENUM('Paid','Pending','Canceled','Expired');

CREATE TABLE payment_details (
  id int NOT NULL AUTO_INCREMENT ,
  amount int NOT NULL,
  payment_method ENUM('Paypal','Credit Card','Cash','Loan','Debit Card') DEFAULT NULL,
  status  ENUM('Paid','Pending','Canceled','Expired') NOT NULL,
  provider varchar(255) NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  modified_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
) ;



-- CREATE TYPE Gender AS ENUM ('Male','Female');

CREATE TABLE users (
  id int NOT NULL AUTO_INCREMENT ,
  email varchar(255) NOT NULL,
  user_name varchar(255) NOT NULL,
  gender ENUM('Male','Female') NOT NULL,
  age int NOT NULL,
  address varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  PRIMARY KEY (id)
) ;






CREATE TABLE sessions (
    id  VARCHAR(255) NOT NULL  PRIMARY KEY ,
    ref_token_id VARCHAR(255) NOT NULL,
    username varchar(255) NOT NULL,
    refresh_token TEXT NOT NULL,
    expires_at timestamp NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp
);





