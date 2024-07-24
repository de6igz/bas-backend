-- +goose Up
-- +goose StatementBegin
DO $$
    DECLARE
        project_id INTEGER;
    BEGIN
        -- Получаем id проекта 'ЖК "Imperial Club""'
        SELECT id INTO project_id FROM bas.public.projects WHERE full_name = 'ЖК "Imperial Club"';

        -- Вставляем данные в таблицу pictures
        INSERT INTO bas.public.pictures (project_id, url)
        VALUES
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/preview.jpeg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/2f783b839fff6e0c6ad74f09ac1a9789.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/7f9c63567869542407bc4cbaac023dc3.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/62d68aa57f5df.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/78f7f811451fcac815567d2058e06fe5.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/97b3c70580bee66f4fc9bcc202d6dbcd.jpg');
    END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DO $$
    DECLARE
        project_id_var INTEGER;
    BEGIN
        -- Получаем id проекта 'ЖК "Imperial Club"'
        SELECT id INTO project_id_var FROM projects WHERE full_name = 'ЖК "Imperial Club"';

        -- Удаляем данные из таблицы pictures
        DELETE FROM pictures WHERE project_id = project_id_var;
    END $$;
-- +goose StatementEnd
