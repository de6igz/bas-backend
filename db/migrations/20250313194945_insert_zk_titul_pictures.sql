-- +goose Up
-- +goose StatementBegin
DO $$
    DECLARE
        project_id INTEGER;
    BEGIN
        -- Получаем id проекта 'ЖК "Титул"'
        SELECT id INTO project_id FROM bas.public.projects WHERE full_name = 'ЖК "Титул"';

        -- Вставляем данные в таблицу pictures
        INSERT INTO bas.public.pictures (project_id, url)
        VALUES
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/preview.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/1.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/2.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/3.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/4.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/5.jpg'),
        (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/6.jpg'),
            (project_id, 'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/7.jpg');
    END $$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DO $$
    DECLARE
        project_id_var INTEGER;
    BEGIN
        -- Получаем id проекта 'ЖК "Парадный Ансамбль"'
        SELECT id INTO project_id_var FROM projects WHERE full_name = 'ЖК "Титул"';

        -- Удаляем данные из таблицы pictures
        DELETE FROM pictures WHERE project_id = project_id_var;
    END $$;
-- +goose StatementEnd
