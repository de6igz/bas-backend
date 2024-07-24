-- +goose Up
-- +goose StatementBegin
DO $$
    DECLARE
        project_id INTEGER;
    BEGIN
        -- Получаем id проекта 'ЖК "Сенат"'
        SELECT id INTO project_id FROM bas.public.projects WHERE full_name = 'ЖК "Сенат"';

        -- Вставляем данные в таблицу pictures
        INSERT INTO bas.public.pictures (project_id, url)
        VALUES
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/preview.jpeg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/2.jpeg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/3.jpeg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/4.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/5.jpeg');
    END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DO $$
    DECLARE
        project_id_var INTEGER;
    BEGIN
        -- Получаем id проекта 'ЖК "Сенат"'
        SELECT id INTO project_id_var FROM projects WHERE full_name = 'ЖК "Сенат"';

        -- Удаляем данные из таблицы pictures
        DELETE FROM pictures WHERE project_id = project_id_var  ;
    END $$;
-- +goose StatementEnd
