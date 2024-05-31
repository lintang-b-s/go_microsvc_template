-- carDeail section ---
-- name: ShowAdsByID :one
SELECT ad.id, ad.car_id, ad.dealer_id, ad.title, ad.description, ad.is_active, ad.posting_date,
     ad.year_production, ad.color, ad.mileage,
     ad.price, ad.status, ad.location, ad.bahan_bakar, ad.trans_type, ad.kapasitas_mesin, ad.stok, ad.thumbnailImage, c.brand,
      c.model, c.car_type,
       d.dealer_name, d.no_telp, d.email, d.location   FROM car_advertisements  ad
    INNER JOIN cars c ON c.id=ad.car_id
    INNER JOIN dealers d ON ad.dealer_id=d.id
    WHERE ad.id=  ?;



-- online ads (checkout page) section ---
-- name: SelectAdsByID :one
SELECT * FROM car_advertisements WHERE id=?  ;




-- name: InsertPaymentDetails :exec
 INSERT INTO payment_details(amount, payment_method, status, provider)
   VALUES (?, ?, ?, ?) ;

-- name: GetInsertedPaymentID :one
SELECT LAST_INSERT_ID();

-- name: UpdateAds :exec
UPDATE car_advertisements SET stok = stok-?
   WHERE id= ?;



-- name: InsertOrder :exec
INSERT INTO orders(user_id, payment_id, price, alamat_pembeli, nomor_telepon_pembeli, emailPembeli, nama_pembeli, dealer_id)
    VALUES (?, ?, ?,?, ?,  ?, ?, ?);



-- name: InsertOrderItems :exec
INSERT INTO order_items(order_id, advertisement_id, quantity)
    VALUES (?, ?, ?);



-- filter car page section ---

-- name: FilterCar :many
SELECT car_advertisements.id,cars.brand, 
    		cars.model,
           car_advertisements.year_production, 
           car_advertisements.price, 
           car_advertisements.mileage,
           car_advertisements.location, 
           car_advertisements.bahan_bakar, 
           car_advertisements.trans_type
    FROM cars
    JOIN car_advertisements ON cars.id = car_advertisements.car_id
    WHERE (? = 'Semua' OR cars.brand = ?)
    AND (car_advertisements.color = ? OR ? = 'Semua')
    AND (car_advertisements.location = ? OR ? = 'Semua')
    AND (car_advertisements.bahan_bakar = ? OR ? = 'Semua')
    AND (car_advertisements.trans_type = ? OR ? = 'Semua')
    AND car_advertisements.year_production BETWEEN ? AND ?
    AND car_advertisements.price BETWEEN ? AND ?
    AND car_advertisements.mileage BETWEEN ? AND ?;


-- name: HomePage :many
SELECT ad.id, ad.car_id, ad.dealer_id, ad.title, ad.description, ad.is_active, ad.posting_date, ad.year_production, ad.color, ad.mileage, ad.price, ad.status, ad.location, ad.bahan_bakar, ad.trans_type, ad.kapasitas_mesin, ad.stok, c.brand, c.model, c.car_type, ad.thumbnailImage   FROM car_advertisements  ad
    INNER JOIN cars c ON c.id=ad.car_id;







-- user management section ---

-- name: InsertUser :exec
INSERT INTO users(
     user_name, email,  gender, password, age, address
)VALUES(?, ?, ?, ?, ?, ?);

    -- ?, ?, ?, ?, ?, ?
    --  ?, ?, ?, ?, ?, ?

-- name: InsertSession :exec
INSERT INTO sessions(
     ID,ref_token_id, username, refresh_token, expires_at
)VALUES(
    UUID(),?, ?, ?, ?
);


-- name: GetUser :one
SELECT id, user_name, email,  gender, age, address
FROM users
WHERE id=?;

-- name: GetUserByEmail :one
SELECT id, user_name, email, password,  gender, age, address
FROM users
WHERE email=?;

-- name: GetSession :one
SELECT id, username, refresh_token, expires_at, created_at
FROM sessions
WHERE id=?;



-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id=?;




-- compare car page section ---

-- name: GetActiveAds :many
SELECT title FROM car_advertisements WHERE is_active = 1;

-- name: GetAdsByTitle :one
SELECT * FROM car_advertisements WHERE title=?

