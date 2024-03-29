
-- Установить временную зону
SET TIMEZONE="Europe/Moscow";

-- создать таблицу номеров
CREATE TABLE books (
                     id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
                     created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
                     updated_at TIMESTAMP NULL,
                     tg_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
                     title VARCHAR (255) NOT NULL,
                     author VARCHAR (255) NOT NULL,
                     number_status INT NOT NULL
);

-- Создать индекс
CREATE INDEX active_numbers ON books (title) WHERE book_status = 1;
