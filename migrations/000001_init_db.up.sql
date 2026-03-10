create table if not exists "user" (
    "id" serial primary key,
    "fullname" varchar (80),
    "email" varchar (80) unique,
    "password" varchar (80),
    "address" varchar (80),
    "phone" varchar (15),
    "picture" varchar (255)
);

create table if not exists "products" (
    "id" serial primary key,
    "name" varchar (80),
    "description" varchar (255),
    "price" int,
    "quantity" int default 0
);

create table if not exists "category" (
    "id" serial primary key,
    "name" varchar (50)
);

create table if not exists product_categories (
    product_id int references products(id) on delete cascade,
    category_id int references category(id) on delete cascade,
    primary key (product_id, category_id)
);

create table if not exists "cart" (
    "id" serial primary key,
    "user_id" int references "user"("id"),
    "quantity" int,
    "product_id" int references "products"("id")
);

create table if not exists "transaction" (
    "id" serial primary key,
    "transaction_id" varchar(20) unique,
    "user_id" int references "user"("id"),
    "delivery_method" varchar(20),
    "subtotal" int,
    "delivery_fee" int, 
    "tax" int,
    "total" int,
    "date" date default current_date,
    "status" varchar, 
    "payment_method" varchar
);

create table if not exists "transaction_product" (
    "id" serial primary key,
    "transaction_id" varchar(20) references "transaction"("transaction_id"),
    "product_id" int references "products"("id"),
    "quantity" int,
    "size" varchar (10),
    "variant" varchar (10),
    "price" int
);

create table if not exists discount (
    "id" serial primary key,
    "product_id" int references "products"("id"),
    "isFlahSale" boolean,
    "discount_rate" float,
    "disc_description" varchar (100)
);

create table if not exists "product_variant" (
    "id" serial primary key,
    "product_id" int references "products"("id"),
    "variant_name" varchar (10),
    "add_price" int default 0
);

create table if not exists "review" (
    "id" serial primary key,
    "product_id" int references "products"("id"),
    "user_id" int references "user"("id"),
    "review_description" varchar (255), 
    "rating" float
);

create table if not exists "product_size" (
    "id" serial primary key,
    "product_id" int references "products"("id"),
    "size_name" varchar (10),
    "add_price" int default 0
);

create table if not exists "product_images" (
    "id" serial primary key,
    "product_id" int references "products"("id"),
    "image_url" varchar (255)
);