CREATE Table company (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(100) NOT NULL,
  description TEXT,
  logo_url VARCHAR(255),
  status VARCHAR(20) default 'active',
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP default CURRENT_TIMESTAMP,
)


CREATE Table services (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  tariffs  text,
  company_id INT not null,
  name VARCHAR(100) not null,
  description TEXT,
  price DECIMAL(10,2),
)

CREATE Table addresses (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id),
  latitude DECIMAL(9,6),
  longitude DECIMAL(9,6),
  created_at TIMESTAMP default CURRENT_TIMESTAMP,
  updated_at TIMESTAMP default CURRENT_TIMESTAMP,
)

CREATE Table orders (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID REFERENCES users(id),
  service_id UUID REFERENCES services(id),
  address_id UUID REFERENCES addresses(id),
  created_at timestamp default CURREENT_TIMESTAMP,
)
