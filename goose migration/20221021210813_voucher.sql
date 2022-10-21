-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE Vouchers(
    id int(11) AUTO_INCREMENT,
    name varchar(255),
    price int(11),
    brand_id int(11),
    PRIMARY KEY (id),
    FOREIGN KEY (brand_id) REFERENCES Brands(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE Vouchers;