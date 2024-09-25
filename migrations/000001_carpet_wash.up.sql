CREATE Table company (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  description TEXT,
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);


CREATE Table services (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tariffs  text,
  company_id INT not null,
  name VARCHAR(100) not null,
  description TEXT,
  price DECIMAL(10,2)
);

CREATE Table addresses (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID,
  latitude VARCHAR(100) not null,
  longitude VARCHAR(100) not null,
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE Table orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID ,
  service_id UUID REFERENCES services(id) ON DELETE CASCADE,
  address_id UUID REFERENCES addresses(id) ON DELETE CASCADE,
  created_at TIMESTAMP default CURRENT_TIMESTAMP
);
