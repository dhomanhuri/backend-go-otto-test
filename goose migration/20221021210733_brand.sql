-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE Brands(
    id int(11) AUTO_INCREMENT,
    name varchar(255),
    PRIMARY KEY(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE Brands;
