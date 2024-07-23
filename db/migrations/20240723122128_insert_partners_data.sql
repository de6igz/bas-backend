-- +goose Up
-- +goose StatementBegin
insert into partners (description, url)
values
    ('Российская компания, специализирующаяся на розничной торговле строительными материалами.',
        'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/2560px-LOGO_Petrovich.svg.png'),
    ('Группа компаний «Базис» известна как оптовый поставщик строительных материалов по всей территории Российской Федерации и в некоторых странах СНГ',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/4214665.png'),
    ('Один из ведущих российских промышленно-строительных холдингов, работающих в сфере разработки, производства, внедрения и поставок комплексных технологических решений для строительства объектов промышленно-гражданского и специального назначения. ',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/a7a2db20a21b52c7a3b3f7e093960483.png'),
    ('Комплексный поставщик инженерных систем',
    'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/etm_logo_white.svg'),
    ('Комплексное снабжение строительных объектов, продажа оптом и в розницу, ремонт и аренда строительного оборудования и инструмента, доставка товара',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/h_new_logo__new.png'),
    ('ООО «Строительные машины» занимается производством и реализацией механических и гидравлических станков для гибки, резки и правки арматурной стали, а так же продвижением и внедрением инновационных решений для разгрузки сыпучих материалов с помощью пневмотранспортного оборудования на строительном рынке России и СНГ.',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/header-logo-full.png'),
    ('Компания «СИРИУС СПб» является прямым представителем по Санкт-Петербургу и Ленинградской области крупнейшего на сегодняшний день производителя спецодежды и средств индивидуальной защиты в России ГК "СИРИУС" с 1998 года, сегодня признана одним из лидеров российского рынка',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo_1.png'),
    ('Российский онлайн-гипермаркет товаров для дома, дачи, стройки и ремонта',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo-filled.svg'),
    ('КРУПНЕЙШИЙ МЕТАЛЛМАРКЕТ РОССИИ',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo-w.svg'),
    ('Поставщик строительных конструкций, осуществляет поставку и аренду опалубочного оборудования.',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.png'),
    ('Поставщики крепежей',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.png83_83.webp'),
    ('Компания по производству крепежных деталей',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.svg'),
    ('ООО «НЕВСКИЙ ТЕХНОЛОГ» – российская торгово-производственная компания, специализирующаяся на производстве и поставках металлоизделий для строительных и производственных нужд как крупным оптовым заказчикам, так и розничным покупателям.',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/LogoNevskyi.png'),
    ('AG-Technologies – крупнейший российский поставщик спецодежды и экипировки для профессионалов',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/mainLogo.png'),
    ('ООО «Титан-Монолит» предлагает широкий ассортимент товаров для железобетонного строительства: около 30 видов фиксаторов арматуры, а также расходные материалы и комплектующие для монолитных и общестроительных работ.',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/new-logo.svg'),
    ('Интернет-магазин для оптовых покупок: более 850 000 электротехнических товаров, удобные способы оплаты и доставки. Кабель, светотехника, низковольтное и щитовое оборудование.',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/rs24.PNG'),
    ('Строительный холдинг Setl Group работает в Петербурге с 1994 года. Setl Group является одним из крупнейших финансово-промышленных объединений Северо-западного региона России',
     'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/sg.09f1d80.svg');


       ;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM partners
WHERE url IN (
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/2560px-LOGO_Petrovich.svg.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/4214665.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/a7a2db20a21b52c7a3b3f7e093960483.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/etm_logo_white.svg',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/h_new_logo__new.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/header-logo-full.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo_1.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo-filled.svg',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo-w.svg',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.png83_83.webp',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.svg',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/LogoNevskyi.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/mainLogo.png',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/new-logo.svg',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/rs24.PNG',
              'https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/sg.09f1d80.svg'
    );
-- +goose StatementEnd
