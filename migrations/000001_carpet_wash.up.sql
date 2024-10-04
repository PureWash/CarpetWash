CREATE Table company (
                         id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                         name VARCHAR(100) NOT NULL,
                         description TEXT,
                         logo_url VARCHAR,
                         created_at TIMESTAMP default CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP,
                         deleted_at int DEFAULT 1     --0 bo'lsa o'chirilgan bo'ladi
);

CREATE Table services (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          tariffs  text,
                          name VARCHAR(100) not null,
                          description TEXT,
                          price DECIMAL(10,2)
);

CREATE Table addresses (
                           id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                           user_id UUID ,
                           latitude VARCHAR(100),
                           longitude VARCHAR(100),
                           created_at TIMESTAMP default CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP,
                           deleted_at int DEFAULT 1
);

CREATE Table orders (
                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                        user_id UUID ,
                        service_id UUID REFERENCES services(id) ON DELETE CASCADE,
                        area float NOT NULL,
                        total_price float,
                        status VARCHAR(50),
                        created_at TIMESTAMP default CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP,
                        deleted_at int DEFAULT 1
);


SELECT
    id,
    (SELECT  jsonb_agg(
             json_build_object(
             'id':u.id,
             'username':u.username,
             'full_name':u.username,
             'full_name':u.phone_number,
             )
             ) FROM users  u WHERE u.id=o.user_id ) AS USER_DETAILS
       ,
    (SELECT  jsonb_agg(
             jsonb_build_object(
              'id':s.id,
             'tariffs':s.tariffs,
             'name':s.name,
             'description':s.description,
             'price':s.price,
             )
             ) FROM  orders WHERE s.id=o.service_id ) AS  SERVICE_DETAILS,
    area,
    total_price,
    status,
    created_at,
    updated_at, deleted_at
    FROM orders o
 WHERE o.deleted_at='1' AND o.area::text ILIKE $1 OR
 o.total_price::text ILIKE $1
 o.status ILIKE $1
 EXISTS (
        SELECT 1
        FROM services s
        WHERE s.id = d.service_id AND (
            s.tariffs ILIKE $1 OR
            s.name ILIKE $1 OR
            s.description:text  ILIKE $1 OR
            s.price::text  ILIKE $1 OR
        )
    )
LIMIT @limit_ OFFSET @offset_
;