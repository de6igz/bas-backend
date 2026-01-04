package dataproviders

import (
	"bas-backend/config"
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func initSQLite(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	dbPath := cfg.Database.Path
	if dbPath == "" {
		dbPath = "./db/data.sqlite"
	}

	if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
		return nil, fmt.Errorf("create sqlite directory: %w", err)
	}

	conn, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on&cache=shared", dbPath))
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	conn.SetMaxOpenConns(1)
	conn.SetMaxIdleConns(1)
	conn.SetConnMaxLifetime(time.Hour)

	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	if err := ensureSchema(ctx, conn); err != nil {
		conn.Close()
		return nil, fmt.Errorf("init sqlite schema: %w", err)
	}

	return conn, nil
}

func ensureSchema(ctx context.Context, db *sql.DB) error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS partners (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			url TEXT NOT NULL,
			description TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS projects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT NOT NULL,
			url TEXT,
			status TEXT,
			short_name TEXT NOT NULL,
			builder_name TEXT NOT NULL,
			body TEXT,
			latitude REAL,
			longitude REAL
		);`,
		`CREATE TABLE IF NOT EXISTS pictures (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			project_id INTEGER NOT NULL,
			url TEXT NOT NULL,
			FOREIGN KEY(project_id) REFERENCES projects(id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS docs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			preview_url TEXT NOT NULL,
			document_url TEXT NOT NULL
		);`,
	}

	for _, stmt := range statements {
		if _, err := db.ExecContext(ctx, stmt); err != nil {
			return fmt.Errorf("run schema statement: %w", err)
		}
	}

	if err := seedPartners(ctx, db); err != nil {
		return err
	}
	if err := seedProjects(ctx, db); err != nil {
		return err
	}
	if err := seedDocs(ctx, db); err != nil {
		return err
	}

	return nil
}

func seedPartners(ctx context.Context, db *sql.DB) error {
	var count int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM partners").Scan(&count); err != nil {
		return fmt.Errorf("count partners: %w", err)
	}
	if count > 0 {
		return nil
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin partners tx: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO partners (description, url) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("prepare partners insert: %w", err)
	}
	defer stmt.Close()

	partners := []struct {
		description string
		url         string
	}{
		{"Российская компания, специализирующаяся на розничной торговле строительными материалами.", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/2560px-LOGO_Petrovich.svg.png"},
		{"Группа компаний «Базис» известна как оптовый поставщик строительных материалов по всей территории Российской Федерации и в некоторых странах СНГ", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/4214665.png"},
		{"Один из ведущих российских промышленно-строительных холдингов, работающих в сфере разработки, производства, внедрения и поставок комплексных технологических решений для строительства объектов промышленно-гражданского и специального назначения. ", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/a7a2db20a21b52c7a3b3f7e093960483.png"},
		{"Комплексный поставщик инженерных систем", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/etm_logo_white.svg"},
		{"Комплексное снабжение строительных объектов, продажа оптом и в розницу, ремонт и аренда строительного оборудования и инструмента, доставка товара", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/h_new_logo__new.png"},
		{"ООО «Строительные машины» занимается производством и реализацией механических и гидравлических станков для гибки, резки и правки арматурной стали, а так же продвижением и внедрением инновационных решений для разгрузки сыпучих материалов с помощью пневмотранспортного оборудования на строительном рынке России и СНГ.", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/header-logo-full.png"},
		{"Компания «СИРИУС СПб» является прямым представителем по Санкт-Петербургу и Ленинградской области крупнейшего на сегодняшний день производителя спецодежды и средств индивидуальной защиты в России ГК \"СИРИУС\" с 1998 года, сегодня признана одним из лидеров российского рынка", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo_1.png"},
		{"Российский онлайн-гипермаркет товаров для дома, дачи, стройки и ремонта", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo-filled.svg"},
		{"КРУПНЕЙШИЙ МЕТАЛЛМАРКЕТ РОССИИ", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo-w.svg"},
		{"Поставщик строительных конструкций, осуществляет поставку и аренду опалубочного оборудования.", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.png"},
		{"Поставщики крепежей", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.png83_83.webp"},
		{"Компания по производству крепежных деталей", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/logo.svg"},
		{"ООО «НЕВСКИЙ ТЕХНОЛОГ» – российская торгово-производственная компания, специализирующаяся на производстве и поставках металлоизделий для строительных и производственных нужд как крупным оптовым заказчикам, так и розничным покупателям.", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/LogoNevskyi.png"},
		{"AG-Technologies – крупнейший российский поставщик спецодежды и экипировки для профессионалов", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/mainLogo.png"},
		{"ООО «Титан-Монолит» предлагает широкий ассортимент товаров для железобетонного строительства: около 30 видов фиксаторов арматуры, а также расходные материалы и комплектующие для монолитных и общестроительных работ.", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/new-logo.svg"},
		{"Интернет-магазин для оптовых покупок: более 850 000 электротехнических товаров, удобные способы оплаты и доставки. Кабель, светотехника, низковольтное и щитовое оборудование.", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/rs24.PNG"},
		{"Строительный холдинг Setl Group работает в Петербурге с 1994 года. Setl Group является одним из крупнейших финансово-промышленных объединений Северо-западного региона России", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/partners/sg.09f1d80.svg"},
	}

	for _, partner := range partners {
		if _, err := stmt.ExecContext(ctx, partner.description, partner.url); err != nil {
			return fmt.Errorf("insert partner: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit partners: %w", err)
	}
	return nil
}

func seedProjects(ctx context.Context, db *sql.DB) error {
	var count int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM projects").Scan(&count); err != nil {
		return fmt.Errorf("count projects: %w", err)
	}
	if count > 0 {
		return nil
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin projects tx: %w", err)
	}
	defer tx.Rollback()

	projectStmt, err := tx.PrepareContext(ctx, `INSERT INTO projects (full_name, url, status, short_name, builder_name, body, latitude, longitude) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("prepare project insert: %w", err)
	}
	defer projectStmt.Close()

	type project struct {
		fullName    string
		url         string
		status      string
		shortName   string
		builderName string
		body        string
		latitude    float64
		longitude   float64
		pictures    []string
	}

	projects := []project{
		{
			fullName:    `ЖК "Сенат"`,
			url:         "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/preview.jpeg",
			status:      "В работе",
			shortName:   "Сенат",
			builderName: `ООО "Сэтл Строй"`,
			body:        "«Сенат в Московском» - роскошный жилой комплекс, созданный Setl Group. Этот прекрасный объект расположен в престижном районе Москвы с развитой социальной и торговой инфраструктурой, включающей супермаркеты, рестораны, спортивные комплексы, детские сады, школы, поликлиники и больницы. До ближайшей станции метро \"Московская\" всего 20 минут ходьбы, а до центра города - 20 минут езды на автомобиле. Рядом находятся прекрасные зеленые зоны, такие как Путиловский парк, Московский парк и дворцово-парковые ансамбли Пушкина и Павловска, до которых можно доехать за полчаса.\n\nПроект ЖК разработан известным архитектурным бюро «А.Лен». Каждый дом в этом жилом комплексе имеет уникальный фасад, который придает зданиям свой особый характер, сохраняя при этом общую композицию. На территории комплекса предусмотрены детские и спортивные площадки, сквер и пешеходный бульвар с изящным воздушным мостом. Особенностью ЖК станет наличие собственного ледового катка, который летом превращается в футбольное поле. Ландшафтный дизайн и озеленение будут выполнены по уникальному авторскому проекту.\n\nВ центральной входной группе жилого комплекса, занимающей более 100 квадратных метров, будут расположены ресепшен, лаундж-зоны и камин, создающие уютную атмосферу. Холлы домов будут отделаны итальянским мрамором с элементами дерева и металла. В каждой парадной будет представлена авторская интерьерная композиция и стильная дизайнерская мебель. Все дома оборудованы системами автоматического сбора данных счетчиков, онлайн-коммуникаций с управляющей компанией, системами IP-домофонии и внутренними датчиками.",
			latitude:    59.846195,
			longitude:   30.293631,
			pictures: []string{
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/preview.jpeg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/2.jpeg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/3.jpeg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/4.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/senat/5.jpeg",
			},
		},
		{
			fullName:    `ЖК "Imperial Club"`,
			url:         "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/preview.jpeg",
			status:      "В работе",
			shortName:   "Imperial Club",
			builderName: `ООО "Сэтл Строй"`,
			body:        "Imperial Club находится в центре Санкт-Петербурга, на первой линии Невы. Архитектурная композиция вписана в окружение с исключительными видами на Васильевский остров. В Imperial Club представлен уникальный формат квартир, двухуровневые сити-виллы, семейные квартиры с возможностью установки камина или зимних садов. Благоустройство подчеркивает особенности исторической части Васильевского острова, а роскошные интерьеры лобби дополняют премиальный сервис.",
			latitude:    59.927811,
			longitude:   30.266258,
			pictures: []string{
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/preview.jpeg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/2f783b839fff6e0c6ad74f09ac1a9789.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/7f9c63567869542407bc4cbaac023dc3.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/62d68aa57f5df.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/78f7f811451fcac815567d2058e06fe5.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/imperial_club/97b3c70580bee66f4fc9bcc202d6dbcd.jpg",
			},
		},
		{
			fullName:    `ЖК "Парадный Ансамбль"`,
			url:         "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/paradniy_ansambl/previewjpg.jpg",
			status:      "В работе",
			shortName:   "Парадный Ансамбль",
			builderName: `ООО "Сэтл Строй"`,
			body:        "«Парадный ансамбль» - проект с европейской философией загородной жизни. Архитектура гармонично вписана в зеленый ландшафт с фасадами из керамогранита и кирпича. Дворы будут закрыты от машин, предусмотрены подземные и наземные паркинги, рядом две станции метро. В проекте множество форматов планировок, благоустроенный пешеходный бульвар с ресторанами, площадями и спортивными зонами, а также четыре школы и семь детских садов.",
			latitude:    59.753372,
			longitude:   30.305025,
			pictures: []string{
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/paradniy_ansambl/previewjpg.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/paradniy_ansambl/1.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/paradniy_ansambl/2.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/paradniy_ansambl/3.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/paradniy_ansambl/4.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/paradniy_ansambl/5.jpg",
			},
		},
		{
			fullName:    `ЖК "Титул"`,
			url:         "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/preview.jpg",
			status:      "В работе",
			shortName:   "Титул",
			builderName: `ООО "Сэтл Строй"`,
			body:        "Жилой комплекс «Титул в Московском» — проект высокого комфорт-класса на Кубинской улице. Архитектура сочетает строгие линии и разнообразные оттенки кирпичной кладки с вечерней подсветкой. Центральные входные группы оформлены в классическом стиле, предусмотрены лаундж-зоны, библиотека и променад с садом камней. Для детей запланированы встроенные детские сады и лицей нового поколения, а также светящиеся качели и озелененные зоны для отдыха.",
			latitude:    59.843141,
			longitude:   30.298462,
			pictures: []string{
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/preview.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/1.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/2.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/3.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/4.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/5.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/6.jpg",
				"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/projects/titul/7.jpg",
			},
		},
	}

	pictureStmt, err := tx.PrepareContext(ctx, `INSERT INTO pictures (project_id, url) VALUES (?, ?)`)
	if err != nil {
		return fmt.Errorf("prepare picture insert: %w", err)
	}
	defer pictureStmt.Close()

	for _, p := range projects {
		res, err := projectStmt.ExecContext(ctx, p.fullName, p.url, p.status, p.shortName, p.builderName, p.body, p.latitude, p.longitude)
		if err != nil {
			return fmt.Errorf("insert project %s: %w", p.fullName, err)
		}
		projectID, err := res.LastInsertId()
		if err != nil {
			return fmt.Errorf("get project id for %s: %w", p.fullName, err)
		}

		for _, pic := range p.pictures {
			if _, err := pictureStmt.ExecContext(ctx, projectID, pic); err != nil {
				return fmt.Errorf("insert picture for %s: %w", p.fullName, err)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit projects: %w", err)
	}

	return nil
}

func seedDocs(ctx context.Context, db *sql.DB) error {
	var count int
	if err := db.QueryRowContext(ctx, "SELECT COUNT(*) FROM docs").Scan(&count); err != nil {
		return fmt.Errorf("count docs: %w", err)
	}
	if count > 0 {
		return nil
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin docs tx: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO docs (preview_url, document_url) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("prepare docs insert: %w", err)
	}
	defer stmt.Close()

	docs := [][2]string{
		{"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/благодарность_гаврилов.jpg", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/2 Благодарность Гаврилов.pdf"},
		{"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/благодарность_грачев.jpg", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/1 Благодарность Грачев.pdf"},
		{"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/благодарность_гретчин.jpg", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/3 Благодарность Гретчин.pdf"},
		{"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/сертификат_партнера_дока_рус.PNG", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/Сертификат партнера ООО Дока Рус.pdf"},
		{"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/СРО_page-0001.jpg", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/сро.pdf"},
		{"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/Necessarily_1.PNG", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/2023-11-14_002.pdf"},
		{"https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/preview/Necessarily_2.PNG", "https://s3.timeweb.cloud/9d596ed3-c39809ab-5902-491e-8139-9af489d08762/docs/documents/2023-11-14_001.pdf"},
	}

	for _, doc := range docs {
		if _, err := stmt.ExecContext(ctx, doc[0], doc[1]); err != nil {
			return fmt.Errorf("insert doc: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit docs: %w", err)
	}
	return nil
}
