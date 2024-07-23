-- +goose Up
-- +goose StatementBegin
insert into docs (preview_url, document_url)
values ('https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/благодарность_гаврилов.jpg',
        'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/2 Благодарность Гаврилов.pdf'),
    ('https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/благодарность_грачев.jpg',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/1 Благодарность Грачев.pdf'),
    ('https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/благодарность_гретчин.jpg',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/3 Благодарность Гретчин.pdf'),
    ('https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/сертификат_партнера_дока_рус.PNG',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/Сертификат партнера ООО Дока Рус.pdf'),
    ('https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/СРО_page-0001.jpg',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/сро.pdf'),
    ('https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/Necessarily_1.PNG',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/2023-11-14_002.pdf'),
    ('https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/Necessarily_2.PNG',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/2023-11-14_001.pdf');

;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
