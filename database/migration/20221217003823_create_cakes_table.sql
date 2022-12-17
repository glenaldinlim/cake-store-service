-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cakes(
    id          BIGINT AUTO_INCREMENT,
    title       VARCHAR(100)            NOT NULL,
    description LONGTEXT                NOT NULL,
    rating      FLOAT(3,1) DEFAULT 0.0  NOT NULL,
    image       VARCHAR(255)            NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW() NULL,
	updated_at  TIMESTAMP DEFAULT NOW() NULL,

    constraint cakes_pk
		primary key (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cakes
-- +goose StatementEnd
