-- +goose Up
-- +goose StatementBegin
insert into bas.public.partners (url, description)
values ('Интернет-магазин для оптовых покупок: более 850 000 электротехнических товаров, удобные способы оплаты и доставки. Кабель, светотехника, низковольтное и щитовое оборудование.','https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/rs24.webp'),
       (),
       ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
