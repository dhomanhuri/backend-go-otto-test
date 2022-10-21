-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE Redemptions(
    id int(11) AUTO_INCREMENT,
    voucher_id int(11),
    transaction_id int(11),
    quantity int(11),
    PRIMARY KEY (id),
    FOREIGN KEY (transaction_id) REFERENCES Transactions(id),
    FOREIGN KEY (voucher_id) REFERENCES Vouchers(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE Redemptions;