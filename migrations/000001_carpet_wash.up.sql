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
  id UUID PRIMARY KEY DEFAULT gen_random_uui  d(),
  tariffs  text,
  name VARCHAR(100) not null,
  description TEXT,  
  price DECIMAL(10,2)
);

CREATE Table addresses (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  latitude DECIMAL(9,6),
  longitude DECIMAL(9,6),
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at int  DEFAULT 1      --0 bo'lsa o'chirilgan bo'ladi
);

CREATE Table orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE,
  service_id UUID REFERENCES services(id) ON DELETE CASCADE,
  area float NOT NULL,
  total_price float, 
  status VARCHAR(50),
  created_at TIMESTAMP default CURREENT_TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at int DEFAULT 1   --0 bo'lsa o'chirilgan bo'ladi 
);

