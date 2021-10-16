CREATE TABLE product
(
  product_id INT NOT NULL,
  product_name VARCHAR NOT NULL,
  product_sku VARCHAR NOT NULL,
  product_image VARCHAR NOT NULL,
  product_price INT NOT NULL,
  product_stock INT NOT NULL,
  product_description VARCHAR NOT NULL,
  PRIMARY KEY (product_id)
);

CREATE TABLE customer
(
  customer_id INT NOT NULL,
  customer_name VARCHAR NOT NULL,
  customer_address VARCHAR NOT NULL,
  customer_phone VARCHAR NOT NULL,
  customer_email VARCHAR NOT NULL,
  PRIMARY KEY (customer_id)
);

CREATE TABLE supplier
(
  supplier_id INT NOT NULL,
  supplier_name VARCHAR NOT NULL,
  supplier_addres VARCHAR NOT NULL,
  PRIMARY KEY (supplier_id)
);

CREATE TABLE purchase_order
(
  purchase_order_id INT NOT NULL,
  purchase_price INT NOT NULL,
  purchase_date DATE NOT NULL,
  supplier_id INT NOT NULL,
  PRIMARY KEY (purchase_order_id),
  FOREIGN KEY (supplier_id) REFERENCES supplier(supplier_id)
);

CREATE TABLE merchant
(
  merchant_id INT NOT NULL,
  merchant_name VARCHAR NOT NULL,
  merchant_address VARCHAR NOT NULL,
  PRIMARY KEY (merchant_id)
);

CREATE TABLE product_supplier_detail
(
  product_id INT NOT NULL,
  supplier_id INT NOT NULL,
  PRIMARY KEY (product_id, supplier_id),
  FOREIGN KEY (product_id) REFERENCES product(product_id),
  FOREIGN KEY (supplier_id) REFERENCES supplier(supplier_id)
);

CREATE TABLE purchase_order_details
(
  quantity INT NOT NULL,
  purchase_order_id INT NOT NULL,
  product_id INT NOT NULL,
  PRIMARY KEY (purchase_order_id, product_id),
  FOREIGN KEY (purchase_order_id) REFERENCES purchase_order(purchase_order_id),
  FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE outlet
(
  outlet_id INT NOT NULL,
  outlet_name VARCHAR NOT NULL,
  outlet_address VARCHAR NOT NULL,
  merchant_id INT NOT NULL,
  PRIMARY KEY (outlet_id),
  FOREIGN KEY (merchant_id) REFERENCES merchant(merchant_id)
);

CREATE TABLE outlet_product_detail
(
  outlet_id INT NOT NULL,
  product_id INT NOT NULL,
  PRIMARY KEY (outlet_id, product_id),
  FOREIGN KEY (outlet_id) REFERENCES outlet(outlet_id),
  FOREIGN KEY (product_id) REFERENCES product(product_id)
);

CREATE TABLE users
(
  user_id INT NOT NULL,
  fullname VARCHAR NOT NULL,
  email VARCHAR NOT NULL,
  phone VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  outlet_id INT NOT NULL,
  merchant_id INT NOT NULL,
  PRIMARY KEY (user_id),
  FOREIGN KEY (outlet_id) REFERENCES outlet(outlet_id),
  FOREIGN KEY (merchant_id) REFERENCES merchant(merchant_id)
);

CREATE TABLE order
(
  order_id INT NOT NULL,
  order_date DATE NOT NULL,
  price INT NOT NULL,
  customer_id INT NOT NULL,
  outlet_id INT NOT NULL,
  user_id INT NOT NULL,
  PRIMARY KEY (order_id),
  FOREIGN KEY (customer_id) REFERENCES customer(customer_id),
  FOREIGN KEY (outlet_id) REFERENCES outlet(outlet_id),
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE order_detail
(
  quantity INT NOT NULL,
  product_id INT NOT NULL,
  order_id INT NOT NULL,
  PRIMARY KEY (product_id, order_id),
  FOREIGN KEY (product_id) REFERENCES product(product_id),
  FOREIGN KEY (order_id) REFERENCES order(order_id)
);